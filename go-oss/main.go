package main

import (
	"golang-leaning/go-oss/router"
	"log"
	"strconv"
	"time"
)

func main() {
	log.Println(".." + strconv.FormatInt(time.Now().Unix(), 10))
	router := router.InitRouter()
	router.Run(":8080")
}
