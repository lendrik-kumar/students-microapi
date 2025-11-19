package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/lendrik-kumar/Go-students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welecome to api"))
	})

	router.HandleFunc("POST /hello", func(w http.ResponseWriter, r *http.Request) {
		
	})

	fmt.Printf("server listining on %s", cfg.HttpServer.Addr)
	server := http.Server{
		Addr:    cfg.HttpServer.Addr,
		Handler: router,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal("failed to start server")
	}

}
