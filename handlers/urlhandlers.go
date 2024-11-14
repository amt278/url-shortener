package handlers

import(
	"fmt"
	"log"
    "net/http"
    "url_shortener/utils"
)

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	original := r.FormValue("url")
	if original == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortened, err := urlshortener.SaveURLToDB(original)
	if err != nil {
		http.Error(w, "Failed to shorten URL", http.StatusInternalServerError)
		return
	}

	log.Printf("shortened URL: %s", shortened)
	fmt.Fprintf(w, "shortened URL: %s", shortened)
}

func originalURLHandler(w http.ResponseWriter, r *http.Request) {
	shortened := r.URL.Path[len("/original/"):]
	if shortened == "" {
		http.Error(w, "shortened URL is required", http.StatusBadRequest)
		return
	}

	original, err := urlshortener.GetURL(shortened)
	if err != nil {
		http.Error(w, "URL Not Found", http.StatusNotFound)
		return
	}

	log.Printf("Original URL: %s\n", original)
	http.Redirect(w, r, original, http.StatusSeeOther)
}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		shortenURLHandler(w, r)
		return
	case http.MethodGet:
		originalURLHandler(w, r)
		return
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
}