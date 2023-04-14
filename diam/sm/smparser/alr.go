package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type ALR struct {
	SessionId                   datatype.UTF8String                      `avp:"Session-Id"`
	Drmp                        *datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId *basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                      `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity                `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity                `avp:"Origin-Realm"`
	DestinationHost             *datatype.DiameterIdentity               `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity                `avp:"Destination-Realm"`
	Msisdn                      *datatype.OctetString                    `avp:"MSISDN,omitempty"`
	UserName                    *datatype.UTF8String                     `avp:"User-Name,omitempty"`
	SmsmiCorrelationId          *basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
	ScAddress                   *datatype.OctetString                    `avp:"SC-Address,omitempty"`
	SmRpMti                     *datatype.Enumerated                     `avp:"SM-RP-MTI,omitempty"`
	SmRpSmea                    *datatype.OctetString                    `avp:"SM-RP-SMEA,omitempty"`
	SrrFlags                    *datatype.Unsigned32                     `avp:"SRR-Flags,omitempty"`
	SmDeliveryNotIntended       *datatype.Enumerated                     `avp:"SM-Delivery-Not-Intended,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                    `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity              `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (alr *ALR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(alr); err != nil {
		return err
	}
	if err := alr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (alr *ALR) sanityCheck() error {
	if len(alr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(alr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(alr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	return nil
}
