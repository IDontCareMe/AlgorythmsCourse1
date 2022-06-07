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
  var module int
  first, second := 0, 1
  
  if n <=1 {
    return n
  }
  for i:= 2; i <= n; i++ {
    first, second = second, first + second
    
    if i == m {
      module = second
    }
    
    if second == module {
      first, second = 0, 1
      i += 2
    }
  }
  return second
}