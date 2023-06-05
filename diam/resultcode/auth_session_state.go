package resultcode

// IETF RFC 6733 [20]
// The Auth-Session-State AVP (AVP Code 277) is of type Enumerated and
// specifies whether state is maintained for a particular session.  The
// client MAY include this AVP in requests as a hint to the server, but
// the value in the server's answer message is binding.  The following
// values are supported:
type AuthSessionState int32

const (
	// This value is used to specify that session state is being
	// maintained, and the access device MUST issue a session termination
	// message when service to the user is terminated.  This is the
	// default value.
	AuthSessionState_STATE_MAINTAINED AuthSessionState = 0
	// This value is used to specify that no session termination messages
	// will be sent by the access device upon expiration of the
	// Authorization-Lifetime.
	AuthSessionState_NO_STATE_MAINTAINED AuthSessionState = 1
)

// Enum value maps for Result.
var (
	AuthSessionState_name = map[int32]string{
		0: "STATE_MAINTAINED",
		1: "NO_STATE_MAINTAINED",
	}
	AuthSessionState_value = map[string]int32{
		"STATE_MAINTAINED":    0,
		"NO_STATE_MAINTAINED": 1,
	}
)

func (a *AuthSessionState) Name() string {
	return AuthSessionState_name[int32(*a)]
}
