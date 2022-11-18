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
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	//fileServer := http.FileServer(http.Dir("./ui/static/"))
	router.StaticFS("/static/", http.Dir("static/"))
	router.GET("/upload", h.uploadForm)
	router.POST("upload", h.upload)

	api := router.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("/:lastname", h.getEmployeeByLastname)
		}
	}
	employees := router.Group("/employees")
	{
		employees.GET("/schedule", h.getSchedule)
	}
	return router
}
