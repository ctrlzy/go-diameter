package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
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
		vendorID = vsa.VendorId.String()
	}

	return fmt.Sprintf("VendorId: %s, AuthApplicationId: %v, AcctApplicationId: %v",
		vendorID, vsa.AuthApplicationId.String(), vsa.AcctApplicationId.String())
}

func (osf *OC_Supported_Features) String() string {
	if osf.OcFeatureVector != nil {
		return fmt.Sprintf("OcFeatureVector: %v", osf.OcFeatureVector.String())
	} else {
		return "OcFeatureVector: nil"
	}
}

func (olr *OC_OLR) String() string {
	redPct := "nil"
	if olr.OcReductionPercentage != nil {
		redPct = fmt.Sprintf("%v", olr.OcReductionPercentage.String())
	}

	valDur := "nil"
	if olr.OcValidityDuration != nil {
		valDur = fmt.Sprintf("%v", olr.OcValidityDuration.String())
	}

	return fmt.Sprintf("OCSequenceNumber: %v, OCReportType: %v, OCReductionPercentage: %v, OCValidityDuration: %v",
		olr.OcSequenceNumber, olr.OcReportType, redPct, valDur)
}

func (info *Proxy_Info) String() string {
	return fmt.Sprintf("ProxyHost: %s, ProxyState: %s", info.ProxyHost.String(), info.ProxyState.String())
}

func (result *Experimental_Result) String() string {
	return fmt.Sprintf("VendorId: %d, ExperimentalResultCode: %d", result.VendorId, result.ExperimentalResultCode)
}

// encode Vendor-Specific-Application-Id struct to grouped AVP
func (vsid *Vendor_Specific_Application_Id) ToDiam() *diam.GroupedAVP {
	a := diam.GroupedAVP{
		AVP: []*diam.AVP{},
	}
	if vsid.VendorId != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.VendorID, avp.Mbit, 0, *vsid.VendorId))
	}
	a.AVP = append(a.AVP, diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, vsid.AuthApplicationId))
	a.AVP = append(a.AVP, diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, vsid.AcctApplicationId))
	return &a
}

// encode Proxy-Info struct to grouped AVP
func (pi *Proxy_Info) ToDiam() *diam.GroupedAVP {
	return &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.ProxyHost, avp.Mbit, 0, pi.ProxyHost),
			diam.NewAVP(avp.ProxyState, avp.Mbit, 0, pi.ProxyState),
		},
	}
}

// encode OC-Supported-Features struct to grouped AVP
func (ocsf *OC_Supported_Features) ToDiam() *diam.GroupedAVP {
	a := diam.GroupedAVP{
		AVP: []*diam.AVP{},
	}
	if ocsf.OcFeatureVector != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.OCFeatureVector, 0, 0, *ocsf.OcFeatureVector))
	}
	return &a
}

// encode OC-OLR struct to grouped AVP
func (olr *OC_OLR) ToDiam() *diam.GroupedAVP {
	a := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.OCSequenceNumber, 0, 0, olr.OcSequenceNumber),
			diam.NewAVP(avp.OCReportType, 0, 0, olr.OcReportType),
		},
	}
	if olr.OcReductionPercentage != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.OCReductionPercentage, 0, 0, *olr.OcReductionPercentage))
	}
	if olr.OcValidityDuration != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.OCValidityDuration, 0, 0, *olr.OcValidityDuration))
	}
	return &a
}

// Encode Experimental-Result struct to grouped AVP
func (er *Experimental_Result) ToDiam() *diam.GroupedAVP {
	return &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, er.VendorId),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, er.ExperimentalResultCode),
		},
	}
}
