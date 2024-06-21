package v1

import (
	"github.com/gin-gonic/gin"

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


			// cars.GET("/", h.getAllCars)
			// cars.PUT("/", h.updateCatalog)
			// cars.GET("/:id", h.getCarById)

		}
	}
	return router
}
