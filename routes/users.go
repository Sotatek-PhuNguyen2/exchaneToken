package routes

import (
	"change_money/admin"
	"change_money/config"
	"change_money/handlers"
	"change_money/repo"
	"change_money/service"

	"github.com/gin-gonic/gin"
)

func ChangeMoneyRoute(router *gin.Engine) {
	route := router.Group("/api/change_money")
	admin := admin.NewAdmin()
	handlers := handlers.NewUserHandler(service.NewUserService(repo.NewUserRepo(config.DB, admin)))
	{
		route.POST("/exchange/:network1/:network2", handlers.ExchangeHandler())

		route.POST("/receive/:network1/:network2", handlers.ReceiveHandler())
	}
}
