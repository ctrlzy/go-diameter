package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFA refers to Mo-Forward-Short-Message-Answer.
// See 3GPP TS 29.338 Clause 6.3.2.4 for details
type OFA struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  *datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.Experimental_Result            `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	SmDeliveryFailureCause      *basetype.SM_Delivery_Failure_Cause      `avp:"SM-Delivery-Failure-Cause,omitempty"`
	SmRpUi                      *datatype.OctetString                    `avp:"SM-RP-UI,omitempty"`
	ExternalIdentifier          *datatype.UTF8String                     `avp:"External-Identifier,omitempty"`
	FailedAvp                   []*diam.AVP                              `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (ofa *OFA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(ofa); err != nil {
		return err
	}
	if err := ofa.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (ofa *OFA) sanityCheck() error {
	if len(ofa.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(ofa.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (r *OFA) String() string {
	result := "OFA { "
	if r != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s",
			r.SessionId, r.AuthSessionState, r.OriginHost, r.OriginRealm)

		if r.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", r.Drmp.String())
		}

		if r.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", r.VendorSpecificApplicationId.String())
		}

		if r.ResultCode != nil {
			result += fmt.Sprintf(", ResultCode: %v", r.ResultCode.String())
		}

		if r.ExperimentalResult != nil {
			result += fmt.Sprintf(", ExperimentalResult: %v", r.ExperimentalResult.String())
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

		if r.SmDeliveryFailureCause != nil {
			result += fmt.Sprintf(", SmDeliveryFailureCause: %v", r.SmDeliveryFailureCause.String())
		}

		if r.SmRpUi != nil {
			result += fmt.Sprintf(", SmRpUi: %v", r.SmRpUi.String())
		}

		if r.ExternalIdentifier != nil {
			result += fmt.Sprintf(", ExternalIdentifier: %v", r.ExternalIdentifier.String())
		}

		if len(r.FailedAvp) > 0 {
			result += ", FailedAvp: ["
			for i, avp := range r.FailedAvp {
				if i > 0 {
					result += ", "
				}
				result += avp.String()
			}
			result += "]"
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
