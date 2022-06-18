// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type SNA struct {
	SessionID                   datatype.UTF8String       `avp:"Session-Id"`
	DRMP                        datatype.Enumerated       `avp:"DRMP"`
	VendorSpecificApplicationId datatype.Grouped          `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated       `avp:"Auth-Session-State"`
	ResultCode                  uint32                    `avp:"Result-Code"`
	ExperimentalResult          datatype.Grouped          `avp:"Experimental-Result"`
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm"`
	WildcardedPublicIdentity    datatype.UTF8String       `avp:"Wildcarded-Public-Identity"`
	WildcardedIMPU              datatype.UTF8String       `avp:"Wildcarded-IMPU"`
	SupportedFeatures           datatype.Grouped          `avp:"Supported-Featrues"`
	UserDataSh                  datatype.OctetString      `avp:"User-Data-Sh"`
	ExpiryTime                  datatype.Time             `avp:"Expiry-Time"`
	OCSupportedFeatures         datatype.Grouped          `avp:"OC-Supported-Features"`
	OCOLR                       datatype.Grouped          `avp:"OC-OLR"`
	FailedAVP                   datatype.Grouped          `avp:"Failed-AVP"`
	ProxyInfo                   datatype.Grouped          `avp:"Proxy-Info"`
	RouteRecord                 datatype.DiameterIdentity `avp:"Route-Record"`
}

// Parse parses the given message.
func (sna *SNA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(sna); err != nil {
		return err
	}
	return nil
}
