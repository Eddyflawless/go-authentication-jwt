package helpers

import (
	"fmt"
	log "go-jwt/pkg/logger"
)

var LoggerInstance *log.Logger

func LoggerInit() {

	// level log.Level, branchName string, version string, outType string, entityOrId string)

	LoggerInstance = log.NewLogger(3, "dev-branch", "1.0.0", "", "tesing")
	fmt.Println("initialized logger")
}
