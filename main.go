package main

import (
	"database/sql"
	"fmt"
	"go-solid/config"
	deliveryHttp "go-solid/delivery/http"
	"go-solid/infrastructure/database"
	"go-solid/usecase"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("unable to load config: %v", err)
	}

	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s host=%s port=%d",
		cfg.Database.User, cfg.Database.Password, cfg.Database.DBName,
		cfg.Database.SSLMode, cfg.Database.Host, cfg.Database.Port,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Failed to close database: %v", err)
		}
	}()

	userRepo := database.NewPostgresUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := deliveryHttp.NewUserHandler(userUsecase)

	http.HandleFunc("/user", userHandler.GetUser)
	http.HandleFunc("/user/create", userHandler.CreateUser)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
