package main

import "international_trade/internal/app"

const configPath = "config/config.yaml"

//	@title			International Trade API
//	@version		1.0.0
//	@description	API Server for test work

//	@host		localhost:8080
//	@BasePath	/

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization
func main() {
	app.Run(configPath)
}
