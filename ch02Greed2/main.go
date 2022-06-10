package main

import (
  "fmt"
  "os"
  "errors"
  "sort"
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
  W, items, err := readInput()
  if err != nil {
    printError(err)
  }
  // sort items by cost per weight
  sortItems(items)
  // fill the bag
  prise := fillBag(W, items)
  // print output
  printResult(prise)
}

// This function read input
// Get count n and max bag weight W in first string
// Get weight and coast of item in next n strings
// Returns readed n, W, []item or error
func readInput()(W int, items []item, err error) {
  var n int
  _, err = fmt.Scan(&n, &W)
  if (err != nil) || (n == 0) {
    err = errors.New("error! Incorrect arguments")
    return 0, []item {}, err
  }
  items = make([]item, n)
  for i,_ := range items {
    _,err = fmt.Scan(&items[i].cost, &items[i].weight)
    if err != nil {
      err = errors.New("error! Incorrect item arguments")
      return 0, []item {}, err
    }
  }
  return
}

// This function sorts slice of item by price per weight from high to low
func sortItems(itm []item)[]item {
  sort.Slice(itm, func(i,j int)bool {
    return float64(itm[i].cost)/float64(itm[i].weight) > float64(itm[j].cost)/float64(itm[j].weight)
  })
  return itm
}

// This function fills the bag items with best price per weight
func fillBag(bagSpace int, items []item)(price float64) {
  for _,itm := range items {
    if bagSpace == 0 {break}
    if bagSpace >= itm.weight {
      bagSpace -= itm.weight
      price += float64(itm.cost)
    } else {
      price += (float64(itm.cost)/float64(itm.weight))*float64(bagSpace)
      bagSpace = 0;
    }
  }
  return price
}

// This function prints result
func printResult(r float64) {
  fmt.Printf("%.3f", r)
}

// This function handle errors
func printError(err error){
  fmt.Println(err)
  os.Exit(1)
}
