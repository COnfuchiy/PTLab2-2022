package main

import "log"

func main() {
	app := NewApp()
	err := app.startApp()
	if err != nil {
		log.Fatalln(err)
	}
}
