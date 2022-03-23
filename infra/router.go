package infra

import (
	"gorm.io/gorm"
	"server/controller"
	"server/middleware/auth"
	repo "server/repository/impl"
	useCase "server/usecase/impl"

	"github.com/gin-gonic/gin"
)

func SetupServer(store *gorm.DB) (*gin.Engine, error) {
	r := gin.New()
	v1 := r.Group("/v1")
	{
		configAuthRouter(v1, store)
	}
	return r, nil
}

func configAuthRouter(v *gin.RouterGroup, store *gorm.DB) {
	rg := v.Group("")
	{
		userRepo := repo.NewUserRepository(store)
		authUseCase := useCase.NewAuthUseCase(userRepo)
		ctrl := controller.NewAuthController(authUseCase)
		rg.POST("/signup", ctrl.Signup)
		rg.POST("/signin", ctrl.Signin)
		rg.POST("/profile", auth.JWTAuthMiddleware(), ctrl.Profile)
		rg.POST("/profile/update", auth.JWTAuthMiddleware(), ctrl.ProfileUpdate)
	}
}
