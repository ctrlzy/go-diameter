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
	//m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestOFR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	//m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestOFR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	//m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestOFR_MissingSCAddress(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	//m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingScAddress)
}

func TestOFR_MissingUserIdentifier(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingUserIdentifier)
}

func TestOFR_MissingSMRPRUI(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	//m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)

	assert.ErrorIs(t, err, smparser.ErrMissingSmRpUi)
}

func TestOFR_OK(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	ofr := smparser.OFR{}
	err := ofr.Parse(m)
	assert.Nil(t, err)
}
