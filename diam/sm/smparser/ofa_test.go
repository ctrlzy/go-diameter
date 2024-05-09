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

func TestOFA_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	ofa := new(smparser.OFA)
	err := ofa.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestOFA_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	ofa := new(smparser.OFA)
	err := ofa.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestOFA_OK(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("test"))
	ofa := new(smparser.OFA)
	err := ofa.Parse(m)
	assert.Nil(t, err)
}

func TestOFA_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id-1234"))
	m.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(22))
	m.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(1)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(2)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("test"))
	m.NewAVP(avp.SMDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.SMEnumeratedDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(12)),
		},
	})
	m.NewAVP(avp.ExternalIdentifier, avp.Mbit|avp.Vbit, 10415, datatype.UTF8String("external-identifier"))
	ofa := new(smparser.OFA)
	err := ofa.Parse(m)
	assert.Nil(t, err)

	assert.Equal(t, ofa.SessionId, datatype.UTF8String("session-id-1234"))
	assert.Nil(t, ofa.Drmp)
	assert.Equal(t, *ofa.ResultCode, datatype.Unsigned32(22))
	assert.Equal(t, ofa.ExperimentalResult.VendorId, datatype.Unsigned32(1))
	assert.Equal(t, ofa.ExperimentalResult.ExperimentalResultCode, datatype.Unsigned32(2))
	assert.Equal(t, ofa.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, ofa.OriginHost, datatype.DiameterIdentity("foobar"))
	assert.Equal(t, ofa.OriginRealm, datatype.DiameterIdentity("test"))
	assert.Equal(t, len(ofa.SupportedFeatures), 0)
	assert.Equal(t, ofa.SmDeliveryFailureCause.SmEnumeratedDeliveryFailureCause, datatype.Enumerated(12))
	assert.Nil(t, ofa.SmDeliveryFailureCause.SmDiagnosticInfo)
	assert.Nil(t, ofa.SmRpUi)
	assert.Equal(t, *ofa.ExternalIdentifier, datatype.UTF8String("external-identifier"))
	assert.Nil(t, ofa.FailedAvp)
	assert.Equal(t, len(ofa.ProxyInfo), 0)
	assert.Equal(t, len(ofa.RouteRecord), 0)
}

func TestOFA_Marshal_OK(t *testing.T) {
	ofa := createStructOFA()
	m1 := diam.NewMessage(diam.MOForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	err := m1.Marshal(ofa)
	assert.Nil(t, err)

	m2 := createDiamOFA()
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}

func createDiamOFA() *diam.Message {
	m2 := diam.NewMessage(diam.MOForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(1001))
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.SMDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.SMEnumeratedDeliveryFailureCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1)),
		},
	})
	return m2
}

func createStructOFA() *smparser.OFA {
	resultCode := datatype.Unsigned32(1001)
	ofa := smparser.OFA{
		SessionId:        datatype.UTF8String("session-id"),
		ResultCode:       &resultCode,
		AuthSessionState: datatype.Enumerated(1),
		OriginHost:       datatype.DiameterIdentity("orig-host"),
		OriginRealm:      datatype.DiameterIdentity("orig-realm"),
		SmDeliveryFailureCause: &basetype.SMDeliveryFailureCause{
			SmEnumeratedDeliveryFailureCause: datatype.Enumerated(1),
		},
	}
	return &ofa
}

func BenchmarkEncodeOFA(b *testing.B) {
	ofa := createStructOFA()
	m1 := diam.NewMessage(diam.MOForwardShortMessage, 0, diam.TGPP_SGD_GDD_APP_ID, 0, 0, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(ofa)
	}
}

func BenchmarkDecodeOFA(b *testing.B) {
	m := createDiamOFA()
	ofa := smparser.OFA{}
	for n := 0; n < b.N; n++ {
		ofa.Parse(m)
	}
}
