package api

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseAPI struct {
	Ctx    *gin.Context
	Errors error
	Logger *zap.SugaredLogger
}
