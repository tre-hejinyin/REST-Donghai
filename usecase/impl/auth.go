package impl

import (
	"crypto/md5"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/domain"
	"server/middleware/cache"
	"server/middleware/constants"
	"server/middleware/jwt"
	"server/repository"
)

type authUseCase struct {
	userRepository repository.UserRepository
}

func NewAuthUseCase(userRepository repository.UserRepository) *authUseCase {
	return &authUseCase{userRepository: userRepository}
}

func (u *authUseCase) Signup(c *gin.Context, user *domain.Users) (*domain.UsersVo, error) {
	user.Password = passwordEncryption(user.Password)
	userCheck, err := u.userRepository.FindByEmail(c, user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if userCheck.Email != "" {
		return nil, errors.New("email already exists")
	}
	err = u.userRepository.Create(c, user)
	if err != nil {
		return nil, err
	}
	var vo = domain.UsersVo{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return &vo, err
}

func (u *authUseCase) Signin(c *gin.Context, user *domain.Users) (string, error) {
	userCheck, err := u.userRepository.FindByEmail(c, user.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", errors.New("email didn't exist")
		}
		return "", err
	}

	if userCheck.Password != passwordEncryption(user.Password) {
		return "", errors.New("password error")
	}
	token, err := createTokenAndToRedis(user.Email)
	if err != nil {
		return "", err
	}

	return token, err
}

func (u *authUseCase) Profile(c *gin.Context) (*domain.UsersVo, error) {
	email, exists := c.Get(constants.EMAIL)
	if !exists || email == nil {
		return nil, errors.New("email didn't exist")
	}
	user, err := u.userRepository.FindByEmail(c, email.(string))
	if err != nil {
		return nil, err
	}
	var vo = domain.UsersVo{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
	return &vo, err
}

func (u *authUseCase) ProfileUpdate(c *gin.Context, user *domain.Users) (*domain.UsersVo, error) {
	email, exists := c.Get(constants.EMAIL)
	if !exists || email == nil {
		return nil, errors.New("email didn't exist")
	}
	var updateUser *domain.Users
	updateUser, err := u.userRepository.UpdateByEmail(c, email.(string), &domain.Users{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
	if err != nil {
		return nil, err
	}
	var voUser = &domain.UsersVo{
		FirstName: updateUser.FirstName,
		LastName:  updateUser.LastName,
		Email:     updateUser.Email,
	}
	return voUser, nil

}

func createTokenAndToRedis(email string) (string, error) {
	// 生成token
	var token string
	token, err := jwt.CreateToken(email)
	if err != nil {
		return "", fmt.Errorf("gen token failed: error=%w", err)
	}
	// 保存至redis
	err = cache.HashSet(constants.CACHE_ACCOUNT_GROUP+email, constants.HASH_TOKEN_KEY, token)
	if err != nil {
		return "", fmt.Errorf("cache hashset failed: error=%w", err)
	}
	return token, err
}

func passwordEncryption(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
