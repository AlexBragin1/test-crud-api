package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	server "test-crud-api"
	"test-crud-api/internal/config"
	"test-crud-api/internal/handler"
	"test-crud-api/internal/repository"
	"test-crud-api/internal/service"
)

func main() {
	cfg := config.LoadDB()

	db := repository.NewDB(repository.Config{
		Host:     cfg.DB.Host,
		Port:     cfg.DB.Port,
		Username: cfg.DB.Username,
		DBName:   cfg.DB.DBname,
		SSLMode:  cfg.DB.Sslmode,
		//Password: os.Getenv("DB_PASSWORD"),
		Password: cfg.DB.Password,
	})

	store := repository.New(db)
	service := service.NewService(store)
	handler := handler.NewHandler(service)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8080", handler.InitRoutes()); err != nil {
			fmt.Printf("error occured while running http server: %s", err.Error())
		}
	}()
	
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT)
	<-ch

	fmt.Print("CRUDAPI Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		fmt.Printf("error occured on server shutting down: %s\n", err.Error())
	}

	if err := db.Close(); err != nil {
		fmt.Printf("error occured on db connection close: %s\n", err.Error())
	}
}
