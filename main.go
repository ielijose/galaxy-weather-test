package main

import (
	"galaxy-weather/cmd"
	_ "galaxy-weather/database"
)

func main() {
	cmd.Galaxy()

	/* server := api.Init()

	var url = fmt.Sprintf(":%s", os.Getenv("API_PORT"))

	server.Logger.Fatal(server.Start(url)) */
}
