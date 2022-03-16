package main

import (
	"ci-perceptron-composer-scientist/pkg/log"
	"github.com/sirupsen/logrus"
)

func init() {
	log.InitLog(logrus.InfoLevel)
}

func main() {
	logrus.Info("Initial commit")
}
