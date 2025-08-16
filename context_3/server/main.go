package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("request iniciada")
	defer log.Println("request finalizada")

	select {
	case <-time.After(time.Second * 5):
		log.Println("request processada com sucesso")
		w.Write([]byte("request processada com sucesso"))
	case <-ctx.Done():
		log.Println("requisicao cancelada pelo client")
	}
}
