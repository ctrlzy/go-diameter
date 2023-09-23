package rdrflags

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS 29.338
// The RDR-Flags AVP is of type Unsigned32 and it shall contain a bit mask. The meaning of the bits shall be as defined in table 5.3.3.21/1:
// Table 5.3.3.21/1: RDR-Flags
// Bit 	Name 						Description
// 0	Single-Attempt-Delivery		This bit if set indicates that only one delivery attempt shall be performed for this particular SM.
// NOTE 1: Bits not defined in this table shall be cleared by the sending entity and discarded by the receiving entity.
const (
	SINGLE_ATTEMPT_DELIVERY = datatype.Unsigned32(0)
)
