package storage

import (
	"context"
	"fmt"
	"go-tube/model"
	"strings"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB interface {
	GetVideos(ctx context.Context) ([]model.Video, error)
	SearchVideos(ctx context.Context, query string) ([]model.Video, error)
	InsertVideo(ctx context.Context, video model.Video) error
}

type MysqlDB struct {
	DB *gorm.DB
}

func NewDB() *MysqlDB {

	dbHost := "mysql"
	dbName := "go_tube"
	dbUser := "go_user"
	dbPassword := "go_password"
	dbPort := "3306"

	source := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(source), &gorm.Config{})
	if err != nil {
		fmt.Println("gorm Db connection ", err)
		panic("Failed to connect to database")
	}

	if err := db.AutoMigrate(&model.Video{}); err != nil {
		fmt.Println("Failed to auto migrate:", err)
		panic("Auto migration failed")
	}

	return &MysqlDB{DB: db}
}

func (db *MysqlDB) GetVideos(ctx context.Context) ([]model.Video, error) {

	var videos []model.Video
	if err := db.DB.WithContext(ctx).Order("published_at desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (db *MysqlDB) SearchVideos(ctx context.Context, query string) ([]model.Video, error) {
	var videos []model.Video
	searchQuery, args := buildSearchQuery(query)

	if err := db.DB.WithContext(ctx).Where(searchQuery, args...).Order("published_at desc").Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}

func (db *MysqlDB) InsertVideo(ctx context.Context, video model.Video) error {
	if video.Title == "" {
		return fmt.Errorf("video title cannot be empty")
	}

	if err := db.DB.WithContext(ctx).Create(&video).Error; err != nil {
		return fmt.Errorf("failed to insert video: %w", err)
	}

	return nil
}

func buildSearchQuery(query string) (string, []interface{}) {
	words := strings.Fields(query)
	var searchQuery strings.Builder
	var args []interface{}

	for i, word := range words {
		if i > 0 {
			searchQuery.WriteString(" OR ")
		}
		searchQuery.WriteString("(LOWER(title) LIKE LOWER(?) OR LOWER(description) LIKE LOWER(?))")
		args = append(args, "%"+word+"%", "%"+word+"%")
	}

	return searchQuery.String(), args
}
