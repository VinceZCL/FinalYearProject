package client

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/VinceZCL/FinalYearProject/app/config"
)

type PostgresClient struct {
	DB *gorm.DB
}

func NewPostgres() (*PostgresClient, error) {

	host := config.Get().Database.Host
	port := config.Get().Database.Port
	name := config.Get().Database.Name
	user := config.Get().Database.User
	password := config.Get().Database.Password

	dsn := fmt.Sprintf("host=%s user=%s password='%s' dbname=%s port=%d sslmode=disable",
		host, user, password, name, port)

	gormLogger := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Millisecond * 200,
			LogLevel:      logger.Info,
			Colorful:      true,
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{Logger: gormLogger})

	if err != nil {
		return nil, err
	}

	return &PostgresClient{
		DB: db,
	}, nil
}
