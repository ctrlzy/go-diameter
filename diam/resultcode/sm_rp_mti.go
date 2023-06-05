package resultcode

// 3GPP TS 29.338 section 5.3.3.2
// The SM-RP-MTI AVP is of type Enumerated and shall contain the RP-Message Type Indicator of the Short Message. The following values are defined:
// – SM_DELIVER (0)
// – SM_STATUS_REPORT (1)
type SmRpMti int32

const (
	SM_DELIVER       SmRpMti = 0
	SM_STATUS_REPORT SmRpMti = 1
)
