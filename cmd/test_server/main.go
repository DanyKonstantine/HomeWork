package main

import (
	"HomeWork/internal/domain/event"
	"HomeWork/internal/domain/person"
	"HomeWork/internal/infa/http"
	"HomeWork/internal/infa/http/controllers"
	"context"
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
)

func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())
	var setings = postgresql.ConnectionURL{
		Database: `Event`,
		Host:     `127.0.0.1:5432`,
		User:     `postgres`,
		Password: `postgres`,
	}
	sess, errConn := postgresql.Open(setings)
	if errConn != nil {
		log.Fatal("Open: ", errConn)
	}
	defer sess.Close()
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
	eventRepository := event.NewRepository(sess)
	eventService := event.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)
	personRepository := person.NewRepository(sess)
	personService := person.NewService(&personRepository)
	personControler := controllers.NewPerconControler(&personService)

	// HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			eventController,
			personControler,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
