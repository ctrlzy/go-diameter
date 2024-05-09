package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// SRA refers to Send-Routing-info-for-SM-Answer
// See 3GPP TS 29.338 Clause 5.3.2.4 for details
type SRA struct {
	SessionId                         datatype.UTF8String                   `avp:"Session-Id"`
	Drmp                              *datatype.Enumerated                  `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId       *basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                        *datatype.Unsigned32                  `avp:"Result-Code,omitempty"`
	ExperimentalResult                *basetype.ExperimentalResult          `avp:"Experimental-Result,omitempty"`
	AuthSessionState                  datatype.Enumerated                   `avp:"Auth-Session-State"`
	OriginHost                        datatype.DiameterIdentity             `avp:"Origin-Host"`
	OriginRealm                       datatype.DiameterIdentity             `avp:"Origin-Realm"`
	UserName                          *datatype.UTF8String                  `avp:"User-Name,omitempty"`
	SupportedFeatures                 []basetype.SupportedFeatures          `avp:"Supported-Features,omitempty"`
	ServingNode                       *basetype.ServingNode                 `avp:"Serving-Node,omitempty"`
	AdditionalServingNode             *basetype.AdditionalServingNode       `avp:"Additional-Serving-Node,omitempty"`
	Smsf3gppAddress                   *basetype.SMSF3GPPAddress             `avp:"SMSF-3GPP-Address,omitempty"`
	SmsfNon3gppAddress                *basetype.SMSFNon3GPPAddress          `avp:"SMSF-Non-3GPP-Address,omitempty"`
	Lmsi                              *datatype.OctetString                 `avp:"LMSI,omitempty"`
	UserIdentifier                    *basetype.UserIdentifier              `avp:"User-Identifier,omitempty"`
	MwdStatus                         *datatype.Unsigned32                  `avp:"MWD-Status,omitempty"`
	MmeAbsentUserDiagnosticSm         *datatype.Unsigned32                  `avp:"MME-Absent-User-Diagnostic-SM,omitempty"`
	MscAbsentUserDiagnosticSm         *datatype.Unsigned32                  `avp:"MSC-Absent-User-Diagnostic-SM,omitempty"`
	SgsnAbsentUserDiagnosticSm        *datatype.Unsigned32                  `avp:"SGSN-Absent-User-Diagnostic-SM,omitempty"`
	Smsf3gppAbsentUserDiagnosticSm    *datatype.Unsigned32                  `avp:"SMSF-3GPP-Absent-User-Diagnostic-SM,omitempty"`
	SmsfNon3gppAbsentUserDiagnosticSm *datatype.Unsigned32                  `avp:"SMSF-Non-3GPP-Absent-User-Diagnostic-SM,omitempty"`
	FailedAvp                         basetype.FailedAVP                    `avp:"Failed-AVP,omitempty"`
	ProxyInfo                         []basetype.ProxyInfo                  `avp:"Proxy-Info,omitempty"`
	RouteRecord                       []datatype.DiameterIdentity           `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (sra *SRA) Parse(m *diam.Message) error {
	if err := m.Unmarshal(sra); err != nil {
		return err
	}
	if err := sra.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (sra *SRA) sanityCheck() error {
	if len(sra.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(sra.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (s *SRA) String() string {
	result := "SRA { "
	if s != nil {
		result += fmt.Sprintf("SessionId: %s, AuthSessionState: %v, OriginHost: %s, OriginRealm: %s",
			s.SessionId, s.AuthSessionState, s.OriginHost, s.OriginRealm)

		if s.Drmp != nil {
			result += fmt.Sprintf(", Drmp: %v", s.Drmp.String())
		}

		if s.VendorSpecificApplicationId != nil {
			result += fmt.Sprintf(", VendorSpecificApplicationId: %v", s.VendorSpecificApplicationId.String())
		}

		if s.ResultCode != nil {
			result += fmt.Sprintf(", ResultCode: %v", s.ResultCode.String())
		}

		if s.ExperimentalResult != nil {
			result += fmt.Sprintf(", ExperimentalResult: %v", s.ExperimentalResult.String())
		}

		if s.UserName != nil {
			result += fmt.Sprintf(", UserName: %s", s.UserName.String())
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

		if s.ServingNode != nil {
			result += fmt.Sprintf(", ServingNode: %v", s.ServingNode.String())
		}

		if s.AdditionalServingNode != nil {
			result += fmt.Sprintf(", AdditionalServingNode: %v", s.AdditionalServingNode.String())
		}

		if s.Smsf3gppAddress != nil {
			result += fmt.Sprintf(", Smsf3gppAddress: %v", s.Smsf3gppAddress.String())
		}

		if s.SmsfNon3gppAddress != nil {
			result += fmt.Sprintf(", SmsfNon3gppAddress: %v", s.SmsfNon3gppAddress.String())
		}

		if s.Lmsi != nil {
			result += fmt.Sprintf(", Lmsi: %v", s.Lmsi.String())
		}

		if s.UserIdentifier != nil {
			result += fmt.Sprintf(", UserIdentifier: %v", s.UserIdentifier.String())
		}

		if s.MwdStatus != nil {
			result += fmt.Sprintf(", MwdStatus: %v", s.MwdStatus.String())
		}

		if s.MmeAbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", MmeAbsentUserDiagnosticSm: %v", s.MmeAbsentUserDiagnosticSm.String())
		}

		if s.MscAbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", MscAbsentUserDiagnosticSm: %v", s.MscAbsentUserDiagnosticSm.String())
		}

		if s.SgsnAbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", SgsnAbsentUserDiagnosticSm: %v", s.SgsnAbsentUserDiagnosticSm.String())
		}

		if s.Smsf3gppAbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", Smsf3gppAbsentUserDiagnosticSm: %v", s.Smsf3gppAbsentUserDiagnosticSm.String())
		}

		if s.SmsfNon3gppAbsentUserDiagnosticSm != nil {
			result += fmt.Sprintf(", SmsfNon3gppAbsentUserDiagnosticSm: %v", s.SmsfNon3gppAbsentUserDiagnosticSm.String())
		}

		if s.FailedAvp != nil {
			result += ", "
			result += s.FailedAvp.String()
		}

		if len(s.ProxyInfo) > 0 {
			result += ", ProxyInfo: ["
			for i, info := range s.ProxyInfo {
				if i > 0 {
					result += ", "
				}
				result += info.String()
			}
			result += "]"
		}

		if len(s.RouteRecord) > 0 {
			result += ", RouteRecord: ["
			for i, record := range s.RouteRecord {
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
