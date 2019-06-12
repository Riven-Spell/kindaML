package kindaml

type Input interface {
	Activate() (float64, error)
}

type Value float64

func (v Value) Activate() (float64, error) {
	return float64(v), nil
}
