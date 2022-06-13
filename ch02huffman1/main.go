package main

import(
  "fmt"
  "io"
  "os"
  "errors"
  "bufio"
)

// Tree node struct node for binar tree contains frq & char
type TreeNode struct {
  frq int
  char rune
  zero *TreeNode
  one *TreeNode
}
func (t *TreeNode) String()string {
  return fmt.Sprintf("Frq: %d, Char: %c, Zero: %x, One: %x", t.frq, t.char, t.zero, t.one)
}

// TreeQueue keep nodes for building a tree
type TreeQueue []TreeNode
// 
func (t *TreeQueue) String()(s string) {
  for _,val := range *t {
    s += fmt.Sprintf("%s: %v\n",string(val.char), val.frq)
  }
  return
}
// This function appends Nodes from map
func (t *TreeQueue) Create(m map[rune]int) {
  // Create all leaves
  for key,value := range m {
    *t = append(*t, TreeNode{value, key, nil, nil})
  }
}
// This function extracts Node from queue and returns it's adreess
func (t *TreeQueue) ExtractMin()*TreeNode {
  min := (*t)[0].frq
  num := 0
  for i,v := range *t {
    if min > v.frq {
      min = v.frq
      num = i
    } 
  }
  // Store Node address
  var p *TreeNode = &(*t)[num]
  // Exclude Node from Queue
  if num == len(*t) - 1 {
    *t = (*t)[:num]  
  } else {
   *t = append((*t)[:num], (*t)[num+1:]...) 
  }
  
  return p
}
// This function insert new node in queue
func (t *TreeQueue)Insert(n TreeNode) {
  *t = append(*t,n)
}


func main() {
  // Read input
  s, err := readInputString()
  if err != nil {
    printErrors(err)
  }
  // search all runes in string
  res := searchRune(s)
  // Create Queue
  var tq TreeQueue = make([]TreeNode, 0)
  tq.Create(res)
  fmt.Println(tq.String())
  var p *TreeNode
  for i:=0; i < 8; i++ {
    p = tq.ExtractMin()
    fmt.Println(p.String())
  }
  // print result
  //printResult(res)
  //
  //p:=createTree(tq)
  //fmt.Println(p.String())
}

// This function reads input and returns string of error
func readInputString()(s string, err error) {
  reader := bufio.NewReader(os.Stdin)
  s, err = reader.ReadString('\n')
  //_, err = fmt.Scan(&s)
    if err == io.EOF { err = nil }
  if len(s) <= 0 {
    err = errors.New("String must contains at least one char")
  }
  return 
}

// This function searching symbols in srtings and returns count of thic symbol as map
func searchRune(s string)(result map[rune]int){
  result = make(map[rune]int)
  for _, value := range s {
    result[value] += 1
  }
  return
}

// This function create binarry tree and and returns root node
func createTree(tq TreeQueue)(node *TreeNode) {
  //for i := 1; i < len(tq); i++ {
  fmt.Println(tq.String(),"Queue lenght:",len(tq))
  for {
    if len(tq) <= 1 {return}
    zero := tq.ExtractMin()
    //fmt.Println("zero frq:",(*zero).frq)
    fmt.Println("zero frq:", zero.String())
    one := tq.ExtractMin()
    fmt.Println("one frq:",one.String())
    node = &TreeNode{
      zero.frq + one.frq,
      0,
      zero,
      one,
    }
    //tq.Insert(*node)
    //fmt.Println("New node:",*node)
    fmt.Println(tq.String(),"Queue lenght:",len(tq))
  }
  //return  
}

// This function prints result
func printResult(m map[rune]int) {
  for key, val := range m {
    fmt.Println(string(key), val)
  }
}

// This function handles errors
func printErrors(err error) {
  fmt.Println(err)
  os.Exit(1)
}