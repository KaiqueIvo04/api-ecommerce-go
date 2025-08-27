package userservice

import (
	"context"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
)

type userService struct {
	userRepo port.IUserRepository
}

func New(userRepo port.IUserRepository) port.IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Add implements port.IUserService.
func (u *userService) Add(ctx context.Context, user *models.User) error {
	user.New(user.GetName(), user.GetEmail(), user.GetPassword(), user.GetType())
	return u.userRepo.Create(ctx, user)
}

// Count implements port.IUserService.
func (u *userService) Count(ctx context.Context) (int64, error) {
	panic("unimplemented")
}

// GetAll implements port.IUserService.
func (u *userService) GetAll(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

// GetByID implements port.IUserService.
func (u *userService) GetByID(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// Remove implements port.IUserService.
func (u *userService) Remove(ctx context.Context, id string) error {
	panic("unimplemented")
}

// Update implements port.IUserService.
func (u *userService) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}
