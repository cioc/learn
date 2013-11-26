package main

import (
  "github.com/cioc/learn/input"
  "github.com/cioc/learn/model/linear"
  "github.com/cioc/learn/optimize"
  "fmt"
)

func main() {
  features, vals, err := input.Read()
  if err != nil {
    fmt.Println(err)
  }
  for i := range(features) {
    fmt.Print(i, features[i])
    fmt.Print(" - ")
    fmt.Println(vals[i])
  }
  //TODO - NUMBER # PARAMS MUST MATCH # FEATURES
  lin := linear.New([]float64{0.0,0.0}, features, vals)
  for i := 0; i < 1500; i++ {
    fmt.Println(lin.Cost())
    params2 := optimize.GradientDescent(lin, 0.01)
    fmt.Println(params2)
    lin.SetParams(params2)
  }
}
