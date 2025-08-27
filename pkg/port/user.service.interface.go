package port

import (
	"context"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
)

type IUserService interface {
	Add(ctx context.Context, user *models.User) error
	GetAll(ctx context.Context) ([]*models.User, error)
	GetByID(ctx context.Context,id string) (*models.User, error)
	Update(ctx context.Context,user *models.User) error
	Remove(ctx context.Context,id string) error
	Count(ctx context.Context) (int64, error)
}
