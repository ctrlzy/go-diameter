// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// refer to 3GPP TS 29.329 6.1.8 Push-Notification-Answer
type PNA struct {
	SessionID                   datatype.UTF8String                  `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                 `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id"`
	ResultCode                  *datatype.Unsigned32                 `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult         `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                  `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity            `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity            `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures         `avp:"Supported-Features,omitempty"`
	FailedAVP                   basetype.FailedAVP                   `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                 `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity          `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (pna *PNA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(pna); err != nil {
		return err
	}
	if err := pna.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (pna *PNA) sanityCheck() error {
	if len(pna.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if pna.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(pna.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(pna.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if pna.ResultCode == nil && pna.ExperimentalResult == nil {
		return ErrMissingResultCode
	}
	return nil
}
