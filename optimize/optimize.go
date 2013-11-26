package optimize

import (
  "github.com/cioc/learn/model"
)

func GradientDescent(model model.Model, alpha float64) ([]float64) {
  params := model.Params()
  lp := len(params)
  output := make([]float64, lp, lp)
  for i := range(params) {
    output[i] = params[i] - alpha * model.Gradient(i)
  }
  return output
}
