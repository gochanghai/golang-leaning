package main

import (
	"golang-leaning/go-oss/router"
)

func main() {
	router := router.InitRouter()
	router.Run(":8080")
}
