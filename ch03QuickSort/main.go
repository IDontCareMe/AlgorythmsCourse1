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

const (
  eq = iota
  gt 
  le 
)

type line struct {
 l,r int 
}

type lines []line
func (theLines lines) Swap(i,j int) { theLines[i], theLines[j] = theLines[j], theLines[i] }
//func (theLines lines) Less(i,j int)bool { return theLines[i].l < theLines[j].l }
func (theLines lines) Len()int { return len(theLines) }

type comparator func(int, int)bool
type comparator3 func(int, int)int

func main() {
  // Read input
  //l, p, err := readInput()
  l, p, err := readInputNew()
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
  /*
  // Sort by left end
  l = quickSort(l, 0, l.Len()-1, func(i, j int)bool{
    return l[i].l < l[j].l
  })
  // Sort by right end
  r = quickSort(r, 0, r.Len()-1, func(i, j int)bool {
    return r[i].r < r[j].r
  })
*/
  // Sort by left end
  l = quickSort3(l, 0, l.Len()-1, func(i, j int)int {
    comp := eq
    if l[i].l < l[j].l {
        comp = le
    } else if l[i].l > l[j].l {
        comp = gt
    }
    return comp
  })
  // Sort by right end
  r = quickSort3(r, 0, r.Len()-1, func(i, j int)int {
    comp := eq
    if r[i].r < r[j].r {
        comp = le
    } else if r[i].r > r[j].r {
        comp = gt
    }
    return comp
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

  return M - N
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}

// Test
func readInputNew()(lines lines, points []int, err error) {
  var n, m int
  fmt.Scan(&n, &m)
  lines = make([]line, n)
  for i,_ := range lines {
    var l, r int
    fmt.Scan(&l, &r)
    lines[i].l = l
    lines[i].r = r
  }
  points = make([]int, m)
  for i,_ := range points {
    var p int
    fmt.Scan(&p)
    points[i] = p
  }
  err = nil
  return
}

// This function sorts slice
func quickSort3(theLines lines, l, r int, Comparator comparator3)lines {
  if l >= r {
    return theLines
  }
  m,z := partition3(theLines, l, r, Comparator)
  quickSort3(theLines, l, m-1, Comparator)
  quickSort3(theLines, z + 1, r, Comparator)
  return theLines
}

//This function parts slice for quick sort func
func partition3(theLines lines, l, r int, Comparator comparator3)(j, z int) {
    //theLines.Swap(l, rand.Intn(theLines.Len() - 1))
  j = l
  z = l
  for i := l + 1; i <= r ; i++ {
    switch Comparator(i,l) {
        case eq: z++; theLines.Swap(i,z);
        case le: z++; theLines.Swap(i,z); j++; theLines.Swap(z,j)
    }
  }
  theLines.Swap(l,j)
  return
}