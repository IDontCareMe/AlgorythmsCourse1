package main

// This function finds minimum quantity of common points in line's slice
func findPoints(lines []line)(n int, output []int) {
  n++
  min := lines[0].right
  output = append(output, min)
  for i := 1; i < len(lines); i++ {
    if lines[i].left > min {
      n++
      min = lines[i].right
      output = append(output, min)
    }
  }
  return
}