package handlers

import (
	"change_money/dto"
	"change_money/errs"
	"change_money/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return UserHandler{
		service: service,
	}
}

func (h UserHandler) ExchangeHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		network1 := ctx.Param("network1")
		network2 := ctx.Param("network2")
		UserResquestExchange := dto.UserResquestExchange{}
		err := ctx.BindJSON(&UserResquestExchange)
		UserResquestExchange.Network1 = network1
		UserResquestExchange.Network2 = network2

		if UserResquestExchange.Value < 10000000000000 {
			WriteError(ctx, errs.BadRequestError("value must be bigger than 10000000000000"))
			return
		}

		if len(UserResquestExchange.PrivateKey) > 64 || len(UserResquestExchange.PrivateKey) < 64 {
			WriteError(ctx, errs.BadRequestError("PrivateKey wrong"))
			return
		}

		if len(UserResquestExchange.ToAddress) > 42 || len(UserResquestExchange.ToAddress) < 42 {
			WriteError(ctx, errs.BadRequestError("Public Address wrong"))
			return
		}

		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}
		response, e := h.service.ExchangeService(UserResquestExchange)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, response)
	}
}

// func (h UserHandler) SendETHToContract() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		UserResquestExchange := dto.UserResquestExchange{}
// 		err := ctx.BindJSON(&UserResquestExchange)
// 		if err != nil {

// 			WriteError(ctx, errs.ErrorReadRequestBody())
// 			return
// 		}
// 		response, e := h.service.SendETHToContract(UserResquestExchange)
// 		if err != nil {

// 			WriteError(ctx, e)
// 			return
// 		}
// 		WriteRespon(ctx, http.StatusOK, response)
// 	}
// }

// func (h UserHandler) SendBSCToContract() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		UserResquestExchange := dto.UserResquestExchange{}
// 		err := ctx.BindJSON(&UserResquestExchange)
// 		if err != nil {

// 			WriteError(ctx, errs.ErrorReadRequestBody())
// 			return
// 		}
// 		response, e := h.service.SendBSCToContract(UserResquestExchange)
// 		if err != nil {

// 			WriteError(ctx, e)
// 			return
// 		}
// 		WriteRespon(ctx, http.StatusOK, response)
// 	}
// }

func (h UserHandler) ReceiveHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		network1 := ctx.Param("network1")
		network2 := ctx.Param("network2")
		UserResquestReceive := dto.UserResquestReceive{}
		UserResquestReceive.Network1 = network1
		UserResquestReceive.Network2 = network2
		err := ctx.BindJSON(&UserResquestReceive)
		if err != nil {

			WriteError(ctx, errs.ErrorReadRequestBody())
			return
		}

		if len(UserResquestReceive.PrivateKey) > 64 || len(UserResquestReceive.PrivateKey) < 64 {
			WriteError(ctx, errs.BadRequestError("PrivateKey wrong"))
			return
		}

		if len(UserResquestReceive.PublicKey) > 42 || len(UserResquestReceive.PublicKey) < 42 {
			WriteError(ctx, errs.BadRequestError("Public Address wrong"))
			return
		}

		response, e := h.service.ReceiveService(UserResquestReceive)
		if e != nil {

			WriteError(ctx, e)
			return
		}
		WriteRespon(ctx, http.StatusOK, response)
	}
}

// func (h UserHandler) ReceiveBSC() gin.HandlerFunc {
// 	return func(ctx *gin.Context) {
// 		UserResquestReceive := dto.UserResquestReceive{}
// 		err := ctx.BindJSON(&UserResquestReceive)
// 		if err != nil {

// 			WriteError(ctx, errs.ErrorReadRequestBody())
// 			return
// 		}
// 		response, e := h.service.ReceiveBSC(UserResquestReceive)
// 		if err != nil {

// 			WriteError(ctx, e)
// 			return
// 		}
// 		WriteRespon(ctx, http.StatusOK, response)
// 	}
// }
