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

func TestALA_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestALA_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestALA_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.Nil(t, err)
}

func TestALA_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))

	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, ala.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, ala.Drmp)
	assert.Nil(t, ala.VendorSpecificApplicationId)
	assert.Equal(t, ala.AuthSessionState, datatype.Enumerated(1))
	assert.Nil(t, ala.ResultCode)
	assert.Nil(t, ala.ExperimentalResult)
	assert.Empty(t, ala.SupportedFeatures)
	assert.Nil(t, ala.FailedAvp)
	assert.Empty(t, ala.ProxyInfo)
	assert.Empty(t, ala.RouteRecord)
}
