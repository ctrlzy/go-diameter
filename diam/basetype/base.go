package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type Vendor_Specific_Application_Id struct {
	VendorId          *datatype.Unsigned32 `avp:"Vendor-Id,omitempty"`
	AuthApplicationId datatype.Unsigned32  `avp:"Auth-Application-Id"`
	AcctApplicationId datatype.Unsigned32  `avp:"Acct-Application-Id"`
}

type OC_Supported_Features struct {
	OcFeatureVector *datatype.Unsigned64 `avp:"OC-Feature-Vector,omitempty"`
}

type OC_OLR struct {
	OcSequenceNumber      datatype.Unsigned64  `avp:"OC-Sequence-Number"`
	OcReportType          datatype.Enumerated  `avp:"OC-Report-Type"`
	OcReductionPercentage *datatype.Unsigned32 `avp:"OC-Reduction-Percentage,omitempty"`
	OcValidityDuration    *datatype.Unsigned32 `avp:"OC-Validity-Duration,omitempty"`
}

type Proxy_Info struct {
	ProxyHost  datatype.DiameterIdentity `avp:"Proxy-Host"`
	ProxyState datatype.OctetString      `avp:"Proxy-State"`
}

type Experimental_Result struct {
	VendorId               datatype.Unsigned32 `avp:"Vendor-Id"`
	ExperimentalResultCode datatype.Unsigned32 `avp:"Experimental-Result-Code"`
}

type Failed_AVP []*diam.AVP

func (vsa *Vendor_Specific_Application_Id) Empty() bool {
	return vsa.AcctApplicationId == 0 && vsa.AuthApplicationId == 0
}

func (vsa *Vendor_Specific_Application_Id) String() string {
	vendorID := "nil"
	if vsa.VendorId != nil {
		vendorID = fmt.Sprintf("%v", *vsa.VendorId)
	}

	return fmt.Sprintf("VendorId: %v, AuthApplicationId: %v, AcctApplicationId: %v",
		vendorID, vsa.AuthApplicationId, vsa.AcctApplicationId)
}

func (osf *OC_Supported_Features) String() string {
	if osf.OcFeatureVector != nil {
		return fmt.Sprintf("OcFeatureVector: %v", *osf.OcFeatureVector)
	} else {
		return "OcFeatureVector: nil"
	}
}

func (olr *OC_OLR) String() string {
	redPct := "nil"
	if olr.OcReductionPercentage != nil {
		redPct = fmt.Sprintf("%v", *olr.OcReductionPercentage)
	}

	valDur := "nil"
	if olr.OcValidityDuration != nil {
		valDur = fmt.Sprintf("%v", *olr.OcValidityDuration)
	}

	return fmt.Sprintf("OCSequenceNumber: %v, OCReportType: %v, OCReductionPercentage: %v, OCValidityDuration: %v",
		olr.OcSequenceNumber, olr.OcReportType, redPct, valDur)
}

func (info *Proxy_Info) String() string {
	return fmt.Sprintf("ProxyHost: %s, ProxyState: %s", info.ProxyHost, info.ProxyState)
}

func (result *Experimental_Result) String() string {
	return fmt.Sprintf("VendorId: %d, ExperimentalResultCode: %d", result.VendorId, result.ExperimentalResultCode)
}
