package resultcode

// TS 3GPP 29.338 5.3.3.19 SM-Delivery-Cause
// SM-Delivery-Cause
// The SM-Delivery-Cause AVP is of type Enumerated and shall indicate the cause of the SMP delivery result. The following values are defined:
// – UE_MEMORY_CAPACITY_EXCEEDED (0)
// – ABSENT_USER (1)
// – SUCCESSFUL_TRANSFER (2)
type SMDeliveryCause int32

const (
	UE_MEMORY_CAPACITY_EXCEEDED SMDeliveryCause = 0
	ABSENT_USER                 SMDeliveryCause = 1
	SUCCESSFUL_TRANSFER         SMDeliveryCause = 2
)
