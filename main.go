package main

import (
	"log"

	DatabaseCon "github.com/jakkritscpe/rest-api-portfolio/database"
	"github.com/jakkritscpe/rest-api-portfolio/routers"

)

func main() {
	
	log.Println("-------------------------------------")
	log.Println("Start API portfolio ... let go !! 👋")
	log.Println("-------------------------------------")

	DatabaseCon.InitDB()

	log.Println("Is Running ....")

	r := routers.SetupRouter()
	r.Run() // listen and serve on 0.0.0.0:8080

}
