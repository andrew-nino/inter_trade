package main

import "international_trade/internal/app"

const configPath = "config/config.yaml"

func main() {
	app.Run(configPath)
}