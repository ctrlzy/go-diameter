// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type UDR struct {
	SessionID                   datatype.UTF8String                     `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                    `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity              `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity               `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Featrues,omitempty"`
	UserIdentity                basetype.User_Identity                  `avp:"User-Identity"`
	WildcardedPublicIdentity    *datatype.UTF8String                    `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                    `avp:"Wildcarded-IMPU,omitempty"`
	ServerName                  *datatype.UTF8String                    `avp:"Server-Name,omitempty"`
	ServiceIndication           []datatype.OctetString                  `avp:"Service-Indication,omitempty"`
	DataReference               []datatype.Enumerated                   `avp:"Data-Reference"`
	IdentitySet                 []datatype.Enumerated                   `avp:"Identity-Set,omitempty"`
	RequestedDomain             *datatype.Enumerated                    `avp:"Requested-Domain,omitempty"`
	CurrentLocation             *datatype.Enumerated                    `avp:"Current-Location,omitempty"`
	DsaiTag                     []datatype.OctetString                  `avp:"DSAI-Tag,omitempty"`
	SessionPriority             *datatype.Enumerated                    `avp:"Session-Priority,omitempty"`
	UserName                    *datatype.UTF8String                    `avp:"User-Name,omitempty"`
	RequestedNodes              *datatype.Unsigned32                    `avp:"Requested-Nodes,omitempty"`
	ServingNodeIndication       *datatype.Enumerated                    `avp:"Serving-Node-Indication,omitempty"`
	PrePagingSupported          *datatype.Enumerated                    `avp:"Pre-paging-Supported,omitempty"`
	LocalTimeZoneIndication     *datatype.Enumerated                    `avp:"Local-Time-Zone-Indication,omitempty"`
	UDRFlags                    *datatype.Unsigned32                    `avp:"UDR-Flags,omitempty"`
	CallReferenceInfo           *basetype.Call_Reference_Info           `avp:"Call-Reference-Info,omitempty"`
	OCSupportedFeatures         *basetype.OC_Supported_Features         `avp:"OC-Supported-Features,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
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
	if len(udr.SessionID) == 0 {
		return ErrMissingSessionID
	}

	if udr.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(udr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(udr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(udr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if udr.UserIdentity.Empty() {
		return ErrMissingUserIdentity
	}
	if len(udr.DataReference) == 0 {
		return ErrMissingDataReference
	}
	return nil
}
