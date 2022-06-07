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

//function fibanocci3 calculates n number of sequence by module m F(n)%F(m)
func fibanocci3(n,m int)int {

  if n <=1 {
    return int(n)
  }

  var period int = -1
  var first, second int = 0, 1
  for i:= 2; i <= n; i++ {
    first, second = second%m, first + second
    //if (first%m == 0) && (second%m == 1) {
    if (first == 0) && (second == 1) {
      period = i-1
      n = n%period
      fmt.Println("Period:", period)
      fmt.Println("New n:", n)
      break
    }
  }
  if period == -1 {
    return second
  }
  
  /*if second == m {
    return 0
  } else if second < m {
    return second
  }*/
  
  if n <=1 {
    return n
  }
  first, second = 0, 1
  for i:= 2; i <= n; i++ {
    first, second = second, first + second
  }
  return second
}
