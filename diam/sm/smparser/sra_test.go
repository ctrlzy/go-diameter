package smparser_test

import (
	"bytes"
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/stretchr/testify/assert"
)

func TestSRA_MissingOriginHost(t *testing.T) {
	m := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	sra := new(smparser.SRA)
	err := sra.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestSRA_MissingOriginRealm(t *testing.T) {
	m := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	sra := new(smparser.SRA)
	err := sra.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestSRA_PARSE_OK(t *testing.T) {
	m := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(11))
	m.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(12)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(0))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	m.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("username"))
	m.NewAVP(avp.ServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMEName, avp.Mbit|avp.Vbit, 10415, datatype.DiameterIdentity("mme-name")),
			diam.NewAVP(avp.MMERealm, avp.Mbit|avp.Vbit, 0, datatype.DiameterIdentity("mme-realm")),
			diam.NewAVP(avp.MMENumberforMTSMS, avp.Mbit|avp.Vbit, 0, datatype.OctetString("mme number")),
		},
	})
	m.NewAVP(avp.AdditionalServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSCNumber, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("msc number")),
		},
	})
	m.NewAVP(avp.MWDStatus, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(4))
	m.NewAVP(avp.MMEAbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(1))

	sra := new(smparser.SRA)
	err := sra.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, sra.SessionId, datatype.UTF8String("session-id"))
	assert.Empty(t, sra.Drmp)
	assert.Empty(t, sra.VendorSpecificApplicationId)
	assert.Equal(t, *sra.ResultCode, datatype.Unsigned32(11))
	assert.Equal(t, sra.ExperimentalResult.VendorId, datatype.Unsigned32(10415))
	assert.Equal(t, sra.ExperimentalResult.ExperimentalResultCode, datatype.Unsigned32(12))
	assert.Equal(t, sra.AuthSessionState, datatype.Enumerated(0))
	assert.Equal(t, sra.OriginHost, datatype.DiameterIdentity("foobar"))
	assert.Equal(t, sra.OriginRealm, datatype.DiameterIdentity("bit"))
	assert.Equal(t, *sra.UserName, datatype.UTF8String("username"))
	assert.Empty(t, sra.SupportedFeatures)
	assert.Equal(t, *sra.ServingNode.MmeName, datatype.DiameterIdentity("mme-name"))
	assert.Equal(t, *sra.ServingNode.MmeRealm, datatype.DiameterIdentity("mme-realm"))
	assert.Equal(t, *sra.ServingNode.MmeNumberForMtSms, datatype.OctetString("mme number"))
	assert.Nil(t, sra.ServingNode.IpsmgwName)
	assert.Nil(t, sra.ServingNode.IpsmgwRealm)
	assert.Nil(t, sra.ServingNode.IpsmgwNumber)
	assert.Nil(t, sra.ServingNode.MscNumber)
	assert.Nil(t, sra.ServingNode.SgsnName)
	assert.Nil(t, sra.ServingNode.SgsnRealm)
	assert.Nil(t, sra.ServingNode.SgsnNumber)
	assert.Equal(t, *sra.AdditionalServingNode.MscNumber, datatype.OctetString("msc number"))
	assert.Nil(t, sra.AdditionalServingNode.MmeName)
	assert.Nil(t, sra.AdditionalServingNode.MmeRealm)
	assert.Nil(t, sra.AdditionalServingNode.MmeNumberForMtSms)
	assert.Nil(t, sra.AdditionalServingNode.SgsnName)
	assert.Nil(t, sra.AdditionalServingNode.SgsnRealm)
	assert.Nil(t, sra.AdditionalServingNode.SgsnNumber)
	assert.Nil(t, sra.Smsf3gppAddress)
	assert.Nil(t, sra.SmsfNon3gppAddress)
	assert.Nil(t, sra.Lmsi)
	assert.Nil(t, sra.UserIdentifier)
	assert.Equal(t, *sra.MwdStatus, datatype.Unsigned32(4))
	assert.Equal(t, *sra.MmeAbsentUserDiagnosticSm, datatype.Unsigned32(1))
	assert.Nil(t, sra.MscAbsentUserDiagnosticSm)
	assert.Nil(t, sra.SgsnAbsentUserDiagnosticSm)
	assert.Nil(t, sra.Smsf3gppAbsentUserDiagnosticSm)
	assert.Nil(t, sra.SmsfNon3gppAbsentUserDiagnosticSm)
	assert.Nil(t, sra.FailedAvp)
	assert.Empty(t, sra.ProxyInfo)
	assert.Empty(t, sra.RouteRecord)
}

func TestSRA_Decode_OK(t *testing.T) {
	sra := createStructSRA()
	m1 := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	err := m1.Marshal(sra)
	assert.Nil(t, err)
	m2 := createDiamSRA()
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}

func createDiamSRA() *diam.Message {
	m2 := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
		},
	})
	m2.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(12))
	m2.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(11)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.UserName, avp.Mbit, 0, datatype.UTF8String("username"))
	m2.NewAVP(avp.ServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMEName, avp.Mbit|avp.Vbit, 10415, datatype.DiameterIdentity("mme-name")),
			diam.NewAVP(avp.MMERealm, avp.Mbit|avp.Vbit, 10415, datatype.DiameterIdentity("mme-realm")),
			diam.NewAVP(avp.MMENumberforMTSMS, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("mme number")),
		},
	})
	m2.NewAVP(avp.AdditionalServingNode, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSCNumber, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("msc number")),
		},
	})
	m2.NewAVP(avp.MWDStatus, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(12))
	m2.NewAVP(avp.MSCAbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(13))
	return m2
}

func createStructSRA() *smparser.SRA {
	authId := datatype.Unsigned32(123)
	vsai := basetype.VendorSpecificApplicationId{
		VendorId:          datatype.Unsigned32(10415),
		AuthApplicationId: &authId,
	}
	resultCode := datatype.Unsigned32(12)
	experimentResult := basetype.ExperimentalResult{
		VendorId:               datatype.Unsigned32(10415),
		ExperimentalResultCode: datatype.Unsigned32(11),
	}
	userName := datatype.UTF8String("username")
	mmeName := datatype.DiameterIdentity("mme-name")
	mmeRealm := datatype.DiameterIdentity("mme-realm")
	mmeNumber := datatype.OctetString("mme number")
	mscNumber := datatype.OctetString("msc number")
	servingNode := basetype.ServingNode{
		MmeName:           &mmeName,
		MmeRealm:          &mmeRealm,
		MmeNumberForMtSms: &mmeNumber,
	}
	additionalServingNode := basetype.AdditionalServingNode{
		MscNumber: &mscNumber,
	}
	mwdStatus := datatype.Unsigned32(12)
	mscAbsentUserDiagnosticSm := datatype.Unsigned32(13)
	sra := &smparser.SRA{
		SessionId:                   "session-id",
		VendorSpecificApplicationId: &vsai,
		ResultCode:                  &resultCode,
		ExperimentalResult:          &experimentResult,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		UserName:                    &userName,
		ServingNode:                 &servingNode,
		AdditionalServingNode:       &additionalServingNode,
		MwdStatus:                   &mwdStatus,
		MscAbsentUserDiagnosticSm:   &mscAbsentUserDiagnosticSm,
	}
	return sra
}

func BenchmarkEncodeSRA(b *testing.B) {
	sra := createStructSRA()
	m1 := diam.NewMessage(diam.SendRoutingInfoforSM, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(sra)
	}
}

func BenchmarkDecodeSRA(b *testing.B) {
	m := createDiamSRA()
	sra := smparser.SRA{}
	for n := 0; n < b.N; n++ {
		sra.Parse(m)
	}
}

func TestUmmarshallSRA(t *testing.T) {
	sraData := []byte{
		0x01, 0x00, 0x01, 0x70, 0x40, 0x80, 0x00, 0x27, 0x01, 0x00, 0x00, 0x60, 0x21, 0x63, 0x29, 0xad, 0x34, 0x04, 0xf7, 0x26, 0x00, 0x00, 0x01, 0x07, 0x40, 0x00, 0x00, 0x2c, 0x36, 0x30, 0x66, 0x63, 0x30, 0x38, 0x36, 0x61, 0x2d, 0x31, 0x62, 0x63, 0x63, 0x2d, 0x34, 0x31, 0x33, 0x61, 0x2d, 0x38, 0x35, 0x30, 0x38, 0x2d, 0x32, 0x39, 0x63, 0x37, 0x64, 0x35, 0x63, 0x33, 0x38, 0x61, 0x62, 0x32, 0x00, 0x00, 0x01, 0x15, 0x40, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x00, 0x01, 0x00, 0x00, 0x09, 0x61, 0xc0, 0x00, 0x00, 0x90, 0x00, 0x00, 0x28, 0xaf, 0x00, 0x00, 0x09, 0x62, 0xc0, 0x00, 0x00, 0x3d, 0x00, 0x00, 0x28, 0xaf, 0x75, 0x73, 0x63, 0x65, 0x6c, 0x6c, 0x2d, 0x6d, 0x6d, 0x65, 0x2d, 0x36, 0x30, 0x33, 0x36, 0x2e, 0x65, 0x70, 0x63, 0x2e, 0x6d, 0x6e, 0x63, 0x35, 0x38, 0x38, 0x2e, 0x6d, 0x63, 0x63, 0x33, 0x31, 0x31, 0x2e, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x09, 0x68, 0xc0, 0x00, 0x00, 0x2d, 0x00, 0x00, 0x28, 0xaf, 0x65, 0x70, 0x63, 0x2e, 0x6d, 0x6e, 0x63, 0x35, 0x38, 0x38, 0x2e, 0x6d, 0x63, 0x63, 0x33, 0x31, 0x31, 0x2e, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x06, 0x6d, 0xc0, 0x00, 0x00, 0x12, 0x00, 0x00, 0x28, 0xaf, 0x81, 0x56, 0x12, 0x36, 0x59, 0xf9, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x40, 0x00, 0x00, 0x17, 0x34, 0x35, 0x34, 0x33, 0x35, 0x31, 0x30, 0x30, 0x30, 0x32, 0x32, 0x39, 0x39, 0x32, 0x36, 0x00, 0x00, 0x00, 0x0c, 0xf0, 0xc0, 0x00, 0x00, 0x10, 0x00, 0x00, 0x28, 0xaf, 0x00, 0x00, 0x00, 0x09, 0x00, 0x00, 0x01, 0x0c, 0x40, 0x00, 0x00, 0x0c, 0x00, 0x00, 0x07, 0xd1, 0x00, 0x00, 0x01, 0x28, 0x40, 0x00, 0x00, 0x29, 0x65, 0x70, 0x63, 0x2e, 0x6d, 0x6e, 0x63, 0x30, 0x33, 0x35, 0x2e, 0x6d, 0x63, 0x63, 0x34, 0x35, 0x34, 0x2e, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01, 0x08, 0x40, 0x00, 0x00, 0x33, 0x6c, 0x61, 0x62, 0x31, 0x2e, 0x68, 0x73, 0x73, 0x31, 0x2e, 0x65, 0x70, 0x63, 0x2e, 0x6d, 0x6e, 0x63, 0x30, 0x33, 0x35, 0x2e, 0x6d, 0x63, 0x63, 0x34, 0x35, 0x34, 0x2e, 0x33, 0x67, 0x70, 0x70, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x6f, 0x72, 0x67, 0x00,
	}
	msg, err := diam.ReadMessage(bytes.NewReader(sraData), dict.Default)
	if err != nil {
		t.Fatal(err)
	}
	sra := &smparser.SRA{}
	err = sra.Parse(msg)
	if err != nil {
		t.Fatal(err)
	}
}
