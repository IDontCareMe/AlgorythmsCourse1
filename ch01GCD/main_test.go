package main

import (
  "fmt"
  "testing"
)

func TestEuclidGCD(t *testing.T) {
  var tests = []struct {
    a,b int
    want int
  } {
    {a:33, b:33, want:33},
    {a:27, b:12, want:3},
    {a:14159572, b:63967072, want:4},
  }

  for _, test := range tests {
    name := fmt.Sprintf("case%d%d:", test.a, test.b)
    t.Run(name, func(t *testing.T) {
      got := EuclidGCD(test.a, test.b)
      if got != test.want {
        t.Errorf("Error! Want: %d; got: %d", test.want, got)
      }
    })
  }
}