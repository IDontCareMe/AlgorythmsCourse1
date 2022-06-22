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
  
}

// This function reads input
func readInputInt()([]int, error) {
  reader := bufio.NewReader(os.Stdin)
  s, err := reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s = strings.TrimSpaces(s)
  n, err := strconv.Atoi(s)
  if err != nil {
    return []int{}, err
  }
  
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}