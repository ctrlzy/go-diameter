package basetype

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

type Serving_Node struct {
	SgsnName          datatype.DiameterIdentity `avp:"SGSN-Name,omitempty"`
	SgsnRealm         datatype.DiameterIdentity `avp:"SGSN-Realm,omitempty"`
	SgsnNumber        datatype.OctetString      `avp:"SGSN-Number,omitempty"`
	MmeName           datatype.DiameterIdentity `avp:"MME-Name,omitempty"`
	MmeRealm          datatype.DiameterIdentity `avp:"MME-Realm,omitempty"`
	MmeNumberForMtSms datatype.OctetString      `avp:"MME-Number-for-MT-SMS,omitempty"`
	MscNumber         datatype.OctetString      `avp:"MSC-Number,omitempty"`
	IpsmgwNumber      datatype.OctetString      `avp:"IP-SM-GW-Number,omitempty"`
	IpsmgwName        datatype.DiameterIdentity `avp:"IP-SM-GW-Name,omitempty"`
	IpsmgwRealm       datatype.DiameterIdentity `avp:"IP-SM-GW-Realm,omitempty"`
}

type Additional_Serving_Node_T4 struct {
	SgsnName          datatype.DiameterIdentity `avp:"SGSN-Name,omitempty"`
	SgsnRealm         datatype.DiameterIdentity `avp:"SGSN-Realm,omitempty"`
	SgsnNumber        datatype.OctetString      `avp:"SGSN-Number,omitempty"`
	MmeName           datatype.DiameterIdentity `avp:"MME-Name,omitempty"`
	MmeRealm          datatype.DiameterIdentity `avp:"MME-Realm,omitempty"`
	MmeNumberForMtSms datatype.OctetString      `avp:"MME-Number-for-MT-SMS,omitempty"`
	MscNumber         datatype.OctetString      `avp:"MSC-Number,omitempty"`
}

type SMSF_SM_Delivery_Outcome struct {
	SmDeliveryCause         datatype.Enumerated `avp:"Absent-User-Diagnostic-SM,omitempty"`
	AbsentUserDiagnosticsSM datatype.Unsigned32 `avp:"Absent-User-Diagnostic-SM,omitempty"`
}

type SMSF_3GPP_Address struct {
	Smsf3gppNumber datatype.OctetString      `avp:"SMSF-3GPP-Number,omitempty"`
	Smsf3gppName   datatype.DiameterIdentity `avp:"SMSF-3GPP-Name,omitempty"`
	Smsf3gppRealm  datatype.DiameterIdentity `avp:"SMSF-3GPP-Realm,omitempty"`
}

type SMSF_Non_3GPP_Address struct {
	SmsfNon3gppNumber datatype.OctetString      `avp:"SMSF-Non-3GPP-Number,omitempty"`
	SmsfNon3gppName   datatype.DiameterIdentity `avp:"SMSF-Non-3GPP-Name,omitempty"`
	SmsfNon3gppRealm  datatype.DiameterIdentity `avp:"SMSF-Non-3GPP-Realm,omitempty"`
}
