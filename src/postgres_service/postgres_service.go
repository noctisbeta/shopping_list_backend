package postgres_service

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

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
	log.Println("Getting postgres service instance")
	if postgresServiceInstance == nil {
		postgresServiceInstance, _ = newPostgresService()
	}
	return postgresServiceInstance
}

func newPostgresService() (*postgresService, error) {
	err := godotenv.Load()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	postgresURL := os.Getenv("POSTGRES_URL")

	log.Println("Connecting to postgres at: " + postgresURL)

	db, err := sql.Open("postgres", postgresURL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &postgresService{DB: db}, nil
}

func (ps *postgresService) GetDB() *sql.DB {
	return ps.DB
}
