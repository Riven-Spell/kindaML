package kindaml

import "math"

type ActivationF func(input float64) (output float64)

func HeavysideStep(input float64) float64 {
	if input == 0 {
		return .5
	} else if input > 0 {
		return 1
	} else {
		return 0
	}
}

func SigNum(input float64) float64 {
	if input == 0 {
		return 0
	} else if input > 0 {
		return 1
	} else {
		return -1
	}
}

func Linear(input float64) float64 {
	return input
}

func PiecewiseLinear(input float64) float64 {
	if input >= .5 {
		return 1
	} else if input <= -.5 {
		return 0
	} else {
		return input + .5
	}
}

func Sigmoid(input float64) float64 {
	return 1 / (1 + math.Pow(math.E, -input))
}

func HyperbolicTangent(input float64) float64 {
	return (math.Pow(math.E, input) - math.Pow(math.E, -input)) / (math.Pow(math.E, input) + math.Pow(math.E, -input))
}

func ReLU(input float64) float64 {
	return math.Max(0, input)
}

func SoftPlus(input float64) float64 {
	return math.Log(1 + math.Pow(math.E, input))
}
