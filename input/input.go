package input

import (
  "bufio"
  "log"
  "os"
  "strings"
  "errors"
  "strconv"
)

//this is no good - lets deal with this later
const MaxSample int = 10000

func extractSample(l string) ([]float64, float64, error) {
  separated := strings.Split(l, "|")
  if len(separated) != 2 {
    return nil, 0, errors.New("incorrectly formatted sample")
  }
  featureStrings := strings.Split(separated[0], ",")
  features := make([]float64, len(featureStrings) + 1, len(featureStrings) + 1)
  features[0] = 1 //bias
  for i, s := range(featureStrings) {
    f, err := strconv.ParseFloat(s, 64)
    if err != nil {
      return nil, 0, errors.New("incorrectly formatted feature")
    }
    features[i + 1] = f
  }
  val, err := strconv.ParseFloat(separated[1], 64)
  return features, val, err
}

func Read() ([][]float64, []float64, error) {
  scanner := bufio.NewScanner(os.Stdin)
  var samples [][]float64
  var vals []float64
  samples = nil
  vals = nil
  for scanner.Scan() {
    features, val, err := extractSample(scanner.Text())
    if err != nil {
      return nil,nil,err
    }
    if samples == nil {
      samples = make([][]float64, 0, MaxSample)
      vals = make([]float64, 0, MaxSample)
    }
    samples = append(samples, features)
    vals = append(vals, val)
  }
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
  return samples, vals, nil
}
