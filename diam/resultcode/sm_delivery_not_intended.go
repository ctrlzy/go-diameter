package resultcode

// 3GPP TS 29.338 section 5.3.3.5
// The SM-Delivery-Not-Intended AVP is of type Enumerated and shall indicate by its presence that delivery of a short message is not intended. It further indicates whether only IMSI or only MCC+MNC with the following values:
// – ONLY_IMSI_REQUESTED (0),
// – ONLY_MCC_MNC_REQUESTED (1).
type SmDeliveryNotIntended int32

const (
	ONLY_IMSI_REQUESTED    SmDeliveryNotIntended = 0
	ONLY_MCC_MNC_REQUESTED SmDeliveryNotIntended = 1
)
