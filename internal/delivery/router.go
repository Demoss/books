package delivery

import (
	"github.com/Demoss/books/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("/by-author", h.getAuthorsBooks)
			books.POST("/", h.addBook)
			books.PUT("/:id", h.updateBook)
			books.DELETE("/delete-book", h.deleteBook)
		}
	}

	return router
}
