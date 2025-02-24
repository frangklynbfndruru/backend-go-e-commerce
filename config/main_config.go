package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frangklynbfndruru/backend-go-e-commerce/models"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/*
Server bertanggung jawab untuk mengatur kompnen dalam menjalankan aplikasi
-Koneksi kedalam DB
-Routing HTTP
Konfigurasi Aplikasi
*/
type Server struct {
	// variabel dan library yang digunakan
	DB        *gorm.DB //menggunakan ORM dari GORM
	Router    *mux.Router
	AppConfig *AppConfig
}

/*
AppConfig bertujuan untuk insialisasi/menyimpan konfigurasi aplikasi sehingga
lebih practice saat dijalankan dan akan disimpan kedalam Server
*/
type AppConfig struct {
	AppName string
	AppPort string
	AppEnv  string
	AppURL  string
}

/*
DbConfig bertujuan untuk insialisasi/menyimpan konfigurasi databse sehingga
lebih practice saat dijalankan dan akan disimpan kedalam Server
*/
type DbConfig struct {
	DbHost     string
	DbUser     string
	DbPassword string
	DbName     string
	DbPort     string
	SSLMode    string
}

func (server *Server) RunPort(address string) {
	fmt.Printf("Listening port on %s", address)
	log.Fatal(http.ListenAndServe(address, server.Router))
}

func (server *Server) Initialize(appConfig AppConfig, dbConfig DbConfig) {
	fmt.Println("Welcome to " + appConfig.AppName)

	server.InitializeDB(dbConfig)
}

func (server *Server) InitializeDB(dbConfig DbConfig) {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbConfig.DbHost, dbConfig.DbUser, dbConfig.DbPassword,
		dbConfig.DbName, dbConfig.DbPort)
	
	server.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Initialize database error!", err)
	}
	for _, model := range models.RegisterModels() {
		err := server.DB.Debug().AutoMigrate(model.Model)

		if err != nil {
			log.Fatal(err)
		}
	}
}
