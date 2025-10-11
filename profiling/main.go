package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
)

func HeavyComputation() {
	var result int
	for i := 0; i < 1e7; i++ {
		result += rand.Intn(100)
	}
}

func MemoryIntensiveOperation() {
	// var bytes []byte
	bytes := make([]byte, 0, 1e6) // -> initialize slice with the capacity

	for i := 0; i < 1e6; i++ {
		bytes = append(bytes, byte(rand.Intn(256)))
	}
}

func Handler(w http.ResponseWriter, r *http.Request) {
	HeavyComputation()
	MemoryIntensiveOperation()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("done!"))
}

func main() {
	http.HandleFunc("/", Handler)

	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	fmt.Println("server is running at port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
