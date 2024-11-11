package handler

import (
	"github.com/bminh1762000/jwt-auth-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body models.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid input body")
		return
	}

	userId, err := h.services.Authorization.Register(user.Username, user.Password)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"user_id": userId,
	})
}

// @Summary SignIn
// @Tags auth
// @Description login
// @ID login
// @Accept  json
// @Produce  json
// @Param input body models.User true "credentials"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.Authorization.Login(user.Username, user.Password)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Generate token
	token, err := h.services.Jwt.GenerateToken(user)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

// @Summary RefreshToken
// @Tags auth
// @Description refresh token
// @ID refresh-token
// @Accept  json
// @Produce  json
// @Param input body models.Token true "refresh token"
// @Success 200 {string} string "token"
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/refresh-token [get]
func (h *Handler) refreshToken(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		newErrorMessage(c, http.StatusBadRequest, "invalid input body")
		return
	}

	user, err := h.services.Jwt.ValidateRefreshToken(token.RefreshToken)
	if err != nil {
		newErrorMessage(c, http.StatusUnauthorized, "invalid refresh token")
		return
	}

	token, err = h.services.Jwt.GenerateToken(user)
	if err != nil {
		newErrorMessage(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"access_token": token.AccessToken,
	})
}
