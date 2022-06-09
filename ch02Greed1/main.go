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

func main() {
  // read input
  lines, err := readInput()
  if err != nil {
    writeError(err)
  }
  // sort slice
  lines = sortSlice(lines)
  n, points := findPoints(lines)
  //output
  fmt.Println(n)
  for _, l := range points {
    fmt.Printf("%d ",l)
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
    err = errors.New("error! The number of lines must be a number, greater than zero.")
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

//this function sorts slice (I don't know algorithm for sorting yet exept bubble sorting)
//sorting by left end of the line
//Think, it will be O(n^2) and it is bad((
func sortSlice(inputSlice []line)([]line) {
  // bubble sorting
  for {
    var changeFlag bool = false
    for i := 0; i < len(inputSlice) - 1; i++ {
      if inputSlice[i].left > inputSlice[i+1].left {
        inputSlice[i],inputSlice[i+1] = inputSlice[i+1], inputSlice[i]
        changeFlag = true
      }
    }
    if changeFlag == false {
      return inputSlice
    }
  }
}

// this function finds common points in line's slice
func findPoints(lines []line)(n int, output []int) {
  curNum := 0
  minBorder := lines[curNum].right
  if len(lines) == 0 {
    return 0, []int{}
  }
  
  for {
    //find least right border for iteration
    for i:= curNum; i < len(lines); i++ {
      if minBorder >= lines[i].right {
        minBorder = lines[i].right
        curNum = i
      }
    }
    //find all lines with left border less than new min border
    for i := curNum; i < len(lines); i++ {
      if minBorder >= lines[i].left {
        curNum = i
      }
    }
    //first point has found
    n++
    output = append(output, minBorder)
    //if all lines are covered, then return
    if curNum >= len(lines) - 1 {
      return
    }
    //new min border is next line border. Get it and go to another step
    minBorder = lines[curNum + 1].right
  }
}

// this function handles errors
func writeError(err error){
  fmt.Println(err)
  os.Exit(1)
}