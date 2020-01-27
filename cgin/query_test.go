package cgin

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testCases := []struct {
		desc            string
		inputQueryKey   string
		inputQueryValue string
		inputDefault    int
		expected        int
	}{
		{
			desc:            "",
			inputQueryKey:   "q",
			inputQueryValue: "1",
			inputDefault:    100,
			expected:        1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "", nil)
			req.URL.RawQuery = fmt.Sprintf("%s=%s", tC.inputQueryKey, tC.inputQueryValue)
			ctx := gin.Context{
				Request: req,
			}
			real := DefaultQueryAsInt(&ctx, tC.inputQueryKey, tC.inputDefault)
			assert.Equal(t, tC.expected, real)
		})
	}
}
