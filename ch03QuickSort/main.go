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
func (theLines lines) LessR(i,j int)bool { return theLines[i].r > theLines[j].r }
func (theLines lines) LessL(i,j int)bool { return theLines[i].l < theLines[j].l }
func (theLines lines) Len()int { return len(theLines) }

func main() {
  l, p, err := readInput()
  if err != nil {
    printError(err)
  }
  //fmt.Println(l)
  //l = quickSort(l, 0, l.Len()-1, true)
  
  for _, val := range p {
    fmt.Printf("%d ", checkContain(val, l))
  }
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

// This function sorts slice
func quickSort(theLines lines, l, r int, left bool)lines {
  if l >= r {
    return theLines
  }
  m := partition(theLines, l, r, left)
  quickSort(theLines, l, m-1, left)
  quickSort(theLines, m + 1, r, left)
  return theLines
}

//This function parts slice for quick sort func
func partition(theLines lines, l, r int, left bool)int {
  j := l
  for i:= l + 1; i <= r ; i++ {
    if left {
      if theLines.LessL(i, l) {
        j++
        theLines.Swap(i,j)
      }
    } else {
      if theLines.LessR(i, l) {
        j++
        theLines.Swap(i,j)
      }
    }
  }
  theLines.Swap(l,j)
  return j
}

// This function checks if ilne contains point
func checkContain(n int, theLines lines)(contains int) {
  fmt.Println("\n n:", n)
  theLines = quickSort(theLines, 0, theLines.Len()-1, false)
  fmt.Println("slice right:", theLines)
  r := 0
  for i,v := range theLines {
    if n > v.r {
      r = i
      fmt.Println("r: ", r)
      break
    }
    r = i
    fmt.Println("r: ", r)
  }
  theLines = quickSort(theLines, 0, r, true)
  fmt.Println("slice left:", theLines)
  for i,v := range theLines {
    if n < v.l {
      contains = i
      break
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