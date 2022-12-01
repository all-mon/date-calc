package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/m0n7h0ff/date-calc/pkg/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	//settings
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static/", http.Dir("static/"))

	//handlers
	router.GET("/upload", h.uploadForm)
	router.POST("/upload", h.upload)

	//view
	employees := router.Group("/employees")
	{
		employees.GET("/schedule", h.getSchedule)
	}

	//api
	api := router.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("/:lastname", h.getEmployeeByLastname)
		}
	}
	return router
}
