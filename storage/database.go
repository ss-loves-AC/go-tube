package storage	

import (
	"go-tube/model"
)

type DB interface {
	GetVideos() ([]model.Video, error)
	SearchVideos(query string) ([]model.Video, error)
}

type Cache interface {
	Get(key string) ([]model.Video, bool)  
	Set(key string, value []model.Video)   
}

type PostgresDB struct{}

func NewDB() *PostgresDB {
	return &PostgresDB{}
}

func (db *PostgresDB) GetVideos() ([]model.Video, error) {
	return nil, nil
}

func (db *PostgresDB) SearchVideos(query string) ([]model.Video, error) {
	return nil, nil
}


type RedisCache struct{}

func NewCache() *RedisCache {
	return &RedisCache{}
}

func (r *RedisCache) Get(key string) ([]model.Video, bool) {
	return nil, false
}

func (r *RedisCache) Set(key string, value []model.Video) {}