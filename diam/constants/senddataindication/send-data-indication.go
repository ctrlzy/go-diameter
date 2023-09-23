package senddataindication

import "github.com/ctrlzy/go-diameter/v4/diam/datatype"

/*
3GPP TS 29.328 6.3.17
The Send-Data-Indication AVP is of type Enumerated. If present it indicates that the sender requests the User-Data. The following values are defined:
USER_DATA_NOT_REQUESTED (0)
USER_DATA_REQUESTED (1)
*/
const (
	USER_DATA_NOT_REQUESTED = datatype.Enumerated(0)
	USER_DATA_REQUESTED     = datatype.Enumerated(1)
)
