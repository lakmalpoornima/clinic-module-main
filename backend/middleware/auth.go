package middleware

import (
	"crypto/rand"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
)

type JWT struct {
	TokenAuth *jwtauth.JWTAuth
}

var Jwt JWT

func init() {
	token := make([]byte, 32)
	rand.Read(token)
	log.Printf("Server Token: %x \n", token)
	Jwt.TokenAuth = jwtauth.New("HS256", token, nil)
}

func EncodeJWT(id string, location string) string {
	expiration := (24 * time.Hour)
	claims := map[string]interface{}{"user_id": id, "location": location, "exp": jwtauth.ExpireIn(expiration)}
	_, tokenString, _ := Jwt.TokenAuth.Encode(claims)
	return tokenString
}

func DecodeJWT(tokenClaim string, r *http.Request) interface{} {
	_, claims, _ := jwtauth.FromContext(r.Context())
	userid := claims[tokenClaim]
	return userid
}
