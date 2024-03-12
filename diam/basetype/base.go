package basetype

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

type Vendor_Specific_Application_Id struct {
	VendorId          datatype.Unsigned32  `avp:"Vendor-Id"`
	AuthApplicationId *datatype.Unsigned32 `avp:"Auth-Application-Id,omitempty"`
	AcctApplicationId *datatype.Unsigned32 `avp:"Acct-Application-Id,omitempty"`
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
	return vsa.AcctApplicationId == nil && vsa.AuthApplicationId == nil
}

func (vsa *Vendor_Specific_Application_Id) String() string {
	if vsa.AcctApplicationId != nil {
		return fmt.Sprintf("VendorId: %s, AcctApplicationId: %v", vsa.VendorId.String(), vsa.AcctApplicationId.String())
	}
	if vsa.AuthApplicationId != nil {
		return fmt.Sprintf("VendorId: %s, AuthApplicationId: %v", vsa.VendorId.String(), vsa.AuthApplicationId.String())
	}
	return vsa.VendorId.String()
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
	a.AVP = append(a.AVP, diam.NewAVP(avp.VendorID, avp.Mbit, 0, vsid.VendorId))

	if vsid.AuthApplicationId != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, vsid.AuthApplicationId))
	}
	if vsid.AcctApplicationId != nil {
		a.AVP = append(a.AVP, diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, vsid.AcctApplicationId))
	}
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
