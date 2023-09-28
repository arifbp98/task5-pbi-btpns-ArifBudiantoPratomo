package main

import (
	"task5-pbi-btpns-ArifBudiantoPratomo/database"
	"task5-pbi-btpns-ArifBudiantoPratomo/router"
)


func main() {
	database.ConnectDB()

	router := router.SetupRouters()

	router.Run(":8088")
}
