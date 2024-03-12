package basetype_test

import (
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/stretchr/testify/assert"
)

func TestVendorSpecificApplicationIdToDiam(t *testing.T) {
	authId := datatype.Unsigned32(123)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          datatype.Unsigned32(10415),
		AuthApplicationId: &authId,
	}
	m1 := vsai.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(0)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestVendorSpecificApplicationIdWithoutOptionalFieldToDiam(t *testing.T) {
	authId := datatype.Unsigned32(123)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          datatype.Unsigned32(10415),
		AuthApplicationId: &authId,
	}
	m1 := vsai.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestOCSupportedFeaturesToDiam(t *testing.T) {
	vector := datatype.Unsigned64(123)
	ocsf := basetype.OC_Supported_Features{
		OcFeatureVector: &vector,
	}
	m1 := ocsf.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.OCFeatureVector, 0, 0, datatype.Unsigned64(123)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestOCOLRToDiam(t *testing.T) {
	per := datatype.Unsigned32(456)
	val := datatype.Unsigned32(789)
	olr := basetype.OC_OLR{
		OcSequenceNumber:      datatype.Unsigned64(123),
		OcReportType:          datatype.Enumerated(1),
		OcReductionPercentage: &per,
		OcValidityDuration:    &val,
	}
	m1 := olr.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.OCSequenceNumber, 0, 0, datatype.Unsigned64(123)),
			diam.NewAVP(avp.OCReportType, 0, 0, datatype.Enumerated(1)),
			diam.NewAVP(avp.OCReductionPercentage, 0, 0, datatype.Unsigned32(456)),
			diam.NewAVP(avp.OCValidityDuration, 0, 0, datatype.Unsigned32(789)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestOCOLRTWithoutOptionalFieldoDiam(t *testing.T) {
	olr := basetype.OC_OLR{
		OcSequenceNumber: datatype.Unsigned64(123),
		OcReportType:     datatype.Enumerated(1),
	}
	m1 := olr.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.OCSequenceNumber, 0, 0, datatype.Unsigned64(123)),
			diam.NewAVP(avp.OCReportType, 0, 0, datatype.Enumerated(1)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestProxyInfoToDiam(t *testing.T) {
	info := basetype.Proxy_Info{
		ProxyHost:  datatype.DiameterIdentity("abc"),
		ProxyState: datatype.OctetString("888"),
	}
	m1 := info.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.ProxyHost, avp.Mbit, 0, datatype.DiameterIdentity("abc")),
			diam.NewAVP(avp.ProxyState, avp.Mbit, 0, datatype.OctetString("888")),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}

func TestExperimentalResultToDiam(t *testing.T) {
	result := basetype.Experimental_Result{
		VendorId:               datatype.Unsigned32(123),
		ExperimentalResultCode: datatype.Unsigned32(456),
	}
	m1 := result.ToDiam()
	m2 := diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	}
	assert.Equal(t, m2.Serialize(), m1.Serialize())
}
