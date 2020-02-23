package response

import routing "github.com/go-ozzo/ozzo-routing"

const contentTypeJson = "application/json"

func MakeResponse(ctx *routing.Context, statusCode int, body []byte) {
	ctx.Response.Header().Add("Content-Type", contentTypeJson)
	ctx.Response.WriteHeader(statusCode)
	ctx.Response.Write(body)
}
