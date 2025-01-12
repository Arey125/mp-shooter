package main

import (
	"fmt"
	"mp-shooter/internal/game"
	"net/http"
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
    port, err := strconv.Atoi(os.Getenv("PORT"))
    if err != nil {
        port = 3000
    }

    fmt.Printf("Starting server on http://127.0.0.1:%d\n", port)

    mux := http.NewServeMux()
    fileServer := http.FileServer(http.Dir("./static"))
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

    game := game.Game{}
    game.RegisterRoutes(mux)

    server := http.Server{
        Addr: fmt.Sprintf(":%d", port),
        Handler: mux,
    }

    err = server.ListenAndServe();
    fmt.Println(err);
}
