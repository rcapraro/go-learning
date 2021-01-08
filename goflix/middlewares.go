package main

import (
	"github.com/form3tech-oss/jwt-go"
	"log"
	"net/http"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
)

func logRequestMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%v] %v", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	}
}

func (s *server) loggedOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		j := jwtmiddleware.New(jwtmiddleware.Options{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte(JwtAppKey), nil
			},SigningMethod: jwt.SigningMethodHS256,
		})
		j.HandlerWithNext(w, r, next)
	}
}
