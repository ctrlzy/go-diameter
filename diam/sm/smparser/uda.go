package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// refer to 3GPP TS 29.329 6.1.2 User-Data-Answer
type UDA struct {
	SessionID                   datatype.UTF8String                  `avp:"Session-Id"`
	DRMP                        *datatype.Enumerated                 `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id"`
	ResultCode                  *datatype.Unsigned32                 `avp:"Result-Code,omitempty"`
	ExperimentalResult          *basetype.ExperimentalResult         `avp:"Experimental-Result,omitempty"`
	AuthSessionState            datatype.Enumerated                  `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity            `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity            `avp:"Origin-Realm"`
	SupportedFeatures           []basetype.SupportedFeatures         `avp:"Supported-Features,omitempty"`
	WildcardedPublicIdentity    *datatype.UTF8String                 `avp:"Wildcarded-Public-Identity,omitempty"`
	WildcardedIMPU              *datatype.UTF8String                 `avp:"Wildcarded-IMPU,omitempty"`
	UserData                    *datatype.OctetString                `avp:"User-Data,omitempty"`
	OCSupportedFeatures         *basetype.OCSupportedFeatures        `avp:"OC-Supported-Features,omitempty"`
	OCOLR                       *basetype.OCOLR                      `avp:"OC-OLR,omitempty"`
	Load                        *basetype.Load                       `avp:"Load,omitempty"`
	FailedAVP                   basetype.FailedAVP                   `avp:"Failed-AVP,omitempty"`
	ProxyInfo                   []basetype.ProxyInfo                 `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity          `avp:"Route-Record,omitempty"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (uda *UDA) Parse(m *diam.Message) error {
	err := m.Unmarshal(uda)
	if err != nil {
		return nil
	}
	if err = uda.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (uda *UDA) sanityCheck() error {
	if len(uda.SessionID) == 0 {
		return ErrMissingSessionID
	}
	if uda.VendorSpecificApplicationId.Empty() {
		return ErrMissingVendorSpecificAppId
	}
	if len(uda.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(uda.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}
