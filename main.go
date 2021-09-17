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

	c := controllers.NewController()

	infrastructure.ConnectDatabase()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts/")
		{
			accounts.GET(":account_id/balance", c.GetSchool)
			accounts.GET("", c.GetSchools)
			accounts.POST("", c.GetSchools)
		}

		login := v1.Group("/login/")
		{
			login.GET(":id", c.GetCity)
			login.GET("", c.GetCities)
		}

		transfers := v1.Group("/transfers/")
		{
			transfers.GET("", c.GetCountries)
			transfers.POST("", c.GetCountries)
		}
	}

	http.ListenAndServe(":8080", r)
}
