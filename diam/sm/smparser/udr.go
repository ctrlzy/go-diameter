// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type UDR struct {
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
	ServiceIndication           datatype.OctetString      `avp:"Service-Indication"`
	ServerName                  datatype.UTF8String       `avp:"Server-Name"`
	RequestedNodes              datatype.Unsigned32       `avp:"Requested-Nodes"`
	ServingNodeIndication       datatype.Enumerated       `avp:"Serving-Node-Indication"`
	PrePagingSupported          datatype.Enumerated       `avp:"Pre-paging-Supported"`
	LocalTimeZoneIndication     datatype.Enumerated       `avp:"Local-Time-Zone-Indication"`
	UDRFlags                    datatype.Unsigned32       `avp:"UDR-Flags"`
	CallReferenceInfo           datatype.Grouped          `avp:"Call-Reference-Info"`
	OCSupportedFeatures         datatype.Grouped          `avp:"OC-Supported-Features"`
	ProxyInfo                   datatype.Grouped          `avp:"Proxy-Info"`
	RouteRecord                 datatype.DiameterIdentity `avp:"Route-Record"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (udr *UDR) Parse(m *diam.Message) error {
	err := m.Unmarshal(udr)
	if err != nil {
		return nil
	}
	if err = udr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (udr *UDR) sanityCheck() error {
	if len(udr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(udr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}
