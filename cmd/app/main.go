package main

import "simple-golang-api/internal/app"

func main() {
	app.Run()

	//toDo: cache requests
	//toDo: add pagination for books methods in handler
	//toDo: work with context
	//toDo: fix dates parsing
	//toDo: add admin user in migrations (seeds)
}
