package main

import (
	"log"
	"net/http"

	"github.com/IbzanJared/sistema-reservas/cmd/server"
	"github.com/IbzanJared/sistema-reservas/internal/config"
)

func main() {
	log.Println("Cargando configuración...")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error cargando configuración: %v", err)
	}

	log.Println("Iniciando servidor en " + cfg.ServerAddress)

	r := server.NewRouter()

	log.Fatal(http.ListenAndServe(cfg.ServerAddress, r))
}
