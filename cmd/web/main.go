package main

import (
	"os"

	"github.com/bejaneps/s3-streaming/cmd/web/sub"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetOutput(os.Stdout)
	//log.SetReportCaller(true)

	if err := sub.Execute(); err != nil {
		log.Fatal(err)
	}
}
