// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package sm

import (
	"errors"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/avp"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smpeer"
)

// handleDWR handles Device-Watchdog-Request messages.
//
// If mandatory AVPs such as Origin-Host or Origin-Realm are missing,
// we ignore the message.
//
// See RFC 6733 section 5.5 for details.
func handlePUR(sm *StateMachine) diam.HandlerFunc {
	return func(c diam.Conn, m *diam.Message) {
		pur := new(smparser.PUR)
		err := pur.Parse(m)
		if err != nil {
			sm.Error(&diam.ErrorReport{
				Conn:    c,
				Message: m,
				Error:   err,
			})
			return
		}
		a := m.Answer(diam.Success)
		// Fix for Same H2H and E2E Identifier in success response
		a.Header.HopByHopID = m.Header.HopByHopID
		a.Header.EndToEndID = m.Header.EndToEndID
		a.NewAVP(avp.OriginHost, avp.Mbit, 0, sm.cfg.OriginHost)
		a.NewAVP(avp.OriginRealm, avp.Mbit, 0, sm.cfg.OriginRealm)
		if sm.cfg.OriginStateID != 0 {
			stateid := datatype.Unsigned32(sm.cfg.OriginStateID)
			m.NewAVP(avp.OriginStateID, avp.Mbit, 0, stateid)
		}
		_, err = a.WriteTo(c)
		if err != nil {
			sm.Error(&diam.ErrorReport{
				Conn:    c,
				Message: m,
				Error:   err,
			})
		}
	}
}

// Create & send PUR Request
func sendPUR(c diam.Conn, cfg *Settings, pur *smparser.PUR) error {
	meta, ok := smpeer.FromContext(c.Context())
	if !ok {
		return errors.New("peer metadata unavailable")
	}
	m := diam.NewRequest(diam.PushNotification, diam.TGPP_SH_APP_ID, dict.Default)
	m.NewAVP(avp.SessionID, avp.Mbit, 0, datatype.UTF8String(getSessionId(cfg.OriginHost)))
	//m.NewAVP(avp.DRMP,avp.Mbit,0,pur.DRMP)
	m.NewAVP(avp.VendorSpecificApplicationID, avp.Mbit, 0, pur.VendorSpecificApplicationId)
	m.NewAVP(avp.AuthSessionState, avp.Mbit, 0, pur.AuthSessionState)
	m.NewAVP(avp.OriginHost, avp.Mbit, 0, cfg.OriginHost)
	m.NewAVP(avp.OriginRealm, avp.Mbit, 0, cfg.OriginRealm)
	m.NewAVP(avp.DestinationRealm, avp.Mbit, 0, meta.OriginRealm)
	m.NewAVP(avp.DestinationHost, avp.Mbit, 0, meta.OriginHost)
	m.NewAVP(avp.SupportedFeatures, avp.Mbit, 0, pur.SupportedFeatures)
	m.NewAVP(avp.DataReference, avp.Mbit|avp.Vbit, 10415, pur.DataReference)
	m.NewAVP(avp.UserDataSh, avp.Mbit|avp.Vbit, 10415, pur.UserDataSh)

	//m.NewAVP(avp.OCSupportedFeatures, avp.Mbit, 0, pur.OCSupportedFeatures)
	//m.NewAVP(avp.ProxyInfo, avp.Mbit, 0, pur.ProxyInfo)
	//m.NewAVP(avp.RouteRecord, avp.Mbit, 0, pur.RouteRecord)
	_, err := m.WriteTo(c)
	return err
}
