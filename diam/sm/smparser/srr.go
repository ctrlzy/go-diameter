package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// SRR refers to Send-Routing-info-for-SM-Request
// See 3GPP TS 29.338 Clause 5.3.2.3 for details
type SRR struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity               `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity                `avp:"Destination-Realm"`
	Msisdn                      *datatype.OctetString                    `avp:"MSISDN,omitempty"`
	UserName                    *datatype.UTF8String                     `avp:"User-Name,omitempty"`
	SmsmiCorrelationId          *basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	ScAddress                   *datatype.OctetString                    `avp:"SC-Address,omitempty"`
	SmRpMti                     *datatype.Enumerated                     `avp:"SM-RP-MTI,omitempty"`
	SmRpSmea                    *datatype.OctetString                    `avp:"SM-RP-SMEA,omitempty"`
	SrrFlags                    *datatype.Unsigned32                     `avp:"SRR-Flags,omitempty"`
	SmDeliveryNotIntended       *datatype.Enumerated                     `avp:"SM-Delivery-Not-Intended,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (srr *SRR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(srr); err != nil {
		return err
	}
	if err := srr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (srr *SRR) sanityCheck() error {
	if len(srr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(srr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(srr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	return nil
}

func (s *SRR) String() string {
	result := "SRR { "
	if s != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s",
			s.SessionId, s.AuthSessionState, s.OriginHost, s.OriginRealm)

		if s.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", s.Drmp.String())
		}

		if s.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", s.VendorSpecificApplicationId.String())
		}

		if s.DestinationHost != nil {
			result += fmt.Sprintf(", DestinationHost: %s", s.DestinationHost.String())
		}

		result += fmt.Sprintf(", DestinationRealm: %s", s.DestinationRealm)

		if s.Msisdn != nil {
			result += fmt.Sprintf(", Msisdn: %v", s.Msisdn.String())
		}

		if s.UserName != nil {
			result += fmt.Sprintf(", UserName: %s", s.UserName.String())
		}

		if s.SmsmiCorrelationId != nil {
			result += fmt.Sprintf(", SmsmiCorrelationId: %v", s.SmsmiCorrelationId.String())
		}

		if s.SupportedFeatures != nil && len(s.SupportedFeatures) > 0 {
			result += ", SupportedFeatures: ["
			for i, feature := range s.SupportedFeatures {
				if i > 0 {
					result += ", "
				}
				result += feature.String()
			}
			result += "]"
		}

		if s.ScAddress != nil {
			result += fmt.Sprintf(", ScAddress: %v", s.ScAddress.String())
		}

		if s.SmRpMti != nil {
			result += fmt.Sprintf(", SmRpMti: %v", s.SmRpMti.String())
		}

		if s.SmRpSmea != nil {
			result += fmt.Sprintf(", SmRpSmea: %v", s.SmRpSmea.String())
		}

		if s.SrrFlags != nil {
			result += fmt.Sprintf(", SrrFlags: %v", s.SrrFlags.String())
		}

		if s.SmDeliveryNotIntended != nil {
			result += fmt.Sprintf(", SmDeliveryNotIntended: %v", s.SmDeliveryNotIntended.String())
		}

		if s.ProxyInfo != nil && len(s.ProxyInfo) > 0 {
			result += ", ProxyInfo: ["
			for i, info := range s.ProxyInfo {
				if i > 0 {
					result += ", "
				}
				result += info.String()
			}
			result += "]"
		}

		if s.RouteRecord != nil && len(s.RouteRecord) > 0 {
			result += ", RouteRecord: ["
			for i, record := range s.RouteRecord {
				if i > 0 {
					result += ", "
				}
				result += record.String()
			}
			result += "]"
		}
	}
	result += " }"
	return result
}
