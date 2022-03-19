package main

import (
	"go_sample_devel/route"
	"log"
)

func main() {
	router := route.SetupRoutes()
	log.Fatalln(router.Run(":9091"))
}

//YOUTUBE_KEY=AIzaSyAEaY6vSNjZ1VXeJYBW_8F8psYQq-X6Kkc;YOUTUBE_CHANNEL_ID=UCV9WL7sW6_KjanYkUUaIDfQ
