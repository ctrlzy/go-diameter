// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type SNA struct {
	SessionID                   datatype.UTF8String                     `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                    `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	ResultCode                  *datatype.Unsigned32                    `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.Experimental_Result           `avp:"Experimental-Result,omitempty"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	WildcardedPublicIdentity    *datatype.UTF8String                    `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                    `avp:"Wildcarded-IMPU,omitempty"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Featrues,omitempty"`
	UserData                    *datatype.OctetString                   `avp:"User-Data,omitempty"`
	ExpiryTime                  *datatype.Time                          `avp:"Expiry-Time,omitempty"`
	OCSupportedFeatures         *basetype.OC_Supported_Features         `avp:"OC-Supported-Features,omitempty"`
	OCOLR                       *basetype.OC_OLR                        `avp:"OC-OLR,omitempty"`
	Load                        []basetype.Load                         `avp:"Load,omitempty"`
	FailedAVP                   *basetype.Failed_AVP                    `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (sna *SNA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(sna); err != nil {
		return err
	}
	if err := sna.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (sna *SNA) sanityCheck() error {
	if len(sna.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if sna.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(sna.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(sna.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if sna.ResultCode == nil && sna.ExperimentalResult == nil {
		return ErrMissingResultCode
	}
	return nil
}
