package main

import (
  "fmt"
  "os"
  "errors"
)

func main() {
  a, b, err := imputInts()
  if err != nil {
    printErrors(err)
  }
  fmt.Println(EuclidGCD(a,b))
}

//Scaning 2 digit from os.Stdin
func imputInts()(a,b int, err error) {
  _, err = fmt.Scan(&a, &b)
  if (a < 0) || (b < 0) {
    err = errors.New("error! incorrect input.")
  }
  return
}

//Print errors
func printErrors(err error){
  fmt.Println(err)
  os.Exit(1)
}

//Calculate the GCD by Euclid
func EuclidGCD(a,b int)int {
  if a == b {
    return a  
  }
  
  if a == 0 {
    return b
  }

  if b == 0 {
    return a
  }

  if a > b {
    return EuclidGCD(a%b, b)
  }

  return EuclidGCD(a, b%a)
}