package main

import (
  "fmt"
  "os"
  "io"
  "errors"
)

func main() {
  
}

func readInputString()(s string) {
  
}

func parseInputString(s string) {
  
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}