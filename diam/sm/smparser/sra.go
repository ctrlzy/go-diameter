package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type SRA struct {
	SessionId                         datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                              *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId       *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                        *datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult                *basetype.Experimental_Result            `avp:"Experimental-Result,omitempty"`
	AuthSessionState                  datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                        datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                       datatype.DiameterIdentity                `avp:"Origin-Realm"`
	UserName                          *datatype.UTF8String                     `avp:"User-Name,omitempty"`
	SupportedFeatures                 []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	ServingNode                       *basetype.Serving_Node                   `avp:"Serving-Node,omitempty"`
	AdditionalServingNode             *basetype.Additional_Serving_Node        `avp:"Additional-Serving-Node,omitempty"`
	Smsf3gppAddress                   *basetype.SMSF_3GPP_Address              `avp:"SMSF-3GPP-Address,omitempty"`
	SmsfNon3gppAddress                *basetype.SMSF_Non_3GPP_Address          `avp:"SMSF-Non-3GPP-Address,omitempty"`
	Lmsi                              *datatype.OctetString                    `avp:"LMSI,omitempty"`
	UserIdentifier                    *basetype.User_Identifier                `avp:"User-Identifier,omitempty"`
	MwdStatus                         *datatype.Unsigned32                     `avp:"MWD-Status,omitempty"`
	MmeAbsentUserDiagnosticSm         *datatype.Unsigned32                     `avp:"MME-Absent-User-Diagnostic-SM,omitempty"`
	MscAbsentUserDiagnosticSm         *datatype.Unsigned32                     `avp:"MSC-Absent-User-Diagnostic-SM,omitempty"`
	SgsnAbsentUserDiagnosticSm        *datatype.Unsigned32                     `avp:"SGSN-Absent-User-Diagnostic-SM,omitempty"`
	Smsf3gppAbsentUserDiagnosticSm    *datatype.Unsigned32                     `avp:"SMSF-3GPP-Absent-User-Diagnostic-SM,omitempty"`
	SmsfNon3gppAbsentUserDiagnosticSm *datatype.Unsigned32                     `avp:"SMSF-Non-3GPP-Absent-User-Diagnostic-SM,omitempty"`
	FailedAvp                         *basetype.Failed_AVP                     `avp:"Failed-AVP,omitempty"`
	ProxyInfo                         []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                       []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
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
