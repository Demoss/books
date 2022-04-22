package handler

import (
	"JMIND/internal/service"
	"github.com/gin-gonic/gin"
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
		block := api.Group("/block/:block")
		{
			block.GET("/total", h.GetInfoByBlock)
		}
	}

	return router
}
