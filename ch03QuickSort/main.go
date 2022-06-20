package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
  "strconv"
)

const debug bool = false

type line struct {
 l,r int 
}

type lines []line
func (theLines lines) Swap(i,j int) { theLines[i], theLines[j] = theLines[j], theLines[i] } 

func main() {
  l, p, err := readInput()
  if err != nil {
    printError(err)
  }
  fmt.Println(l)
  l.Swap(0,1)
  fmt.Println(l)
  fmt.Println(p)
}

// This function reads input
func readInput()(lines lines, points []int, err error) {
  // Create reader
  reader := bufio.NewReader(os.Stdin)
  // Read number of lines & number of points
  s, err := reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return
  }
  s = strings.TrimSpace(s)
  str := strings.Fields(s)
  // Extract number of lines
  n, err := strconv.Atoi(str[0])
  if err != nil {
    return
  }
  if debug { fmt.Println(n) }
  // Extract number of points
  m, err := strconv.Atoi(str[1])
  if err != nil {
    return
  }
  if debug { fmt.Println(m) }
  // Read lines & fill slice
  lines = make([]line, n)
  for i, _ := range lines {
    s, err = reader.ReadString('\n')
    if (err != nil) && (err != io.EOF) {
      return
    }
    s = strings.TrimSpace(s)
    str = strings.Fields(s)
    lines[i].l, err = strconv.Atoi(str[0])
    lines[i].r, err = strconv.Atoi(str[1])
    if (err != nil) {
      return
    }
  }
  // Read points
  points = make([]int, m)
  s, err = reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return
  }
  s = strings.TrimSpace(s)
  str = strings.Fields(s)
  for i, _ := range points {
    points[i], err = strconv.Atoi(str[i])
    if err != nil {
      return
    }
  }
  return
}

// This function checks if ilne contains point
func checkContain(n int, lines []line)(contains int) {
  return
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}