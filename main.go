package main

import (
	"fmt"

	route2 "github.com/mazin97/imersaofsfc2-simulator/application/route"
)

func main() {
	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringjson, _ := route.ExportJsonPositions()
	fmt.Println((stringjson[0]))
}
