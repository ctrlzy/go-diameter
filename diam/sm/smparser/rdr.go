package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type RDR struct {
	SessionId                   datatype.UTF8String                     `avp:"Session-Id"`
	Drmp                        datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity               `avp:"Destination-Host,omitempty"`
	DestinationRealm            datatype.DiameterIdentity               `avp:"Destination-Realm"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	UserIdentifier              basetype.User_Identifier                `avp:"User-Identifier"`
	SmsmiCorrelationId          basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	ScAddress                   datatype.OctetString                    `avp:"SC-Address"`
	SmDeliveryOutcome           basetype.SM_Delivery_Outcome            `avp:"SM-Delivery-Outcome"`
	RdrFlags                    datatype.Unsigned32                     `avp:"RDR-Flags,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (rdr *RDR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(rdr); err != nil {
		return err
	}
	if err := rdr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (rdr *RDR) sanityCheck() error {
	if len(rdr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(rdr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(rdr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if rdr.UserIdentifier.Empty() {
		return ErrMissingUserIdentifier
	}
	if len(rdr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	/*  TODO: delivery-outcome is hard to check whether is empty or not
	the SM-Delivery-Cause is Enum which has 0 value.
	if rdr.SmDeliveryOutcome.Empty() {
		return ErrMissingSmDe
	}
	*/
	return nil
}
