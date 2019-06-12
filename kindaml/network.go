package kindaml

import "errors"

type Network struct {
	Outputs []*Input
	Inputs  []*Input
	Layers  [][]*Input
}

type LayerDescriptor struct {
	AFunc       ActivationF
	Size        int
	ExtraInputs []*Input //I *think* this will be useful for a RNN.
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
		outputs[k], err = (*n.Outputs[k]).Activate()

		if err != nil {
			break
		}
	}

	return
}

func InitNetwork(inputCount int, Layers []LayerDescriptor) Network {
	n := Network{}

	ins := make(map[*Input]bool)

	n.Inputs = make([]*Input, inputCount)
	for k := range n.Inputs {
		var v Input = Value(0)
		n.Inputs[k] = &v
		ins[&v] = true
	}

	n.Layers = make([][]*Input, len(Layers))
	for k, v := range Layers {
		n.Layers[k] = make([]*Input, v.Size)
		n.Layers[k-1] = append(n.Layers[k-1], v.ExtraInputs...)

		for x := range n.Layers[k] {
			var i Input = Node{
				Activation: v.AFunc,
				Biases:     make([]float64, len(n.Layers[k-1])),
				Weights:    make([]float64, len(n.Layers[k-1])),
				Inputs:     n.Layers[k-1],
				memos:      make(map[float64]float64),
			}

			n.Layers[k][x] = &i
		}
	}

	n.Outputs = n.Layers[len(n.Layers)-1]

	return n
}
