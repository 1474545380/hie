package main

import (
	"hie/main/router"
)

func main() {
	r := router.Router()
	r.Run(":8000")
}
