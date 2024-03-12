package smparser_test

import (
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/basetype"
	"github.com/ctrlzy/go-diameter/v4/diam/constants/authsessionstate"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/stretchr/testify/assert"
)

func TestALA_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestALA_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestALA_OK(t *testing.T) {
	m := diam.NewRequest(diam.AlertServiceCenter, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))
	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.Nil(t, err)
}

func TestALA_PARSE_OK(t *testing.T) {
	m := diam.NewMessage(diam.AlertServiceCenter, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bit"))

	ala := new(smparser.ALA)
	err := ala.Parse(m)
	assert.Nil(t, err)
	assert.Equal(t, ala.SessionId, datatype.UTF8String("session-id"))
	assert.Nil(t, ala.Drmp)
	assert.Nil(t, ala.VendorSpecificApplicationId)
	assert.Equal(t, ala.AuthSessionState, datatype.Enumerated(1))
	assert.Nil(t, ala.ResultCode)
	assert.Nil(t, ala.ExperimentalResult)
	assert.Empty(t, ala.SupportedFeatures)
	assert.Nil(t, ala.FailedAvp)
	assert.Empty(t, ala.ProxyInfo)
	assert.Empty(t, ala.RouteRecord)
}

func TestALA_Decode_OK(t *testing.T) {
	ala := createStructALA()

	m1 := diam.NewMessage(diam.AlertServiceCenter, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	err := m1.Marshal(ala)
	assert.Nil(t, err)
	m2 := createDiamALA()

	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}

// create a diam.message ALA
func createDiamALA() *diam.Message {
	m2 := diam.NewMessage(diam.AlertServiceCenter, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	})
	m2.NewAVP(avp.ResultCode, avp.Mbit, 0, datatype.Unsigned32(10))
	m2.NewAVP(avp.ExperimentalResult, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.VendorID, avp.Mbit, 0, datatype.Unsigned32(10415)),
			diam.NewAVP(avp.ExperimentalResultCode, avp.Mbit, 0, datatype.Unsigned32(20)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	return m2
}

func createStructALA() *smparser.ALA {
	acctId := datatype.Unsigned32(456)
	vsai := basetype.Vendor_Specific_Application_Id{
		VendorId:          datatype.Unsigned32(10415),
		AcctApplicationId: &acctId,
	}
	resultCode := datatype.Unsigned32(10)
	experimentalResult := basetype.Experimental_Result{
		VendorId:               datatype.Unsigned32(10415),
		ExperimentalResultCode: datatype.Unsigned32(20),
	}

	ala := &smparser.ALA{
		SessionId:                   "session-id",
		VendorSpecificApplicationId: &vsai,
		ResultCode:                  &resultCode,
		ExperimentalResult:          &experimentalResult,
		AuthSessionState:            authsessionstate.NO_STATE_MAINTAINED,
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
	}
	return ala
}

func BenchmarkEncodeALA(b *testing.B) {
	ala := createStructALA()
	m1 := diam.NewMessage(diam.AlertServiceCenter, 0, diam.TGPP_S6C_APP_ID, 0, 0, dict.Default)
	for n := 0; n < b.N; n++ {
		m1.Marshal(ala)
	}
}

func BenchmarkDecodeALA(b *testing.B) {
	m := createDiamALA()
	ala := smparser.ALA{}
	for n := 0; n < b.N; n++ {
		ala.Parse(m)
	}
}
