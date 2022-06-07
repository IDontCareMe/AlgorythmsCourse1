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
    {9, 2, 0},
    {10, 2, 1},
    {1025, 55, 5},
    {12589, 369, 89},
    {1598753, 25897, 20305},
    {60282445765134413, 2263, 974 },
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