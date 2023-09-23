package smdeliverycause

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// TS 3GPP 29.338 5.3.3.19 SM-Delivery-Cause
// SM-Delivery-Cause
// The SM-Delivery-Cause AVP is of type Enumerated and shall indicate the cause of the SMP delivery result. The following values are defined:
// – UE_MEMORY_CAPACITY_EXCEEDED (0)
// – ABSENT_USER (1)
// – SUCCESSFUL_TRANSFER (2)

const (
	UE_MEMORY_CAPACITY_EXCEEDED = datatype.Enumerated(0)
	ABSENT_USER                 = datatype.Enumerated(1)
	SUCCESSFUL_TRANSFER         = datatype.Enumerated(2)
)
