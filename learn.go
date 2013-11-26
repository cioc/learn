package main

import (
  "github.com/cioc/learn/input"
  "github.com/cioc/learn/model/linear"
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
  lin := linear.New([]float64{1.0,1.0,1.0}, features, vals)
  fmt.Println(lin.Cost())
}
