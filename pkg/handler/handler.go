package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		employees := api.Group("/employees")
		{
			employees.GET("/:fio",h.getEmployeeByLastname)
		}
	}
	return router
}
