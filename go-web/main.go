package main

import (
	orm "golang-leaning/go-web/init"
	"golang-leaning/go-web/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":8080")
}
