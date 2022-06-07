package main

import (
  "testing"
  "fmt"
)

func TestFibanocci3(t *testing.T) {
  var tests = []struct{
    n,m int
    want int
  } {
    {0, 17, 0},
    {20, 2, 1},
    {10, 2, 1},
    {6, 10, 8},
  }
  for _,test := range tests {
    name := fmt.Sprintf("case F(%d)%%F(%d):", test.n, test.m)
    t.Run(name, func(t *testing.T){
      got := fibanocci3(test.n, test.m)
      if got != test.want {
        t.Errorf("Error! want: %d; got : %d", test.want, got)
      }
    })
  }
}