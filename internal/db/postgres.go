package db

import "context"

type Pool interface {
	Close()
}

func OpenPostgres(ctx context.Context, databaseURL string) (Pool, error) {
	_ = ctx
	_ = databaseURL
	return nil, nil
}
