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

func TestSRR_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	srr := new(smparser.SRR)
	err := srr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestSRR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	srr := new(smparser.SRR)
	err := srr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestSRR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	srr := new(smparser.SRR)
	err := srr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestSRR_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345678"))
	m.NewAVP(avp.SMRPMTI, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1))
	m.NewAVP(avp.SMRPSMEA, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("smea"))
	m.NewAVP(avp.SRRFlags, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(12))
	m.NewAVP(avp.SMDeliveryNotIntended, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1))
	srr := new(smparser.SRR)
	err := srr.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, srr.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, srr.Drmp)
	assert.Nil(t, srr.VendorSpecificApplicationId)
	assert.Equal(t, srr.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, srr.OriginHost, datatype.DiameterIdentity("foobar"))
	assert.Equal(t, srr.OriginRealm, datatype.DiameterIdentity("bit"))
	assert.Nil(t, srr.DestinationHost)
	assert.Equal(t, srr.DestinationRealm, datatype.DiameterIdentity("dest-realm"))
	assert.Equal(t, *srr.Msisdn, datatype.OctetString("12345678"))
	assert.Nil(t, srr.UserName)
	assert.Nil(t, srr.SmsmiCorrelationId)
	assert.Empty(t, srr.SupportedFeatures)
	assert.Nil(t, srr.ScAddress)
	assert.Equal(t, *srr.SmRpMti, datatype.Enumerated(1))
	assert.Equal(t, *srr.SmRpSmea, datatype.OctetString("smea"))
	assert.Equal(t, *srr.SrrFlags, datatype.Unsigned32(12))
	assert.Equal(t, *srr.SmDeliveryNotIntended, datatype.Enumerated(1))
}

func TestSRR_Decode_OK(t *testing.T) {
	srr := createStructSRR()
	m1 := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	err := m1.Marshal(srr)
	assert.Nil(t, err)
	m2 := createDiamSRR()
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}

func createDiamSRR() *diam.Message {
	m2 := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m2.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m2.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345"))
	m2.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("username"))
	m2.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m2.NewAVP(avp.SMRPMTI, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(2))
	m2.NewAVP(avp.SMRPSMEA, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sm-rp-smea"))
	m2.NewAVP(avp.SMDeliveryNotIntended, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(0))
	return m2
}

func createStructSRR() *smparser.SRR {
	authId := datatype.Unsigned32(123)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          datatype.Unsigned32(10415),
		AuthApplicationId: &authId,
	}
	destHost := datatype.DiameterIdentity("dest-host")
	msisdn := datatype.OctetString("12345")
	userName := datatype.UTF8String("username")
	scAddr := datatype.OctetString("sc-addr")
	smrpmti := datatype.Enumerated(2)
	smrpsmea := datatype.OctetString("sm-rp-smea")
	smDeliveryNotIntended := datatype.Enumerated(0)
	srr := &smparser.SRR{
		SessionId:                   "session-id",
		VendorSpecificApplicationId: &vsai,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		DestinationHost:             &destHost,
		DestinationRealm:            datatype.DiameterIdentity("dest-realm"),
		Msisdn:                      &msisdn,
		UserName:                    &userName,
		ScAddress:                   &scAddr,
		SmRpMti:                     &smrpmti,
		SmRpSmea:                    &smrpsmea,
		SmDeliveryNotIntended:       &smDeliveryNotIntended,
	}
	return srr
}

func BenchmarkEncodeSRR(b *testing.B) {
	srr := createStructSRR()
	m1 := diam.NewRequest(diam.SendRoutingInfoforSM, diam.TGPP_S6C_APP_ID, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(srr)
	}
}

func BenchmarkDecodeSRR(b *testing.B) {
	m := createDiamSRR()
	srr := smparser.SRR{}
	for n := 0; n < b.N; n++ {
		srr.Parse(m)
	}
}
