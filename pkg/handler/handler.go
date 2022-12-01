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

// InitRoutes настройка Gin роутера(uri, файловый сервер, папка с шаблонами )
func (h *Handler) InitRoutes() *gin.Engine {
	//settings
	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	router.StaticFS("/static/", http.Dir("static/"))

	//view handlers
	router.GET("/upload", h.uploadForm)
	router.POST("/upload", h.upload)

	employees := router.Group("/employees")
	{
		employees.GET("/schedule", h.getSchedule)
	}

	//api handlers
	api := router.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("/:lastname", h.getEmployeeByLastname)
		}
	}
	return router
}
