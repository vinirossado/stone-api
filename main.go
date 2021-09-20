package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"stone.com/api/controllers"
	"stone.com/api/infrastructure"
)

func main() {
	r := gin.New()

	connection := os.Getenv("SQLSERVER_CONNECTION")
	if connection == "" {
		//TODO: Implementar a lib de Log
		os.Exit(1)
	}

	infrastructure.ConnectDatabase()
	c := controllers.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts/")
		{
			accounts.GET(":account_id/balance", c.GetBalance)
			accounts.GET("", c.GetAccounts)
			accounts.POST("", c.NewAccount)
		}

		login := v1.Group("/login/")
		{
			login.POST("", c.Login)
		}

		transfers := v1.Group("/transfers/")
		{
			transfers.GET("", c.GetTransfersByUser)
			transfers.POST("", c.Transfer)
		}
	}

	http.ListenAndServe(":8080", r)
}
