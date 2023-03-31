package main

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("my_secret_key")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Super Secret Information")
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", login)
	http.ListenAndServe(":8080", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	// 检验账号密码是否正确
	if username == "admin" && password == "password" {
		// 创建token
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["authorized"] = true
		claims["user"] = username

		// 设置过期时间为1小时
		tokenString, _ := token.SignedString(mySigningKey)

		// 返回token
		fmt.Fprintf(w, tokenString)
	} else {
		fmt.Fprintf(w, "Wrong username or password")
	}
}

func main() {
	// http://localhost:8080
	// http://localhost:8080/login?username=admin&password=password
	handleRequests()
}
