package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"net/http"
	"time"
)

var jwtKeys = []byte("my_secret_key")

var users = map[string]string{
	"lwgg": "12345678",
	"cjgg": "90abcdef",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Sign(w http.ResponseWriter, r *http.Request)  {
	var cred Credentials
	err := json.NewDecoder(r.Body).Decode(&cred)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	exceptedPassword,OK := users[cred.Username]
	if !OK || exceptedPassword != cred.Password{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: cred.Username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeys)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	})
}

func Welcome(w http.ResponseWriter,r *http.Request)  {
	c,err := r.Cookie("token")
	if err != nil{
		if err == http.ErrNoCookie{
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := c.Value
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error){
		return jwtKeys,nil
	})
	if err != nil{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if !token.Valid{
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.Write([]byte(fmt.Sprintf("Welcomes, %s!", claims.Username)))
}


func Refresh(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tknStr := c.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKeys, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeys)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}


func main()  {
	http.HandleFunc("/sign",Sign)
	http.HandleFunc("/welcome",Welcome)
	http.HandleFunc("/refresh",Refresh)
	log.Fatal(http.ListenAndServe(":8000",nil))
}
