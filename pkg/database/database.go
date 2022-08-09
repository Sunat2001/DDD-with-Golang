package database

import (
	"CRUD/pkg/domain/users/Entities"
	"github.com/kelseyhightower/envconfig"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type dBConfig struct {
	Connection string `default:"mysql"`
	Host       string `default:"127.0.0.1"`
	Database   string `default:"root"`
	Port       string `default:"3306"`
	User       string `default:"root"`
	Password   string `default:"root"`
}

//Connector variable used for connect to DB operation's
var Connector *gorm.DB

//InitDB creates DB Connection
func InitDB() {
	var err error

	Connector, err = gorm.Open(mysql.Open(prepareConnectionString()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Connector.Migrator().AutoMigrate(&Entities.User{})

	log.Println("Successfully connected to Database!!")
}

func prepareConnectionString() string {
	var dbConfig dBConfig

	err := envconfig.Process("DB", &dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	dsn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.Database + "?charset=utf8mb4&parseTime=True&loc=Local"

	return dsn
}
