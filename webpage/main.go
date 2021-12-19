package webpage

import (
	"fmt"
	"runtime"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func stats(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
    runtime.ReadMemStats(&m)
	fmt.Fprint(w, m.TotalAlloc / 1024 / 1024)
}

func Start() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/stats", stats)
	http.ListenAndServe(":7070", nil)
}