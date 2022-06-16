// https://golang-blog.blogspot.com/2020/08/package-heap-in-golang.html
package main

import (
  "fmt"
  "os"
  "io"
  "container/heap"
  "bufio"
  "strings"
  "strconv"
)

type theHeap []int
//Implements sort.Interface
func (h theHeap) Len()int { return len(h) }
func (h theHeap) Less (i,j int)bool { return h[i] > h[j] }
func (h theHeap) Swap (i,j int) { h[i], h[j] = h[j], h[i] }
//Implements heap.Interface
func (h *theHeap) Push(x interface{}) {
  *h = append(*h, x.(int))
}
func (h *theHeap) Pop()interface{} {
  old := *h
  n := len(old)
  x := old[n-1]
  *h = old[:n-1]
  return x
}

func main() {
  var n int
  _,err := fmt.Scanln(&n)
  if err != nil {
    printError(err)
  }
  h := &theHeap{}
  heap.Init(h)
  reader := bufio.NewReader(os.Stdin)
  for i := 0; i < n; i++ {
    s, err := readInputString(reader)
    if err != nil {
      printError(err)
    }
    com, arg, err := parseInputString(s)
    if err != nil {
      printError(err)
    }
    switch com {
      case "Insert": heap.Push(h, arg)
      case "ExtractMax": fmt.Println(heap.Pop(h))
      default: fmt.Println("error!")
    }
  }
}

// This function reads input as string
func readInputString(r *bufio.Reader)(s string, err error) {
  //reader := bufio.NewReader(os.Stdin)
  s, err = r.ReadString('\n')
  if err != nil {
    if err != io.EOF {
      return "", err
    }
  }
  s = strings.TrimSpace(s)
  //fmt.Println(s)
  return
}

// This function parse string
func parseInputString(s string)(com string, arg int, err error) {
  commands := strings.Split(s, " ")
  //fmt.Printf("%q\n",commands)
  com = commands[0]
  if len(commands) <= 1 {
    return com, 0, nil
  }
  arg, err = strconv.Atoi(commands[1])
  return
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}