package main

import (
	"github.com/Irvanlas/task-5-pbi-btpns-irvan/route"
)

func main() {
	r := gin.Default()
	models.ConnectDatabase()
	r.Run()
}
