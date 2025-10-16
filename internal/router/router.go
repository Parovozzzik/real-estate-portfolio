package router

import (
    "net/http"

    "github.com/go-chi/chi/v5"
)

var validAPIKeys = map[string]bool{
	"mysecretapikey123": true,
	"anotherapikeyabc":  true,
}

func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key") // Common header for API keys

		if apiKey == "" {
			http.Error(w, "API Key is missing", http.StatusUnauthorized)
			return
		}

		if !validAPIKeys[apiKey] {
			http.Error(w, "Invalid API Key", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func GetRouter() *chi.Mux {
    r := chi.NewRouter()

    r.Use(APIKeyAuthMiddleware)

    r.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("hi"))
    })

    r.Mount("/users", usersRouter())
    r.Mount("/estate-types", estateTypesRouter())

    return r
}