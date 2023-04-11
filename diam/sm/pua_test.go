package sm_test

import (
	"testing"

	"github.com/ctrlzy/go-diameter/v4/diam"
	"github.com/ctrlzy/go-diameter/v4/diam/dict"
	"github.com/ctrlzy/go-diameter/v4/diam/sm/smparser"
	"github.com/stretchr/testify/assert"
)

func TestMarshalPUA(t *testing.T) {
	pua := &smparser.PUA{
		SessionID: "test-session-id",
		VendorSpecificApplicationId: smparser.VendorSpecificApplicationId{
			VendorId:          10405,
			AuthApplicationId: 123,
			AcctApplicationId: 456,
		},
		DRMP:             1,
		ResultCode:       1,
		AuthSessionState: 1,
		OriginHost:       "test-orig-host",
		OriginRealm:      "test-orig-realm",
	}

	msg := diam.NewRequest(diam.ProfileUpdate, diam.TGPP_SH_APP_ID, dict.Default)
	err := msg.Marshal(pua)
	assert.Nil(t, err)

}
