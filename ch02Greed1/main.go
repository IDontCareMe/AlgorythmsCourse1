package main

import (
  "fmt"
  "os"
  "errors"
)

// line structure
type line struct {
  left, right int  
}
func (l *line) String()string {
  return fmt.Sprintf("Left: %d, Right: %d", l.left, l.right)
}

func main() {
  // read input
  lines, err := readInput()
  if err != nil {
    writeError(err)
  }
  // sort slice
  //lines = sortSliceLib(lines)
  lines = SortSlice(lines)
  
  n, points := findPoints(lines)
  // output
  fmt.Println(n)
  for _, l := range points {
    fmt.Printf("%d ",l)
  }
}

// This functions reads input
func readInput()(lines []line, err error){
  // Get the number of lines
  var n int
  _, err = fmt.Scan(&n)
  if err != nil {
    return []line{}, err
  }
  if n <= 0 {
    err = errors.New("error! The number of lines must be a number, greater than zero.")
  }
  // Create slice of lines & fill it
  lines = make([]line, n)
  for i,_ := range lines {
    _,err = fmt.Scan(&lines[i].left, &lines[i].right)
    if err != nil {
      return []line{}, err
    }
    // Swap ends of line if necessary
    if lines[i].left > lines[i].right {
      lines[i].left, lines[i].right = lines[i].right, lines[i].left
    }
  }
  return lines, err
}

// This function handles errors
func writeError(err error){
  fmt.Println(err)
  os.Exit(1)
}