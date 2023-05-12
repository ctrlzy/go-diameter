package smparser_test

import (
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/stretchr/testify/assert"
)

func TestRDA_MissingOriginHost(t *testing.T) {
	m := diam.NewMessage(diam.ReportSMDeliveryStatus, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	rda := new(smparser.RDA)
	err := rda.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestRDA_MissingOriginRealm(t *testing.T) {
	m := diam.NewMessage(diam.ReportSMDeliveryStatus, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	rda := new(smparser.RDA)
	err := rda.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestRDA_PARSE_OK(t *testing.T) {
	m := diam.NewMessage(diam.ReportSMDeliveryStatus, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(11))
	m.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(12)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))

	rda := new(smparser.RDA)
	err := rda.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, rda.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, rda.Drmp)
	assert.Nil(t, rda.VendorSpecificApplicationId)
	assert.Equal(t, rda.AuthSessionState, datatype.Enumerated(1))
	assert.Equal(t, *rda.ResultCode, datatype.Unsigned32(11))
	assert.Equal(t, rda.ExperimentalResult.VendorId, datatype.Unsigned32(10415))
	assert.Equal(t, rda.ExperimentalResult.ExperimentalResultCode, datatype.Unsigned32(12))
	assert.Empty(t, rda.SupportedFeatures)
	assert.Nil(t, rda.FailedAvp)
	assert.Empty(t, rda.ProxyInfo)
	assert.Empty(t, rda.RouteRecord)
}

func TestRDA_Decode_OK(t *testing.T) {
	vsai := basetype.Vendor_Specific_Application_Id{
		AuthApplicationId: 123,
		AcctApplicationId: 456,
	}
	resultCode := datatype.Unsigned32(10)
	experimentalResult := basetype.Experimental_Result{
		VendorId:               datatype.Unsigned32(10415),
		ExperimentalResultCode: datatype.Unsigned32(20),
	}

	rda := &smparser.RDA{
		SessionId:                   "session-id",
		VendorSpecificApplicationId: &vsai,
		ResultCode:                  &resultCode,
		ExperimentalResult:          &experimentalResult,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
	}

	m1 := diam.NewMessage(diam.ReportSMDeliveryStatus, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	err := m1.Marshal(rda)
	assert.Nil(t, err)

	m2 := diam.NewMessage(diam.ReportSMDeliveryStatus, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	})
	m2.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(10))
	m2.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(20)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))

	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}
