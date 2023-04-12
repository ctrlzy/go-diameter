package basetype

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

type User_Identity struct {
	PublicIdentity     datatype.UTF8String  `avp:"Public-Identity,omitempty"`
	MSISDN             datatype.OctetString `avp:"MSISDN,omitempty"`
	ExternalIdentifier datatype.UTF8String  `avp:"External-Identifier,omitempty"`
}

type Repository_Data_ID struct {
	ServiceIndication datatype.OctetString `avp:"Service-Indication"`
	SequenceNumber    datatype.Unsigned32  `avp:"Sequence-Number"`
}

type Call_Reference_Info struct {
	CallReferenceNumber datatype.OctetString `avp:"Call-Reference-Number"`
	AsNumber            datatype.OctetString `avp:"AS-Number"`
}

type Supported_Applications struct {
	AuthApplicationId           datatype.Unsigned32            `avp:"Auth-Application-Id,omitempty"`
	AcctApplicationId           datatype.Unsigned32            `avp:"Acct-Application-Id,omitempty"`
	VendorSpecificApplicationId Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
}

func (ui User_Identity) Empty() bool {
	return (len(ui.PublicIdentity) + len(ui.MSISDN) + len(ui.ExternalIdentifier)) == 0
}
