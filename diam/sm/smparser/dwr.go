// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/datatype"
)

// DWR is a Device-Watchdog-Request message.
// See RFC 6733 section 5.5.1 for details.
type DWR struct {
	OriginHost    datatype.DiameterIdentity `avp:"Origin-Host"`
	OriginRealm   datatype.DiameterIdentity `avp:"Origin-Realm"`
	OriginStateID *diam.AVP                 `avp:"Origin-State-Id"`
}

// Parse parses and validates the given message, and returns nil when
// all AVPs are ok.
func (dwr *DWR) Parse(m *diam.Message) error {
	err := m.Unmarshal(dwr)
	if err != nil {
		return nil
	}
	if err = dwr.sanityCheck(); err != nil {
		return err
	}
	return nil
}

// sanityCheck ensures all mandatory AVPs are present.
func (dwr *DWR) sanityCheck() error {
	if len(dwr.OriginHost) == 0 {
		return ErrMissingOriginHost
	}
	if len(dwr.OriginRealm) == 0 {
		return ErrMissingOriginRealm
	}
	return nil
}

func (r *DWR) String() string {
	result := "DWR { "
	if r != nil {
		result += fmt.Sprintf("OriginHost: %s, OriginRealm: %s", r.OriginHost, r.OriginRealm)
		if r.OriginStateID != nil {
			result += fmt.Sprintf(", OriginStateID: %v", r.OriginStateID.String())
		}
	} else {
		result += "nil"
	}
	result += " }"
	return result
}
