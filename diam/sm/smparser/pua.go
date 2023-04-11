// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type PUA struct {
	SessionID                   datatype.UTF8String                     `avp:"Session-Id"`
	DRMP                        datatype.Enumerated                     `avp:"DRMP"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id"`
	ResultCode                  datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult          datatype.Grouped                        `avp:"Experimental-Result"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	WildcardedPublicIdentity    datatype.UTF8String                     `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              datatype.UTF8String                     `avp:"Wildcarded-IMPU,omitempty"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	FailedAVP                   datatype.Grouped                        `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   datatype.Grouped                        `avp:"Proxy-Info,omitempty"`
	RouteRecord                 datatype.DiameterIdentity               `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (pua *PUA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(pua); err != nil {
		return err
	}
	return nil
}
