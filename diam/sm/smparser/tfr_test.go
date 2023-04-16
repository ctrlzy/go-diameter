package smparser_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)

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
	tfr := smparser.TFR{}
	err := tfr.Parse(m)
	assert.Nil(t, err)
}

func TestTFR_PARSE_OK(t *testing.T) {
	ti := time.Now()
	m := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("user-name"))
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMRPUI, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-ui"))
	m.NewAVP(avp.MMENumberforMTSMS, avp.Vbit, 10415, datatype.OctetString("mme-num-for-mt-sms"))
	m.NewAVP(avp.TFRFlags, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(1))
	m.NewAVP(avp.SMDeliveryTimer, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(1))
	m.NewAVP(avp.SMDeliveryStartTime, avp.Mbit|avp.Vbit, 10415, datatype.Time(ti))
	m.NewAVP(avp.MaximumRetransmissionTime, avp.Vbit, 10415, datatype.Time(ti))
	tfr := smparser.TFR{}
	err := tfr.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, tfr.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, tfr.Drmp)
	assert.Nil(t, tfr.VendorSpecificApplicationId)
	assert.Equal(t, tfr.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, tfr.OriginHost, datatype.DiameterIdentity("orig-host"))
	assert.Equal(t, tfr.OriginRealm, datatype.DiameterIdentity("orig-realm"))
	assert.Equal(t, tfr.DestinationHost, datatype.DiameterIdentity("dest-host"))
	assert.Equal(t, tfr.DestinationRealm, datatype.DiameterIdentity("dest-realm"))
	assert.Equal(t, tfr.UserName, datatype.UTF8String("user-name"))
	assert.Empty(t, tfr.SupportedFeatures)
	assert.Nil(t, tfr.SmsmiCorrelationId)
	assert.Equal(t, tfr.ScAddress, datatype.OctetString("sc-addr"))
	assert.Equal(t, tfr.SmRpUi, datatype.OctetString("sm-rp-ui"))
	assert.Equal(t, *tfr.MmeNumberForMtSms, datatype.OctetString("mme-num-for-mt-sms"))
	assert.Nil(t, tfr.SgsnNumber)
	assert.Equal(t, *tfr.TfrFlags, datatype.Unsigned32(1))
	assert.Equal(t, *tfr.SmDeliveryTimer, datatype.Unsigned32(1))
	assert.Equal(t, *tfr.SmDeliveryStartTime, datatype.Time(ti))
	assert.Equal(t, *tfr.MaximumRetransmissionTime, datatype.Time(ti))
	assert.Nil(t, tfr.SmsGmscAddress)
	assert.Empty(t, tfr.ProxyInfo)
	assert.Empty(t, tfr.RouteRecord)
}

func TestTFR_Decode_OK(t *testing.T) {
	drmp := datatype.Enumerated(1)
	vsai := basetype.Vendor_Specific_Application_Id{
		AuthApplicationId: 123,
		AcctApplicationId: 456,
	}
	sgsnNumber := datatype.OctetString("sgsn-num")
	tfr := &smparser.TFR{
		SessionId:                   "session-id",
		Drmp:                        &drmp,
		VendorSpecificApplicationId: &vsai,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		DestinationHost:             datatype.DiameterIdentity("dest-host"),
		DestinationRealm:            datatype.DiameterIdentity("dest-realm"),
		UserName:                    datatype.UTF8String("user-name"),
		ScAddress:                   datatype.OctetString("sc-addr"),
		SmRpUi:                      datatype.OctetString("sm-rp-ui"),
		SgsnNumber:                  &sgsnNumber,
	}
	msg := diam.NewRequest(diam.MTForwardShortMessage, diam.TGPP_SGD_GDD_APP_ID, dict.Default)
	err := msg.Marshal(tfr)
	assert.Nil(t, err)
	fmt.Print(msg)
}