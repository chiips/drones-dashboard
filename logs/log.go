package logs

import (
	"os"

	log "github.com/sirupsen/logrus"
)

//Log is our logger type
type Log struct {
	*log.Logger
}

//NewLogger sets up logrus with the given filename
func NewLogger(filename string) (*Log, error) {

	logger := log.New()

	//currently no .env var so defaults to setting up for terminal output
	env := os.Getenv("environment")

	if env == "production" {
		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			log.Println("Could not open log file:", filename, err.Error())
			return nil, err
		}
		logger.SetOutput(f)
		logger.SetFormatter(&log.JSONFormatter{})
	} else {
		formatter := &log.TextFormatter{}
		formatter.TimestampFormat = "02-01-2006 15:04:05" //use timestamp instead of unique number
		formatter.FullTimestamp = true
		logger.SetFormatter(formatter)

	}

	return &Log{logger}, nil

}