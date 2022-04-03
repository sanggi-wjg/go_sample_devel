package main

import (
	"go_sample_devel/pkg/database"
	"go_sample_devel/route"
	"log"
)

func init() {
	if err := database.Setup(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	router := route.SetupRoutes()
	log.Fatalln(router.Run(":9091"))
}
