package main

import (
	"fmt"
	"github.com/pquerna/otp/totp"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Starting server on port :8787 ...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/totp/" {
			http.Redirect(w, r, "/totp", http.StatusFound)
			return
		} else if r.URL.Path != "/" {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		if r.Method == "GET" {
			html := `
			<!DOCTYPE html>
			<html>
			<head>
				<title>OTP Generator</title>
			</head>
			<body>
				<form action="/totp" method="GET">
					<label for="secret">Secret:</label>
					<input type="text" id="secret" name="secret">
					<input type="submit" value="Generate">
				</form>
			</body>
			</html>
			`
			w.Write([]byte(html))
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/totp", func(w http.ResponseWriter, r *http.Request) {
		secret := r.FormValue("secret")
		if secret == "" {
			http.Error(w, "secret is required", http.StatusBadRequest)
			return
		}

		otp, err := totp.GenerateCode(secret, time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(otp))
	})

	log.Fatal(http.ListenAndServe(":8787", nil))
}
