package main

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/smgt/drone-datocms-callback/plugin"
)

func main() {
	logrus.SetFormatter(new(formatter))

	var args plugin.Args
	if err := envconfig.Process("", &args); err != nil {
		logrus.Fatalln(err)
	}

	switch args.Level {
	case "debug":
		logrus.SetFormatter(textFormatter)
		logrus.SetLevel(logrus.DebugLevel)
	case "trace":
		logrus.SetFormatter(textFormatter)
		logrus.SetLevel(logrus.TraceLevel)
	}

	if err := plugin.Exec(context.Background(), args); err != nil {
		logrus.Fatalln(err)
	}
}

// default formatter that writes logs without including timestamp
// or level information.
type formatter struct{}

func (*formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

// text formatter that writes logs with level information
var textFormatter = &logrus.TextFormatter{
	DisableTimestamp: true,
}
