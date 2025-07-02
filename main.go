package main

import (
	"log"
	"os"
	"runtime"
	"time"

	"app/internal/cmd"
	"core/config"
)

// @title			API Gateway Documentation
// @version		    1.0.0
// @description	    API Gateway documentation
// @host			localhost:3000
// @BasePath		/api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description     Provide your Bearer token in the format: Bearer <token>
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// always use UTC time
	time.Local = time.UTC

	wd, _ := os.Getwd()
	config.SetBasePath(wd)

	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
