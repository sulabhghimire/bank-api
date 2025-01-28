package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sulabhghimire/bank-api/internals/db/sqlc"
	"github.com/sulabhghimire/bank-api/internals/util"
)

type transferRequest struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,currency"`
}

func (server *Server) createTransfer(ctx *gin.Context) {
	var body transferRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if !server.validAccount(ctx, body.FromAccountID, body.Currency) {
		return
	}

	if !server.validAccount(ctx, body.ToAccountID, body.Currency) {
		return
	}

	arg := db.TransferTxParams{
		FromAccountID: body.FromAccountID,
		ToAccountID:   body.ToAccountID,
		Amount:        body.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	msg := "Amount transferred successfully"
	ctx.JSON(http.StatusCreated, util.SuccessResponse(msg, result.Transfer))
	return
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64, currency string) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return false
	}

	if account.Currency != currency {
		err := fmt.Errorf("account [%d] currency mismatch: %s vs %s", accountID, account.Currency, currency)
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return false

	}

	return true
}
