package kindaml

import "errors"

type Node struct {
	Activation ActivationF
	Biases     []float64
	Weights    []float64
	Inputs     []Input //This neural network executes backwards, effectively.

	cachedValue *float64
	memos       map[float64]float64
}

func (n Node) Activate() (float64, error) {
	if len(n.Inputs) != len(n.Weights) || len(n.Weights) != len(n.Biases) { //All must be the same length.
		return 0, errors.New("number of inputs, weights, and biases must be same")
	}

	input := float64(0)

	for k, v := range n.Inputs {
		av, err := v.Activate()

		if err != nil {
			return 0, err
		}

		input += (av * n.Weights[k]) + n.Biases[k]
	}

	if v, ok := n.memos[input]; ok {
		return v, nil
	}

	x := n.Activation(input)
	n.cachedValue = &x
	return x, nil
}
