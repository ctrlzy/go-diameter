package basetype

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

type Vendor_Specific_Application_Id struct {
	VendorId          datatype.Unsigned32 `avp:"Vendor-Id,omitempty"`
	AuthApplicationId datatype.Unsigned32 `avp:"Auth-Application-Id"`
	AcctApplicationId datatype.Unsigned32 `avp:"Acct-Application-Id"`
}

type OC_Supported_Features struct {
	OcFeatureVector datatype.Unsigned64 `avp:"OC-Feature-Vector,omitempty"`
}

type OC_OLR struct {
	OcSequenceNumber      datatype.Unsigned64 `avp:"OC-Sequence-Number"`
	OcReportType          datatype.Enumerated `avp:"OC-Report-Type"`
	OcReductionPercentage datatype.Unsigned32 `avp:"OC-Reduction-Percentage,omitempty"`
	OcValidityDuration    datatype.Unsigned32 `avp:"OC-Validity-Duration,omitempty"`
}

type Proxy_Info struct {
	ProxyHost  datatype.DiameterIdentity `avp:"Proxy-Host"`
	ProxyState datatype.OctetString      `avp:"Proxy-State"`
}

type Experimental_Result struct {
	VendorId               datatype.Unsigned32 `avp:"Vendor-Id"`
	ExperimentalResultCode datatype.Unsigned32 `avp:"Experimental-Result-Code"`
}

type Failed_AVP struct {
}

func (vsai *Vendor_Specific_Application_Id) Empty() bool {
	return vsai.AcctApplicationId == 0 && vsai.AuthApplicationId == 0
}
