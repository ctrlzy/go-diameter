package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFR refers to Mo-Forward-Short-Message-Answer.
// See 3GPP TS 29.338 Clause 6.2.2 for details
type OFA struct {
	SessionId                   datatype.UTF8String                     `avp:"Session-Id"`
	Drmp                        datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult          basetype.Experimental_Result            `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	SmDeliveryFailureCause      basetype.SM_Delivery_Failure_Cause      `avp:"SM-Delivery-Failure-Cause,omitempty"`
	SmRpUi                      datatype.OctetString                    `avp:"SM-RP-UI,omitempty"`
	ExternalIdentifier          datatype.UTF8String                     `avp:"External-Identifier,omitempty"`
	FailedAvp                   basetype.Failed_AVP                     `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
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
