package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/nikitamirzani323/isbpanel_movie2/routers"
)

func main() {
	var errenv = godotenv.Load()
	if errenv != nil {
		panic("Failed to load env file")
	}
	app := routers.Init()
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "5000"
		}
		if err := app.Listen(":" + port); err != nil {
			log.Panic(err)
		}
	}()
	c := make(chan os.Signal, 1)                    // Create channel to signify a signal being sent
	signal.Notify(c, os.Interrupt, syscall.SIGTERM) // When an interrupt or termination signal is sent, notify the channel

	_ = <-c // This blocks the main thread until an interrupt is received
	log.Println("Gracefully shutting down...")
	_ = app.Shutdown()

	log.Println("Running cleanup tasks...")

	log.Println("Fiber was successful shutdown.")
}