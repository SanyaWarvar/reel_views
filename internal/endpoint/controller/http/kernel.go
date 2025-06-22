package http

import (
	v1 "rv/internal/endpoint/controller/http/api/v1"
	"rv/pkg/applogger"

	"rv/pkg/response"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Kernel struct {
	logInputParamOnErr bool

	logger  applogger.Logger
	builder *response.Builder

	dispatcher *v1.Dispatcher
}

func NewKernel(
	logInputParamOnErr bool,

	logger applogger.Logger,
	builder *response.Builder,
	dispatcher *v1.Dispatcher,
) *Kernel {
	return &Kernel{
		logInputParamOnErr: logInputParamOnErr,

		logger:  logger,
		builder: builder,

		dispatcher: dispatcher,
	}
}

func (k *Kernel) Init() *gin.Engine {
	router := gin.New()
	gin.SetMode(gin.DebugMode)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		gin.Recovery(),
		HeaderCtxHandler(),
		LoggerHandler(k.logger, k.logInputParamOnErr),
		CORSHandler(),
		ErrorHandler(k.builder),
	)
	k.initApi(router.Group("/rl", RequestIdValidationHandler))
	return router
}

func (k *Kernel) initApi(router *gin.RouterGroup) {
	k.dispatcher.Init(router.Group("api"))
}
