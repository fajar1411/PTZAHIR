package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	MYSQL_USER     string
	MYSQL_PASSWORD string
	MYSQL_HOST     string
	MYSQL_PORT     uint
	MYSQL_DBNAME   string
	SERVER_PORT    uint
}

var lock = &sync.Mutex{}
var appConfig *AppConfig

func GetConfig() *AppConfig {
	lock.Lock()
	defer lock.Unlock()

	if appConfig == nil {
		appConfig = initConfig()
	}

	return appConfig
}

func initConfig() *AppConfig {
	var defaultConfig AppConfig

	if _, exist := os.LookupEnv("SECRET"); !exist {
		if err := godotenv.Load("local.env"); err != nil {
			log.Println(err)
		}
	}

	cnvServerPort, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	fmt.Println("ini port", cnvServerPort)
	if err != nil {
		log.Fatal("Cannot parse Server Port variable")
		return nil
	}
	defaultConfig.SERVER_PORT = uint(cnvServerPort)
	defaultConfig.MYSQL_DBNAME = os.Getenv("MYSQL_DBNAME")
	defaultConfig.MYSQL_USER = os.Getenv("MYSQL_USER")
	defaultConfig.MYSQL_PASSWORD = os.Getenv("MYSQL_PASSWORD")
	defaultConfig.MYSQL_HOST = os.Getenv("MYSQL_HOST")

	cnvDBPort, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		log.Fatal("Cannot parse DB Port variable")
		return nil
	}
	defaultConfig.MYSQL_PORT = uint(cnvDBPort)

	return &defaultConfig
}
