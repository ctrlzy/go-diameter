// Copyright 2023 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFR refers to Mo-Forward-Short-Message-Request.
// See 3GPP TS 29.338 Clause 6.3.2.3 for details
type OFR struct {
	SessionId                   datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity             `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity            `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity             `avp:"Destination-Realm"`
	ScAddress                   datatype.OctetString                  `avp:"SC-Address"`
	OfrFlags                    *datatype.Unsigned32                  `avp:"OFR-Flags,omitempty"`
	SupportedFeatures           []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	UserIdentifier              basetype.UserIdentifier               `avp:"User-Identifier"`
	SmRpUi                      datatype.OctetString                  `avp:"SM-RP-UI"`
	SmsmiCorrelationId          *basetype.SMSMICorrelationID          `avp:"SMSMI-Correlation-ID,omitempty"`
	SmDeliveryOutcome           *basetype.SMDeliveryOutcome           `avp:"SM-Delivery-Outcome,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
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

func (r *OFR) String() string {
	result := "OFR { "
	if r != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s, DestinationRealm: %s, ScAddress: %s, UserIdentifier: %v, SmRpUi: %s",
			r.SessionId, r.AuthSessionState, r.OriginHost, r.OriginRealm, r.DestinationRealm, r.ScAddress, r.UserIdentifier, r.SmRpUi)

		if r.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", r.Drmp.String())
		}

		if r.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", r.VendorSpecificApplicationId.String())
		}

		if r.DestinationHost != nil {
			result += fmt.Sprintf(", DestinationHost: %s", r.DestinationHost.String())
		}

		if r.OfrFlags != nil {
			result += fmt.Sprintf(", OfrFlags: %v", r.OfrFlags.String())
		}

		if len(r.SupportedFeatures) > 0 {
			result += ", SupportedFeatures: ["
			for i, feature := range r.SupportedFeatures {
				if i > 0 {
					result += ", "
				}
				result += feature.String()
			}
			result += "]"
		}

		if r.SmsmiCorrelationId != nil {
			result += fmt.Sprintf(", SmsmiCorrelationId: %v", r.SmsmiCorrelationId.String())
		}

		if r.SmDeliveryOutcome != nil {
			result += fmt.Sprintf(", SmDeliveryOutcome: %v", r.SmDeliveryOutcome.String())
		}

		if len(r.ProxyInfo) > 0 {
			result += ", ProxyInfo: ["
			for i, info := range r.ProxyInfo {
				if i > 0 {
					result += ", "
				}
				result += info.String()
			}
			result += "]"
		}

		if len(r.RouteRecord) > 0 {
			result += ", RouteRecord: ["
			for i, record := range r.RouteRecord {
				if i > 0 {
					result += ", "
				}
				result += record.String()
			}
			result += "]"
		}
	} else {
		result += "nil"
	}
	result += " }"
	return result
}
