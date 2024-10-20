package main

import (
	"fmt"
	"time"

	"github.com/CUBS-sources-code/CUBS-coin/handler"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	initConfig()
    initTimeZone()
    db := initDatabase()

    // Mock repository
    // userRepositoryMock := repository.NewUserRepositoryMock()
	// _ = userRepositoryMock
	
	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	app := fiber.New()

    // Routes
	app.Get("/users", userHandler.GetUsers)
	app.Get("/user/:student_id", userHandler.GetUser)

    // Start server
	app.Listen(fmt.Sprintf(":%v", viper.GetInt("app.port")))
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

func initDatabase() *gorm.DB {
    dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.database"))

	dial := mysql.Open(dsn)
	db, err := gorm.Open(dial)
	if err != nil {
		panic(err)
	}

    return db
}

func initTimeZone() {
    ict, err := time.LoadLocation("Asia/Bangkok")
    if err != nil {
        panic(err)
    }

    time.Local = ict
}