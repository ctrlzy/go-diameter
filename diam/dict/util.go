// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// Dictionary parser, helper functions.  Part of go-diameter.

package dict

import (
	"errors"
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// parentAppIds map allows for hierarchical AVP search dependencies
// If an AVP code was not found in the app's dictionary, the search will continue in the parent app
// dictionary and only then in base diameter dictionary
// Note: care must be taken to avoid creating parent-child loops
var parentAppIds map[uint32]uint32 = map[uint32]uint32{
	16777251: 4,
	16777238: 4,
	4:        1,
}

// Apps return a list of all applications loaded in the Parser object.
// Apps must never be called concurrently with LoadFile or Load.
func (p *Parser) Apps() []*App {
	//p.mu.Lock()
	//defer p.mu.Unlock()
	var apps []*App
	for _, f := range p.file {
		for _, app := range f.App {
			apps = append(apps, app)
		}
	}
	return apps
}

// App returns a dictionary application for the given application code
// if exists. App must never be called concurrently with LoadFile or Load.
func (p *Parser) App(code uint32, typ ...string) (*App, error) {
	var app *App
	if len(typ) > 0 {
		app = p.apptype[appIdTypeIdx{code, typ[0]}]
	}
	if app != nil {
		return app, nil
	}
	app = p.appcode[code]
	if app != nil && (len(app.Type) == 0 || len(typ) == 0 || app.Type == typ[0]) {
		return app, nil
	}
	return nil, ErrApplicationUnsupported
}

// ErrApplicationUnsupported indicates that the application requested
// does not exist in the dictionary Parser object.
var ErrApplicationUnsupported = errors.New("application unsupported")

func MakeUnknownAVP(appid, code, vendorID uint32) *AVP {
	return &AVP{
		Name:     fmt.Sprintf("Unknown-%d-%d", code, vendorID),
		Code:     code,
		VendorID: vendorID,
		Data: Data{
			Type:     datatype.UnknownType,
			TypeName: "Unknown",
		},
		App: &App{
			ID:     appid,
			Vendor: []*Vendor{&Vendor{ID: vendorID}},
		},
	}
}

// FindAVPWithVendor is a helper function that returns a pre-loaded AVP from the Parser, considering vendorID as filter.
// For no vendorID filter, use UndefinedVendorID constant
// If the AVP code is not found for the given appid it tries with appid=0
// before returning an error.
// Code can be either the AVP code (int, uint32) or name (string).
//
// FindAVPWithVendor must never be called concurrently with LoadFile or Load.
func (p *Parser) FindAVPWithVendor(appid uint32, code interface{}, vendorID uint32) (*AVP, error) {
	return p.findAVPWithVendorRecursive(appid, code, vendorID, appid)
}

// find AVP recursively in the dictionary
func (p *Parser) findAVPWithVendorRecursive(appid uint32, code interface{}, vendorID uint32, origAppID uint32) (*AVP, error) {
	var (
		avp *AVP
		ok  bool
	)

	switch codeVal := code.(type) {
	case string:
		avp, ok = p.avpname[nameIdx{appid, codeVal, vendorID}]
	case uint32:
		avp, ok = p.avpcode[codeIdx{appid, uint32(codeVal), vendorID}]
	case int:
		avp, ok = p.avpcode[codeIdx{appid, uint32(codeVal), vendorID}]
	default:
		return nil, fmt.Errorf("unsupported AVP code type %T(%#v)", codeVal, code)
	}

	// return avp if found
	if ok {
		return avp, nil
	}

	// return unknown AVP or error if not found in base app
	if !ok && appid == 0 {
		if codeU32, isUint32 := code.(uint32); isUint32 {
			return MakeUnknownAVP(origAppID, codeU32, vendorID), nil
		}
		return nil, fmt.Errorf("could not find AVP %v for Vendor: %d", code, vendorID)
	}

	// didn't find in current app, go to parent app for search
	nextApp := uint32(0)
	if parentAppId, isScopedApp := parentAppIds[appid]; isScopedApp {
		nextApp = parentAppId
	}

	// didn't find the parent app, go to base app for search
	return p.findAVPWithVendorRecursive(nextApp, code, vendorID, origAppID)
}

// FindAVP is a helper function that returns a pre-loaded AVP from the Parser.
// If the AVP code is not found for the given appid it tries with appid=0
// before returning an error.
// Code can be either the AVP code (int, uint32) or name (string).
//
// FindAVP must never be called concurrently with LoadFile or Load.
func (p *Parser) FindAVP(appid uint32, code interface{}) (*AVP, error) {
	return p.FindAVPWithVendor(appid, code, UndefinedVendorID)
}

// ScanAVP is a helper function that returns a pre-loaded AVP from the Dict.
// It's similar to FindAPI except that it scans the list of available AVPs
// instead of looking into one specific appid.
//
// ScanAVP is 20x or more slower than FindAVP. Use with care.
// Code can be either the AVP code (uint32) or name (string).
//
// ScanAVP must never be called concurrently with LoadFile or Load.
func (p *Parser) ScanAVP(code interface{}) (*AVP, error) {
	//p.mu.Lock()
	//defer p.mu.Unlock()
	switch codeVal := code.(type) {
	case string:
		for idx, avp := range p.avpname {
			if idx.name == codeVal {
				return avp, nil
			}
		}
		return nil, fmt.Errorf("could not find AVP %s", code.(string))
	case uint32:
		for idx, avp := range p.avpcode {
			if idx.code == codeVal {
				return avp, nil
			}
		}
		return nil, fmt.Errorf("could not find AVP code %d", code.(uint32))
	case int:
		for idx, avp := range p.avpcode {
			if idx.code == uint32(codeVal) {
				return avp, nil
			}
		}
		return nil, fmt.Errorf("could not find AVP code %d", code.(int))
	}
	return nil, fmt.Errorf("unsupported AVP code type %#v", code)
}

// FindCommand returns a pre-loaded Command from the Parser.
//
// FindCommand must never be called concurrently with LoadFile or Load.
func (p *Parser) FindCommand(appid, code uint32) (*Command, error) {
	//p.mu.Lock()
	//defer p.mu.Unlock()
	if cmd, ok := p.command[codeIdx{appid, code, UndefinedVendorID}]; ok {
		return cmd, nil
	}
	// Always fall back to base dict.
	if cmd, ok := p.command[codeIdx{0, code, UndefinedVendorID}]; ok {
		return cmd, nil
	}
	return nil, fmt.Errorf("could not find preloaded Command with code %d", code)
}

// Enum is a helper function that returns a pre-loaded Enum item for the
// given AVP appid, code and n. (n is the enum code in the dictionary)
//
// Enum must never be called concurrently with LoadFile or Load.
func (p *Parser) Enum(appid, code uint32, n int32) (*Enum, error) {
	avp, err := p.FindAVP(appid, code)
	if err != nil {
		return nil, err
	}
	if avp.Data.Type != datatype.EnumeratedType {
		return nil, fmt.Errorf("data of AVP %s (%d) data is not Enumerated", avp.Name, avp.Code)
	}
	for _, item := range avp.Data.Enum {
		if item.Code == n {
			return item, nil
		}
	}
	return nil, fmt.Errorf("could not find preload Enum %d for AVP %s (%d)", n, avp.Name, avp.Code)
}

// Rule is a helper function that returns a pre-loaded Rule item for the
// given AVP code and name.
//
// Rule must never be called concurrently with LoadFile or Load.
func (p *Parser) Rule(appid, code uint32, n string) (*Rule, error) {
	avp, err := p.FindAVP(appid, code)
	if err != nil {
		return nil, err
	}
	if avp.Data.Type != datatype.GroupedType {
		return nil, fmt.Errorf("data of AVP %s (%d) data is not Grouped", avp.Name, avp.Code)
	}
	for _, item := range avp.Data.Rule {
		if item.AVP == n {
			return item, nil
		}
	}
	return nil, fmt.Errorf("could not find preload Rule for %s for AVP %s (%d)", n, avp.Name, avp.Code)
}
