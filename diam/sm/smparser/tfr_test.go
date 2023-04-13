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

func TestTFR_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	//m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestTFR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	//m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestTFR_MissingDestinationHost(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	//m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingDestHost)
}

func TestTFR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	//m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestTFR_MissingUserName(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	//m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingUserName)
}

func TestTFR_MissingSCAddress(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	//m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingScAddress)
}

func TestTFR_MissingSMRPRUI(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	//m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingSmRpUi)
}

func TestTFR_OK(t *testing.T) {
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.TFR{}
	err := ofr.Parse(m)
	assert.Nil(t, err)
}
