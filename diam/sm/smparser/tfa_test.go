package smparser_test

import (
	"fmt"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/stretchr/testify/assert"
)

func TestTFA_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	tfr := new(smparser.TFA)
	err := tfr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestTFA_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	tfr := new(smparser.TFA)
	err := tfr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestTFA_OK(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("test"))
	tfr := new(smparser.TFA)
	err := tfr.Parse(m)
	assert.Nil(t, err)
}

func TestTFA_PARSE_OK(t *testing.T) {
	m := createDiamTFA()
	tfa := smparser.TFA{}
	err := tfa.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, tfa.SessionId, datatype.UTF8String("session-id"))
	assert.Empty(t, tfa.Drmp)
	assert.Empty(t, tfa.VendorSpecificApplicationId)
	assert.Empty(t, tfa.ResultCode)
	assert.Equal(t, tfa.ExperimentalResult.VendorId, datatype.Unsigned32(10415))
	assert.Equal(t, tfa.ExperimentalResult.ExperimentalResultCode, datatype.Unsigned32(11))
	assert.Equal(t, tfa.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, tfa.OriginHost, datatype.DiameterIdentity("orig-host"))
	assert.Equal(t, tfa.OriginRealm, datatype.DiameterIdentity("orig-realm"))
	assert.Empty(t, tfa.SupportedFeatures)
	assert.Equal(t, *tfa.AbsentUserDiagnosticSm, datatype.Unsigned32(12))
	assert.Equal(t, tfa.SmDeliveryFailureCause.SmEnumeratedDeliveryFailureCause, datatype.Enumerated(0))
	assert.Empty(t, tfa.SmDeliveryFailureCause.SmDiagnosticInfo)
	assert.Empty(t, tfa.SmRpUi)
	assert.Empty(t, tfa.RequestedRetransmissionTime)
	assert.Empty(t, tfa.UserIdentifier)
	assert.Empty(t, tfa.FailedAvp)
	assert.Empty(t, tfa.ProxyInfo)
	assert.Empty(t, tfa.RouteRecord)
}

func TestTFA_Decode_OK(t *testing.T) {
	tfa := createStructTFA()
	msg := diam.NewMessage(diam.MTForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	err := msg.Marshal(tfa)
	assert.Nil(t, err)
	fmt.Print(msg)
}

func createDiamTFA() *diam.Message {
	m := diam.NewMessage(diam.MTForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(11)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.AbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(12))
	m.NewAVP(avp.SMDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.SMEnumeratedDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(0)),
		},
	})
	return m
}

func createStructTFA() *smparser.TFA {
	drmp := datatype.Enumerated(1)
	vsai := basetype.Vendor_Specific_Application_Id{
		AuthApplicationId: 123,
		AcctApplicationId: 456,
	}

	resultCode := datatype.Unsigned32(1)
	experimentalResult := basetype.Experimental_Result{
		VendorId:               datatype.Unsigned32(10415),
		ExperimentalResultCode: datatype.Unsigned32(123),
	}
	tfa := &smparser.TFA{
		SessionId:                   "session-id",
		Drmp:                        &drmp,
		VendorSpecificApplicationId: &vsai,
		ResultCode:                  &resultCode,
		ExperimentalResult:          &experimentalResult,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
	}
	return tfa
}

func BenchmarkEncodeTFA(b *testing.B) {
	tfa := createStructTFA()
	m1 := diam.NewMessage(diam.MTForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(tfa)
	}
}

func BenchmarkDecodeTFA(b *testing.B) {
	m := createDiamTFA()
	tfa := smparser.TFA{}
	for n := 0; n < b.N; n++ {
		tfa.Parse(m)
	}
}
