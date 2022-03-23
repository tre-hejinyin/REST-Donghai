package impl

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"server/domain"
)

type userRepository struct {
	db *gorm.DB
}

func (u userRepository) FindByEmail(ctx context.Context, email string) (result domain.Users, err error) {
	err = u.db.WithContext(ctx).Where(&domain.Users{Email: email}).First(&result).Error
	return
}

func (u userRepository) Create(ctx context.Context, user *domain.Users) error {
	err := u.db.WithContext(ctx).Create(&user).Error
	return err
}

func (u userRepository) UpdateByEmail(ctx context.Context, email string, user *domain.Users) (*domain.Users, error) {
	err := u.db.WithContext(ctx).
		Model(&user).
		Clauses(clause.Returning{}).
		Where(&domain.Users{Email: email}).
		Updates(&user).Error
	return user, err
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
