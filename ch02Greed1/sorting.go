package main

import (
  "sort"
)

// byRight implements sort.Interface for []line based on the right field
type byRight []line

func (b byRight)Len()int { return len(b) }
func (b byRight)Less(i,j int)bool { return b[i].right < b[j].right }
func (b byRight)Swap(i,j int) { b[i], b[j] = b[j], b[i] }

// This function sorts slice of lines by right end using sort.Sort()
func SortSliceLib(s []line)[]line {
  sort.Sort(byRight(s))
  return s
}

// This function sorts slice of lines by right end using sort.Slice
func SortSliceLib2(s []line)[]line {
  sort.Slice(s,func(i,j int)bool{
    return s[i].right < s[j].right
  })
  return s
}

// This is handmade function. It sorts slice (I don't know algorithm for sorting yet exept bubble sorting)
//sorting by right end of the line
//Think, it will be O(n^2) and it is bad((
func SortSlice(inputSlice []line)([]line) {
  // bubble sorting
  for {
    var changeFlag bool = false
    for i := 0; i < len(inputSlice) - 1; i++ {
      if inputSlice[i].right > inputSlice[i+1].right {
        inputSlice[i],inputSlice[i+1] = inputSlice[i+1], inputSlice[i]
        changeFlag = true
      }
    }
    if changeFlag == false {
      return inputSlice
    }
  }
}