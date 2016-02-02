package config

import (
	"log"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/vrischmann/envconfig"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var conf struct {
	DB struct {
		Connection string `envconfig:"default=root:toor@/radio-go"`
	}
}

// initialized env configs
var Env = initializeConfig()

// initialize config and turns it to map
func initializeConfig() map[string]string {
	err := envconfig.Init(&conf)
	if err != nil {
		log.Panic("Error on env config initialize! ", err.Error())
	}

	return map[string]string{
		"dbConnection": conf.DB.Connection,
	}
}

func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func GetConnection() gorm.DB {
	db, err := gorm.Open("mysql", Env["dbConnection"])
	if err != nil {
		log.Panicf("Error on database connection! Err: %s", err.Error())
	}

	return db
}
