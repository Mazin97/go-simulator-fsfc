package main

import (
	"fmt"
	"log"

	dotenv "github.com/joho/godotenv"
	route2 "github.com/mazin97/imersaofsfc2-simulator/application/route"
)

func init() {
	err := dotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println((stringjson[0]))
}
