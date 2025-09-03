package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewConnection(ctx context.Context, connectionURI string) (*pgxpool.Pool, error) {
	ctx, cancel := context.WithTimeout(ctx, 10 * time.Second)
	defer cancel()

	conn, err := pgxpool.Connect(ctx, connectionURI)
	if err != nil {
		return nil, err
	}

	return conn, nil
}