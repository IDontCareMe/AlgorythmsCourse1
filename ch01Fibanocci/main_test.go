package main

import (
  "testing"
  "fmt"
)

func TestFibanocci(t *testing.T) {
  var tests = []struct {
    n int
    want int
  } {
    {0, 0},
    {1, 1},
    {8, 21},
    {20, 6765},
  }
  for _,test := range tests {
    name := fmt.Sprintf("case %d", test.n)
    t.Run(name, func(t *testing.T) {
      got := fibanocci1(test.n)
      if got != test.want {
        t.Errorf("want:%d; got:%d", test.want, got)
      }
    })
  }
}

func TestFibanocci2(t *testing.T) {
  var tests = []struct {
    n int
    want int
  } {
    {0, 0},
    {1, 1},
    {8, 1},
    {20, 5},
    {841645, 5},
  }
  for _,test := range tests {
    name := fmt.Sprintf("case %d", test.n)
    t.Run(name, func(t *testing.T) {
      got := fibanocci2(test.n)
      if got != test.want {
        t.Errorf("want:%d; got:%d", test.want, got)
      }
    })
  }
}