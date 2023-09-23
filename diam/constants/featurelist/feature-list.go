package featurelist

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS 29.338 section 5.3.3.13
// Feature-List AVP
// The syntax of this AVP is defined in 3GPP TS 29.229 [5]. A null value indicates that there is no feature used by the application.
// For the S6c application, the meaning of the bits shall be as defined in table 5.3.3.13/1 for the Feature-List-ID 1.
// Table 5.3.3.13/1: Features of Feature-List-ID 1 used in S6c

// Feature bit Feature M/O Description
// 0 SMSF-Support O SMSF-Support This feature is applicable for the SRR/SRA command pair.
// If the SMS-GMSC or IP-SM-GW or SMS-Router does not support this feature,
// the HSS shall not return SMSF related AVPs (SMSF-3GPP-Address, SMSF-Non-3GPP-Address, SMSF-3GPP-Absent-User-Diagnostic-SM, SMSF-Non-3GPP-Absent-User-Diagnostic-SM) in SRA,
// and when the UE is known not to be reachable for SMS via MSC/MME and/or SGSN,
// the HSS may populate AVPs within the Serving-Node AVP and within the Additional-Serving-Node AVP with available SMSF address information.
const (
	SMSFSupport = datatype.Unsigned32(1)
)
