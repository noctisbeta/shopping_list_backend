package postgres_service

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type IPostgresService interface {
	GetDB() *sql.DB
}

type postgresService struct {
	DB *sql.DB
}

var postgresServiceInstance *postgresService

func GetPostgresServiceInstance() *postgresService {
	if postgresServiceInstance == nil {
		postgresServiceInstance, _ = newPostgresService()
	}
	return postgresServiceInstance
}

func newPostgresService() (*postgresService, error) {
	db, err := sql.Open("postgres", "postgres://default:DyMaum9fovO1@ep-misty-union-83632081.eu-central-1.postgres.vercel-storage.com:5432/verceldb")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &postgresService{DB: db}, nil
}

func (ps *postgresService) GetDB() *sql.DB {
	return ps.DB
}
