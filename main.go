package main

import (
	"nida/vmwareapi/handlers"
	"nida/vmwareapi/repository"
	"nida/vmwareapi/services"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()
	// db := initDatabase()

	cusRepo := repository.NewCustomerRepositoryConn()
	cusService := services.NewCustomerService(cusRepo)
	cusHandle := handlers.NewCustomerHandler(cusService)

	router := fiber.New()

	router.Get("/customer", cusHandle.GetCustomers)
	router.Get("/customer/:id", cusHandle.GetCustomer)

	router.Listen(":8000")
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	t, err := time.LoadLocation("Asia/Bangkok")

	if err != nil {
		panic(err)
	}

	time.Local = t
}

// func initDatabase() *sqlx.DB {

// 	dsn := fmt.Sprintf("%v:%v@%v:%v/%v",
// 		viper.GetString("db.username"),
// 		viper.GetString("db.password"),
// 		viper.GetString("db.host"),
// 		viper.GetInt("db.port"),
// 		viper.GetString("db.database"),
// 	)

// 	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)

// 	if err != nil {
// 		panic(err)
// 	}

// 	db.SetConnMaxLifetime(1 * time.Minute)
// 	db.SetMaxOpenConns(10)
// 	db.SetMaxIdleConns(10)

// 	return db

// }
