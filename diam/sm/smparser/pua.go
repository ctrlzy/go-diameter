// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// refer to 3GPP TS 29.329 6.1.4 Profile-Update-Answer
type PUA struct {
	SessionID                   datatype.UTF8String                  `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                 `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id"`
	ResultCode                  *datatype.Unsigned32                 `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult         `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                  `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity            `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity            `avp:"Origin-Realm"`
	WildcardedPublicIdentity    *datatype.UTF8String                 `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                 `avp:"Wildcarded-IMPU,omitempty"`
	RepositoryDataID            *basetype.RepositoryDataID           `avp:"Repository-Data-ID,omitempty"`
	DataReference               datatype.Enumerated                  `avp:"Data-Reference,omitempty"`
	SupportedFeatures           []basetype.SupportedFeatures         `avp:"Supported-Features,omitempty"`
	OCSupportedFeatures         *basetype.OCSupportedFeatures        `avp:"OC-Supported-Features,omitempty"`
	OCOLR                       *basetype.OCOLR                      `avp:"OC-CLR,omitempty"`
	Load                        *basetype.Load                       `avp:"Load,omitempty"`
	FailedAVP                   basetype.FailedAVP                   `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                 `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity          `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (pua *PUA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(pua); err != nil {
		return err
	}
	if err := pua.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (pua *PUA) sanityCheck() error {
	if len(pua.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if pua.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(pua.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(pua.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if pua.ResultCode == nil && pua.ExperimentalResult == nil {
		return ErrMissingResultCode
	}
	return nil
}
