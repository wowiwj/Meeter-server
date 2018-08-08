package main

import (
	"Meeter/router"
	orm "Meeter/database"
)

func main() {

	router := router.Init()
	defer orm.DB.Close()
	router.Run(":8085")
}
