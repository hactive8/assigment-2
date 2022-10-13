package main

import "hactive/assigment-2/infrastructure/server"

func main() {
	app := server.Start()

	err := app.Listen(":8080")

	if err != nil {
		panic(err)
	}
}
