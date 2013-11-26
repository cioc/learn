package model

type Model interface {
  Cost() float64
  Gradient(int) float64
  Params() []float64
  SetParams([]float64)
}
