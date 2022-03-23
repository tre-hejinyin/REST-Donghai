package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"server/domain"
	"server/middleware/response"
	"server/usecase"
)

type auth struct {
	useCase usecase.AuthUseCase
}

func NewAuthController(usecase usecase.AuthUseCase) *auth {
	return &auth{useCase: usecase}
}

func (u *auth) Signup(c *gin.Context) {
	var user domain.Users
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	err = u.useCase.Signup(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
	})
}

func (u *auth) Signin(c *gin.Context) {
	var user domain.Users
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	token, err := u.useCase.Signin(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  http.StatusOK,
		Data: token,
	})
}

func (u *auth) Profile(c *gin.Context) {
	vo, err := u.useCase.Profile(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: vo,
	})
}

func (u *auth) ProfileUpdate(c *gin.Context) {
	var user domain.Users
	err := c.ShouldBind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  http.StatusText(http.StatusBadRequest),
		})
		return
	}
	vo, err := u.useCase.ProfileUpdate(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, response.Body{
		Code: http.StatusOK,
		Msg:  http.StatusText(http.StatusOK),
		Data: vo,
	})
}
