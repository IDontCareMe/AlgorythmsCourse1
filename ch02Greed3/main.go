package main

import (
  "fmt"
  "os"
  "errors"
)

func main() {
  
}

func readInputInt()(n int, err error) {
  _, err = fmt.Scan(&n)
  if (err != nil) || (n == 0) {
    return 0, errors.New(" error! Incorrect input!")
  }
  return
}

func printError(err error) {
  fmt.Println(err)
  os.Exit(1)
}