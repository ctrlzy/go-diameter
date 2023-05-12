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

func TestRDR_MissingOriginHost(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginHost)
}

func TestRDR_MissingOriginRealm(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingOriginRealm)
}

func TestRDR_MissingDestinationRealm(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingDestRealm)
}

func TestRDR_MissingUserIdentifier(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingUserIdentifier)
}

func TestRDR_MissingSCAddress(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingScAddress)
}

func TestRDR_MissingSMDeliveryOutcome(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))

	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.ErrorIs(t, err, smparser.ErrMissingSmDeliveryOutcome)
}

func TestRDR_MandantoryFields(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMESMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.SMDeliveryCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1)),
					diam.NewAVP(avp.AbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(11)),
				},
			}),
		},
	})
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)
	assert.Nil(t, err)
}

func TestRDR_PARSE_OK(t *testing.T) {
	m := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m.NewAVP(avp.DRMP, 0, 0, datatype.Enumerated(1))
	m.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	})
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(2))
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("foobar"))
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("bar"))
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m.NewAVP(avp.SMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMESMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.SMDeliveryCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1)),
					diam.NewAVP(avp.AbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(11)),
				},
			}),
		},
	})
	m.NewAVP(avp.RDRFlags, avp.Vbit, 10415, datatype.Unsigned32(1))
	rdr := new(smparser.RDR)
	err := rdr.Parse(m)

	assert.Nil(t, err)
	assert.Equal(t, rdr.SessionId, datatype.UTF8String("session-id"))
	assert.Equal(t, *rdr.Drmp, datatype.Enumerated(1))
	assert.Nil(t, rdr.VendorSpecificApplicationId.VendorId)
	assert.Equal(t, rdr.VendorSpecificApplicationId.AuthApplicationId, datatype.Unsigned32(123))
	assert.Equal(t, rdr.VendorSpecificApplicationId.AcctApplicationId, datatype.Unsigned32(456))
	assert.Equal(t, rdr.AuthSessionState, datatype.Enumerated(2))
	assert.Equal(t, rdr.OriginHost, datatype.DiameterIdentity("foobar"))
	assert.Equal(t, rdr.OriginRealm, datatype.DiameterIdentity("bar"))
	assert.Equal(t, *rdr.DestinationHost, datatype.DiameterIdentity("dest-host"))
	assert.Equal(t, rdr.DestinationRealm, datatype.DiameterIdentity("dest-realm"))
	assert.Empty(t, rdr.SupportedFeatures)
	assert.Equal(t, *rdr.UserIdentifier.Msisdn, datatype.OctetString("12345"))
	assert.Nil(t, rdr.UserIdentifier.UserName)
	assert.Nil(t, rdr.UserIdentifier.Lmsi)
	assert.Nil(t, rdr.UserIdentifier.ExternalIdentifier)
	assert.Nil(t, rdr.SmsmiCorrelationId)
	assert.Equal(t, rdr.ScAddress, datatype.OctetString("sc-addr"))
	assert.Equal(t, *rdr.SmDeliveryOutcome.MmeSmDeliveryOutcome.SmDeliveryCause, datatype.Enumerated(1))
	assert.Equal(t, *rdr.SmDeliveryOutcome.MmeSmDeliveryOutcome.AbsentUserDiagnosticSm, datatype.Unsigned32(11))
	assert.Nil(t, rdr.SmDeliveryOutcome.IpsmgwSmDeliveryOutcome)
	assert.Nil(t, rdr.SmDeliveryOutcome.MscSmDeliveryOutcome)
	assert.Nil(t, rdr.SmDeliveryOutcome.SgsnSmDeliveryOutcome)
	assert.Equal(t, *rdr.RdrFlags, datatype.Unsigned32(1))
	assert.Empty(t, rdr.ProxyInfo)
	assert.Empty(t, rdr.RouteRecord)
}

func TestRDR_Decode_OK(t *testing.T) {
	vsai := basetype.Vendor_Specific_Application_Id{
		AuthApplicationId: 123,
		AcctApplicationId: 456,
	}
	destHost := datatype.DiameterIdentity("dest-host")
	msisdn := datatype.OctetString("12345")
	userIdentifier := basetype.User_Identifier{
		Msisdn: &msisdn,
	}
	deliverCause := datatype.Enumerated(1)
	absentUserDiagnosticSm := datatype.Unsigned32(11)
	smDeliveryOutcome := basetype.SM_Delivery_Outcome{
		MmeSmDeliveryOutcome: &basetype.Delivery_Outcome{
			SmDeliveryCause:        &deliverCause,
			AbsentUserDiagnosticSm: &absentUserDiagnosticSm,
		},
	}
	rdrFlags := datatype.Unsigned32(1)
	rdr := &smparser.RDR{
		SessionId:                   "session-id",
		VendorSpecificApplicationId: &vsai,
		AuthSessionState:            datatype.Enumerated(1),
		OriginHost:                  datatype.DiameterIdentity("orig-host"),
		OriginRealm:                 datatype.DiameterIdentity("orig-realm"),
		DestinationHost:             &destHost,
		DestinationRealm:            datatype.DiameterIdentity("dest-realm"),
		UserIdentifier:              userIdentifier,
		ScAddress:                   datatype.OctetString("sc-addr"),
		SmDeliveryOutcome:           smDeliveryOutcome,
		RdrFlags:                    &rdrFlags,
	}

	m1 := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	err := m1.Marshal(rdr)
	assert.Nil(t, err)

	m2 := diam.NewRequest(diam.ReportSMDeliveryStatus, diam.TGPP_S6C_APP_ID, dict.Default)
	m2.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String("session-id"))
	m2.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.AuthApplicationID, avp.Mbit, 0, datatype.Unsigned32(123)),
			diam.NewAVP(avp.AcctApplicationID, avp.Mbit, 0, datatype.Unsigned32(456)),
		},
	})
	m2.NewAVP(avp.AuthSessionState, avp.Mbit, 0, datatype.Enumerated(1))
	m2.NewAVP(avp.OriginHost, avp.Mbit, 0, datatype.DiameterIdentity("orig-host"))
	m2.NewAVP(avp.OriginRealm, avp.Mbit, 0, datatype.DiameterIdentity("orig-realm"))
	m2.NewAVP(avp.DestinationHost, avp.Mbit, 0, datatype.DiameterIdentity("dest-host"))
	m2.NewAVP(avp.DestinationRealm, avp.Mbit, 0, datatype.DiameterIdentity("dest-realm"))
	m2.NewAVP(avp.UserIdentifier, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MSISDN, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("12345")),
		},
	})
	m2.NewAVP(avp.SCAddress, avp.Mbit|avp.Vbit, 10415, datatype.OctetString("sc-addr"))
	m2.NewAVP(avp.SMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
		AVP: []*diam.AVP{
			diam.NewAVP(avp.MMESMDeliveryOutcome, avp.Mbit|avp.Vbit, 10415, &diam.GroupedAVP{
				AVP: []*diam.AVP{
					diam.NewAVP(avp.SMDeliveryCause, avp.Mbit|avp.Vbit, 10415, datatype.Enumerated(1)),
					diam.NewAVP(avp.AbsentUserDiagnosticSM, avp.Mbit|avp.Vbit, 10415, datatype.Unsigned32(11)),
				},
			}),
		},
	})
	m2.NewAVP(avp.RDRFlags, avp.Vbit, 10415, datatype.Unsigned32(1))
	m2.Header.HopByHopID = m1.Header.HopByHopID
	m2.Header.EndToEndID = m1.Header.EndToEndID

	assert.Equal(t, m1.String(), m2.String())
}
