package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/hamzabaazaoui/simplebank/db/sqlc"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var CreateAccountRequest createAccountRequest
	err := ctx.ShouldBindJSON(&CreateAccountRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, errorResponse(err))
	}

	acc, err := server.store.CreateAccount(ctx, db.CreateAccountParams{
		Owner:    CreateAccountRequest.Owner,
		Currency: CreateAccountRequest.Currency,
		Balance:  0,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, acc)
}
