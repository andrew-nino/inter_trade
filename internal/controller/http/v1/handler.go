package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	_ "international_trade/docs"

	"international_trade/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api/v1", h.userIdentity)
	{
		hash := api.Group("/hash")
		{
			hash.POST("/:type", h.addingHash)
			hash.DELETE("/:type", h.deleteHash)
			hash.GET("/:type", h.getValue)
		}
	}
	return router
}
