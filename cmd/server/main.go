package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/lautarojayat/e_shop/api/http/users"
	"github.com/lautarojayat/e_shop/config"
	"github.com/lautarojayat/e_shop/logger"
	"github.com/lautarojayat/e_shop/persistence/files"
	"github.com/lautarojayat/e_shop/server"
)

func listenAndServe(s *http.Server) {
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("error while serving http. error=%q", err)
	}
}

func main() {
	l := logger.New()

	cfgFile, err := files.OpenFile("config/default.yaml")
	l.Println("about to read configs")

	if err != nil {
		l.Fatalf("fatal: couldn't read cfg file. error=%q", err)
	}

	l.Println("about to generate configs object")
	cfg, err := config.FromYAML(cfgFile)

	if err != nil {
		l.Fatalf("fatal: couldn't create config object. error=%q", err)
	}

	l.Println("about to generate http endpoints")
	mux := api.MakeHTTPEndpoints(l)

	l.Println("about to listen and serve")
	s := server.NewServer(cfg, mux)
	go listenAndServe(s)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigs
	log.Printf("received signal %q", sig)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(15)*time.Second)

	defer func() {
		cancel()
	}()

	err = s.Shutdown(ctx)

	if err != nil {
		log.Fatalf("fatal: error while shutting down. error=%q", err)
	}

	log.Println("Terminating process")

}
