// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type PNR struct {
	SessionID                   datatype.UTF8String                     `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                    `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity               `avp:"Destination-Host"`
	DestinationRealm            datatype.DiameterIdentity               `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Featrues,omitempty"`
	UserIdentity                basetype.User_Identity                  `avp:"User-Identity"`
	WildcardedPublicIdentity    *datatype.UTF8String                    `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                    `avp:"Wildcarded-IMPU,omitempty"`
	UserName                    *datatype.UTF8String                    `avp:"User-Name,omitempty"`
	UserData                    datatype.OctetString                    `avp:"User-Data"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (pnr *PNR) Parse(m *diam.Message) error {
	err := m.Unmarshal(pnr)
	if err != nil {
		return nil
	}
	if err = pnr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (pnr *PNR) sanityCheck() error {
	if len(pnr.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if pnr.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(pnr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(pnr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(pnr.DestinationHost) == 0 {
		return ErrMissingDestHost
	}
	if len(pnr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if pnr.UserIdentity.Empty() {
		return ErrMissingUserIdentity
	}
	if len(pnr.UserData) == 0 {
		return ErrMissingUserData
	}
	return nil
}
