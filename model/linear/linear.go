package linear

import (
  "fmt"
)

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
  for i := range(as) {
    sum += as[i] * bs[i]
  }
  return sum
}

func (model Linear) PrintData() {
  for i := range(model.samples) {
    fmt.Print(model.samples[i])
    fmt.Print(" - ")
    fmt.Println(model.labels[i])
  }
}

func (model Linear) ScaleFeatures() (means []float64, ranges []float64) {
  featureCount := len(model.params)
  sampleCount := len(model.labels)
  means = make([]float64, featureCount, featureCount)
  ranges = make([]float64, featureCount, featureCount)
  for i := 1; i < featureCount; i++ {
    //determine mean and range
    var mean float64
    max := model.samples[0][i]
    min := model.samples[0][i]
    for j := 0; j < sampleCount; j++ {
      mean += (1 / float64(sampleCount)) * model.samples[j][i]
      if model.samples[j][i] > max {
        max = model.samples[j][i]
      }
      if model.samples[j][i] < min {
        min = model.samples[j][i]
      }
    }
    for j := 0; j < sampleCount; j++ {
      model.samples[j][i] -= mean
      model.samples[j][i] /= (max - min)
    }
    means[i] = mean
    ranges[i] = (max - min)
  }
  return means, ranges
}

func (model Linear) Cost() (float64) {
  m := float64(len(model.labels))
  var accum float64
  for i := range(model.samples) {
    err := dot(model.samples[i], model.params) - model.labels[i]
    err *= err
    accum += (1 / (2 * m)) * err
  }
  return accum
}

func (model Linear) Params() ([]float64) {
  return model.params
}

func (model Linear) SetParams(params []float64) {
  copy(model.params, params)
}

func (model Linear) Gradient(delta int) (float64) {
  m := float64(len(model.labels))
  var accum float64
  for i := range(model.samples) {
    err := dot(model.samples[i], model.params) - model.labels[i]
    err *= (1 / m)
    err *= model.samples[i][delta]
    accum += err
  }
  return accum
}
