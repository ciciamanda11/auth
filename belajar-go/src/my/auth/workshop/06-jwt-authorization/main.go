package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Authorization menggunakan claim JWT
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: answer here

		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah
		expectedPassword, ok := users[creds.Username]
		if !ok || expectedPassword.Password != creds.Password {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("user credential invalid"))
			return
		}
		// TODO: answer here

		// Task: 1. Deklarasi expiry time untuk token jwt
		// 		 2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//       3. Expiry time menggunakan time millisecond
		expirationTime := time.Now().Add(5 * time.Minute)
		claims := &Claims{
			Username: creds.Username,
			Role:     expectedPassword.Role,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		// TODO: answer here

		// Task: Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		// TODO: answer here

		// Task: 1. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//       2. Buat return internal error ketika ada kesalahan ketika pembuatan JWT string
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// TODO: answer here

		// Task: Set token string kedalam cookie response
		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		// TODO: answer here
	})

	mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: answer here

		// Task: Ambil value dari cookie token
		tokenString := c.Value
		// TODO: answer here

		// Task: Deklarasi variable claim
		claims := &Claims{}
		// TODO: answer here

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		// TODO: answer here

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// TODO: answer here

		// Task: return unauthorized ketika role user tidak sesuai dengan role admin
		if claims.Role != "admin" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// TODO: answer here

		// Task: return data dalam claim, seperti username yang telah didefinisikan
		w.Write([]byte(fmt.Sprintf("Welcome Admin %s!", claims.Username)))
		// TODO: answer here
	})

	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		// TODO: answer here

		// Task: Ambil value dari cookie token
		tokenString := c.Value
		// TODO: answer here

		// Task: Deklarasi variable claim
		claims := &Claims{}
		// TODO: answer here

		// Task: 1. Parse JWT token ke dalam claim
		//       2. return unauthorized ketika ada kesalahan ketika parsing token
		//	     3. return bad request ketika field token tidak ada
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		// TODO: answer here

		// Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)
		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		// TODO: answer here

		// Task: return data dalam claim, seperti username yang telah didefinisikan
		w.Write([]byte(fmt.Sprintf("Welcome Admin %s!", claims.Username)))
		// TODO: answer here
	})

	return mux
}
