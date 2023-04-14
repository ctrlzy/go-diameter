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

func TestALR_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestALR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestALR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestALR_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.Nil(t, err)
}

func TestALR_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345678"))
	m.NewAVP(avp.SMRPMTI, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1))
	m.NewAVP(avp.SMRPSMEA, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("smea"))
	m.NewAVP(avp.SRRFlags, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(12))
	m.NewAVP(avp.SMDeliveryNotIntended, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1))
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, alr.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, alr.Drmp)
	assert.Nil(t, alr.VendorSpecificApplicationId)
	assert.Equal(t, alr.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, alr.OriginHost, datatype.DiameterIdentity("foobar"))
	assert.Equal(t, alr.OriginRealm, datatype.DiameterIdentity("bit"))
	assert.Nil(t, alr.DestinationHost)
	assert.Equal(t, alr.DestinationRealm, datatype.DiameterIdentity("dest-realm"))
	assert.Equal(t, *alr.Msisdn, datatype.OctetString("12345678"))
	assert.Nil(t, alr.UserName)
	assert.Nil(t, alr.SmsmiCorrelationId)
	assert.Empty(t, alr.SupportedFeatures)
	assert.Nil(t, alr.ScAddress)
	assert.Equal(t, *alr.SmRpMti, datatype.Enumerated(1))
	assert.Equal(t, *alr.SmRpSmea, datatype.OctetString("smea"))
	assert.Equal(t, *alr.SrrFlags, datatype.Unsigned32(12))
	assert.Equal(t, *alr.SmDeliveryNotIntended, datatype.Enumerated(1))
}
