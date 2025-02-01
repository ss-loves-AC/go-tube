package storage	

import (
	"go-tube/model"
)

type DB interface {
	GetVideos() ([]model.Video, error)
	SearchVideos(query string) ([]model.Video, error)
	InsertVideo(video model.Video) error
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

func (db *PostgresDB) InsertVideo(video model.Video) error {
	return nil
}


