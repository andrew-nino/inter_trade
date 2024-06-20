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

	// api := router.Group("/api", h.userIdentity)
	// {
	// 	cars := api.Group("/cars")
	// 	{
	// 		cars.POST("/", h.addingCars)
	// 		cars.GET("/", h.getAllCars)
	// 		cars.PUT("/", h.updateCatalog)
	// 		cars.GET("/:id", h.getCarById)
	// 		cars.DELETE("/:id", h.deleteCar)
	// 	}
	// }
	return router
}
