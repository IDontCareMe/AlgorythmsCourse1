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
    answer := search(value, sorted)
    if answer != -1 { answer++ }
    fmt.Printf("%d ", answer)
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
  index = -1
  l := 0
  r := len(slice) - 1
  for {
    //fmt.Println(n,":",l,r)
    if l > r { break }
    m := l + (r-l)/2
    if slice[m] == n {
      index = m
      break
    } else if slice[m] > n {
      r = m-1
    } else {
      l = m + 1
    }
  }
  return
}

func printError(err error) {
  fmt.Println(err)
  os.Exit(1)
}