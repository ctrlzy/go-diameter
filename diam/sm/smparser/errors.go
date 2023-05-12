// Copyright 2013-2015 go-diameter authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

package smparser

import (
	"errors"
	"fmt"

	"github.com/ctrlzy/go-diameter/v4/diam"
)

var (
	// ErrMissingResultCode is returned by Parse when
	// the message does nt contain a Result-Code AVP.
	ErrMissingResultCode = errors.New("missing Result-Code")

	// ErrMissingOriginHost is returned by Parse when
	// the message does not contain an Origin-Host AVP.
	ErrMissingOriginHost = errors.New("missing Origin-Host")

	// ErrMissingOriginRealm is returned by Parse when
	// the message does not contain an Origin-Realm AVP.
	ErrMissingOriginRealm = errors.New("missing Origin-Realm")

	// ErrMissingDestHost is returned by Parse when
	// the message does not contain an Destination-Host AVP.
	ErrMissingDestHost = errors.New("missing Destination-Host")

	// ErrMissingDestRealm is returned by Parse when
	// the message does not contain an Destination-Realm AVP.
	ErrMissingDestRealm = errors.New("missing Destination-Realm")

	// ErrMissingUserName is returned by Parse when
	// the message does not contain an User-Name AVP.
	ErrMissingUserName = errors.New("missing User-Name")

	// ErrMissingUserName is returned by Parse when
	// the message does not contain an User-Identity AVP.
	ErrMissingUserIdentity = errors.New("missing User-Identity")

	// ErrMissingUserName is returned by Parse when
	// the message does not contain an Server-Name AVP.
	ErrMissingServerName = errors.New("missing Server-Name")

	// ErrMissingUserName is returned by Parse when
	// the message does not contain an Data-Reference AVP.
	ErrMissingDataReference = errors.New("missing Data-Reference")

	// ErrMissingScAddress is returned by Parse when
	// the message does not contain an SC-Address AVP.
	ErrMissingScAddress = errors.New("missing SC-Address")

	// ErrMissingSmDeliveryOutcome is returned by Parse when
	// the message does not contain an SM-Delivery-Outcome AVP.
	ErrMissingSmDeliveryOutcome = errors.New("missing SM-Delivery-Outcome")

	// ErrMissingSmRpUi is returned by Parse when
	// the message does not contain an SM-RP-UI AVP.
	ErrMissingSmRpUi = errors.New("missing SM-RP-UI")

	// ErrMissingUserIdentifier is returned by Parse when
	// the message does not contain an User-Identifier AVP.
	ErrMissingUserIdentifier = errors.New("missing User-Identifier")

	// ErrMissingApplication is returned by Parse when
	// the CER does not contain any Acct-Application-Id or
	// Auth-Application-Id, or their embedded versions in
	// the Vendor-Specific-Application-Id AVP.
	ErrMissingApplication = errors.New("missing application")

	// ErrNoCommonSecurity is returned by Parse when
	// the CER contains the Inband-Security-Id.
	// We currently don't support that.
	ErrNoCommonSecurity = errors.New("no common security")

	// ErrNoCommonApplication is returned by Parse when the
	// application IDs in the CER don't match the applications
	// defined in our dictionary.
	ErrNoCommonApplication = errors.New("no common application")
)

// ErrUnexpectedAVP is returned by Parse when the code of the AVP passed
// as AcctApplicationID, AuthApplicationID or VendorSpecificApplicationID
// and its embedded AVPs do not match their names.
type ErrUnexpectedAVP struct {
	AVP *diam.AVP
}

// Error implements the error interface.
func (e *ErrUnexpectedAVP) Error() string {
	return fmt.Sprintf("unexpected AVP: %s", e.AVP)
}
