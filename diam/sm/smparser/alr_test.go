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
	ofa := new(smparser.ALR)
	err := ofa.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestALR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	ofa := new(smparser.ALR)
	err := ofa.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestALR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	ofa := new(smparser.ALR)
	err := ofa.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestALR_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	ofa := new(smparser.ALR)
	err := ofa.Parse(m)
	assert.Nil(t, err)
}
