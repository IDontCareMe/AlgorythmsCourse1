package main

import (
  "fmt"
  "os"
)

func main() {
  // Read sorted
  sorted, err := readInputInt()
  if err != nil {
    printError(err)
  }
  // Read numbers
  numbers, err := readInputInt()
  if err != nil {
    printError(err)
  }
  // Search numbers in sorted
  for _,value := range numbers {
    fmt.Printf("%d ", search(value, sorted) + 1)
  }
}

func readInputInt()(slice []int, err error) {
  var n int
  _, err = fmt.Scan(&n)
  if err != nil {
    return
  }
  slice = make([]int, n)
  for i,_ := range slice {
    _, err = fmt.Scan(&n)
    if err != nil {
      return
    }
    slice[i] = n
  }
  return
}

func search(n int, slice []int)(index int) {
  if len(slice) < 1 {
    return -2
  }
  index = -2
  l := 0
  r := len(slice) - 1
  m := l + (r-l)/2
  if slice[m] == n {
    index = m// index from 1
  } else if slice[m] > n {
    index = search(n, slice[:m])
  } else {
    index = search(n, slice[m+1:])
    if index > 0 { index += m }
  }
  return
}

func printError(err error) {
  fmt.Println(err)
  os.Exit(1)
}