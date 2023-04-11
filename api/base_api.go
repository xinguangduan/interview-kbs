package api

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/lightsoft/interview-knowledge-base/global"
	"github.com/lightsoft/interview-knowledge-base/utils"
	"go.uber.org/zap"
)

type BaseApi struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}

func NewBaseApi() BaseApi {
	return BaseApi{
		Logger: global.Logger,
	}
}

type BuildRequestOption struct {
	Ctx     *gin.Context
	DTO     any
	BindUri bool
	BindAll bool
}

func (m *BaseApi) BuildRequest(option BuildRequestOption) *BaseApi {
	var errResult error

	// 绑定请求上下文
	m.Ctx = option.Ctx

	// 绑定请求数据
	if option.DTO != nil {
		if option.BindAll || option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBindUri(option.DTO))
		}

		if option.BindAll || !option.BindUri {
			errResult = utils.AppendError(errResult, m.Ctx.ShouldBind(option.DTO))
		}

		if errResult != nil {
			errResult = m.ParseValidateErrors(errResult, option.DTO)
			m.AddError(errResult)
			m.Fail(ResponseMessage{
				Msg: m.GetError().Error(),
			})
		}
	}

	return m
}

func (m *BaseApi) AddError(errNew error) {
	m.Errors = utils.AppendError(m.Errors, errNew)
}

func (m *BaseApi) GetError() error {
	return m.Errors
}

func (m *BaseApi) ParseValidateErrors(errs error, target any) error {
	var errResult error

	errValidation, ok := errs.(validator.ValidationErrors)
	if !ok {
		return errs
	}

	// 通过反射获取指针指向元素的类型对象
	fields := reflect.TypeOf(target).Elem()
	for _, fieldErr := range errValidation {
		field, _ := fields.FieldByName(fieldErr.Field())
		errMessageTag := fmt.Sprintf("%s_err", fieldErr.Tag())
		errMessage := field.Tag.Get(errMessageTag)
		if errMessage == "" {
			errMessage = field.Tag.Get("message")
		}

		if errMessage == "" {
			errMessage = fmt.Sprintf("%s: %s Error", fieldErr.Field(), fieldErr.Tag())
		}

		errResult = utils.AppendError(errResult, errors.New(errMessage))
	}

	return errResult
}

func (m *BaseApi) Fail(resp ResponseMessage) {
	HttpResponse(m.Ctx, buildStatus(resp, http.StatusBadRequest), resp)
}

func (m *BaseApi) OK(resp ResponseMessage) {
	HttpResponse(m.Ctx, buildStatus(resp, http.StatusOK), resp)
}

func (m *BaseApi) ServerFail(resp ResponseMessage) {
	HttpResponse(m.Ctx, buildStatus(resp, http.StatusInternalServerError), resp)
}

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
