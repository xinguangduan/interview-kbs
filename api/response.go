package api

import (
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Status int    `json:"-"`
	Code   int    `json:"code,omitempty"`
	Msg    string `json:"msg,omitempty"`
	Data   any    `json:"data,omitempty"`
	Total  int    `json:"total,omitempty"`
}

func (m ResponseMessage) IsEmpty() bool {
	return reflect.DeepEqual(m, ResponseMessage{})
}
func HttpResponse(ctx *gin.Context, status int, resp ResponseMessage) {
	if resp.IsEmpty() {
		ctx.AbortWithStatus(status)
		return
	}

	ctx.AbortWithStatusJSON(status, resp)
}

func buildStatus(resp ResponseMessage, nDefaultStatus int) int {
	if 0 == resp.Status {
		return nDefaultStatus
	}

	return resp.Status
}

func OK(ctx *gin.Context, resp ResponseMessage) {
	HttpResponse(ctx, buildStatus(resp, http.StatusOK), resp)
}

func Fail(ctx *gin.Context, resp ResponseMessage) {
	HttpResponse(ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func ServerFail(ctx *gin.Context, resp ResponseMessage) {
	HttpResponse(ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}
