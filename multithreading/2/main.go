package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	// "sync"
)

// teste utilizando apache benchmark
// sudo apt update
// sudo apt install apache2-utils
// ab -V
// ab -n 10000 -c 100 http://localhost:3000/

var number uint64 = 0

func main() {
	// m := sync.Mutex{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		// number++
		// m.Unlock()

		atomic.AddUint64(&number, 1)

		w.Write([]byte(fmt.Sprintf("vc teve acesso a essa página %d vezes", number)))
	})
	http.ListenAndServe(":3000", nil)
}
