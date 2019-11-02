package main

import (
	orm "./init"
	"./router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	_ = router.Run(":8080")
}