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
func (theLines lines) Less(i,j int)bool { return theLines[i].r > theLines[j].r }
func (theLines lines) Len()int { return len(theLines) }

func main() {
  l, p, err := readInput()
  if err != nil {
    printError(err)
  }
  l = quickSort(l, 0, l.Len()-1)
  
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
func quickSort(theLines lines, l, r int)lines {
  if l >= r {
    return theLines
  }
  m := partition(theLines, l, r)
  quickSort(theLines, l, m-1)
  quickSort(theLines, m + 1, r)
  return theLines
}

//This function parts slice for quick sort func
func partition(theLines lines, l, r int)int {
  j := l
  for i:= l + 1; i <= r ; i++ {
    if theLines.Less(i, l) {
      j++
      theLines.Swap(i,j)
    }
  }
  theLines.Swap(l,j)
  return j
}

// This function checks if ilne contains point
func checkContain(n int, theLines lines)(contains int) {
  for _,v := range theLines {
    if n <= v.r {
      if n >= v.l {
        contains++
      }
    } else {
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