package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi"
	"github.com/jzherran/track-price-go/trackitem/register"
)

type (
	// Middleware struct to manage middlewares added in router http.
	Middleware func(http.Handler) http.Handler
)

// NewHTTPRouter init the router http instance.
func NewHTTPRouter(middlewares ...Middleware) chi.Router {
	router := chi.NewRouter()

	log.Printf("Registering %d middlewares into the http server", len(middlewares))

	for _, middleware := range middlewares {
		router.Use(middleware)
	}

	return router
}

// StartHTTPServer init http server.
func StartHTTPServer(ctx context.Context, router chi.Router) {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}
}

func main() {
	// create dependencies
	middleware := []Middleware{
		NewTimeoutMiddleware(),
		NewHTTPRecoverMiddleware(),
		NewRequestIDMiddleware(),
	}
	router := NewHTTPRouter(middleware...)
	/* ds, err := google.NewGSheetService()
	if err != nil {
		panic(err)
	}*/

	// Configure our application modules
	register.SetupModule(router, nil)

	// start the application
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-done
		cancel()
	}()
	StartHTTPServer(ctx, router)
}

/*

// List of items with urls
	data := make(map[string]string)

	data["nevera:alkosto"] = "https://www.alkosto.com/nevera-samsung-394-litros-negro-rt38k5992bs"

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("#product-price-88274 > span:nth-child(1) > span:nth-child(2)", func(e *colly.HTMLElement) {
		// Print link
		fmt.Printf("Price found: %q\n", e.Text)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	for _, url := range data {
		fmt.Println("Iterate")
		c.Visit(url)
	}

*/
