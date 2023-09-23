package absentuserdiagnosticsm

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS 29.338 5.3.3.20
// The Absent-User-Diagnostic-SM AVP is of type Unsigned32 and shall indicate the diagnostic explaining the absence of the subscriber. The values are defined in 3GPP TS 23.040 [3] clause 3.3.2.
// 3GPP TS 23.040 3.3.2
/*
Table 1a: Assignment of values to reasons for absence
(values must be in the range of 0 to 255, see 3GPP TS 29.002 [15])
Values	Reason for absence
0		– no paging response via the MSC
1		– IMSI detached
2		– roaming restriction
3		– deregistered in the HLR for non GPRS
4		– MS purged for non GPRS
5		– no paging response via the SGSN
6		– GPRS detached
7		– deregistered in the HLR for GPRS
8		– MS purged for GPRS
9		– Unidentified subscriber via the MSC
10		– Unidentified subscriber via the SGSN
11		– deregistered in the HSS/HLR for IMS
12		– no response via the IP-SM-GW
13		- the MS is temporarily unavailable
All ‘non GPRS’ reasons (except for roaming restriction) can be combined with all ‘GPRS’ reasons and vice-versa
All other integer values are reserved.
*/
const (
	NO_PAGING_RESPONSE_VIA_MSC       = datatype.Unsigned32(0)
	IMSI_DETACHED                    = datatype.Unsigned32(1)
	ROAMING_RESTRICTION              = datatype.Unsigned32(2)
	DEREGISTERED_IN_HLR_FOR_NON_GPRS = datatype.Unsigned32(3)
	MS_PURGED_FOR_NON_GPRS           = datatype.Unsigned32(4)
	NO_PAGING_RESPONSE_VIA_SGSN      = datatype.Unsigned32(5)
	GPRS_DETACHED                    = datatype.Unsigned32(6)
	DEREGISTERED_IN_HLR_FOR_GPRS     = datatype.Unsigned32(7)
	MS_PURGED_FOR_GPRS               = datatype.Unsigned32(8)
	UNIDENTIFIED_SUBSCRIBER_VIA_MSC  = datatype.Unsigned32(9)
	UNIDENTIFIED_SUBSCRIBER_VIA_SGSN = datatype.Unsigned32(10)
	DEREGISTERED_IN_HSS_FOR_IMS      = datatype.Unsigned32(11)
	NO_RESPONSE_VIA_IP_SM_GW         = datatype.Unsigned32(12)
	MS_TEMPORARILY_UNAVAILABLE       = datatype.Unsigned32(13)
)
