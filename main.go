package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"stone.com/api/controllers"
	"stone.com/api/infrastructure"
	"stone.com/api/repository"
)

func main() {
	r := gin.New()

	connection := os.Getenv("SQLSERVER_CONNECTION")
	if connection == "" {
		//TODO: Implementar a lib de Log
		os.Exit(1)
	}

	db := infrastructure.ConnectDatabase()

	account := controllers.NewAccount(repository.NewAccountRepository(db))

	r.GET("accounts", account.GetAccounts)
	r.GET("accounts/:account_id/balance", account.GetBalance)
	r.POST("", account.NewAccount)

	// r.POST("", c.Login)

	// r.GET("", c.GetTransfersByUser)
	// r.POST("", c.Transfer)

	http.ListenAndServe(":8080", r)
}
