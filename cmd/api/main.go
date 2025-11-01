package main

import (
    "fmt"
    "log"
    "net/http"
    
    "github.com/denerkrebs/WorkflowEngine/internal/infrastructure/router"
)

func main() {
    // Configurações
    port := ":8080"
    
    // Criar router
    r := router.NewRouter()
    
    // Logs de inicialização
    fmt.Println("Starting")
    fmt.Printf("Server running at http://localhost%s\n", port)
    
    // Iniciar servidor
    if err := http.ListenAndServe(port, r); err != nil {
        log.Fatal("Error while starting server:", err)
    }
}