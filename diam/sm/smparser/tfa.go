package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// TFA refers to Mt-Forward-Short-Message-Answer.
// See 3GPP TS 29.338 Clause 6.3.2.6 for details
type TFA struct {
	SessionId                   datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  *datatype.Unsigned32                  `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult          `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity             `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	AbsentUserDiagnosticSm      *datatype.Unsigned32                  `avp:"Absent-User-Diagnostic-SM,omitempty"`
	SmDeliveryFailureCause      *basetype.SMDeliveryFailureCause      `avp:"SM-Delivery-Failure-Cause,omitempty"`
	SmRpUi                      *datatype.OctetString                 `avp:"SM-RP-UI,omitempty"`
	RequestedRetransmissionTime *datatype.Time                        `avp:"Requested-Retransmission-Time,omitempty"`
	UserIdentifier              *basetype.UserIdentifier              `avp:"User-Identifier,omitempty"`
	FailedAvp                   basetype.FailedAVP                    `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (tfa *TFA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(tfa); err != nil {
		return err
	}
	if err := tfa.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (tfa *TFA) sanityCheck() error {
	if len(tfa.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(tfa.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (t *TFA) String() string {
	result := "TFA { "
	if t != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s",
			t.SessionId, t.AuthSessionState, t.OriginHost, t.OriginRealm)

		if t.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", t.Drmp.String())
		}

		if t.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", t.VendorSpecificApplicationId.String())
		}

		if t.ResultCode != nil {
			result += fmt.Sprintf(", ResultCode: %v", t.ResultCode.String())
		}

		if t.ExperimentalResult != nil {
			result += fmt.Sprintf(", ExperimentalResult: %v", t.ExperimentalResult.String())
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

		if t.AbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", AbsentUserDiagnosticSm: %v", t.AbsentUserDiagnosticSm.String())
		}

		if t.SmDeliveryFailureCause != nil {
			result += fmt.Sprintf(", SmDeliveryFailureCause: %v", t.SmDeliveryFailureCause.String())
		}

		if t.SmRpUi != nil {
			result += fmt.Sprintf(", SmRpUi: %v", t.SmRpUi.String())
		}

		if t.RequestedRetransmissionTime != nil {
			result += fmt.Sprintf(", RequestedRetransmissionTime: %v", t.RequestedRetransmissionTime.String())
		}

		if t.UserIdentifier != nil {
			result += fmt.Sprintf(", UserIdentifier: %v", t.UserIdentifier.String())
		}

		if t.FailedAvp != nil && len(t.FailedAvp) > 0 {
			result += ", FailedAvp: ["
			for i, avp := range t.FailedAvp {
				if i > 0 {
					result += ", "
				}
				result += avp.String()
			}
			result += "]"
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
