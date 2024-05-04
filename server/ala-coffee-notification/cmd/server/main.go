package main

import (
	"ala-coffee-notification/configs"
	externalservices "ala-coffee-notification/external-services"
	"ala-coffee-notification/models"
	"ala-coffee-notification/utils"
	"ala-coffee-notification/utils/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var Env models.Env

func main() {
	ctx := context.Background()

	// Load config
	err := configs.InitConfig()
	if err != nil {
		log.Fatalf("InitConfig error occurred. Err: %s", err)
	}

	// Load extend Validations
	err = utils.RegisterValidation()
	if err != nil {
		log.Fatalf("Some error occurred. Err: %s", err)
	}

	go externalservices.SetupConsumerGroup(ctx)

	// Init repo
	// repos := repositories.InitRepositories(store.PostgresConfig)

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	r := router.New(http.NewServeMux())
	routeGlobal := r.Mount(fmt.Sprintf(`/%s`, configs.Env.GlobalPrefix))
	configs.RouterConfig(routeGlobal)

	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", configs.Env.Port),
		Handler:           r,
		ReadHeaderTimeout: 0,
	}

	go func() {
		log.Printf("Server is listening on %s\n", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe error: %s", err)
		}
	}()

	// Shutdown the server gracefully
	s := <-signalCh
	if s != nil {
		if err := srv.Shutdown(ctx); err != nil {
			log.Printf("[Server] Shutdown Failed: %v\n", err)
			return
		}

		log.Println("[Server] Shutdown Gracefully! 🚀")
	}
}
