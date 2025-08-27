package port

import (
	"context"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
)

type IUserRepository interface {
	GetById(ctx context.Context, id string) (*models.User, error)
	GetAll(ctx context.Context) ([]*models.User, error)
	Create(ctx context.Context, user *models.User) error
	Update(ctx context.Context, user *models.User) error
	Delete(ctx context.Context, id string) (bool, error)
}