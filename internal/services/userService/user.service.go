package userService

import (
	"context"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
)

type UserService struct {
	userRepo port.IUserRepository
}

func New(userRepo port.IUserRepository) port.IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

// Add implements port.IUserService.
func (u *UserService) Create(ctx context.Context, user models.User) (models.User, error) {
	return u.userRepo.Insert(ctx, user)
}

// Count implements port.IUserService.
func (u *UserService) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// GetAll implements port.IUserService.
func (u *UserService) GetAll(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

// GetByID implements port.IUserService.
func (u *UserService) GetByID(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// Remove implements port.IUserService.
func (u *UserService) Remove(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Update implements port.IUserService.
func (u *UserService) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}
