package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my kickass webapp</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch email me at <a href=\"mailto:josh@mail.com\">josh@mail.com</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	pathParam := chi.URLParam(r, "slug")
	if pathParam != "" {
		fmt.Fprintf(w, "<p>This is the article path param: %s</p>", pathParam)

	}
	queryParam := r.URL.Query().Get("id")
	if queryParam != "" {
		fmt.Fprintf(w, "<p>This is the query param: %s</p>", queryParam)
	}

	fmt.Fprint(w, "<h1>Welcome to the FAQs page")
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faqs", faqHandler)
	r.Get("/faqs/{slug}", faqHandler)

	// page not found
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
