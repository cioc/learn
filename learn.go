package main

import (
  "github.com/cioc/learn/input"
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
}
