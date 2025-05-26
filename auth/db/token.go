package db

import (
	"context"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type TokenRepository struct {
	cache *redis.Client
}

func NewTokenRepository(cache *redis.Client) *TokenRepository {
	return &TokenRepository{cache: cache}
}

func (r *TokenRepository) Add(userID uuid.UUID, token string) error {
	return r.cache.Set(context.Context(context.Background()), userID.String(), token, 0).Err()
}

func (r *TokenRepository) Get(userID uuid.UUID) (string, error) {
	return r.cache.Get(context.Context(context.Background()), userID.String()).Result()
}
