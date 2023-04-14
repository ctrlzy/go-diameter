// Copyright 2023 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFR refers to Mo-Forward-Short-Message-Request.
// See 3GPP TS 29.338 Clause 6.2.1 for details
type OFR struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity               `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity                `avp:"Destination-Realm"`
	ScAddress                   datatype.OctetString                     `avp:"SC-Address"`
	OfrFlags                    *datatype.Unsigned32                     `avp:"OFR-Flags,omitempty"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	UserIdentifier              basetype.User_Identifier                 `avp:"User-Identifier"`
	SmRpUi                      datatype.OctetString                     `avp:"SM-RP-UI"`
	SmsmiCorrelationId          *basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	SmDeliveryOutcome           *basetype.SM_Delivery_Outcome            `avp:"SM-Delivery-Outcome,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (ofr *OFR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(ofr); err != nil {
		return err
	}
	if err := ofr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (ofr *OFR) sanityCheck() error {
	if len(ofr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(ofr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(ofr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if len(ofr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	if ofr.UserIdentifier.Empty() {
		return ErrMissingUserIdentifier
	}
	if len(ofr.SmRpUi) == 0 {
		return ErrMissingSmRpUi
	}
	return nil
}
