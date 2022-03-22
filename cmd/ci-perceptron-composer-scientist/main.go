package main

import (
	"ci-perceptron-composer-scientist/internal/perceptron"
	"ci-perceptron-composer-scientist/pkg/log"
	"github.com/sirupsen/logrus"
)

func init() {
	log.InitLog(logrus.InfoLevel)
}

func main() {
	p := perceptron.NewInstance(2, 1)

	p.AddTest([]int{0, 0}, 0)
	p.AddTest([]int{0, 1}, 0)
	p.AddTest([]int{1, 0}, 1)
	p.AddTest([]int{1, 1}, 1)

	p.Train()
}
