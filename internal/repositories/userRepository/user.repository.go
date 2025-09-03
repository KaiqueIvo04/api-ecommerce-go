package userRepository

import (
	"context"
	"log"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
	"github.com/jackc/pgx/v4/pgxpool"
)

type UserRepository struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) port.IUserRepository {
	return &UserRepository{
		conn: conn,
	}
}

// Create implements port.IUserRepository.
func (ur *UserRepository) Insert(ctx context.Context, user models.User) (models.User, error) {
	_, err := ur.conn.Exec(
		ctx,
		"INSERT INTO users (id, name, email, password, type) VALUES ($1, $2, $3, $4, $5)",
		user.Id,
		user.Name,
		user.Email,
		user.Password,
		user.Type,
	)
	if err != nil {
		log.Print(err)
		return models.User{}, err
	}

	return user, err
}

// Delete implements port.IUserRepository.
func (u *UserRepository) Delete(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// GetAll implements port.IUserRepository.
func (u *UserRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

// GetById implements port.IUserRepository.
func (u *UserRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// Update implements port.IUserRepository.
func (u *UserRepository) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}
