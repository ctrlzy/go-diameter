package smparser

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// OFR refers to Mt-Forward-Short-Message-Request.
// See 3GPP TS 29.338 Clause 6.3.2.5 for details
type TFR struct {
	SessionId                   datatype.UTF8String                     `avp:"Session-Id"`
	Drmp                        datatype.Enumerated                     `avp:"DRMP,omitempty"`
	VendorSpecificApplicationId basetype.Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
	AuthSessionState            datatype.Enumerated                     `avp:"Auth-Session-State"`
	OriginHost                  datatype.DiameterIdentity               `avp:"Origin-Host"`
	OriginRealm                 datatype.DiameterIdentity               `avp:"Origin-Realm"`
	DestinationHost             datatype.DiameterIdentity               `avp:"Destination-Host"`
	DestinationRealm            datatype.DiameterIdentity               `avp:"Destination-Realm"`
	UserName                    datatype.UTF8String                     `avp:"User-Name"`
	SupportedFeatures           []basetype.Supported_Features           `avp:"Supported-Features,omitempty"`
	SmsmiCorrelationId          basetype.SMSMI_Correlation_ID           `avp:"SMSMI-Correlation-ID,omitempty"`
	ScAddress                   datatype.OctetString                    `avp:"SC-Address"`
	SmRpUi                      datatype.OctetString                    `avp:"SM-RP-UI"`
	MmeNumberForMtSms           datatype.OctetString                    `avp:"MME-Number-for-MT-SMS,omitempty"`
	SgsnNumber                  datatype.OctetString                    `avp:"SGSN-Number,omitempty"`
	TfrFlags                    datatype.Unsigned32                     `avp:"TFR-Flags,omitempty"`
	SmDeliveryTimer             datatype.Unsigned32                     `avp:"SM-Delivery-Timer,omitempty"`
	SmDeliveryStartTime         datatype.Time                           `avp:"SM-Delivery-Start-Time,omitempty"`
	MaximumRetransmissionTime   datatype.Time                           `avp:"Maximum-Retransmission-Time,omitempty"`
	SmsGmscAddress              datatype.OctetString                    `avp:"SMS-GMSC-Address,omitempty"`
	ProxyInfo                   []basetype.Proxy_Info                   `avp:"Proxy-Info,omitempty"`
	RouteRecord                 []datatype.DiameterIdentity             `avp:"Route-Record,omitempty"`
}

// Parse parses the given message.
func (tfr *TFR) Parse(m *diam.Message) error {
	if err := m.Unmarshal(tfr); err != nil {
		return err
	}
	if err := tfr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

func (tfr *TFR) sanityCheck() error {
	if len(tfr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(tfr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	if len(tfr.DestinationHost) == 0 {
		return ErrMissingDestHost
	}
	if len(tfr.DestinationRealm) == 0 {
		return ErrMissingDestRealm
	}
	if len(tfr.UserName) == 0 {
		return ErrMissingUserName
	}
	if len(tfr.ScAddress) == 0 {
		return ErrMissingScAddress
	}
	if len(tfr.SmRpUi) == 0 {
		return ErrMissingSmRpUi
	}
	return nil
}
