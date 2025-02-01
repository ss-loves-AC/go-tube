package storage

import "go-tube/model"

type Cache interface {
	Get(key string) ([]model.Video, bool)
	Set(key string, value []model.Video)
}

type RedisCache struct{}

func NewCache() *RedisCache {
	return &RedisCache{}
}

func (r *RedisCache) Get(key string) ([]model.Video, bool) {
	return nil, false
}

func (r *RedisCache) Set(key string, value []model.Video) {}
