package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type User_Identity struct {
	PublicIdentity     *datatype.UTF8String  `avp:"Public-Identity,omitempty"`
	MSISDN             *datatype.OctetString `avp:"MSISDN,omitempty"`
	ExternalIdentifier *datatype.UTF8String  `avp:"External-Identifier,omitempty"`
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
	AuthApplicationId           *datatype.Unsigned32            `avp:"Auth-Application-Id,omitempty"`
	AcctApplicationId           *datatype.Unsigned32            `avp:"Acct-Application-Id,omitempty"`
	VendorSpecificApplicationId *Vendor_Specific_Application_Id `avp:"Vendor-Specific-Application-Id,omitempty"`
}

func (ui User_Identity) Empty() bool {
	return (ui.PublicIdentity == nil) && (ui.MSISDN == nil) && (ui.ExternalIdentifier == nil)
}

func (u *User_Identity) String() string {
	result := "User_Identity { "
	if u.PublicIdentity != nil {
		result += fmt.Sprintf("PublicIdentity: %s, ", u.PublicIdentity.String())
	}
	if u.MSISDN != nil {
		result += fmt.Sprintf("MSISDN: %s, ", u.MSISDN.String())
	}
	if u.ExternalIdentifier != nil {
		result += fmt.Sprintf("ExternalIdentifier: %s, ", u.ExternalIdentifier.String())
	}
	result += "}"
	return result
}

func (r *Repository_Data_ID) String() string {
	return fmt.Sprintf("Repository_Data_ID { ServiceIndication: %s, SequenceNumber: %d }", r.ServiceIndication.String(), r.SequenceNumber)
}

func (c *Call_Reference_Info) String() string {
	return fmt.Sprintf("Call_Reference_Info { CallReferenceNumber: %s, AsNumber: %s }", c.CallReferenceNumber.String(), c.AsNumber.String())
}

func (s *Supported_Applications) String() string {
	result := "Supported_Applications { "
	if s.AuthApplicationId != nil {
		result += fmt.Sprintf("AuthApplicationId: %s, ", s.AuthApplicationId.String())
	}
	if s.AcctApplicationId != nil {
		result += fmt.Sprintf("AcctApplicationId: %s, ", s.AcctApplicationId.String())
	}
	if s.VendorSpecificApplicationId != nil {
		result += fmt.Sprintf("VendorSpecificApplicationId: %s, ", s.VendorSpecificApplicationId.String())
	}
	result += "}"
	return result
}
