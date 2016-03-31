package example

import (
	"github.com/appleboy/gofight"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestBasicHelloWorld(t *testing.T) {
	r := gofight.New()
	version := "0.0.1"

	r.GET("/").
		// trun on the debug mode.
		SetDebug(true).
		SetHeader(gofight.H{
			"X-Version": version,
		}).
		Run(BasicEngine(), func(r gofight.HttpResponse, rq gofight.HttpRequest) {

			assert.Equal(t, version, rq.Header.Get("X-Version"))
			assert.Equal(t, "Hello World", r.Body.String())
			assert.Equal(t, http.StatusOK, r.Code)
		})
}