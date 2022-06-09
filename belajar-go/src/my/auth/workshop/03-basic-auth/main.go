package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	// encrypt token dari username menggunakan bcrypt kemudian kirim ke user kedalam cookie
	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		cookieName := "token"
		var creds Credentials

		// Task:  Buat JSON body diconvert menjadi credential struct & return bad request ketika terjadi kesalahan decoding json
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: answer here

		// Task: Ambil password dari username yang dipakai untuk login
		expectedPassword, ok := users[creds.Username]
		// TODO: answer here

		// Task: return unauthorized jika password salah
		if !ok || expectedPassword != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user credential invalid"))
			return
		}
		// TODO: answer here

		// Task: 1. Buat token string menggunakan bcrypt dari credential username
		// 		 2. return internal server error ketika terjadi kesalahan encrypting token
		tokenString, err := bcrypt.GenerateFromPassword([]byte(creds.Username), bcrypt.MinCost)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO: answer here

		// Task: Set token string kedalam cookie response
		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,
			Value:   string(tokenString),
			Expires: time.Now().Add(5 * time.Minute),
		})
		// TODO: answer here
	})

	return mux
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
