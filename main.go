package main

import (
	"online-judge/routers"
	"online-judge/models"
)

func main() {
	models.Initdb()
	router := routers.InitRouter()
	router.Run(":8080")
}
