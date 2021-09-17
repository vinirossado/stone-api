package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	connection := os.Getenv("SQLSERVER_CONNECTION")

	if connection == "" {
		//TODO: Implementar a lib de Log
		os.Exit(1)
	}

	http.ListenAndServe(":8080", r)
}
