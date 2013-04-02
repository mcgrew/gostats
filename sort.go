package statistics

import (
  "sort"
  "math"
  )

type Doublet struct {
  Data []float64
  Companion []float64
}

type Triplet struct {
  Data []float64
  Companion []float64
  Companion2 []float64
}

// TandemSort sorts 2 arrays.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func TandemSort(arr, companion []float64) {
  data := new(Doublet)
  data.Data = arr
  data.Companion = companion
  sort.Sort(data)
}

// TripletSort sorts 2 arrays.
// It makes one call to data.Len to determine n, and O(n*log(n)) calls to
// data.Less and data.Swap. The sort is not guaranteed to be stable.
func TripletSort(arr, companion, companion2 []float64) {
  data := new(Triplet)
  data.Data = arr
  data.Companion = companion
  data.Companion2 = companion2
  sort.Sort(data)
}

func (p Doublet) Len() int {
  return len(p.Data)
}
func (p Doublet) Less(i, j int) bool {
  return p.Data[i] < p.Data[j] || math.IsNaN(p.Data[i]) && !math.IsNaN(p.Data[j])
}
func (p Doublet) Swap(i, j int) {
  p.Data[i], p.Data[j] = p.Data[j], p.Data[i]
  p.Companion[i], p.Companion[j] = p.Companion[j], p.Companion[i]
}

func (p Triplet) Len() int {
  return len(p.Data)
}
func (p Triplet) Less(i, j int) bool {
  return p.Data[i] < p.Data[j] || math.IsNaN(p.Data[i]) && !math.IsNaN(p.Data[j])
}
func (p Triplet) Swap(i, j int) {
  p.Data[i], p.Data[j] = p.Data[j], p.Data[i]
  p.Companion[i], p.Companion[j] = p.Companion[j], p.Companion[i]
  p.Companion2[i], p.Companion2[j] = p.Companion2[j], p.Companion2[i]
}

