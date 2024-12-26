package main

import (
	"binary_horizon/db"
	"binary_horizon/router"
)

func main() {
	db.InitPostgresDB()
	router.InitRouter().Run()
}
