package main

import (
	"fmt"
	"net/http"

	"github.com/Asri-Sirait/Golang-Session-Login-and-Register-with-MySQL-/tree/main/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
