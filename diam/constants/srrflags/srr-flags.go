package srrflags

// 3GPP TS 29.338 section 5.3.3.4
// The SRR-Flags AVP is of type Unsigned32 and it shall contain a bit mask. The meaning of the bits shall be as defined in table 5.3.3.4./1:
// 0 GPRS-Indicator This bit shall be set if the SMS-GMSC supports receiving of two serving nodes addresses from the HSS.
// 1 SM-RP-PRI This bit shall be set if the delivery of the short message shall be attempted when a service centre address is already contained in the Message Waiting Data file
// 2 Single-Attempt-Delivery This bit if set indicates that only one delivery attempt shall be performed for this particular SM.

const (
	GPRSIndicator         = 1
	SmRpPri               = 1 << 1
	SingleAttemptDelivery = 1 << 2
)
