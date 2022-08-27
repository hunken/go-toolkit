package mysql

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strings"
	"time"
)

type DB interface {
	GetConfig(database string) string
	Connect(dsn string) *gorm.DB
}

type Connection struct {
}

func (conn Connection) GetConfig(database string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database = strings.ToUpper(database)

	dbUser := os.Getenv(database + "_DB_USER")
	dbPw := os.Getenv(database + "_DB_PW")
	dbHost := os.Getenv(database + "_DB_HOST")
	dbName := os.Getenv(database + "_DB_NAME")
	dbPort := os.Getenv(database + "_DB_PORT")

	dsn := dbUser + ":" + dbPw + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println("DSN " + database + "->" + string(dsn))
	return dsn
}

func (conn Connection) Connect(dsn string) *gorm.DB {
	loggerus := logrus.New()
	loggerus.Formatter = &logrus.JSONFormatter{}

	newLogger := logger.New(
		loggerus, // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		Logger:                 newLogger,
	})
	if err != nil {
		log.Fatal("Has an unexpected error : " + err.Error())
	}
	return db
}
