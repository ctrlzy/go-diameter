package onetimenotification

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

/*
3GPP TS 29.338 6.3.22
The One-Time-Notification AVP is of type Enumerated. If present it indicates that the sender requests to be notified only one time. The following values are defined:
ONE_TIME_NOTIFICATION_REQUESTED (0)
This AVP is only applicable to UE reachability for IP (25)
*/
const (
	ONE_TIME_NOTIFICATION_REQUESTED = datatype.Enumerated(0)
)
