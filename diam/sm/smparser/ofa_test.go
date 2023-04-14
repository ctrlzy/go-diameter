package smparser_test

import (
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
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
