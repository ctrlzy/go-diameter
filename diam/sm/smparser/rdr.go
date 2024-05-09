package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// RDR refers to Report-SM-Delivery-Status-Request.
// See 3GPP TS 29.338 Clause 5.3.2.7 for details
type RDR struct {
	SessionId                   datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity             `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity            `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity             `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	UserIdentifier              basetype.UserIdentifier               `avp:"User-Identifier"`
	SmsmiCorrelationId          *basetype.SMSMICorrelationID          `avp:"SMSMI-Correlation-ID,omitempty"`
	ScAddress                   datatype.OctetString                  `avp:"SC-Address"`
	SmDeliveryOutcome           basetype.SMDeliveryOutcome            `avp:"SM-Delivery-Outcome"`
	RdrFlags                    *datatype.Unsigned32                  `avp:"RDR-Flags,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (rdr *RDR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(rdr); err != nil {
		return err
	}
	if err := rdr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (rdr *RDR) sanityCheck() error {
	if len(rdr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(rdr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(rdr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if rdr.UserIdentifier.Empty() {
		return ErrMissingUserIdentifier
	}
	if len(rdr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	if rdr.SmDeliveryOutcome.Empty() {
		return ErrMissingSmDeliveryOutcome
	}
	return nil
}

func (r *RDR) String() string {
	result := "RDR { "
	if r != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s",
			r.SessionId, r.AuthSessionState, r.OriginHost, r.OriginRealm)

		if r.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", r.Drmp.String())
		}

		if r.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", r.VendorSpecificApplicationId.String())
		}

		if r.DestinationHost != nil {
			result += fmt.Sprintf(", DestinationHost: %s", r.DestinationHost.String())
		}

		result += fmt.Sprintf(", DestinationRealm: %s, UserIdentifier: %v, ScAddress: %s, SmDeliveryOutcome: %v",
			r.DestinationRealm, r.UserIdentifier, r.ScAddress, r.SmDeliveryOutcome)

		if r.SupportedFeatures != nil && len(r.SupportedFeatures) > 0 {
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

		if r.RdrFlags != nil {
			result += fmt.Sprintf(", RdrFlags: %v", r.RdrFlags.String())
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
