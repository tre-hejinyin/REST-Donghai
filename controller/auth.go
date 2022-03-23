package controller

import (
	"errors"
	"net/http"
	"regexp"
	"strings"

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
	err = validationSignupData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
		})
		return
	}
	result, err := u.useCase.Signup(c, &user)
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
		Data: result,
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
	err = validationSigninData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
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
	err = validationProfileUpdateData(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Body{
			Code: http.StatusBadRequest,
			Msg:  err.Error(),
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

func validationSignupData(user *domain.Users) (err error) {
	user.Email = strings.TrimSpace(user.Email)
	validationEmail, err := regexp.MatchString(`^.{1,}$`, user.Email)
	if err != nil {
		return err
	}
	if !validationEmail {
		return errors.New("email cannot be empty")
	}
	err = validationUserName(user)
	if err != nil {
		return err
	}
	validationPassword, err := regexp.MatchString(`^[A-Za-z0-9]{6,16}$`, user.Password)
	if err != nil {
		return err
	}
	if !validationPassword {
		return errors.New("password has to be minimum 6 characters,maximum 16 characters and alphanumeric")
	}
	return nil
}

func validationSigninData(user *domain.Users) (err error) {
	if strings.TrimSpace(user.Email) == "" {
		return errors.New("email cannot be empty")
	}
	if strings.TrimSpace(user.Password) == "" {
		return errors.New("password cannot be empty")
	}
	return err
}

func validationProfileUpdateData(user *domain.Users) (err error) {
	return validationUserName(user)
}

func validationUserName(user *domain.Users) (err error) {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	validationFirstName, err := regexp.MatchString(`^.{1,64}$`, user.FirstName)
	if err != nil {
		return err
	}
	if !validationFirstName {
		return errors.New("firstName cannot be empty and more than 64 characters")
	}
	validationLastName, err := regexp.MatchString(`^.{1,64}$`, user.LastName)
	if err != nil {
		return err
	}
	if !validationLastName {
		return errors.New("lastName cannot be empty and more than 64 characters")
	}
	return nil
}
