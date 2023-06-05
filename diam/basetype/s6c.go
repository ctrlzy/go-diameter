package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type Serving_Node struct {
	SgsnName          *datatype.DiameterIdentity `avp:"SGSN-Name,omitempty"`
	SgsnRealm         *datatype.DiameterIdentity `avp:"SGSN-Realm,omitempty"`
	SgsnNumber        *datatype.OctetString      `avp:"SGSN-Number,omitempty"`
	MmeName           *datatype.DiameterIdentity `avp:"MME-Name,omitempty"`
	MmeRealm          *datatype.DiameterIdentity `avp:"MME-Realm,omitempty"`
	MmeNumberForMtSms *datatype.OctetString      `avp:"MME-Number-for-MT-SMS,omitempty"`
	MscNumber         *datatype.OctetString      `avp:"MSC-Number,omitempty"`
	IpsmgwNumber      *datatype.OctetString      `avp:"IP-SM-GW-Number,omitempty"`
	IpsmgwName        *datatype.DiameterIdentity `avp:"IP-SM-GW-Name,omitempty"`
	IpsmgwRealm       *datatype.DiameterIdentity `avp:"IP-SM-GW-Realm,omitempty"`
}

type Additional_Serving_Node struct {
	SgsnName          *datatype.DiameterIdentity `avp:"SGSN-Name,omitempty"`
	SgsnRealm         *datatype.DiameterIdentity `avp:"SGSN-Realm,omitempty"`
	SgsnNumber        *datatype.OctetString      `avp:"SGSN-Number,omitempty"`
	MmeName           *datatype.DiameterIdentity `avp:"MME-Name,omitempty"`
	MmeRealm          *datatype.DiameterIdentity `avp:"MME-Realm,omitempty"`
	MmeNumberForMtSms *datatype.OctetString      `avp:"MME-Number-for-MT-SMS,omitempty"`
	MscNumber         *datatype.OctetString      `avp:"MSC-Number,omitempty"`
}

type SMSF_SM_Delivery_Outcome struct {
	SmDeliveryCause         *datatype.Enumerated `avp:"Absent-User-Diagnostic-SM,omitempty"`
	AbsentUserDiagnosticsSM *datatype.Unsigned32 `avp:"Absent-User-Diagnostic-SM,omitempty"`
}

type SMSF_3GPP_Address struct {
	Smsf3gppNumber *datatype.OctetString      `avp:"SMSF-3GPP-Number,omitempty"`
	Smsf3gppName   *datatype.DiameterIdentity `avp:"SMSF-3GPP-Name,omitempty"`
	Smsf3gppRealm  *datatype.DiameterIdentity `avp:"SMSF-3GPP-Realm,omitempty"`
}

type SMSF_Non_3GPP_Address struct {
	SmsfNon3gppNumber *datatype.OctetString      `avp:"SMSF-Non-3GPP-Number,omitempty"`
	SmsfNon3gppName   *datatype.DiameterIdentity `avp:"SMSF-Non-3GPP-Name,omitempty"`
	SmsfNon3gppRealm  *datatype.DiameterIdentity `avp:"SMSF-Non-3GPP-Realm,omitempty"`
}

func (s *Serving_Node) String() string {
	result := "Serving_Node { "
	if s.SgsnName != nil {
		result += fmt.Sprintf("SgsnName: %s, ", s.SgsnName.String())
	}
	if s.SgsnRealm != nil {
		result += fmt.Sprintf("SgsnRealm: %s, ", s.SgsnRealm.String())
	}
	if s.SgsnNumber != nil {
		result += fmt.Sprintf("SgsnNumber: %s, ", s.SgsnNumber.String())
	}
	if s.MmeName != nil {
		result += fmt.Sprintf("MmeName: %s, ", s.MmeName.String())
	}
	if s.MmeRealm != nil {
		result += fmt.Sprintf("MmeRealm: %s, ", s.MmeRealm.String())
	}
	if s.MmeNumberForMtSms != nil {
		result += fmt.Sprintf("MmeNumberForMtSms: %s, ", s.MmeNumberForMtSms.String())
	}
	if s.MscNumber != nil {
		result += fmt.Sprintf("MscNumber: %s, ", s.MscNumber.String())
	}
	if s.IpsmgwNumber != nil {
		result += fmt.Sprintf("IpsmgwNumber: %s, ", s.IpsmgwNumber.String())
	}
	if s.IpsmgwName != nil {
		result += fmt.Sprintf("IpsmgwName: %s, ", s.IpsmgwName.String())
	}
	if s.IpsmgwRealm != nil {
		result += fmt.Sprintf("IpsmgwRealm: %s, ", s.IpsmgwRealm.String())
	}
	result += "}"
	return result
}

func (a *Additional_Serving_Node) String() string {
	result := "Additional_Serving_Node { "
	if a.SgsnName != nil {
		result += fmt.Sprintf("SgsnName: %s, ", a.SgsnName.String())
	}
	if a.SgsnRealm != nil {
		result += fmt.Sprintf("SgsnRealm: %s, ", a.SgsnRealm.String())
	}
	if a.SgsnNumber != nil {
		result += fmt.Sprintf("SgsnNumber: %s, ", a.SgsnNumber.String())
	}
	if a.MmeName != nil {
		result += fmt.Sprintf("MmeName: %s, ", a.MmeName.String())
	}
	if a.MmeRealm != nil {
		result += fmt.Sprintf("MmeRealm: %s, ", a.MmeRealm.String())
	}
	if a.MmeNumberForMtSms != nil {
		result += fmt.Sprintf("MmeNumberForMtSms: %s, ", a.MmeNumberForMtSms.String())
	}
	if a.MscNumber != nil {
		result += fmt.Sprintf("MscNumber: %s, ", a.MscNumber.String())
	}
	result += "}"
	return result
}

func (s *SMSF_SM_Delivery_Outcome) String() string {
	result := "SMSF_SM_Delivery_Outcome { "
	if s.SmDeliveryCause != nil {
		result += fmt.Sprintf("SmDeliveryCause: %s, ", s.SmDeliveryCause.String())
	}
	if s.AbsentUserDiagnosticsSM != nil {
		result += fmt.Sprintf("AbsentUserDiagnosticsSM: %s, ", s.AbsentUserDiagnosticsSM.String())
	}
	result += "}"
	return result
}

func (s *SMSF_3GPP_Address) String() string {
	result := "SMSF_3GPP_Address { "
	if s.Smsf3gppNumber != nil {
		result += fmt.Sprintf("Smsf3gppNumber: %s, ", s.Smsf3gppNumber.String())
	}
	if s.Smsf3gppName != nil {
		result += fmt.Sprintf("Smsf3gppName: %s, ", s.Smsf3gppName.String())
	}
	if s.Smsf3gppRealm != nil {
		result += fmt.Sprintf("Smsf3gppRealm: %s, ", s.Smsf3gppRealm.String())
	}
	result += "}"
	return result
}

func (s *SMSF_Non_3GPP_Address) String() string {
	result := "SMSF_Non_3GPP_Address { "
	if s.SmsfNon3gppNumber != nil {
		result += fmt.Sprintf("SmsfNon3gppNumber: %s, ", s.SmsfNon3gppNumber.String())
	}
	if s.SmsfNon3gppName != nil {
		result += fmt.Sprintf("SmsfNon3gppName: %s, ", s.SmsfNon3gppName.String())
	}
	if s.SmsfNon3gppRealm != nil {
		result += fmt.Sprintf("SmsfNon3gppRealm: %s, ", s.SmsfNon3gppRealm.String())
	}
	result += "}"
	return result
}
