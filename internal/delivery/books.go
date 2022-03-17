package delivery

import (
	"github.com/Demoss/books/internal/delivery/request/addBook"
	"github.com/Demoss/books/internal/delivery/request/deleteBook"
	"github.com/Demoss/books/internal/delivery/response/getAuthorsBooks"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addBook(c *gin.Context) {
	var req addBook.Request
	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Books.AddBook(addBook.MapToDomain(req))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")
}

func (h *Handler) getAuthorsBooks(c *gin.Context) {
	var req getAuthorsBooks.Request
	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	books, err := h.services.Books.GetAuthorsBooks(getAuthorsBooks.MapToQuery(req))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAuthorsBooks.MapToResponse(books))
}
func (h *Handler) getBookById(c *gin.Context) {

}
func (h *Handler) updateBook(c *gin.Context) {

}
func (h *Handler) deleteBook(c *gin.Context) {
	var req deleteBook.Request
	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Books.DeleteBook(deleteBook.MapToCommand(req))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, "OK")
}
