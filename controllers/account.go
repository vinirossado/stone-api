package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"stone.com/api/application"
)

type Account struct {
	ap application.AccountAppInterface
}

func NewAccount(ap application.AccountAppInterface) *Account {
	return &Account{
		ap: ap,
	}
}

func (a *Account) GetAccounts(ctx *gin.Context) {

	accounts, err := a.ap.GetAccounts()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": accounts,
	})
}

func (a *Account) GetBalance(ctx *gin.Context) {
	accountId, err := strconv.ParseUint(ctx.Param("account_id"), 10, 64)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"data": err.Error(),
		})
		return

	}
	account, err := a.ap.GetBalance((accountId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"data": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": account,
	})

}

func (c *Account) NewAccount(ctx *gin.Context) {

}
