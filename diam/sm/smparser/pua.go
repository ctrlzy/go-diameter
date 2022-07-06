// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type PUA struct {
	SessionID                   datatype.UTF8String       `avp:"Session-Id"`
	DRMP                        datatype.Enumerated       `avp:"DRMP"`
	VendorSpecificApplicationId datatype.Grouped          `avp:"Vendor-Specific-Application-Id"`
	ResultCode                  uint32                    `avp:"Result-Code"`
	ExperimentalResult          datatype.Grouped          `avp:"Experimental-Result"`
	AuthSessionState            datatype.Enumerated       `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm"`
	WildcardedPublicIdentity    datatype.UTF8String       `avp:"Wildcarded-Public-Identity"`
	WildcardedIMPU              datatype.UTF8String       `avp:"Wildcarded-IMPU"`
	SupportedFeatures           datatype.Grouped          `avp:"Supported-Featrues"`
	FailedAVP                   datatype.Grouped          `avp:"Failed-AVP"`
	ProxyInfo                   datatype.Grouped          `avp:"Proxy-Info"`
	RouteRecord                 datatype.DiameterIdentity `avp:"Route-Record"`
}

// Parse parses the given message.
func (pua *PUA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(pua); err != nil {
		return err
	}
	return nil
}
