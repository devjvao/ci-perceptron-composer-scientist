package perceptron

import (
	"ci-perceptron-composer-scientist/pkg/log"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

type Test struct {
	Inputs         []int
	ExpectedOutput int
}

type Input struct {
	Value  int
	Weight float32
}

func (i *Input) Multiply() float32 {
	return float32(i.Value) * i.Weight
}

func (i Input) CalculateNewWeight(errorValue int, learningRate float32) float32 {
	return i.Weight + (float32(errorValue) * learningRate * float32(i.Value))
}

type Model struct {
	Bias            Input
	Inputs          []*Input
	Tests           []Test
	LearningRate    float32
	ExpectedOutput  int
	IteratorCounter int
}

func NewInstance(inputNumber int, learningRate float32) Model {
	model := Model{
		LearningRate:    learningRate,
		Bias:            Input{},
		IteratorCounter: 0,
	}

	for i := 0; i < inputNumber; i++ {
		model.Inputs = append(model.Inputs, &Input{
			Value:  0,
			Weight: 0,
		})
	}

	return model
}

func (m *Model) SetNewInputValues(newValues []int, expectedOutput int) {
	for index, input := range m.Inputs {
		input.Value = newValues[index]
	}
	m.ExpectedOutput = expectedOutput
}

func (m *Model) Sum() int {
	result := m.Bias.Multiply()

	for _, input := range m.Inputs {
		result += input.Multiply()
	}

	if result > 0 {
		return 1
	}

	return 0
}

func (m *Model) GenerateNewWeights(errorValue int) {
	for _, input := range m.Inputs {
		input.Weight = input.CalculateNewWeight(errorValue, m.LearningRate)
	}
	m.PrintWeights("New weights")
}

func (m *Model) AddTest(inputs []int, expectedOutput int) {
	m.Tests = append(m.Tests, Test{
		Inputs:         inputs,
		ExpectedOutput: expectedOutput,
	})
}

func (m Model) PrintWeights(prefix string) {
	var weights []float32
	for _, input := range m.Inputs {
		weights = append(weights, input.Weight)
	}
	logrus.Info(fmt.Sprintf("%s: %v", prefix, weights))
}

func (m *Model) Train() {
	rand.Shuffle(len(m.Tests), func(i, j int) {
		m.Tests[i], m.Tests[j] = m.Tests[j], m.Tests[i]
	})

	for _, test := range m.Tests {
		m.IteratorCounter++
		m.SetNewInputValues(test.Inputs, test.ExpectedOutput)

		generatedOutput := m.Sum()

		logrus.Info(log.Separator)
		logrus.Info(fmt.Sprintf("Iteration %d: Inputs → %v\tExpected output: %d\tGenerated output: %d", m.IteratorCounter, test.Inputs, test.ExpectedOutput, generatedOutput))

		if generatedOutput != test.ExpectedOutput {
			errorValue := test.ExpectedOutput - generatedOutput
			logrus.Info(fmt.Sprintf("Error value: %d", errorValue))

			m.GenerateNewWeights(errorValue)
			m.Train()

			return
		}
	}

	logrus.Info(log.Separator)
	m.PrintWeights(fmt.Sprintf("Final weights in %d iterations", m.IteratorCounter))
	logrus.Info(log.Separator)
}
