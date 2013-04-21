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

func TestMedian(t *testing.T) {
  slice1 := []float64{ 3.5, 10.0, 93.0, 12.0, 54.0}
  slice2 := []float64{ 54.0, 93.0, 87.2, 3.5, 10.0, 12.0 }
  result := Median(slice1)
  expect := 12.0
  if result != expect {
    t.Errorf("Median failed for odd length input. Expected %f, got %f", expect,
             result)
  }
  expect = 33.0
  result = Median(slice2)
  if result != expect {
    t.Errorf("Median failed for even length input. Expected %f, got %f", expect,
             result)
  }
}

func TestMean(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  result := Mean(slice1)
  expect := 43.25
  if result != expect {
    t.Errorf("Mean failed. Expected %f, got %f", expect, result)
  }
}

func TestSum(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  result := Sum(slice1)
  expect := 259.5
  if result != expect {
    t.Errorf("Sum failed. Expected %f, got %f", expect, result)
  }
}

func TestMin(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  result := Min(slice1)
  expect := 3.5
  if result != expect {
    t.Errorf("Min failed. Expected %f, got %f", expect, result)
  }
}

func TestFirstQuartile(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  result := FirstQuartile(slice1)
  expect :=  8.375
  if result != expect {
    t.Errorf("FirstQuartile failed. Expected %f, got %f", expect, result)
  }
}

func TestThirdQuartile(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  result := ThirdQuartile(slice1)
  expect :=  88.5
  if result != expect {
    t.Errorf("ThirdQuartile failed. Expected %f, got %f", expect, result)
  }
}

func TestRegularRange(t *testing.T) {
  slice1 := []float64{ 54.0, 93.0, 87.0, 3.5, 10.0, 12.0 }
  lower, upper := RegularRange(slice1)
  expect := -111.8125
  if lower != expect {
    t.Errorf("RegularRange failed. Expected lower bound %f, got %f", expect,
             lower)
  }
  lower = MinRegular(slice1)
  if lower != expect {
    t.Errorf("MinRegular failed. Expected lower bound %f, got %f", expect,
             lower)
  }
  expect = 208.6875
  if upper != expect {
    t.Errorf("RegularRange failed. Expected upper bound %f, got %f", expect,
             upper)
  }
  upper = MaxRegular(slice1)
  if upper != expect {
    t.Errorf("MaxRegular failed. Expected upper bound %f, got %f", expect,
             upper)
  }
}
