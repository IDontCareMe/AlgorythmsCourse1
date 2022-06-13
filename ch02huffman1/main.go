package main

import(
  "fmt"
  "io"
  "os"
  "errors"
  "bufio"
  "sort"
  "strings"
)

// Tree node struct node for binar tree contains frq & char
type TreeNode struct {
  frq int
  char rune
}
func (t *TreeNode) String()string {
  return fmt.Sprintf("Frq: %d, Char: %c", t.frq, t.char)
}

func main() {
  // Read input
  s, err := readInputString()
  if err != nil {
    printErrors(err)
  }
  // search all runes in string
  res := searchRune(s)
  // Create code table
  table := codeTable(res)
  // Code string
  answer := code(s, table)
  // print result
  printResult(answer, table)
}

// This function reads input and returns string of error
func readInputString()(s string, err error) {
  reader := bufio.NewReader(os.Stdin)
  s, err = reader.ReadString('\n')
    if err == io.EOF { err = nil }
  if len(s) <= 0 {
    err = errors.New("String must contains at least one char")
  }
  s = strings.TrimSpace(s)
  return 
}

// This function searching symbols in srtings and returns count of it as a map
func searchRune(s string)(result map[rune]int){
  result = make(map[rune]int)
  for _, value := range s {
    result[value] += 1
  }
  return
}

// This function returns sorted code table
func codeTable(m map[rune]int)(table map[rune]string) {
  // Create table as slice
  var t []TreeNode = make([]TreeNode, 0)
  for key, value := range m {
    t = append(t, TreeNode {value, key})
  }
  // Sort table slice
  sort.Slice(t,func(i,j int)bool{
    return t[i].frq > t[j].frq
  })
  // Code table
  table = make(map[rune]string)
  if len(t) <= 1 {
    table[t[0].char] = "0"
    return
  }
  for i,val := range t {
    c := strings.Repeat("1",i)
    if i != len(t)-1 {
      c += "0"
    }
    table[val.char] = c
  }
  return
}

// This function codes input string
func code(s string, t map[rune]string)(c string) {
  for _,v := range s {
    c += t[v] 
  }
  return
}

// This function prints result
func printResult(s string, t map[rune]string) {
  fmt.Printf("%d %d\n",len(t), len(s))
  for key, val := range t {
    fmt.Printf("%s: %s\n",string(key), val)
  }
  fmt.Println(s)
}

// This function handles errors
func printErrors(err error) {
  fmt.Println(err)
  os.Exit(1)
}