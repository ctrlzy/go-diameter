// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type SNR struct {
	SessionID                   datatype.UTF8String                     `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                    `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity              `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity               `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	UserIdentity                basetype.User_Identity                  `avp:"User-Identity"`
	WildcardedPublicIdentity    *datatype.UTF8String                    `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                    `avp:"Wildcarded-IMPU,omitempty"`
	ServiceIndication           []datatype.OctetString                  `avp:"Service-Indication,omitempty"`
	SendDataIndication          *datatype.Enumerated                    `avp:"Send-Data-Indication,omitempty"`
	ServerName                  *datatype.UTF8String                    `avp:"Server-Name,omitempty"`
	SubsReqType                 datatype.Enumerated                     `avp:"Subs-Req-Type"`
	DataReference               []datatype.Enumerated                   `avp:"Data-Reference"`
	IdentitySet                 []datatype.Enumerated                   `avp:"Identity-Set,omitempty"`
	ExpiryTime                  *datatype.Time                          `avp:"Expiry-Time,omitempty"`
	DSAITag                     []datatype.OctetString                  `avp:"DSAI-Tag,omitempty"`
	OneTimeNotification         *datatype.Enumerated                    `avp:"One-Time-Notification,omitempty"`
	UserName                    *datatype.UTF8String                    `avp:"User-Name,omitempty"`
	OCSupportedFeatures         *basetype.OC_Supported_Features         `avp:"OC-Supported-Features,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (snr *SNR) Parse(m *diam.Message) error {
	err := m.Unmarshal(snr)
	if err != nil {
		return nil
	}
	if err = snr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (snr *SNR) sanityCheck() error {
	if len(snr.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if snr.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(snr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(snr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(snr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if snr.UserIdentity.Empty() {
		return ErrMissingUserIdentity
	}
	if len(snr.DataReference) == 0 {
		return ErrMissingDataReference
	}
	return nil
}
