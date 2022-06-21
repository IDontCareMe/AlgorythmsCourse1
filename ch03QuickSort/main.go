package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
  "strings"
  "strconv"
)

const debug bool = true

type line struct {
 l,r int 
}

type lines []line
func (theLines lines) Swap(i,j int) { theLines[i], theLines[j] = theLines[j], theLines[i] }
//func (theLines lines) Less(i,j int)bool { return theLines[i].l < theLines[j].l }
func (theLines lines) Len()int { return len(theLines) }

type comparator func(int, int)bool

func main() {
  // Read input
  l, p, err := readInput()
  if err != nil {
    printError(err)
  }
  // Duplicate l-Slice to sort by right end
  r := make(lines, len(l))
  copy(r, l)
  if debug {
    fmt.Println("unsorted l:", l)
    fmt.Println("unsorted r:", r)
  }
  // Sort by left end
  l = quickSort(l, 0, l.Len()-1, func(i, j int)bool{
    return l[i].l < l[j].l
  })
  // Sort by right end
  r = quickSort(r, 0, r.Len()-1, func(i, j int)bool {
    return r[i].r < r[j].r
  })
  if debug {
    fmt.Println("sorted l:", l)
    fmt.Println("sorted r:", r)
  }
  
  answer := ""
  for _, val := range p {
    answer += fmt.Sprintf("%d ", checkContain(val, l, r))
  }
  fmt.Println(answer)
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
  if debug { fmt.Println("Number of lines:", n) }
  // Extract number of points
  m, err := strconv.Atoi(str[1])
  if err != nil {
    return
  }
  if debug { fmt.Println("Number of points:", m) }
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
func quickSort(theLines lines, l, r int, Comparator comparator)lines {
  if l >= r {
    return theLines
  }
  m := partition(theLines, l, r, Comparator)
  quickSort(theLines, l, m-1, Comparator)
  quickSort(theLines, m + 1, r, Comparator)
  return theLines
}

//This function parts slice for quick sort func
func partition(theLines lines, l, r int, Comparator comparator)int {
  j := l
  for i:= l + 1; i <= r ; i++ {
    //if theLines.Less(i, l) {
    if Comparator(i, l) {
      j++
      theLines.Swap(i,j)
    }
  }
  theLines.Swap(l,j)
  return j
}

// This function checks if ilne contains point
func checkContain(n int, left, right lines)(contains int) {
  M:=0
  for m, l, r := 0, 0, left.Len()-1; l <= r; {
    m = l + (r-l)/2
    if n >= left[m].l {
      l = m + 1
    } else {
      r = m - 1
    }
    M = l
  }
  if debug { fmt.Println("Cont. left:", M) }
  
  /*
  for _,v := range left[:M] {
    if n <= v.r {
      contains++
    }
  } 
  */
  
  // For right
  N:=0
  for m, l, r := 0, 0, right.Len()-1; l <= r; {
    m = l + (r-l)/2
    if n <= right[m].r {
      r = m - 1
    } else {
      l = m + 1
    }
    N = r + 1
  }
  if debug { fmt.Println("Cont. right:", N) }

  contains = M - N
  
  /*
  for _,v := range theLines {
    if n >= v.l {
      if n >= v.r {
        contains++
      }
    } else {
        break
    }
  }
  return
*/
  return
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}