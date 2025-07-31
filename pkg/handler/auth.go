package handler

import (
	"ToDoApp"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary      signUp
// @Description  create acc
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param       input body ToDoApp.User true "acc info"
// @Success      200  {integer}   integer 1
// @Failure      400  {object}  errorResponse
// @Failure      404  {object} errorResponse
// @Failure      500  {object} errorResponse
// @Failure default {object} errorResponse
// @Router       /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input ToDoApp.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// @Summary      signIn
// @Description  login
// @Tags         auth
// @ID login
// @Accept       json
// @Produce      json
// @Param       input body signInInput true "login info"
// @Success      200  {string}   string "token"
// @Failure      400  {object}  errorResponse
// @Failure      404  {object} errorResponse
// @Failure      500  {object} errorResponse
// @Failure default {object} errorResponse
// @Router       /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
