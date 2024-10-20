package main

import (
	"fmt"
	"time"

	"github.com/CUBS-sources-code/CUBS-coin/handler"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {

	// Load config
	initConfig()
	// Set timezone
	initTimeZone()
	// Connect database
	db := initDatabase()
	// Init fiber app
	app := initApp()
    api := app.Group("/api")

	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

    transactionRepository := repository.NewTransactionRepositoryDB(db)
    transactionService := service.NewTransactionService(transactionRepository, userRepository)
    transactionHandler := handler.NewTransactionHandler(transactionService)

	// Routes
	api.Get("/users", userHandler.GetUsers)
	api.Get("/user/:student_id", userHandler.GetUser)
	api.Post("/user/create", userHandler.CreateUser)

    api.Get("/transactions", transactionHandler.GetTransactions)
	api.Get("/transaction/:id", transactionHandler.GetTransaction)
    api.Post("/transaction/create", transactionHandler.CreateTransaction)

	// Start server
	logs.Info("CUBS coin service started at port " + viper.GetString("app.port"))
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
	db, err := gorm.Open(dial, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
        TranslateError: true,
	})
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

func initApp() *fiber.App {
    return fiber.New(fiber.Config{
		DisableStartupMessage: true,
		Prefork:               viper.GetBool("app.prefork"),
		AppName:               viper.GetString("app.name"),
	})
}