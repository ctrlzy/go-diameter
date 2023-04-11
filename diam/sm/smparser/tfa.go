package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFR refers to Mt-Forward-Short-Message-Answer.
// See 3GPP TS 29.338 Clause 6.3.2.6 for details
type TFA struct {
	SessionId                   datatype.UTF8String                     `avp:"Session-Id"`
	Drmp                        datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult          basetype.Experimental_Result            `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	AbsentUserDiagnosticSm      datatype.Unsigned32                     `avp:"Absent-User-Diagnostic-SM,omitempty"`
	SmDeliveryFailureCause      basetype.SM_Delivery_Failure_Cause      `avp:"SM-Delivery-Failure-Cause,omitempty"`
	SmRpUi                      datatype.OctetString                    `avp:"SM-RP-UI,omitempty"`
	RequestedRetransmissionTime datatype.Time                           `avp:"Requested-Retransmission-Time,omitempty"`
	UserIdentifier              basetype.User_Identifier                `avp:"User-Identifier,omitempty"`
	FailedAvp                   basetype.Failed_AVP                     `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
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
