package main

import (
	"logging/route"
)

func main() {
	route := route.NewRoute()
	route.Run(":3000")
}
