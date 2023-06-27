package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type Supported_Features struct {
	VendorId      datatype.Unsigned32 `avp:"Vendor-Id"`
	FeatureListId datatype.Unsigned32 `avp:"Feature-List-ID"`
	FeatureList   datatype.Unsigned32 `avp:"Feature-List"`
}

type SMSMI_Correlation_ID struct {
	HssId             *datatype.UTF8String `avp:"HSS-ID,omitempty"`
	OriginatingSipUri *datatype.UTF8String `avp:"Originating-SIP-URI,omitempty"`
	DestinationSipUri *datatype.UTF8String `avp:"Destination-SIP-URI,omitempty"`
}

type User_Identifier struct {
	UserName           *datatype.UTF8String  `avp:"User-Name,omitempty"`
	Msisdn             *datatype.OctetString `avp:"MSISDN,omitempty"`
	ExternalIdentifier *datatype.UTF8String  `avp:"External-Identifier,omitempty"`
	Lmsi               *datatype.OctetString `avp:"LMSI,omitempty"`
}

type SM_Delivery_Failure_Cause struct {
	SmEnumeratedDeliveryFailureCause datatype.Enumerated   `avp:"SM-Enumerated-Delivery-Failure-Cause"`
	SmDiagnosticInfo                 *datatype.OctetString `avp:"SM-Diagnostic-Info,omitempty"`
}

type SM_Delivery_Outcome struct {
	MmeSmDeliveryOutcome    *Delivery_Outcome `avp:"MME-SM-Delivery-Outcome,omitempty"`
	MscSmDeliveryOutcome    *Delivery_Outcome `avp:"MSC-SM-Delivery-Outcome,omitempty"`
	SgsnSmDeliveryOutcome   *Delivery_Outcome `avp:"SGSN-SM-Delivery-Outcome,omitempty"`
	IpsmgwSmDeliveryOutcome *Delivery_Outcome `avp:"IP-SM-GW-SM-Delivery-Outcome,omitempty"`
}

type Delivery_Outcome struct {
	SmDeliveryCause        *datatype.Enumerated `avp:"SM-Delivery-Cause,omitempty"`
	AbsentUserDiagnosticSm *datatype.Unsigned32 `avp:"Absent-User-Diagnostic-SM,omitempty"`
}

func (ui *User_Identifier) Empty() bool {
	return (ui.UserName == nil) && (ui.Msisdn == nil) && (ui.ExternalIdentifier == nil) && (ui.Lmsi == nil)
}

func (do *Delivery_Outcome) Empty() bool {
	return (do.SmDeliveryCause == nil) && (do.AbsentUserDiagnosticSm == nil)
}

func (sdo *SM_Delivery_Outcome) Empty() bool {
	return (sdo.MmeSmDeliveryOutcome == nil || sdo.MmeSmDeliveryOutcome.Empty()) && (sdo.SgsnSmDeliveryOutcome == nil || sdo.SgsnSmDeliveryOutcome.Empty()) && (sdo.IpsmgwSmDeliveryOutcome == nil || sdo.IpsmgwSmDeliveryOutcome.Empty()) && (sdo.MscSmDeliveryOutcome == nil || sdo.MscSmDeliveryOutcome.Empty())
}

func (sf *Supported_Features) String() string {
	return fmt.Sprintf("VendorId: %d, FeatureListId: %d, FeatureList: %d",
		sf.VendorId, sf.FeatureListId, sf.FeatureList)
}

func (smi *SMSMI_Correlation_ID) String() string {
	result := "SMSMI_Correlation_ID{"
	result += fmt.Sprintf("HssId: %s, ", smi.HssId.String())

	if smi.OriginatingSipUri != nil {
		result += fmt.Sprintf("OriginatingSipUri: %s, ", smi.OriginatingSipUri.String())
	}
	if smi.DestinationSipUri != nil {
		result += fmt.Sprintf("DestinationSipUri: %s", smi.DestinationSipUri.String())
	}
	result += "}"
	return result
}

func (ui *User_Identifier) String() string {
	result := "User_Identifier{"
	if ui.UserName != nil {
		result += fmt.Sprintf("UserName: %s, ", ui.UserName.String())
	}
	if ui.Msisdn != nil {
		result += fmt.Sprintf("Msisdn: %s, ", ui.Msisdn.String())
	}
	if ui.ExternalIdentifier != nil {
		result += fmt.Sprintf("ExternalIdentifier: %s, ", ui.ExternalIdentifier.String())
	}
	if ui.Lmsi != nil {
		result += fmt.Sprintf("Lmsi: %s", ui.Lmsi.String())
	}
	result += "}"
	return result
}

func (smdf *SM_Delivery_Failure_Cause) String() string {
	result := "SM_Delivery_Failure_Cause{"
	result += fmt.Sprintf("SmEnumeratedDeliveryFailureCause: %s", smdf.SmEnumeratedDeliveryFailureCause.String())
	if smdf.SmDiagnosticInfo != nil {
		result += fmt.Sprintf(", SmDiagnosticInfo: %s", smdf.SmDiagnosticInfo.String())
	}
	result += "}"
	return result
}

func (smo *SM_Delivery_Outcome) String() string {
	result := "SM_Delivery_Outcome{"
	if smo.MmeSmDeliveryOutcome != nil {
		result += fmt.Sprintf("MmeSmDeliveryOutcome: %s, ", smo.MmeSmDeliveryOutcome.String())
	}
	if smo.MscSmDeliveryOutcome != nil {
		result += fmt.Sprintf("MscSmDeliveryOutcome: %s, ", smo.MscSmDeliveryOutcome.String())
	}
	if smo.SgsnSmDeliveryOutcome != nil {
		result += fmt.Sprintf("SgsnSmDeliveryOutcome: %s, ", smo.SgsnSmDeliveryOutcome.String())
	}
	if smo.IpsmgwSmDeliveryOutcome != nil {
		result += fmt.Sprintf("IpsmgwSmDeliveryOutcome: %s", smo.IpsmgwSmDeliveryOutcome.String())
	}
	result += "}"
	return result
}

func (d *Delivery_Outcome) String() string {
	result := "Delivery_Outcome{"
	if d.SmDeliveryCause != nil {
		result += fmt.Sprintf("SmDeliveryCause: %s, ", d.SmDeliveryCause.String())
	}
	if d.AbsentUserDiagnosticSm != nil {
		result += fmt.Sprintf("AbsentUserDiagnosticSm: %s", d.AbsentUserDiagnosticSm.String())
	}
	result += "}"
	return result
}

// encode Supported-Features struct to grouped AVP
func (sf *Supported_Features) ToDiam() *diam.GroupedAVP {
	return &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, sf.VendorId),
			diam.NewAVP(avp.FeatureListID, avp.Vbit, 10415, sf.FeatureListId),
			diam.NewAVP(avp.FeatureList, avp.Vbit, 10415, sf.FeatureList),
		},
	}
}

// encode SMSMI-Correlation-ID struct to grouped AVP
func (smiId *SMSMI_Correlation_ID) ToDiam() *diam.GroupedAVP {
	a := diam.GroupedAVP{
		AVP: []*diam.AVP{},
	}
	if smiId.HssId != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.HSSID, avp.Vbit, 10415, *smiId.HssId))
	}
	if smiId.OriginatingSipUri != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.OriginatingSIPURI, avp.Vbit, 10415, *smiId.OriginatingSipUri))
	}
	if smiId.DestinationSipUri != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.DestinationSIPURI, avp.Vbit, 10415, *smiId.DestinationSipUri))
	}
	return &a
}

// encode User-Identifier struct to grouped AVP
func (ui *User_Identifier) ToDiam() *diam.GroupedAVP {
	a := diam.GroupedAVP{
		AVP: []*diam.AVP{},
	}
	if ui.UserName != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.UserName, avp.Mbit, 0, *ui.UserName))
	}
	if ui.Msisdn != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, *ui.Msisdn))
	}
	if ui.ExternalIdentifier != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.ExternalIdentifier, avp.Mbit|avp.Vbit, 10415, *ui.ExternalIdentifier))
	}
	if ui.Lmsi != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.LMSI, avp.Mbit|avp.Vbit, 10415, *ui.Lmsi))
	}
	return &a
}
