package main

import (
	"context"
	"fmt"
	"go-ddd-api/internal/config"
	"go-ddd-api/internal/infra/service"
	"go-ddd-api/internal/infra/store"
	"go-ddd-api/internal/interface/web"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()

	logger := logrus.New()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Port)

	errCh := make(chan error)

	db, err := sqlx.Open("postgres", cfg.PgConnURI())

	if err != nil {
		log.Fatal(err)
	}

	logger.Printf("Connected to database: %s", cfg.DbName)

	st := store.NewStore(db)

	srv := service.NewService(logger, st)

	api := web.NewAPI(cfg, logger, srv)

	shutdownCh := make(chan os.Signal, 1)

	signal.Notify(shutdownCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		errCh <- api.Run()
	}()

	select {
	case err := <-errCh:
		logger.Error(err)
	case <-shutdownCh:
		ctx, cancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeout)
		defer cancel()
		err := api.Shutdown(ctx)

		if err != nil {
			logger.Error(err)
			err = api.Close()
		}

		if err != nil {
			log.Fatal(err)
		}
	}

}
