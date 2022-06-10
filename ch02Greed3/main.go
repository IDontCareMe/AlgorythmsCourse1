package main

import (
  "fmt"
  "os"
  "errors"
)

func main() {
  // Read input
  n, err := readInputInt()
  if err != nil {
    printError(err)
  }
  // Find manys of numbers
  res := findMany(n)
  // Print result
  printResult(res)
}

// This function reads inpet and returns int and error(if is)
func readInputInt()(n int, err error) {
  _, err = fmt.Scan(&n)
  if (err != nil) || (n == 0) {
    return 0, errors.New(" error! Incorrect input!")
  }
  return
}

// This function calculates maximum manys of numbers with sum eq n
func findMany(n int)(many []int) {
  var i int = 1
  for {
    if i < (n - i) {
      n -=i
      many = append(many, i)
    } else {
      many = append(many, n)
      break
    }
    i++
  }
  return
}

// This function printsvresult
func printResult(res []int) {
  fmt.Println(len(res))
  for _, val := range res {
    fmt.Printf("%d ", val)
  }
  fmt.Print("\n")
}

// This function handle and print errors
func printError(err error) {
  fmt.Println(err)
  os.Exit(1)
}