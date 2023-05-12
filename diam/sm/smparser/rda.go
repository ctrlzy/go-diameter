package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type RDA struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	ResultCode                  *datatype.Unsigned32                     `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.Experimental_Result            `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	UserIdentifier              *basetype.User_Identifier                `avp:"User-Identifier,omitempty"`
	FailedAvp                   *basetype.Failed_AVP                     `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
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
