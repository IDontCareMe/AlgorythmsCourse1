package main

import (
  "fmt"
  "os"
  "io"
)

func main() {
  
}

func readInput() {
  
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}