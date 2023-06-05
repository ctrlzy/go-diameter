package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// TFR refers to Mt-Forward-Short-Message-Request.
// See 3GPP TS 29.338 Clause 6.3.2.5 for details
type TFR struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity                `avp:"Destination-Host"`
	DestinationRealm            datatype.DiameterIdentity                `avp:"Destination-Realm"`
	UserName                    datatype.UTF8String                      `avp:"User-Name"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	SmsmiCorrelationId          *basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	ScAddress                   datatype.OctetString                     `avp:"SC-Address"`
	SmRpUi                      datatype.OctetString                     `avp:"SM-RP-UI"`
	MmeNumberForMtSms           *datatype.OctetString                    `avp:"MME-Number-for-MT-SMS,omitempty"`
	SgsnNumber                  *datatype.OctetString                    `avp:"SGSN-Number,omitempty"`
	TfrFlags                    *datatype.Unsigned32                     `avp:"TFR-Flags,omitempty"`
	SmDeliveryTimer             *datatype.Unsigned32                     `avp:"SM-Delivery-Timer,omitempty"`
	SmDeliveryStartTime         *datatype.Time                           `avp:"SM-Delivery-Start-Time,omitempty"`
	MaximumRetransmissionTime   *datatype.Time                           `avp:"Maximum-Retransmission-Time,omitempty"`
	SmsGmscAddress              *datatype.OctetString                    `avp:"SMS-GMSC-Address,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (tfr *TFR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(tfr); err != nil {
		return err
	}
	if err := tfr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (tfr *TFR) sanityCheck() error {
	if len(tfr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(tfr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(tfr.DestinationHost) == 0 {
		return ErrMissingDestHost
	}
	if len(tfr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if len(tfr.UserName) == 0 {
		return ErrMissingUserName
	}
	if len(tfr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	if len(tfr.SmRpUi) == 0 {
		return ErrMissingSmRpUi
	}
	return nil
}

func (t *TFR) String() string {
	result := "TFR { "
	if t != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s, DestinationHost: %s, DestinationRealm: %s, UserName: %s, ScAddress: %s, SmRpUi: %s",
			t.SessionId, t.AuthSessionState, t.OriginHost, t.OriginRealm, t.DestinationHost, t.DestinationRealm, t.UserName, t.ScAddress, t.SmRpUi)

		if t.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", t.Drmp.String())
		}

		if t.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", t.VendorSpecificApplicationId.String())
		}

		if t.SupportedFeatures != nil && len(t.SupportedFeatures) > 0 {
			result += ", SupportedFeatures: ["
			for i, feature := range t.SupportedFeatures {
				if i > 0 {
					result += ", "
				}
				result += feature.String()
			}
			result += "]"
		}

		if t.SmsmiCorrelationId != nil {
			result += fmt.Sprintf(", SmsmiCorrelationId: %v", t.SmsmiCorrelationId.String())
		}

		if t.MmeNumberForMtSms != nil {
			result += fmt.Sprintf(", MmeNumberForMtSms: %v", t.MmeNumberForMtSms.String())
		}

		if t.SgsnNumber != nil {
			result += fmt.Sprintf(", SgsnNumber: %v", t.SgsnNumber.String())
		}

		if t.TfrFlags != nil {
			result += fmt.Sprintf(", TfrFlags: %v", t.TfrFlags.String())
		}

		if t.SmDeliveryTimer != nil {
			result += fmt.Sprintf(", SmDeliveryTimer: %v", t.SmDeliveryTimer.String())
		}

		if t.SmDeliveryStartTime != nil {
			result += fmt.Sprintf(", SmDeliveryStartTime: %v", t.SmDeliveryStartTime.String())
		}

		if t.MaximumRetransmissionTime != nil {
			result += fmt.Sprintf(", MaximumRetransmissionTime: %v", t.MaximumRetransmissionTime.String())
		}

		if t.SmsGmscAddress != nil {
			result += fmt.Sprintf(", SmsGmscAddress: %v", t.SmsGmscAddress.String())
		}

		if t.ProxyInfo != nil && len(t.ProxyInfo) > 0 {
			result += ", ProxyInfo: ["
			for i, info := range t.ProxyInfo {
				if i > 0 {
					result += ", "
				}
				result += info.String()
			}
			result += "]"
		}

		if t.RouteRecord != nil && len(t.RouteRecord) > 0 {
			result += ", RouteRecord: ["
			for i, record := range t.RouteRecord {
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
