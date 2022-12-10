package logger

import (
	"encoding/json"
	"os"

	"fmt"

	log "github.com/sirupsen/logrus"
)

type Logger struct {
	branchName *string
	version    *string
	outType    string
	entity     *string
}

var Log *log.Entry

func NewLogger(level log.Level, branchName string, version string, outType string, entityOrId string) *Logger {

	// environment := os.Getenv("GO_ENVIRONMENT")

	log.SetFormatter(&log.TextFormatter{
		ForceColors: true,
	})

	// if environment == "production" || environment == "prod" {
	// 	log.SetFormatter(&log.TextFormatter{
	// 		ForceColors: true,
	// 	})
	// } else {
	// 	log.SetFormatter(&log.JSONFormatter{})
	// }

	log.SetReportCaller(true)

	// log.SetLevel(level)

	log.SetOutput(os.Stdout)

	l := &Logger{
		branchName: &branchName,
		version:    &version,
		outType:    outType,
		entity:     &entityOrId,
	}

	// Log = ParseLog(l)

	return l

}

func (l *Logger) MError(msg string, err error, data interface{}) {

	addError(err)
	Log.Error(msg)

}

func (l *Logger) Info(msg string, data interface{}) {

	// var err error

	if data != nil {
		msg, _ = createCustomMessage(msg, data)
	}

	ParseLog(l)
	fmt.Printf("should log message %v \n", msg)

	log.Info(msg)
}

func (l *Logger) Warn(msg string, err error, data interface{}) {

	addError(err)
	log.Error(msg)

}

func (l *Logger) Debug(msg string, data interface{}) {

	Log.Debug(msg)
}

// func ParseLog(l *Logger) *log.Entry {

// 	return log.WithFields(
// 		log.Fields{
// 			"branchName": *l.branchName,
// 			"version":    *l.version,
// 			"entity":     l.entity,
// 		},
// 	)

// }

func ParseLog(l *Logger) {

	log.WithFields(
		log.Fields{
			"branchName": *l.branchName,
			"version":    *l.version,
			"entity":     l.entity,
		},
	)

}

func createCustomMessage(msg string, data interface{}) (string, error) {

	var err error

	if data != nil {

		data, err = json.Marshal(data)

		if err != nil {
			fmt.Printf("Error: %v", err)
			return "", err
		}

		msg = fmt.Sprintf("%v | meta: %v", msg, data)

	}

	return msg, err
}

func addError(err error) {

	if err != nil {

		Log.WithFields(
			log.Fields{
				"error": err.Error(),
			},
		)
	}
}
