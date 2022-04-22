package main

import (
	"JMIND"
	"JMIND/config"
	"JMIND/internal/handler"
	"JMIND/internal/repository"
	"JMIND/internal/service"
	"JMIND/pkg/db/mongodb"
	"context"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := config.GetConfig()

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	conn := mongodb.ParamsToConnect{
		User:     cfg.MongoDB.Username,
		Host:     cfg.MongoDB.Host,
		Port:     cfg.MongoDB.Port,
		Password: cfg.MongoDB.Password,
		Database: cfg.MongoDB.Database,
		AuthDB:   cfg.MongoDB.AuthDB,
	}

	client, err := mongodb.NewMongo(ctx, conn)
	if err != nil {
		return err
	}

	repos := repository.NewRepository(client)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(JMIND.Server)
	go func() {
		if err := srv.Run(cfg.PORT, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Info("App Started")

	<-ctx.Done()

	stop()

	logrus.Info("shutting down gracefully")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err.Error())
	}

	if err = client.Client().Disconnect(ctx); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
	logrus.Info("App Shutting Down")

	return nil
}
