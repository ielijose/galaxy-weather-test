package main

import (
	"fmt"
	"galaxy-weather/api"
	_ "galaxy-weather/database"
	"os"
)

func main() {

	server := api.Init()

	var url = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	server.Logger.Fatal(server.Start(url))
}
