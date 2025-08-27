package userrepository

import (
	"context"

	"github.com/KaiqueIvo04/api-ecommerce-go/internal/models"
	"github.com/KaiqueIvo04/api-ecommerce-go/pkg/port"
	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepository struct {
	conn *pgxpool.Pool
}

func New(conn *pgxpool.Pool) port.IUserRepository {
	return &userRepository{
		conn: conn,
	}
}

// Create implements port.IUserRepository.
func (ur *userRepository) Create(ctx context.Context, user *models.User) error {
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	_, err := ur.conn.Exec(
		ctx,
		"INSERT INTO users (id, name, email, password, _type) VALUES ($1, $2, $3, $4, $5)",
		user.GetId(),
		user.GetName(),
		user.GetEmail(),
		user.GetPassword(),
		user.GetType(),
	)

	return err
}

// Delete implements port.IUserRepository.
func (u *userRepository) Delete(ctx context.Context, id string) (bool, error) {
	panic("unimplemented")
}

// GetAll implements port.IUserRepository.
func (u *userRepository) GetAll(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

// GetById implements port.IUserRepository.
func (u *userRepository) GetById(ctx context.Context, id string) (*models.User, error) {
	panic("unimplemented")
}

// Update implements port.IUserRepository.
func (u *userRepository) Update(ctx context.Context, user *models.User) error {
	panic("unimplemented")
}
