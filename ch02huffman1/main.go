//https://golang-blog.blogspot.com/2020/08/package-heap-in-golang.html
package main

import(
  "fmt"
  "io"
  "os"
  "errors"
  "bufio"
  //"sort"
  "strings"
  "container/heap"
)

// Tree node struct node for binar tree 
type TreeNode struct {
  char rune
  frq int
  index int // index in Queue, uses in heap.Interface
  code string
  left, right *TreeNode
}
func (t *TreeNode)buildCode(s string, m *map[rune]string){
  t.code = s
  //if not leafs
  if (t.left != nil) && (t.right != nil) {
    t.left.buildCode(s + "0", m)
    t.right.buildCode(s + "1", m)
  } else {
    //fmt.Printf("%c: %s\n", t.char, t.code)
    (*m)[t.char] = t.code
  }
}
func (t *TreeNode) String()string {
  return fmt.Sprintf("%c: %s", t.char, t.code)
}


// Priority queue implements heap.Interface and contains TreeNode
type PriorityQueue []*TreeNode
// Implements sort.Interface
func (pq PriorityQueue) Len()int { return len(pq) }
func (pq PriorityQueue) Less(i, j int)bool { return pq[i].frq < pq[j].frq }//less priority first
func (pq PriorityQueue) Swap(i, j int) {
  pq[i], pq[j] = pq[j], pq[i]
  pq[i].index = i
  pq[j].index = j
}
// Implements heap.Interface
func (pq *PriorityQueue) Push (x interface{}) {
  n := len(*pq)
  item := x.(*TreeNode)
  item.index = n
  *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop()interface {} {
  old := *pq
  n := len(old)
  item := old[n - 1]
  old[n - 1] = nil
  item.index = -1
  *pq = old[0:n-1]
  return item
}
//func (pq *PriorityQueue) Update(item *TreeNode, )

func main() {
  // Read input
  s, err := readInputString()
  if err != nil {
    printErrors(err)
  }
  // search all runes in string
  dict := searchRune(s)
  // Create code table
  table := codeTable(dict)
  
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
func codeTable(m map[rune]int)(t map[rune]string) {
  // Create Priority Queue
  pq := make(PriorityQueue, len(m))
  i:=0
  for char, frq := range m {
    pq[i] = &TreeNode{
      char: char,
      frq: frq,
      index: i,
      //code: "",
    }
    i++
  }
  heap.Init(&pq)

  // Create binary tree
  for pq.Len() > 1 {
    left := heap.Pop(&pq).(*TreeNode)
    right := heap.Pop(&pq).(*TreeNode)
    node := &TreeNode{
      frq: left.frq + right.frq,
      left: left,
      right: right,
    }
    heap.Push(&pq, node)
    //fmt.Println(pq)
  }
  root := heap.Pop(&pq).(*TreeNode)
  t = make(map[rune]string)
  root.buildCode("", &t)
  //fmt.Println(root.frq)

  //Create code table
  
/*
  item := &TreeNode{
    char: rune('X'),
      frq: 10,
      code: "",
  }
  heap.Push(&pq, item)

  for pq.Len() > 0 {
    item := heap.Pop(&pq).(*TreeNode)
    fmt.Printf("%c: %d\n", item.char, item.frq)
  }*/
  /*
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
*/
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