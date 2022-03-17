package delivery

import (
	sign_in "github.com/Demoss/books/internal/delivery/request/sign-in"
	sign_up "github.com/Demoss/books/internal/delivery/request/sign-up"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	var req sign_up.Request
	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.Authorization.CreateUser(sign_up.MapToDomain(req))
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, "OK")

}

func (h *Handler) signIn(c *gin.Context) {
	var req sign_in.Request
	if err := c.BindJSON(&req); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid input body")
	}
	token, err := h.services.Authorization.GenerateToken(req.Username, req.Password)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]string{"token": token})
}
