package main

import (
	"vitica/data"
	"vitica/web"
)

func main() {
	data.DB()
	web.StartWebServer()
}
