package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// RDA refers to Report-SM-Delivery-Status-Answer.
// See 3GPP TS 29.338 Clause 5.3.2.8 for details
type RDA struct {
	SessionId                   datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  *datatype.Unsigned32                  `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult          `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity             `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	UserIdentifier              *basetype.UserIdentifier              `avp:"User-Identifier,omitempty"`
	FailedAvp                   basetype.FailedAVP                    `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (rda *RDA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(rda); err != nil {
		return err
	}
	if err := rda.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (rda *RDA) sanityCheck() error {
	if len(rda.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(rda.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (r *RDA) String() string {
	result := "RDA { "
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

		if r.UserIdentifier != nil {
			result += fmt.Sprintf(", UserIdentifier: %v", r.UserIdentifier.String())
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
