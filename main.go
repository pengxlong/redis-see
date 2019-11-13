package main

import (
	"admin-redis/app"
)

func main() {
	app := app.NewApp("redis-manager")
	app.Run()
}