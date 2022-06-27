package main

import (
  "fmt"
  "io"
  "os"
  "bufio"
  "strconv"
  "strings"
)

func main() {
  
}

//  This function reads input from os.Stdin
func readInputInt() (num []int, err error) {
  reader := bufio.NewReader(os.Stdin)
  str, err := reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  n, err := strconv.Atoi(str)
  if err != nil {
    return []int{}, err
  }
  str, err = reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s := strings.Fields(str)
  num = make([]int, n)
  for i, _ := range num {
    num[i], err =  strconv.Atoi(s[i])
    if err != nil {
      return []int{}, err
    }
  }
  return
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}