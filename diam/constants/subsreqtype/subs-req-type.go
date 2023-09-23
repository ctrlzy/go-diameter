package subsreqtype

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS 29.328 6.3.6
// The Subs-Req-Type AVP is of type Enumerated, and indicates the type of the subscription-to-notifications request. The following values are defined:
// Subscribe (0) 	This value is used by an AS to subscribe to notifications of changes in data.
// Unsubscribe (1)	This value is used by an AS to unsubscribe to notifications of changes in data.
const (
	SUBSCRIBE   = datatype.Enumerated(0)
	UNSUBSCRIBE = datatype.Enumerated(1)
)
