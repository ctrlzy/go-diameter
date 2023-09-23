package smrpmti

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS 29.338 section 5.3.3.2
// The SM-RP-MTI AVP is of type Enumerated and shall contain the RP-Message Type Indicator of the Short Message. The following values are defined:
// – SM_DELIVER (0)
// – SM_STATUS_REPORT (1)
type SmRpMti datatype.Enumerated

const (
	SM_DELIVER       = datatype.Enumerated(0)
	SM_STATUS_REPORT = datatype.Enumerated(1)
)
