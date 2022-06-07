package main

import (
  "fmt"
  "os"
)

func main() {
  n,m,err := inputInts()
  if err != nil {
    printError(err)
    os.Exit(1)
  }
  fmt.Println(fibanocci3(n, m))
}

func inputInts() (n,m int, err error) {
  _, err = fmt.Scan(&n, &m)
  return
}

func printError(err error) {
  fmt.Println(err)
}

//function fibanocci3 calculates n number of sequence by module m
func fibanocci3(n,m int)int {
  
}