package app

import (
	"change_money/config"
	"change_money/event_logs"
	"change_money/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func SetUpApp() {
	config.InitDatabase()
	subBSC := event_logs.NewBSCSubscription(config.DB)
	subETh := event_logs.NewETHSubscription(config.DB)
	go subBSC.WatchEventContract()
	go subETh.WatchSendETHToContract()
	go subETh.WatchReceivedETHToWallet()
	router := gin.Default()
	routes.ChangeMoneyRoute(router)
	log.Println("Server is running on PORT ", ":8080")
	router.Run(config.ConnectPort())
}
