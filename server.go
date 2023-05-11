package main

import (
	"e-commerce/db"
	"e-commerce/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
