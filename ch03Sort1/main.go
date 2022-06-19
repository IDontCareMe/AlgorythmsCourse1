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
  _, inv := mergeSort(unsorted)
  fmt.Println(inv)
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

// This function...
func mergeSort(unsorted []int)([]int, int) {
  if len(unsorted) > 1 {
    m := len(unsorted)/2
    l,i1 := mergeSort(unsorted[:m])
    r, i2 := mergeSort(unsorted[m:])
    sorted, n := merge(l, r, i1, i2)
    //sorted, n := merge(mergeSort(unsorted[:m]), mergeSort(unsorted[m:]))
    return sorted, n
  }
  return unsorted, 0
}

// This function merge two slices to one
func merge(left,right []int, invl, invr int)(merged []int, n int) {
  merged = make([]int, len(left) + len(right))
  l, r := 0, 0
  for i, _ := range merged {
    if (l >= len(left)) && (r < len(right)) {
      merged[i] = right[r]
      r++
    } else if (l < len(left)) && (r >= len(right)) {
      merged[i] = left[l]
      l++
    } else if (l < len(left))&& (r < len(right)){
      if left[l] < right[r] {
        merged[i] = left[l]
        l++
      } else {
        merged[i] = right[r]
        r++
      }
    }
  }
  n = invl + invr + findInversionNumber(left, right)
  //fmt.Println(n)
  return
}

// This function finds all inversions
func findInversionNumber(left, right []int)(n int) {
  r := 0
  for index, value := range left {
    for i := r; i < len(right); i++ {
      if value > right[i] {
        n += len(left) - index
        r++
      } else {
        break
      }
    } 
  }
  return
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println("error!", err)
    os.Exit(1)
  }
}