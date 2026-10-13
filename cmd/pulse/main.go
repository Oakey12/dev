package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func main() {
	http.HandleFunc("/health", healthHandler)
	fmt.Println("Сервер запущен")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
