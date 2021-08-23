package main

import (
	"go-training-restful/config"
	"go-training-restful/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	e.Logger.Fatal(e.Start(":8000"))
}
