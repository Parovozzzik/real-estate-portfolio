package router

import (
    "context"
    "net/http"
    "strings"

    "github.com/go-chi/chi/v5"
    "github.com/golang-jwt/jwt/v5"

    "github.com/Parovozzzik/real-estate-portfolio/internal/config"
)

var validAPIKeys = map[string]bool{
	"mysecretapikey123": true,
	"anotherapikeyabc":  true,
}

func APIKeyAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-Key")

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

func JWTMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, `{"error": "Authorization header required"}`, http.StatusUnauthorized)
            return
        }

        // Формат: "Bearer {token}"
        parts := strings.Split(authHeader, " ")
        if len(parts) != 2 || parts[0] != "Bearer" {
            http.Error(w, `{"error": "Authorization format: Bearer {token}"}`, http.StatusUnauthorized)
            return
        }

        tokenString := parts[1]

        cfg := config.GetConfig()
        jwtSecret := []byte(cfg.JwtSecret)

        // Парсим и валидируем токен
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            // Проверяем алгоритм подписи
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, jwt.ErrSignatureInvalid
            }
            return jwtSecret, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, `{"error": "Invalid token"}`, http.StatusUnauthorized)
            return
        }

        // Извлекаем claims и добавляем в контекст
        if claims, ok := token.Claims.(jwt.MapClaims); ok {
            userID := claims["user_id"]
            ctx := context.WithValue(r.Context(), "userID", userID)
            next.ServeHTTP(w, r.WithContext(ctx))
        } else {
            http.Error(w, `{"error": "Invalid token claims"}`, http.StatusUnauthorized)
        }
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
    r.Mount("/estates", estatesRouter())

    return r
}