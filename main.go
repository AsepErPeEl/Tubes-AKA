package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var tpl = template.Must(template.ParseFiles("templates/index.html"))

const repeat = 100000

func gcdIter(a, b int64) {
	for b != 0 {
		a, b = b, a%b
	}
}

func gcdRec(a, b int64) {
	if b == 0 {
		return
	}
	gcdRec(b, a%b)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func manualHandler(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.ParseInt(r.URL.Query().Get("a"), 10, 64)
	b, _ := strconv.ParseInt(r.URL.Query().Get("b"), 10, 64)

	start := time.Now()
	for i := 0; i < repeat; i++ {
		gcdIter(a, b)
	}
	iterMs := float64(time.Since(start).Nanoseconds()) / 1e6 / repeat

	start = time.Now()
	for i := 0; i < repeat; i++ {
		gcdRec(a, b)
	}
	recMs := float64(time.Since(start).Nanoseconds()) / 1e6 / repeat

	json.NewEncoder(w).Encode(map[string]any{
		"a":       a,
		"b":       b,
		"iter_ms": iterMs,
		"rec_ms":  recMs,
	})
}

func autoHandler(w http.ResponseWriter, r *http.Request) {
	n, _ := strconv.ParseInt(r.URL.Query().Get("n"), 10, 64)

	a := n
	b := n - 1

	start := time.Now()
	for i := 0; i < repeat; i++ {
		gcdIter(a, b)
	}
	iterMs := float64(time.Since(start).Nanoseconds()) / 1e6 / repeat

	start = time.Now()
	for i := 0; i < repeat; i++ {
		gcdRec(a, b)
	}
	recMs := float64(time.Since(start).Nanoseconds()) / 1e6 / repeat

	json.NewEncoder(w).Encode(map[string]any{
		"n":       n,
		"iter_ms": iterMs,
		"rec_ms":  recMs,
	})
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/api/manual", manualHandler)
	http.HandleFunc("/api/auto", autoHandler)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
