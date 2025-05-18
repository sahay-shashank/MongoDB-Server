package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/sahay-shashank/mongodb-server/internal/core/auth"
	"github.com/sahay-shashank/mongodb-server/internal/database"
	"github.com/sahay-shashank/mongodb-server/internal/web/router"
)

func Server() {
	databaseConnectionResult := database.InitDatabase()
	if databaseConnectionResult.Error {
		databaseConnectionResult.LogToStderr()
		os.Exit(1)
	}
	databaseConnectionResult.LogToStdout()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-quit
		databaseCloseResult := database.CloseDatabase()
		if databaseCloseResult.Error {
			databaseConnectionResult.LogToStderr()
			os.Exit(1)
		}
		databaseCloseResult.LogToStdout()
		log.Println("Server shutting down...")
		os.Exit(0)
	}()
	port, found := os.LookupEnv("PORT")
	var serverPort string
	if found && isInt(port) {
		serverPort = port
	} else {
		serverPort = "8080"
	}
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		panic("JWT_SECRET not set")
	}
	auth.SetJWTSecret(secret)
	svr := &http.Server{
		Addr:    fmt.Sprintf(":%v", serverPort),
		Handler: router.NewRouter(),
	}
	log.Printf("Server started on port %v", serverPort)
	if err := svr.ListenAndServe(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func isInt(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
