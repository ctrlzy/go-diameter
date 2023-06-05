package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// ALR is a Alert-Service-Centre-Request message.
// See 3GPP TS 29.338 section 5.3.2.5 for details.
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

func (r *ALR) String() string {
	result := "ALR { "
	result += fmt.Sprintf("SessionId: %s, ", r.SessionId)
	if r.Drmp != nil {
		result += fmt.Sprintf("Drmp: %s, ", r.Drmp.String())
	}
	if r.VendorSpecificApplicationId != nil {
		result += fmt.Sprintf("VendorSpecificApplicationId: %s, ", r.VendorSpecificApplicationId.String())
	}
	result += fmt.Sprintf("AuthSessionState: %s, ", r.AuthSessionState)
	result += fmt.Sprintf("OriginHost: %s, ", r.OriginHost)
	result += fmt.Sprintf("OriginRealm: %s, ", r.OriginRealm)
	if r.DestinationHost != nil {
		result += fmt.Sprintf("DestinationHost: %s, ", r.DestinationHost.String())
	}
	result += fmt.Sprintf("DestinationRealm: %s, ", r.DestinationRealm)
	result += fmt.Sprintf("ScAddress: %s, ", r.ScAddress)
	result += fmt.Sprintf("UserIdentifier: %s, ", r.UserIdentifier.String())
	if r.SmsmiCorrelationId != nil {
		result += fmt.Sprintf("SmsmiCorrelationId: %s, ", r.SmsmiCorrelationId.String())
	}
	if r.MaximumUeAvailabilityTime != nil {
		result += fmt.Sprintf("MaximumUeAvailabilityTime: %s, ", r.MaximumUeAvailabilityTime.String())
	}
	if r.SmsGmscAlertEvent != nil {
		result += fmt.Sprintf("SmsGmscAlertEvent: %d, ", *r.SmsGmscAlertEvent)
	}
	if r.ServingNode != nil {
		result += fmt.Sprintf("ServingNode: %s, ", r.ServingNode.String())
	}
	if len(r.SupportedFeatures) > 0 {
		result += "SupportedFeatures: ["
		for i, feature := range r.SupportedFeatures {
			result += feature.String()
			if i < len(r.SupportedFeatures)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	if len(r.ProxyInfo) > 0 {
		result += "ProxyInfo: ["
		for i, info := range r.ProxyInfo {
			result += info.String()
			if i < len(r.ProxyInfo)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	if len(r.RouteRecord) > 0 {
		result += "RouteRecord: ["
		for i, record := range r.RouteRecord {
			result += record.String()
			if i < len(r.RouteRecord)-1 {
				result += ", "
			}
		}
		result += "], "
	}
	result += "}"
	return result
}
