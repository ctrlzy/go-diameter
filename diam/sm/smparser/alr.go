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
	ScAddress                   datatype.OctetString                     `avp:"SC-Address"`
	UserIdentifier              basetype.User_Identifier                 `avp:"User-Identifier"`
	SmsmiCorrelationId          *basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	MaximumUeAvailabilityTime   *datatype.Time                           `avp:"Maximum-UE-Availability-Time,omitempty"`
	SmsGmscAlertEvent           *datatype.Unsigned32                     `avp:"SMS-GMSC-Alert-Event,omitempty"`
	ServingNode                 *basetype.Serving_Node                   `avp:"Serving-Node,omitempty"`
	SupportedFeatures           []basetype.Supported_Features            `avp:"Supported-Features,omitempty"`
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
	if len(alr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	if alr.UserIdentifier.Empty() {
		return ErrMissingUserIdentifier
	}
	return nil
}
