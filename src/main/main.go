package main

import (
	"flag"
	"os"

	"github.com/google/logger"
)

const logPath = "../../log/example.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func main() {
	flag.Parse()

	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	logger.Init("LoggerExample", *verbose, true, lf)

	logger.Info("Info  I'm about to do something!")
	logger.Error("Error I'm about to do something!")
	logger.Fatal("Fatal I'm about to do something!")
}
