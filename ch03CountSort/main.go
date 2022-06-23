package main

import (
  "fmt"
  "io"
  "os"
  "bufio"
  "strings"
  "strconv"
)

const debug bool = true

func main() {
  arr, err := readInputInt()
  if err != nil {
    printError(err)
  }
  answer := countSort(arr)
  printResult(answer)
}

// This function reads input
func readInputInt()([]int, error) {
  reader := bufio.NewReader(os.Stdin)
  s, err := reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s = strings.TrimSpace(s)
  n, err := strconv.Atoi(s)
  if err != nil {
    return []int{}, err
  }
  // Read array
  s, err = reader.ReadString('\n')
  if (err != nil) && (err != io.EOF) {
    return []int{}, err
  }
  s = strings.TrimSpace(s)
  str := strings.Fields(s)
  numbers := make([]int,n)
  for i, _ := range numbers {
    numbers[i], err = strconv.Atoi(str[i])
    if err != nil {
      return []int{}, err
    }
  }
  if debug { fmt.Println("Inpun:", numbers) }
  return numbers, err
}

// This function sorts Slice by count Sort
func countSort(numbers []int)(sorted []int) {
  var counts [11]int
  for _, value := range numbers {
    counts[value]++
  }
  if debug { fmt.Println("Counts:", counts) }
  for i := 1; i < len(counts); i++ {
    counts[i] += counts[i-1]
  }
  if debug { fmt.Println("Weights:", counts) }
  sorted = make([]int, len(numbers))
  for j := (len(numbers) - 1); j >= 0; j-- {
    sorted[counts[numbers[j]]-1] = numbers[j]
    counts[numbers[j]]--
  }
  return
}

// This function prints result
func printResult(result []int) {
  for _,val := range result {
    fmt.Printf("%d ", val)
  }
}

// This function handles errors
func printError(err error) {
  if err != io.EOF {
    fmt.Println(err)
    os.Exit(1)
  }
}