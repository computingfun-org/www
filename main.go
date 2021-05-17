package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

const port = ":8080"

func main() {

	server := fiber.New(fiber.Config{
		IdleTimeout:      time.Minute,
		WriteTimeout:     time.Minute,
		ReadTimeout:      time.Minute,
		DisableKeepalive: true,
		Views:            html.New("./views", ".html"),
	})

	server.Get("/", func(c *fiber.Ctx) error {
		return nil
	})

	go func() {
		if err := server.Listen(port); err != nil {
			panic(err)
		}
	}()
	defer fmt.Println("🛑 server shutdown - error:", server.Shutdown())

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	s := <-signals
	fmt.Println("🏁 signal:", s)
}
