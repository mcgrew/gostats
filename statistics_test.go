//  Copyright 2013 Thomas McGrew
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package statistics

import (
  "testing"
  "sort"
  )

func TestTripletSort(t *testing.T) {
  slice1 := make([]float64, 0, 1000)
  slice2 := make([]float64, 0, 1000)
  slice3 := make([]float64, 0, 1000)
  for i := 0; i < 1000; i++ {
    slice1 = append(slice1, float64(1000 - i))
    slice2 = append(slice2, float64(1000 - i))
    slice3 = append(slice3, float64(1000 - i))
  }
  TripletSort(slice1, slice2, slice3)
  if !sort.Float64sAreSorted(slice1) || !sort.Float64sAreSorted(slice1) ||
    !sort.Float64sAreSorted(slice3) {
   t.Errorf("Triplet sort failed")
  }
}

func TestTandemSort(t *testing.T) {
  slice1 := make([]float64, 0, 1000)
  slice2 := make([]float64, 0, 1000)
  for i := 0; i < 1000; i++ {
    slice1 = append(slice1, float64(1000 - i))
    slice2 = append(slice2, float64(1000 - i))
  }
  TandemSort(slice1, slice2)
  if !sort.Float64sAreSorted(slice1) || !sort.Float64sAreSorted(slice1) {
   t.Errorf("Triplet sort failed")
  }
}

func TestGetRank(t *testing.T) {
  slice1 := make([]float64, 0, 1000)
  for i := 0; i < 1000; i++ {
    slice1 = append(slice1, float64(1000 - i))
  }
  result := getRank(slice1)
  for i := 1; i < 1000; i++ {
    if result[i-1] < result[i] {
      t.Errorf("getRank failed")
    }
  }
  slice1 = make([]float64, 0, 1000)
  for i := 0; i < 1000; i++ {
    slice1 = append(slice1, float64(999 - i))
  }
  slice1[200] = 800
  result = getRank(slice1)
  if result[200] != 799.5 || result[199] != 799.5 {
    t.Errorf("getRank failed")
  }
}

