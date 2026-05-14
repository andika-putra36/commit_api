package main

import (
	"commit_api/internal/config"
)

func main() {
	router := config.InitializeEverything()
	router.Run(":8888")
}
