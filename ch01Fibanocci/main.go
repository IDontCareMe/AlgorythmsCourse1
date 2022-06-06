package main

import (
  "fmt"
  "os"
)

func main() {
  n,err := inputInt()
  if err != nil {
    printError(err)
    os.Exit(1)
  }
  //fmt.Println(fibanocci1(n))
  fmt.Println(fibanocci2(n))
}

func inputInt()(int, error) {
  var n int
  _, err := fmt.Scan(&n)
  return n, err
}

func printError(err error) {
  fmt.Println(err)
}

//this function returns n number fibanocci sequence
func fibanocci1(n int)int {
  if n <=1 {
    return n
  }
  var (
    first int = 0
    second int = 1
  )
  for i := 2; i <= n; i++ {
    first, second = second, first + second
  }
  return second
}

// this function returns last character
//of n number fibanocci 
func fibanocci2(n int)int {
  if n <= 1 {
    return n
  }
  var (
    first int = 0
    second int = 1
  )
  for i := 2; i <= n; i++ {
    first, second = second, (first + second)%10
  }
  return second
}