package main

import (
  "fmt"
  "os"
)

func main() {
  a, b, err := imputInts()
  if err != nil {
    printErrors(err)
  }
}

//Scaning 2 digit from os.Stdin
func imputInts()(a,b int, err error) {
  _, err = fmt.Scan(&a, &b)
  return
}

//Print errors
func printErrors(err error){
  fmt.Println(err)
  os.Exit(1)
}