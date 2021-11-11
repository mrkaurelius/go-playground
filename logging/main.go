package main

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func initLogger() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.TraceLevel)

	customFormatter := new(log.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	log.SetFormatter(customFormatter)
}

func main() {
	fmt.Println("Merhaba yalan dunya")

	// bunun seviyesi ne
	log.Print("Log deneme")
	log.Info("Log info deneme")
	log.Debug("Log Debug deneme")
	log.Warn("Log Warn deneme")
	log.Fatal("Log fatality")
}
