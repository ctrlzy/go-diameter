package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
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

func (tfr *TFR) ToDiam() *diam.Message {
	// TODO: change dict.Default to base and SGD/GDD？
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, tfr.SessionId)
	if tfr.Drmp != nil {
		m.NewAVP(avp.DRMP, 0, 0, *tfr.Drmp)
	}
	if tfr.VendorSpecificApplicationId != nil {
		m.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, tfr.VendorSpecificApplicationId.ToDiam())
	}
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, tfr.AuthSessionState)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, tfr.OriginHost)
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, tfr.OriginRealm)
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, tfr.DestinationHost)
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, tfr.DestinationRealm)
	m.NewAVP(avp.UserName, avp.Mbit, 0, tfr.UserName)
	if len(tfr.SupportedFeatures) > 0 {
		for _, sf := range tfr.SupportedFeatures {
			m.NewAVP(avp.SupportedFeatures, avp.Vbit, 10415, sf.ToDiam())
		}
	}
	if tfr.SmsmiCorrelationId != nil {
		m.NewAVP(avp.SMSMICorrelationID, avp.Vbit, 10415, tfr.SmsmiCorrelationId.ToDiam())
	}
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, tfr.ScAddress)
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, tfr.SmRpUi)
	if tfr.MmeNumberForMtSms != nil {
		m.NewAVP(avp.MMENumberforMTSMS, avp.Vbit, 10415, *tfr.MmeNumberForMtSms)
	}
	if tfr.SgsnNumber != nil {
		m.NewAVP(avp.SGSNNumber, avp.Vbit, 10415, *tfr.SgsnNumber)
	}
	if tfr.TfrFlags != nil {
		m.NewAVP(avp.TFRFlags, avp.Mbit|avp.Vbit, 10415, *tfr.TfrFlags)
	}
	if tfr.SmDeliveryTimer != nil {
		m.NewAVP(avp.SMDeliveryTimer, avp.Mbit|avp.Vbit, 10415, *tfr.SmDeliveryTimer)
	}
	if tfr.SmDeliveryStartTime != nil {
		m.NewAVP(avp.SMDeliveryStartTime, avp.Mbit|avp.Vbit, 10415, *tfr.SmDeliveryStartTime)
	}
	if tfr.MaximumRetransmissionTime != nil {
		m.NewAVP(avp.MaximumRetransmissionTime, avp.Vbit, 10415, *tfr.MaximumRetransmissionTime)
	}
	if tfr.SmsGmscAddress != nil {
		m.NewAVP(avp.SMSGMSCAddress, avp.Vbit, 10415, *tfr.SmsGmscAddress)
	}
	if len(tfr.ProxyInfo) > 0 {
		for _, pi := range tfr.ProxyInfo {
			m.NewAVP(avp.ProxyInfo, avp.Mbit, 0, pi.ToDiam())
		}
	}
	if len(tfr.RouteRecord) > 0 {
		for _, rr := range tfr.RouteRecord {
			m.NewAVP(avp.RouteRecord, avp.Mbit, 0, rr)
		}
	}
	return m
}
