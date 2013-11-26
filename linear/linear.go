package linear

type Linear struct {
  params []float64
  samples [][]float64
  labels []float64
}

func New(params []float64, samples [][]float64, labels []float64) (Linear) {
  return Linear{params, samples, labels}
}

func dot(as []float64, bs []float64) (float64) {
  var sum float64
  sum = 0
  for i := range(as) {
    sum += as[i] * bs[i]
  }
  return sum
}

func (model Linear) Cost() (float64) {
  m := float64(len(model.labels))
  var accum float64
  accum = 0
  for i := range(model.samples) {
    err := dot(model.samples[i], model.params) - model.labels[i]
    err = err * err
    accum += err
  }
  return (1 / (2 * m)) * accum
}

func (model Linear) Params() ([]float64) {
  return model.params
}

func (model Linear) SetParams(params []float64) {
  model.params = params
}

func (model Linear) Gradient(delta int) (float64) {
  m :=  float64(len(model.labels))
  var accum float64
  accum = 0
  for i := range(model.samples) {
    err := dot(model.samples[i], model.params) - model.labels[i]
    err *= model.samples[i][delta]
    accum += err
  }
  return (1 / m) * accum
}
