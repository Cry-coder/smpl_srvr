package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/Cry-coder/smpl_srvr/internal/domain/event"
	"github.com/Cry-coder/smpl_srvr/internal/infra/http"
	"github.com/Cry-coder/smpl_srvr/internal/infra/http/controllers"
)

// @title                       Test Server
// @version                     0.1.0
// @description                 Test Server boilerplate
func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	// Event
	eventRepository := event.NewRepository()                         // створення нового репозиторію для зберігання даних
	eventService := event.NewService(&eventRepository)               // створюємо сервіс який буде обробляти запити і передаємо туди репозиторій
	eventController := controllers.NewEventController(&eventService) // передаємо сервісб івентконтроллер обробляє вхіні запити (з інтернету)

	// HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			eventController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
	if err = event.Session.Close(); err != nil {
		log.Fatal("Error While closing db:", err)
	}
}
