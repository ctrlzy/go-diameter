// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package sm

import (
	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
)

var puaACK = struct{}{}

// handlePUA handles Push-Notification-Answer messages.
func handlePUA(sm *StateMachine, puac chan struct{}) diam.HandlerFunc {
	return func(c diam.Conn, m *diam.Message) {
		pua := new(smparser.PUA)
		if err := pua.Parse(m); err != nil {
			sm.Error(&diam.ErrorReport{
				Conn:    c,
				Message: m,
				Error:   err,
			})
			return
		}
		if pua.ResultCode != diam.Success {
			return
		}
		select {
		case puac <- puaACK:
		default:
		}
	}
}
