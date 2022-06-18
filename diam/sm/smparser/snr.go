// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type SNR struct {
	SessionID                   datatype.UTF8String       `avp:"Session-Id"`
	DRMP                        datatype.Enumerated       `avp:"DRMP"`
	VendorSpecificApplicationId datatype.Grouped          `avp:"Vendor-Specific-Application-Id"`
	AuthSessionState            datatype.Enumerated       `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity `avp:"Destination-Host"`
	DestinationRealm            datatype.DiameterIdentity `avp:"Destination-Realm"`
	SupportedFeatures           datatype.Grouped          `avp:"Supported-Features"`
	UserIdentity                datatype.Grouped          `avp:"User-Identity"`
	WildcardedPublicIdentity    datatype.UTF8String       `avp:"Wildcarded-Public-Identity"`
	WildcardedIMPU              datatype.UTF8String       `avp:"Wildcarded-IMPU"`
	ServiceIndication           datatype.OctetString      `avp:"Service-Indication"`
	SendDataIndication          datatype.Enumerated       `avp:"Send-Data-Indication"`
	ServerName                  datatype.UTF8String       `avp:"Server-Name"`
	SubsReqType                 datatype.Enumerated       `avp:"Subs-Req-Type"`
	DataReference               datatype.Enumerated       `avp:"Data-Reference"`
	IdentitySet                 datatype.Enumerated       `avp:"Identity-Set"`
	ExpiryTime                  datatype.Time             `avp:"Expiry-Time"`
	DSAITag                     datatype.OctetString      `avp:"DSAI-Tag"`
	OneTimeNotification         datatype.Enumerated       `avp:"One-Time-Notification"`
	UserName                    string                    `avp:"User-Name"`
	OCSupportedFeatures         datatype.Grouped          `avp:"OC-Supported-Features"`
	ProxyInfo                   datatype.Grouped          `avp:"Proxy-Info"`
	RouteRecord                 datatype.DiameterIdentity `avp:"Route-Record"`
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
	if len(snr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(snr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}
