package identityset

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

/*
3GPP TS 29.328 6.3.10
The Identity-Set AVP is of type Enumerated and indicates the requested set of IMS Public Identities. The following values are defined:
ALL_IDENTITIES (0)
REGISTERED_IDENTITIES (1)
IMPLICIT_IDENTITIES (2)
ALIAS_IDENTITIES (3)
*/
const (
	ALL_IDENTITIES        = datatype.Enumerated(0)
	REGISTERED_IDENTITIES = datatype.Enumerated(1)
	IMPLICIT_IDENTITIES   = datatype.Enumerated(2)
	ALIAS_IDENTITIES      = datatype.Enumerated(3)
)
