package main

import(
	"fmt"
    "log"
    "net/http"
    "url_shortener/handlers"
    "url_shortener/db"
)

func main() {
	if err := db.InitDatabase(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	http.HandleFunc("/shorten", handlers.URLHandler)
	http.HandleFunc("/original/", handlers.URLHandler)

	const port = ":8080"

	fmt.Printf("Server running on port %s...\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}