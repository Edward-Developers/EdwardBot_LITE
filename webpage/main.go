package webpage

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello World! %s", time.Now())
	if err != nil {
		return
	}
}

func stats(w http.ResponseWriter, r *http.Request) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	_, err := fmt.Fprint(w, m.TotalAlloc/1024/1024)
	if err != nil {
		return
	}
}

func Start() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/stats", stats)
	err := http.ListenAndServe(":7070", nil)
	if err != nil {
		return
	}
}
