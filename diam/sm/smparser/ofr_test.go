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

func TestOFR_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	ofr := smparser.OFR{}
	err := ofr.Parse(m)
	assert.NotNil(t, err)
}

func TestOFR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)
	assert.NotNil(t, err)
}

func TestOFR_OK(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("test"))
	m.NewAVP(avp.OriginStateID, avp.Mbit, 0, datatype.Unsigned32(1))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)
	assert.Nil(t, err)
}
