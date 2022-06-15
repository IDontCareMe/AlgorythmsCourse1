package main

import (
  "fmt"
  "os"
  "errors"
  "bufio"
  "strings"
  "io"
)

type Dictionary map[string]string

func main() {
  // Read input
  d, s, err := readInputString()
  if err != nil {
    printError(err)
  }

  /*
  for key,value := range d {
    fmt.Printf("%s: %s\n", key, value)
  }
  fmt.Println(s)*/
  
  // Decode
  answer := decode(s, d)
  // Print result
  fmt.Println(answer)
}

// This function reads input from io.Stdin
func readInputString()(dictionary Dictionary, code string, err error){
  // Read first string
  var k, l int
  _, err = fmt.Scan(&k, &l)
  if (err != nil) || (k == 0) || (l == 0) {
    err = errors.New("error: Wrong input! k and l must be a number > 0")
    return
  }
 
  // Read k strings & store it in map
  reader := bufio.NewReader(os.Stdin)
  dictionary = make(Dictionary)
  for i := 0; i < k; i++ {
    s, er := reader.ReadString('\n')
    if (er != nil) {
      err = er
      return
    }
    arr := strings.Split(s, ":")
    dictionary[strings.TrimSpace(arr[1])] = strings.TrimSpace(arr[0]) 
  }
  // Read coded string
  code, err = reader.ReadString('\n')
  if err == io.EOF {
    err = nil
  }
  return
}

// This function decode string using dictionary
func decode(code string, d Dictionary)(s string) {
  letter := ""
  for _,val := range code {
    letter += string(val)
    if char, ok := d[letter]; ok {
      s += char
      letter = ""
    }
  }
  return
}

// This function handles errors
func printError(err error) {
  fmt.Println(err)
  os.Exit(1)
}