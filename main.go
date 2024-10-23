package main

import (
	"fmt"
	"time"

	"github.com/CUBS-sources-code/CUBS-coin/handler"
	"github.com/CUBS-sources-code/CUBS-coin/logs"
	"github.com/CUBS-sources-code/CUBS-coin/repository"
	"github.com/CUBS-sources-code/CUBS-coin/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	app.Use(cors.New())
	api := app.Group("/api")
	
	userRepository := repository.NewUserRepositoryDB(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

    transactionRepository := repository.NewTransactionRepositoryDB(db)
    transactionService := service.NewTransactionService(transactionRepository, userRepository)
    transactionHandler := handler.NewTransactionHandler(transactionService)

	authService := service.NewAuthService(userRepository)
	authHandler := handler.NewAuthHandler(authService)

	// Public
	api.Get("/transactions", transactionHandler.GetTransactions)
	api.Get("/transaction/:id", transactionHandler.GetTransaction)
	api.Post("/signup", authHandler.SignUp)
	api.Post("/signin", authHandler.SignIn)

	// Private
	api.Use(authHandler.AuthorizationRequired())
	api.Get("/user", userHandler.GetMyUser)
	api.Post("/transfer", transactionHandler.Transfer)

	// Admin
	api.Use(authHandler.IsAdmin)
	api.Post("/transaction/create", transactionHandler.CreateTransaction)
	api.Post("/user/create", userHandler.CreateUser)
	api.Get("/users", userHandler.GetUsers)
	api.Get("/user/:student_id", userHandler.GetUser)
	api.Patch("/changetoadmin/:student_id", userHandler.ChangeRoleToAdmin)
	api.Patch("/changetomember/:student_id", userHandler.ChangeRoleToMember)

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

