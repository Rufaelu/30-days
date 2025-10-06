package main

import (
	"30days/routes"
)

func main() {
	e := routes.NewRoute()

	e.Start(":8080")

}
