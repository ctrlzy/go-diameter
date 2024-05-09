package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type UserIdentity struct {
	PublicIdentity     *datatype.UTF8String  `avp:"Public-Identity,omitempty"`
	MSISDN             *datatype.OctetString `avp:"MSISDN,omitempty"`
	ExternalIdentifier *datatype.UTF8String  `avp:"External-Identifier,omitempty"`
}

type RepositoryDataID struct {
	ServiceIndication datatype.OctetString `avp:"Service-Indication"`
	SequenceNumber    datatype.Unsigned32  `avp:"Sequence-Number"`
}

type CallReferenceInfo struct {
	CallReferenceNumber datatype.OctetString `avp:"Call-Reference-Number"`
	AsNumber            datatype.OctetString `avp:"AS-Number"`
}

type SupportedApplications struct {
	AuthApplicationId           *datatype.Unsigned32         `avp:"Auth-Application-Id,omitempty"`
	AcctApplicationId           *datatype.Unsigned32         `avp:"Acct-Application-Id,omitempty"`
	VendorSpecificApplicationId *VendorSpecificApplicationId `avp:"Vendor-Specific-Application-Id,omitempty"`
}

func (ui UserIdentity) Empty() bool {
	return (ui.PublicIdentity == nil) && (ui.MSISDN == nil) && (ui.ExternalIdentifier == nil)
}

func (u *UserIdentity) String() string {
	result := "UserIdentity { "
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

func (r *RepositoryDataID) String() string {
	return fmt.Sprintf("Repository_Data_ID { ServiceIndication: %s, SequenceNumber: %d }", r.ServiceIndication.String(), r.SequenceNumber)
}

func (c *CallReferenceInfo) String() string {
	return fmt.Sprintf("Call_Reference_Info { CallReferenceNumber: %s, AsNumber: %s }", c.CallReferenceNumber.String(), c.AsNumber.String())
}

func (s *SupportedApplications) String() string {
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
