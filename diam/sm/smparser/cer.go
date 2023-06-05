// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// CER is a Capabilities-Exchange-Request message.
// See RFC 6733 section 5.3.1 for details.
type CER struct {
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm"`
	OriginStateID               *diam.AVP                 `avp:"Origin-State-Id"`
	InbandSecurityID            *diam.AVP                 `avp:"Inband-Security-Id"`
	AcctApplicationID           []*diam.AVP               `avp:"Acct-Application-Id"`
	AuthApplicationID           []*diam.AVP               `avp:"Auth-Application-Id"`
	VendorSpecificApplicationID []*diam.AVP               `avp:"Vendor-Specific-Application-Id"`
	appID                       []uint32                  // List of supported application IDs.
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok, and all accounting or authentication applications
// in the CER match the applications in our dictionary. If one or more
// mandatory AVPs are missing, it returns a nil failedAVP and a proper
// error. If all mandatory AVPs are present but no common application
// is found, then it returns the failedAVP (with the application that
// we don't support in our dictionary) and an error. Another cause
// for error is the presence of Inband Security, we don't support that.
func (cer *CER) Parse(m *diam.Message, localRole Role) (failedAVP *diam.AVP, err error) {
	if err = m.Unmarshal(cer); err != nil {
		return nil, err
	}
	if err = cer.sanityCheck(); err != nil {
		return nil, err
	}
	if cer.InbandSecurityID != nil {
		if v := cer.InbandSecurityID.Data.(datatype.Unsigned32); v != 0 {
			return nil, ErrNoCommonSecurity
		}
	}
	app := &Application{
		AcctApplicationID:           cer.AcctApplicationID,
		AuthApplicationID:           cer.AuthApplicationID,
		VendorSpecificApplicationID: cer.VendorSpecificApplicationID,
	}
	if failedAVP, err = app.Parse(m.Dictionary(), localRole); err != nil {
		return failedAVP, err
	}
	cer.appID = app.ID()
	return nil, nil
}

// sanityCheck ensures mandatory AVPs are present.
func (cer *CER) sanityCheck() error {
	if len(cer.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(cer.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

// Applications return a list of supported Application IDs.
func (cer *CER) Applications() []uint32 {
	return cer.appID
}

func (r *CER) String() string {
	result := "CER { "
	result += fmt.Sprintf("OriginHost: %s, ", r.OriginHost)
	result += fmt.Sprintf("OriginRealm: %s, ", r.OriginRealm)
	if r.OriginStateID != nil {
		result += fmt.Sprintf("OriginStateID: %s, ", r.OriginStateID.String())
	}
	if r.InbandSecurityID != nil {
		result += fmt.Sprintf("InbandSecurityID: %s, ", r.InbandSecurityID.String())
	}
	result += "AcctApplicationID: ["
	for i, avp := range r.AcctApplicationID {
		if i > 0 {
			result += ", "
		}
		result += avp.String()
	}
	result += "], "
	result += "AuthApplicationID: ["
	for i, avp := range r.AuthApplicationID {
		if i > 0 {
			result += ", "
		}
		result += avp.String()
	}
	result += "], "
	result += "VendorSpecificApplicationID: ["
	for i, avp := range r.VendorSpecificApplicationID {
		if i > 0 {
			result += ", "
		}
		result += avp.String()
	}
	result += "], "
	result += fmt.Sprintf("appID: %v, ", r.appID)
	result += "}"
	return result
}
