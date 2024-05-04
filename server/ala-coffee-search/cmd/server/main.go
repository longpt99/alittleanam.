package main

import (
	"ala-coffee-search/configs"
	"ala-coffee-search/utils"
	"ala-coffee-search/utils/router"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	elasticsearch "github.com/elastic/go-elasticsearch/v8"
)

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

	// Init repo
	// repos := repositories.InitRepositories(store.PostgresConfig)

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Some error occurred when connect es. Err: %s", err)
	}

	// Create a channel to receive signals
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	r := router.New(http.NewServeMux())
	routeGlobal := r.Mount(fmt.Sprintf(`/%s`, configs.Env.GlobalPrefix))
	configs.RouterConfig(routeGlobal)

	routeGlobal.HandleFunc("GET /search", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())

		// q := r.URL.Query().Get("q")

		read, write := io.Pipe()

		go func() {
			defer write.Close()
			res, err := client.Search(
				client.Search.WithContext(r.Context()),
				client.Search.WithIndex("products"),
				// client.Search.WithBody(buildQuery(q)),
				// client.Search.WithTrackTotalHits(true),
			)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				defer res.Body.Close()
				io.Copy(write, res.Body)
			}
		}()
		io.Copy(w, read)
	})

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
