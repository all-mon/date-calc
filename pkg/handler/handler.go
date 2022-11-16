package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m0n7h0ff/date-calc/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("/:lastname", h.getEmployeeByLastname)
		}
	}
	return router
}
