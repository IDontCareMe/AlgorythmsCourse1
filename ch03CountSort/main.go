package main

import (
  "fmt"
  "io"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  arr, err := readInputInt()
  if err != nil {
    printError(err)
  }
  fmt.Println(arr)
}

// This function reads input
func readInputInt()([]int, error) {
  reader := bufio.NewReader(os.Stdin)
  s, err := reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s = strings.TrimSpace(s)
  n, err := strconv.Atoi(s)
  if err != nil {
    return []int{}, err
  }
  // Read array
  s, err = reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s = strings.TrimSpace(s)
  str := strings.Fields(s)
  numbers := make([]int,n)
  for i, _ := range numbers {
    numbers[i], err = strconv.Atoi(str[i])
    if err != nil {
      return []int{}, err
    }
  }
  return numbers, err
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}