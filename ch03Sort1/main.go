package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  unsorted, err := readInput()
  if err != nil {
    printError(err)
  }
  fmt.Println(unsorted)
}

// this function read os.Stdin and returns []int slice
func readInput()(unsorted []int, err error) {
  var n int
  _, err = fmt.Scan(&n)
  if err != nil {
    return []int{}, err
  }
  unsorted = make([]int, n)
  for index, _ := range unsorted {
    _, err = fmt.Scan(&unsorted[index])
    if err != nil {
      return []int{}, err
    }
  }
  return
}

func merge(left,right []int)(merged []int) {
  merged = make([]int, len(left) + len(right))
  l, r := 0, 0
  for i, _ := range merged {
    if l >= len(left) {
      merged = right[r:]
    }
    if left[l] < right[r] {
      merged[i] = left[l]
      l++
    }
  }
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println("error!", err)
    os.Exit(1)
  }
}