package main

import (
	"fmt"
	"gosql/endpoints"

	"github.com/gin-gonic/gin"

	"gosql/config"
	"gosql/dsl"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	bind, port, dslPath, err := config.ParseArgs()

	if err != nil {
		println("Error: " + err.Error())
		return
	}

	dslConfig, err := dsl.ParseDSL(dslPath)

	if err != nil {
		println("Error: " + err.Error())
		return
	}

	fmt.Println("Collected DSL configuration")

	endpoints.MakeEndpoints(dslConfig, r)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	fullAddress := fmt.Sprintf("%s:%d", bind, port)

	var ginError = r.Run(fullAddress)
	if ginError != nil {
		println(ginError.Error())
	}
}
