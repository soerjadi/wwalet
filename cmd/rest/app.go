package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/soerjadi/wwalet/internal/config"
	"github.com/soerjadi/wwalet/internal/delivery/rest"
	trxHdl "github.com/soerjadi/wwalet/internal/delivery/rest/transaction"
	userHdl "github.com/soerjadi/wwalet/internal/delivery/rest/user"
	"github.com/soerjadi/wwalet/internal/repository/transaction"
	"github.com/soerjadi/wwalet/internal/repository/user"
	trxUcs "github.com/soerjadi/wwalet/internal/usecase/transaction"
	userUcs "github.com/soerjadi/wwalet/internal/usecase/user"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("[Config] error reading config from file. err = %v", err)
		return
	}

	// open database connection
	dataSource := fmt.Sprintf("user=%s password=%s	host=%s port=%s dbname=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sqlx.Open(cfg.Database.Driver, dataSource)
	if err != nil {
		fmt.Printf("cannot connect to db. err = %v", err)
		return
	}

	handlers, err := initiateHandler(cfg, db)
	if err != nil {
		fmt.Printf("unable to initiate handler. err = %v", err)
		return
	}

	r := mux.NewRouter()
	rest.RegisterHandlers(r, handlers...)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	fmt.Printf("Server running in port : %s", cfg.Server.Port)

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Errorf("error running apps. err = %v", err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait  for.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.WaitTimeout())
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	fmt.Println("shutting down")
	os.Exit(0)
}

func initiateHandler(cfg *config.Config, db *sqlx.DB) ([]rest.API, error) {
	userRepository, err := user.GetRepository(db)
	if err != nil {
		return nil, fmt.Errorf("unable to initiate user repository. err : %v", err)
	}
	transactionRepository, err := transaction.GetRepository(db)
	if err != nil {
		return nil, fmt.Errorf("unable to initiate trx repository. err : %v", err)
	}

	userUsecase := userUcs.GetUsecase(userRepository, cfg)
	transactionUsecase := trxUcs.GetUsecase((transactionRepository))

	userHandler := userHdl.NewHandler(userUsecase, cfg)
	transactionHandler := trxHdl.NewHandler(transactionUsecase, userUsecase, cfg)

	return []rest.API{
		userHandler,
		transactionHandler,
	}, nil
}
