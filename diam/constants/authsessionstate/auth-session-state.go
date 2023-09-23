package authsessionstate

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// IETF RFC 6733 [20]
// The Auth-Session-State AVP (AVP Code 277) is of type Enumerated and
// specifies whether state is maintained for a particular session.  The
// client MAY include this AVP in requests as a hint to the server, but
// the value in the server's answer message is binding.  The following
// values are supported:

const (
	// This value is used to specify that session state is being
	// maintained, and the access device MUST issue a session termination
	// message when service to the user is terminated.  This is the
	// default value.
	STATE_MAINTAINED = datatype.Enumerated(0)
	// This value is used to specify that no session termination messages
	// will be sent by the access device upon expiration of the
	// Authorization-Lifetime.
	NO_STATE_MAINTAINED = datatype.Enumerated(1)
)
