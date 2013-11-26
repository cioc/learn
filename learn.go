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
  //TODO - CLEAN UP ALL THIS GARBAGE
  lin := linear.New([]float64{0.0,0.0,0.0}, features, vals)
  means, ranges := lin.ScaleFeatures()
  
  //fmt.Println(means)
  //fmt.Println(ranges)
  lin.PrintData()
  for i := 0; i < 100000; i++ {
    fmt.Println(lin.Cost())
    params2 := optimize.GradientDescent(lin, 0.01)
    fmt.Println(params2)
    lin.SetParams(params2)
  }
  finalParams := lin.Params()
  fmt.Println(finalParams)
  fmt.Println(means)
  fmt.Println(ranges)
  for i := range(finalParams) {
    fmt.Println(finalParams[i] / ranges[i])
  }
  fmt.Println(finalParams[0] + (finalParams[1] * ((1650.0 - means[1]) / ranges[1])) + (finalParams[2] * ((3.0 - means[2]) / ranges[2])))
}
