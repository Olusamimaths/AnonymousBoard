package database

import (
	"fmt"
	"net/url"

	"github.com/olusamimaths/AnonymousBoard/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseConnection interface {
	Get() *gorm.DB
}

type databaseConnection struct {
	DB *gorm.DB
}

func (d *databaseConnection) Get() *gorm.DB {
	return d.DB
}

func NewDatabaseConnection(c *config.Config) DatabaseConnection {
	config := c.Get()

	user := config.GetString("db.username")
	password := config.GetString("db.password")
	database := config.GetString("db.database")
	host := config.GetString("db.host")
	port := config.GetInt("db.port")

	var enableLogging logger.Interface
	if config.GetBool("db.log") {
		enableLogging = logger.Default
	}

	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     database,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}

	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{Logger: enableLogging})

	if err != nil {
		panic("databse connecton failed")
	}

	if config.GetBool("db.sync") {
		synchronize(db)
	}
	return &databaseConnection{DB: db}
}
