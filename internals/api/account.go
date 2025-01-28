package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sulabhghimire/bank-api/internals/db/sqlc"
	"github.com/sulabhghimire/bank-api/internals/util"
)

type createAccountRequest struct {
	Owner    string `db:"owner" binding:"required"`
	Currency string `db:"currency" binding:"required,currency"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var body createAccountRequest
	if err := ctx.ShouldBind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    body.Owner,
		Currency: body.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	msg := "Account Created Successfully"
	ctx.JSON(http.StatusCreated, util.SuccessResponse(msg, &account))
	return
}

type GetAccountByIDParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccountByID(ctx *gin.Context) {
	var params GetAccountByIDParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, params.ID)
	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, util.ErrorResponse(err))
		} else {
			ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		}
		return
	}

	msg := "Account Fetched Successfully"
	ctx.JSON(http.StatusOK, util.SuccessResponse(msg, &account))
	return
}

type ListAccountsQueryParams struct {
	Page  int32 `form:"page" binding:"omitempty,min=1"`
	Limit int32 `form:"limit" binding:"omitempty,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var query ListAccountsQueryParams
	if err := ctx.ShouldBindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, util.ErrorResponse(err))
		return
	}

	if query.Page == 0 {
		query.Page = 1
	}
	if query.Limit == 0 {
		query.Limit = 5
	}

	arg := db.ListAccountsParams{
		Limit:  query.Limit,
		Offset: (query.Page - 1) * query.Limit,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, util.ErrorResponse(err))
		return
	}

	msg := "Accounts fetched successfully"
	ctx.JSON(http.StatusOK, util.SuccessResponse(msg, accounts))
	return
}
