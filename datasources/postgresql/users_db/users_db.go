package usersdb

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

type Config struct {
	HOST           string `mapstructure:"DB_HOST"`
	PORT           string `mapstructure:"DB_PORT"`
	USER           string `mapstructure:"DB_USER"`
	PASSWORD       string `mapstructure:"DB_PASS"`
	DBNAME         string `mapstructure:"DB_NAME"`
	SSLMODE        string `mapstructure:"DB_SSLMODE"`
	DB_DRIVER      string `mapstructure:"DB_DRIVER"`
	DB_SOURCE      string `mapstructure:"DB_SOURCE"`
	SERVER_ADDRESS string `mapstructure:"SERVER_ADDRESS"`
}

const (
	postgresql_users_host     = "DB_HOST"
	postgresql_users_port     = "DB_PORT"
	postgresql_users_user     = "DB_USER"
	postgresql_users_pass     = "DB_PASS"
	postgresql_users_db_name  = "DB_NAME"
	postgresql_users_ssl_mode = "DB_SSLMODE"
)

var (
	Client   *sql.DB
	host     = os.Getenv(postgresql_users_host)
	port     = os.Getenv(postgresql_users_port)
	user     = os.Getenv(postgresql_users_user)
	password = os.Getenv(postgresql_users_pass)
	db_name  = os.Getenv(postgresql_users_db_name)
	ssl_mode = os.Getenv(postgresql_users_ssl_mode)
)

func init() {
	var err error
	dataSourceName := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, password, db_name, ssl_mode,
	)

	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully connfigured")
}
