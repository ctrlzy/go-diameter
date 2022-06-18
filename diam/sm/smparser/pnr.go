// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type PNR struct {
	SessionID                   datatype.UTF8String       `avp:"Session-Id"`
	DRMP                        datatype.Enumerated       `avp:"DRMP"`
	VendorSpecificApplicationId datatype.Grouped          `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated       `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity `avp:"Destination-Host"`
	DestinationRealm            datatype.DiameterIdentity `avp:"Destination-Realm"`
	SupportedFeatures           datatype.Grouped          `avp:"Supported-Featrues"`
	UserIdentity                datatype.Grouped          `avp:"User-Identity"`
	WildcardedPublicIdentity    datatype.UTF8String       `avp:"Wildcarded-Public-Identity"`
	WildcardedIMPU              datatype.UTF8String       `avp:"Wildcarded-IMPU"`
	UserName                    string                    `avp:"User-Name"`
	UserDataSh                  datatype.OctetString      `avp:"User-Data-Sh"`
	ProxyInfo                   datatype.Grouped          `avp:"Proxy-Info"`
	RouteRecord                 datatype.DiameterIdentity `avp:"Route-Record"`
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
	if len(pnr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(pnr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}
