package usecase

import (
	"github.com/gin-gonic/gin"
	"server/domain"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock_$GOFILE -package=mock
type AuthUseCase interface {
	Signup(c *gin.Context, users *domain.Users) error
	Signin(c *gin.Context, d *domain.Users) (string, error)
	Profile(c *gin.Context) (*domain.UsersVo, error)
	ProfileUpdate(c *gin.Context, d *domain.Users) (*domain.UsersVo, error)
}
