package requesteddomain

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

// 3GPP TS29.328 6.3.7
// The Requested-Domain AVP is of type Enumerated, and indicates the access domain for which certain data (e.g. user state) are requested. The following values are defined:
// CS-Domain (0)	The requested data apply to the CS domain.
// PS-Domain (1)	The requested data apply to the PS domain.
const (
	CS_DOMAIN = datatype.Enumerated(0)
	PS_DOMAIN = datatype.Enumerated(1)
)
