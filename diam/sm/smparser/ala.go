package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// ALA is a Alert-Service-Centre-Answer message.
// See 3GPP TS 29.338 section 5.3.2.6 for details.
type ALA struct {
	SessionId                   datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  *datatype.Unsigned32                  `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult          `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity             `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	FailedAvp                   basetype.FailedAVP                    `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (ala *ALA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(ala); err != nil {
		return err
	}
	if err := ala.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (ala *ALA) sanityCheck() error {
	if len(ala.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(ala.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (a *ALA) String() string {
	result := "ALA { "
	result += fmt.Sprintf("SessionId: %s, ", a.SessionId)
	if a.Drmp != nil {
		result += fmt.Sprintf("Drmp: %s, ", a.Drmp.String())
	}
	if a.VendorSpecificApplicationId != nil {
		result += fmt.Sprintf("VendorSpecificApplicationId: %s, ", a.VendorSpecificApplicationId.String())
	}
	if a.ResultCode != nil {
		result += fmt.Sprintf("ResultCode: %d, ", *a.ResultCode)
	}
	if a.ExperimentalResult != nil {
		result += fmt.Sprintf("ExperimentalResult: %s, ", a.ExperimentalResult.String())
	}
	result += fmt.Sprintf("AuthSessionState: %s, ", a.AuthSessionState)
	result += fmt.Sprintf("OriginHost: %s, ", a.OriginHost)
	result += fmt.Sprintf("OriginRealm: %s, ", a.OriginRealm)
	if len(a.SupportedFeatures) > 0 {
		result += "SupportedFeatures: ["
		for i, feature := range a.SupportedFeatures {
			result += feature.String()
			if i < len(a.SupportedFeatures)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	if len(a.FailedAvp) > 0 {
		result += "FailedAvp: ["
		for i, avp := range a.FailedAvp {
			result += avp.String()
			if i < len(a.FailedAvp)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	if len(a.ProxyInfo) > 0 {
		result += "ProxyInfo: ["
		for i, info := range a.ProxyInfo {
			result += info.String()
			if i < len(a.ProxyInfo)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	if len(a.RouteRecord) > 0 {
		result += "RouteRecord: ["
		for i, record := range a.RouteRecord {
			result += record.String()
			if i < len(a.RouteRecord)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	result += "}"
	return result
}
