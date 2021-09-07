package main

import (
	"fmt"
	CONFIG "go_image_processing_cli/config"
)

var environmentType = CONFIG.GetEnvByKey("ENVIRONMENT")

func main() {
	fmt.Printf ("Hello World + %s", environmentType)
}
