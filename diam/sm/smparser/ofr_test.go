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

func TestOFR_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id-adfa"))
	m.NewAVP(avp.DRMP, 0, 0, datatype.Enumerated(14))
	m.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(1001)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(1002)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.OFRFlags, avp.Vbit, 10415, datatype.Unsigned32(1))
	m.NewAVP(avp.SupportedFeatures, 0, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.FeatureListID, avp.Vbit, 10415, datatype.Unsigned32(456)),
			diam.NewAVP(avp.FeatureList, avp.Vbit, 10415, datatype.Unsigned32(789)),
		},
	})
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	m.NewAVP(avp.SMSMICorrelationID, avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.HSSID, avp.Vbit, 10415, datatype.UTF8String("hss-id")),
			diam.NewAVP(avp.OriginatingSIPURI, avp.Vbit, 10415, datatype.UTF8String("orig-sip-uri")),
			diam.NewAVP(avp.DestinationSIPURI, avp.Vbit, 10415, datatype.UTF8String("dest-sip-uri")),
		},
	})
	m.NewAVP(avp.SMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMESMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.SMDeliveryCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1)),
					diam.NewAVP(avp.AbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(23)),
				},
			}),
		},
	})
	ofr := smparser.OFR{}
	err := ofr.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, ofr.SessionId, datatype.UTF8String("session-id-adfa"))
	assert.Equal(t, *ofr.Drmp, datatype.Enumerated(14))
	assert.Nil(t, ofr.VendorSpecificApplicationId.VendorId)
	assert.Equal(t, ofr.VendorSpecificApplicationId.AuthApplicationId, datatype.Unsigned32(1002))
	assert.Equal(t, ofr.VendorSpecificApplicationId.AcctApplicationId, datatype.Unsigned32(1001))
	assert.Equal(t, ofr.AuthSessionState, datatype.Enumerated(1))
	assert.Equal(t, ofr.OriginHost, datatype.DiameterIdentity("orig-host"))
	assert.Equal(t, ofr.OriginRealm, datatype.DiameterIdentity("orig-realm"))
	assert.Equal(t, *ofr.DestinationHost, datatype.DiameterIdentity("dest-host"))
	assert.Equal(t, ofr.DestinationRealm, datatype.DiameterIdentity("dest-realm"))
	assert.Equal(t, ofr.ScAddress, datatype.OctetString("sc-addr"))
	assert.Equal(t, *ofr.OfrFlags, datatype.Unsigned32(1))
	assert.Equal(t, len(ofr.SupportedFeatures), 1)
	assert.Equal(t, ofr.SupportedFeatures[0].VendorId, datatype.Unsigned32(123))
	assert.Equal(t, ofr.SupportedFeatures[0].FeatureListId, datatype.Unsigned32(456))
	assert.Equal(t, ofr.SupportedFeatures[0].FeatureList, datatype.Unsigned32(789))
	assert.Equal(t, *ofr.UserIdentifier.Msisdn, datatype.OctetString("12345"))
	assert.Nil(t, ofr.UserIdentifier.ExternalIdentifier)
	assert.Nil(t, ofr.UserIdentifier.Lmsi)
	assert.Nil(t, ofr.UserIdentifier.UserName)
	assert.Equal(t, ofr.SmRpUi, datatype.OctetString("sm-rp-ui"))
	assert.Equal(t, ofr.SmsmiCorrelationId.HssId, datatype.UTF8String("hss-id"))
	assert.Equal(t, *ofr.SmsmiCorrelationId.OriginatingSipUri, datatype.UTF8String("orig-sip-uri"))
	assert.Equal(t, *ofr.SmsmiCorrelationId.DestinationSipUri, datatype.UTF8String("dest-sip-uri"))
	assert.Equal(t, *ofr.SmDeliveryOutcome.MmeSmDeliveryOutcome.SmDeliveryCause, datatype.Enumerated(1))
	assert.Equal(t, *ofr.SmDeliveryOutcome.MmeSmDeliveryOutcome.AbsentUserDiagnosticSm, datatype.Unsigned32(23))
	assert.Nil(t, ofr.SmDeliveryOutcome.MscSmDeliveryOutcome)
	assert.Nil(t, ofr.SmDeliveryOutcome.IpsmgwSmDeliveryOutcome)
	assert.Nil(t, ofr.SmDeliveryOutcome.SgsnSmDeliveryOutcome)
	assert.Equal(t, len(ofr.ProxyInfo), 0)
	assert.Equal(t, len(ofr.RouteRecord), 0)
}
