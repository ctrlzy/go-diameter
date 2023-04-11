package basetype

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

type Supported_Features struct {
	VendorId      datatype.Unsigned32 `avp:"Vendor-Id"`
	FeatureListId datatype.Unsigned32 `avp:"Feature-List-ID"`
	FeatureList   datatype.Unsigned32 `avp:"Feature-List"`
}

type SMSMI_Correlation_ID struct {
	HssId             datatype.OctetString `avp:"HSS-ID,omitempty"`
	OriginatingSipUri datatype.UTF8String  `avp:"Originating-SIP-URI,omitempty"`
	DestinationSipUri datatype.UTF8String  `avp:"Destination-SIP-URI,omitempty"`
}

type User_Identifier struct {
	UserName           datatype.UTF8String  `avp:"User-Name,omitempty"`
	Msisdn             datatype.OctetString `avp:"MSISDN,omitempty"`
	ExternalIdentifier datatype.UTF8String  `avp:"External-Identifier,omitempty"`
	Lmsi               datatype.OctetString `avp:"LMSI,omitempty"`
}

func (ui *User_Identifier) Empty() bool {
	return (len(ui.UserName) + len(ui.Msisdn) + len(ui.ExternalIdentifier) + len(ui.Lmsi)) == 0
}

type SM_Delivery_Failure_Cause struct {
	SmEnumeratedDeliveryFailureCause datatype.Enumerated  `avp:"SM-Enumerated-Delivery-Failure-Cause"`
	SmDiagnosticInfo                 datatype.OctetString `avp:"SM-Diagnostic-Info,omitempty"`
}

type SM_Delivery_Outcome struct {
	MmeSmDeliveryOutcome    Delivery_Outcome `avp:"MME-SM-Delivery-Outcome,omitempty"`
	MscSmDeliveryOutcome    Delivery_Outcome `avp:"MSC-SM-Delivery-Outcome,omitempty"`
	SgsnSmDeliveryOutcome   Delivery_Outcome `avp:"SGSN-SM-Delivery-Outcome,omitempty"`
	IpsmgwSmDeliveryOutcome Delivery_Outcome `avp:"IP-SM-GW-SM-Delivery-Outcome,omitempty"`
}

type Delivery_Outcome struct {
	SmDeliveryCause        datatype.Enumerated `avp:"SM-Delivery-Cause,omitempty"`
	AbsentUserDiagnosticSm datatype.Unsigned32 `avp:"Absent-User-Diagnostic-SM,omitempty"`
}
