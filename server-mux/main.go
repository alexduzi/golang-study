package main

import "net/http"

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books/{id}", GetBookHandler)
	mux.HandleFunc("GET /books/dir/{d...}", GetBookHandler2)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Test"))
	})
	http.ListenAndServe(":8080", mux)
}

func GetBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	w.Write([]byte("Test " + id))
}

func GetBookHandler2(w http.ResponseWriter, r *http.Request) {
	d := r.PathValue("d")
	w.Write([]byte("Test " + d))
}
