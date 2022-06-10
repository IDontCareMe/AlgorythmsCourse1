package main

import (
  "fmt"
  "os"
  "errors"
)

type item struct {
  weight int
  cost int
  //cpw float64 //cost per weight = c/w
}
func (i *item)String()string {
  return fmt.Sprintf("Cost: %d, Weight: %d", i.cost, i.weight)
}

func main() {
  // read input
  n, W, items, err := readInput()
  if err != nil {
    printError(err)
  }
  // sort items by cpw
  // fill the bag
  
  // print output
  printResult(n, W, items)
}

// This function read input
// Get count n and max bag weight W in first string
// Get weight and coast of item in next n strings
// Returns readed n, W, []item or error
func readInput()(n, W int, items []item, err error) {
  _, err = fmt.Scan(&n, &W)
  if (err != nil) || (n == 0) {
    err = errors.New("error! Incorrect arguments")
    return 0, 0, []item {}, err
  }
  items = make([]item, n)
  for i,_ := range items {
    _,err = fmt.Scan(&items[i].cost, &items[i].weight)
    if err != nil {
      err = errors.New("error! Incorrect item arguments")
      return 0, 0, []item {}, err
    }
  }
  return
}

// This function prints result
func printResult(n, W int, i []item) {
  fmt.Println(n)
  fmt.Println(W)
  for _,v := range i {
    fmt.Println(v.String())
  }
}

// This function handle errors
func printError(err error){
  fmt.Println(err)
  os.Exit(1)
}
