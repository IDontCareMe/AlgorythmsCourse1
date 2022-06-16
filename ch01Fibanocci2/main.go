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

//function fibanocci3 calculates n number of sequence by module m F(n)%m
func fibanocci3(n,m int)int {

  if n <=1 {
    return n%m
  }

  var period int = -1
  arr := make([]int, m*6+2)
  var first, second int = 0, 1
  arr[0] = first
  arr[1] = second
  for i:= 2; i <= n; i++ {
    first, second = second, (first + second)%m
    arr[i] = second
    //fmt.Println(arr[i])
    if (first == 0) && (second == 1) {
      period = i-1
      //fmt.Println("Period:", period)
      n = n%period
      break
    }
  }
  if period == -1 {
    return second
  }

  /*
  if n <=1 {
    return n%m
  }
  first, second = 0, 1
  for i:= 2; i <= n; i++ {
    first, second = second, (first + second)%m
  }
  //return second
  */
  /*for _,val := range arr {
    fmt.Println(val)
  }*/
  //fmt.Println("arr =", arr[n])
  return arr[n]
}
