package smparser_test

import (
	"testing"
	"time"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/constants/authsessionstate"
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
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	alr := new(smparser.ALR)
	err := alr.Parse(m)
	assert.Nil(t, err)
}

func TestALR_PARSE_OK(t *testing.T) {
	ti := time.Now()
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("scaddress"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SMSMICorrelationID, avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.HSSID, avp.Vbit, 10415, datatype.UTF8String("hss-id")),
			diam.NewAVP(avp.OriginatingSIPURI, avp.Vbit, 10415, datatype.UTF8String("orig-sip-uri")),
			diam.NewAVP(avp.DestinationSIPURI, avp.Vbit, 10415, datatype.UTF8String("dest-sip-uri")),
		},
	})
	m.NewAVP(avp.MaximumUEAvailabilityTime, avp.Vbit, 10415, datatype.Time(ti))
	m.NewAVP(avp.SMSGMSCAlertEvent, avp.Vbit, 10415, datatype.Unsigned32(1))
	m.NewAVP(avp.ServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMEName, avp.Mbit|avp.Vbit, 10415, datatype.DiameterIdentity("mme-name")),
			diam.NewAVP(avp.MMERealm, avp.Vbit, 10415, datatype.DiameterIdentity("mme-realm")),
			diam.NewAVP(avp.MMENumberforMTSMS, avp.Vbit, 10415, datatype.OctetString("mme number")),
		},
	})
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
	assert.Equal(t, alr.ScAddress, datatype.OctetString("scaddress"))
	assert.Equal(t, *alr.UserIdentifier.Msisdn, datatype.OctetString("12345"))
	assert.Empty(t, alr.UserIdentifier.UserName)
	assert.Empty(t, alr.UserIdentifier.Lmsi)
	assert.Empty(t, alr.UserIdentifier.ExternalIdentifier)
	assert.Equal(t, *alr.SmsmiCorrelationId.HssId, datatype.UTF8String("hss-id"))
	assert.Equal(t, *alr.SmsmiCorrelationId.DestinationSipUri, datatype.UTF8String("dest-sip-uri"))
	assert.Equal(t, *alr.SmsmiCorrelationId.OriginatingSipUri, datatype.UTF8String("orig-sip-uri"))
	assert.Equal(t, *alr.MaximumUeAvailabilityTime, datatype.Time(ti))
	assert.Equal(t, *alr.SmsGmscAlertEvent, datatype.Unsigned32(1))
	assert.Empty(t, alr.ServingNode.IpsmgwName)
	assert.Empty(t, alr.ServingNode.IpsmgwNumber)
	assert.Empty(t, alr.ServingNode.IpsmgwRealm)
	assert.Empty(t, alr.ServingNode.MscNumber)
	assert.Empty(t, alr.ServingNode.SgsnName)
	assert.Empty(t, alr.ServingNode.SgsnNumber)
	assert.Empty(t, alr.ServingNode.SgsnRealm)
	assert.Equal(t, *alr.ServingNode.MmeName, datatype.DiameterIdentity("mme-name"))
	assert.Equal(t, *alr.ServingNode.MmeRealm, datatype.DiameterIdentity("mme-realm"))
	assert.Equal(t, *alr.ServingNode.MmeNumberForMtSms, datatype.OctetString("mme number"))
	assert.Empty(t, alr.SupportedFeatures)
	assert.Empty(t, alr.ProxyInfo)
	assert.Empty(t, alr.RouteRecord)
}

func TestALR_Decode_OK(t *testing.T) {
	alr := createStructALR()

	m1 := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	err := m1.Marshal(alr)
	assert.Nil(t, err)
	m2 := createDiamALR()
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}

func createDiamALR() *diam.Message {

	m2 := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.DRMP, 0, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, authsessionstate.NO_STATE_MAINTAINED)
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m2.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m2.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m2.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m2.NewAVP(avp.ServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMEName, avp.Mbit|avp.Vbit, 10415, datatype.DiameterIdentity("mme-name")),
			diam.NewAVP(avp.MMERealm, avp.Vbit, 10415, datatype.DiameterIdentity("mme-realm")),
			diam.NewAVP(avp.MMENumberforMTSMS, avp.Vbit, 10415, datatype.OctetString("mme number")),
		},
	})
	return m2
}

func createStructALR() *smparser.ALR {
	drmp := datatype.Enumerated(1)
	authId := datatype.Unsigned32(123)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          datatype.Unsigned32(10415),
		AuthApplicationId: &authId,
	}
	msisdn := datatype.OctetString("12345")
	userIdentifier := basetype.User_Identifier{
		Msisdn: &msisdn,
	}
	mmeName := datatype.DiameterIdentity("mme-name")
	mmeRealm := datatype.DiameterIdentity("mme-realm")
	mmeNumber := datatype.OctetString("mme number")
	servingNode := basetype.Serving_Node{
		MmeName:           &mmeName,
		MmeRealm:          &mmeRealm,
		MmeNumberForMtSms: &mmeNumber,
	}
	destHost := datatype.DiameterIdentity("dest-host")
	alr := &smparser.ALR{
		SessionId:                   "session-id",
		Drmp:                        &drmp,
		VendorSpecificApplicationId: &vsai,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		DestinationHost:             &destHost,
		DestinationRealm:            datatype.DiameterIdentity("dest-realm"),
		ScAddress:                   datatype.OctetString("sc-addr"),
		UserIdentifier:              userIdentifier,
		ServingNode:                 &servingNode,
	}
	return alr
}

func BenchmarkEncodeALR(b *testing.B) {
	alr := createStructALR()
	m1 := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(alr)
	}
}

func BenchmarkDecodeALR(b *testing.B) {
	m := createDiamALR()
	alr := smparser.ALR{}
	for n := 0; n < b.N; n++ {
		alr.Parse(m)
	}
}
