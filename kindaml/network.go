package kindaml

import "errors"

type Network struct {
	Outputs []*Node
	Inputs  []*Input
	Layers  [][]*Input
}

func (n Network) Run(inputs []float64) (outputs []float64, err error) {
	outputs = make([]float64, len(n.Outputs))

	if len(inputs) != len(n.Inputs) {
		err = errors.New("number of inputs must be same as nn inputs")
		return
	}

	for k := range n.Inputs {
		*n.Inputs[k] = Value(inputs[k])
	}

	for k := range outputs {
		outputs[k], err = n.Outputs[k].Activate()

		if err != nil {
			break
		}
	}

	return
}
