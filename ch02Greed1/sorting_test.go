package main

import (
  "testing"
)

var test []line = []line {
  {3, 5},
  {14, 25},
  {0, 8},
  {4,17},
  {2, 9},
}
var want []line = []line {
  {3, 5},
  {0, 8},
  {2, 9},
  {4,17},
  {14, 25},
}

//tests

func TestSortSliceLib(t *testing.T) {
  got := SortSliceLib(test)
  for i,v := range got {
    if (v.left != want[i].left) && (v.right != want[i].right) {
      t.Errorf("want %d;%d got %d;%d", want[i].left, want[i].right, v.left, v.right)
    }
  }
}

func TestSortSliceLib2(t *testing.T) {
  got := SortSliceLib2(test)
  for i,v := range got {
    if (v.left != want[i].left) && (v.right != want[i].right) {
      t.Errorf("want %d;%d got %d;%d", want[i].left, want[i].right, v.left, v.right)
    }
  }
}

func TestSortSlice(t *testing.T) {
  got := SortSlice(test)
  for i,v := range got {
    if (v.left != want[i].left) && (v.right != want[i].right) {
      t.Errorf("want %d;%d got %d;%d", want[i].left, want[i].right, v.left, v.right)
    }
  }
}

//benchmarks
func BenchmarkSortSliceLib(b *testing.B) {
  for n := 0; n < b.N; n++ {
    SortSliceLib(test)
  }
}

func BenchmarkSortSliceLib2(b *testing.B) {
  for n := 0; n < b.N; n++ {
    SortSliceLib2(test)
  }
}

func BenchmarkSortSlice(b *testing.B) {
  for n := 0; n < b.N; n++ {
    SortSlice(test)
  }
}