package main

import (
  "fmt"
  "os"
  "errors"
)

type line struct {
  left, right int  
}

func main() {
  // read input
  lines, err := readInput()
  if err != nil {
    writeError(err)
  }
  // sort slice

  //output
  for _, l := range lines {
    fmt.Println(l.left, l.right)
  }
}

// this functions reads input
func readInput()(lines []line, err error){
  //Get number of lines
  var n int
  _, err = fmt.Scan(&n)
  if err != nil {
    return []line{}, err
  }
  if n <= 0 {
    err = errors.New("error! The number of lines must be a number greater than zero.")
  }
  //Create slice of lines & fill it
  lines = make([]line, n)
  for i,_ := range lines {
    _,err = fmt.Scan(&lines[i].left, &lines[i].right)
    if err != nil {
      return []line{}, err
    }
    //Swap ends of line if necessary
    if lines[i].left > lines[i].right {
      lines[i].left, lines[i].right = lines[i].right, lines[i].left
    }
  }
  return lines, err
}

// this function handles errors
func writeError(err error){
  fmt.Println(err)
  os.Exit(1)
}