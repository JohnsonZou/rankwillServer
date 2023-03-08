package main

import (
	_ "github.com/PuerkitoBio/goquery"
	_ "github.com/mitchellh/mapstructure"
	"rankwillServer/server"

	_ "rankwillServer/server"
)

func main() {
	server.GinRun()
}
