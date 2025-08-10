package main

import (
	"log"
	"net/http"

	"github.com/Thi0x40go/gin-postgres/internal/routers"
)

func main() {
	routers.RegisterUserRoutes()

	log.Println("Servidor iniciado na porta :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
