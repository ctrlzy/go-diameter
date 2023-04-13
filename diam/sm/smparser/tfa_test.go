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
