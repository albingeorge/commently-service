package main

import "github.com/albingeorge/commently-service/app/router"

func main() {
	r := router.SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
