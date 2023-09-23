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
	assert.Equal(t, *ofr.SmsmiCorrelationId.HssId, datatype.UTF8String("hss-id"))
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

func TestOFR_Marshal_OK(t *testing.T) {
	ofr := createStructOFR()
	m1 := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	err := m1.Marshal(ofr)
	assert.Nil(t, err)

	m2 := createDiamOFR()
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID
	assert.Equal(t, m1.String(), m2.String())
}

func createDiamOFR() *diam.Message {
	m2 := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session_id"))
	m2.NewAVP(avp.DRMP, 0, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_SGD_GDD_APP_ID)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(diam.TGPP_SGD_GDD_APP_ID)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m2.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m2.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("1238888888"))
	m2.NewAVP(avp.SupportedFeatures, 0, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.FeatureListID, avp.Vbit, 10415, datatype.Unsigned32(1)),
			diam.NewAVP(avp.FeatureList, avp.Vbit, 10415, datatype.Unsigned32(2)),
		},
	})
	m2.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345678")),
		},
	})
	m2.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	return m2
}

func createStructOFR() *smparser.OFR {
	drmp := datatype.Enumerated(1)
	vendorId := datatype.Unsigned32(10415)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          &vendorId,
		AuthApplicationId: datatype.Unsigned32(diam.TGPP_SGD_GDD_APP_ID),
		AcctApplicationId: datatype.Unsigned32(diam.TGPP_SGD_GDD_APP_ID),
	}
	destHost := datatype.DiameterIdentity("dest-host")
	supportedFeatures := basetype.Supported_Features{
		VendorId:      datatype.Unsigned32(10415),
		FeatureListId: datatype.Unsigned32(1),
		FeatureList:   datatype.Unsigned32(2),
	}
	msisdn := datatype.OctetString("12345678")
	userIdentifier := basetype.User_Identifier{
		Msisdn: &msisdn,
	}
	ofr := smparser.OFR{
		SessionId:                   datatype.UTF8String("session_id"),
		Drmp:                        &drmp,
		VendorSpecificApplicationId: &vsai,
		AuthSessionState:            datatype.Enumerated(0),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		DestinationHost:             &destHost,
		DestinationRealm:            datatype.DiameterIdentity("dest-realm"),
		ScAddress:                   datatype.OctetString("1238888888"),
		SupportedFeatures:           []basetype.Supported_Features{supportedFeatures},
		UserIdentifier:              userIdentifier,
		SmRpUi:                      datatype.OctetString("sm-rp-ui"),
	}
	return &ofr
}

func BenchmarkEncodeOFR(b *testing.B) {
	ofr := createStructOFR()
	m1 := diam.NewRequest(diam.MOForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(ofr)
	}
}

func BenchmarkDecodeOFR(b *testing.B) {
	m := createDiamOFR()
	ofr := smparser.OFR{}
	for n := 0; n < b.N; n++ {
		ofr.Parse(m)
	}
}
