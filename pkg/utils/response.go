package utils

import (
	"encoding/json"
	"net/http"

	"github.com/andreylm/bmlabs/pkg/logger"
	routing "github.com/qiangxue/fasthttp-routing"
)

const contentTypeJSON = "application/json"

// MakeJSONResponseFromBytes - makes response for fasthttp-routing
func MakeJSONResponseFromBytes(ctx *routing.Context, statusCode int, body []byte) {
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", contentTypeJSON)
	ctx.Response.SetBody(body)
}

// MakeJSONResponseFromMap - makes response for fasthttp-routing
func MakeJSONResponseFromMap(ctx *routing.Context, statusCode int, body map[string]string) {
	data, err := json.Marshal(body)
	if err != nil {
		logger.Get().Error(err)
		ctx.Response.SetStatusCode(http.StatusInternalServerError)
		ctx.Response.Header.Add("Content-Type", contentTypeJSON)
		return
	}
	
	ctx.Response.SetStatusCode(statusCode)
	ctx.Response.Header.Add("Content-Type", contentTypeJSON)
	ctx.Response.SetBody(data)
}
