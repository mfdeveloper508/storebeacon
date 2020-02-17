package database

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/validations"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Database string
	Port     int
}

var (
	Conn *gorm.DB
)

var dbConfig DatabaseConfig

func Configure(tmpConfig DatabaseConfig) {
	dbConfig = tmpConfig
}

func GetConfig() DatabaseConfig {
	return dbConfig
}

func Close() {
	Conn.Close()
}

func Override(conn *gorm.DB) {
	Conn = conn
}

var Config DatabaseConfig

func Connect() {
	connectionUrl := "host=" + dbConfig.Host + " user=" + dbConfig.User + " dbname=" + dbConfig.Database + " password=" + dbConfig.Password + " sslmode=disable"

	var err error
	Conn, err = gorm.Open("postgres", connectionUrl)
	//Conn.LogMode(true)

	validations.RegisterCallbacks(Conn)

	if err != nil {
		panic(err)
	}
}
