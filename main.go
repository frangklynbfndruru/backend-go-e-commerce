package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/frangklynbfndruru/backend-go-e-commerce/config"
	"github.com/joho/godotenv"
)

func main() {
	var server = config.Server{}
	var appConfig = config.AppConfig{}
	var dbConfig = config.DbConfig{}

	err := godotenv.Load()
	if err != nil {
		fmt.Errorf("Error at .env")
		// fmt.Println("Error at .env")
		// log.Fatal(err)
	}

	appConfig.AppName = os.Getenv("APP_NAME")
	appConfig.AppPort = os.Getenv("APP_PORT")
	appConfig.AppEnv = os.Getenv("APP_ENV")

	dbConfig.DbHost = os.Getenv("DB_HOST")
	dbConfig.DbUser = os.Getenv("DB_USER")
	dbConfig.DbPort = os.Getenv("DB_PORT")
	dbConfig.DbPassword = os.Getenv("DB_PASSWORD")
	dbConfig.DbName = os.Getenv("DB_NAME")

	flag.Parse() //untuk menerima command go run dari terminal
	// arg := flag.Arg(0) //mengambil argumen pertama dari command line. contoh `go run db:migrate`

	server.Initialize(appConfig, dbConfig)
	server.RunPort(": " + appConfig.AppPort)

}
