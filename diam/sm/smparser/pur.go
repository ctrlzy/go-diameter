// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// refer to 3GPP TS 29.329 6.1.3 Profile-Update-Request
type PUR struct {
	SessionID                   datatype.UTF8String                  `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                 `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated                  `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity            `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity            `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity           `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity            `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures         `avp:"Supported-Features,omitempty"`
	UserIdentity                basetype.UserIdentity                `avp:"User-Identity"`
	WildcardedPublicIdentity    *datatype.UTF8String                 `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                 `avp:"Wildcarded-IMPU,omitempty"`
	UserName                    *datatype.UTF8String                 `avp:"User-Name,omitempty"`
	DataReference               []datatype.Enumerated                `avp:"Data-Reference"`
	UserData                    datatype.OctetString                 `avp:"User-Data"`
	OCSupportedFeatures         *basetype.OCSupportedFeatures        `avp:"OC-Supported-Features,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                 `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity          `avp:"Route-Record,omitempty"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (pur *PUR) Parse(m *diam.Message) error {
	err := m.Unmarshal(pur)
	if err != nil {
		return nil
	}
	if err = pur.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (pur *PUR) sanityCheck() error {
	if len(pur.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if pur.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(pur.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(pur.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(pur.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if pur.UserIdentity.Empty() {
		return ErrMissingUserIdentity
	}
	if len(pur.DataReference) == 0 {
		return ErrMissingDataReference
	}
	if len(pur.UserData) == 0 {
		return ErrMissingUserData
	}
	return nil
}
