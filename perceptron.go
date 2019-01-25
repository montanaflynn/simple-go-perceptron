package perceptron

type Perceptron struct {
	weights      []float64
	learningRate float64
	iterations   int
}

func (p *Perceptron) Predict(inputs []float64) float64 {
	activation := p.weights[0]
	for i := 0; i < len(inputs); i++ {
		activation += p.weights[i] * inputs[i]
	}
	if activation >= 0 {
		return 1.0
	}
	return 0.0
}

func (p *Perceptron) Fit(input [][]float64, target []int) {
	var it int
	var prediction, error float64
	for it < p.iterations {
		for i := 0; i < len(input); i++ {
			prediction = p.Predict(input[i])
			error = float64(target[i]) - prediction
			p.weights[0] += p.learningRate * error
			for j := 0; j < len(input[i]); j++ {
				p.weights[j] += p.learningRate * error * input[i][j]
			}
		}
		it++
	}
}

func New(size int, learningRate float64, iterations int) Perceptron {
	if learningRate == 0 {
		learningRate = 0.01
	}
	if iterations == 0 {
		iterations = 10
	}
	return Perceptron{
		weights:      make([]float64, size),
		learningRate: learningRate,
		iterations:   iterations,
	}
}
