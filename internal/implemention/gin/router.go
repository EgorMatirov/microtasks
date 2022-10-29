package ginrouter

import (
	"github.com/EgorMatirov/microtasks/internal/usecase"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RouterHandler struct {
	ucHandler usecase.Handler
	logger    *zap.Logger
}

func NewRouter(i usecase.Handler, logger *zap.Logger) RouterHandler {
	return RouterHandler{
		ucHandler: i,
		logger:    logger,
	}
}

func (rH RouterHandler) SetRoutes(r *gin.Engine) {
	r.GET("hello/:name", rH.sayHello)
}
