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
	"fmt"
	"math"
	"sort"
)

// Returns the standard deviation of a set of values.
//
// @param values The set to find the standard deviation of.
// @return The standard deviation as a double.
func StandardDeviation(values []float64) float64 {
	var sum, sumSq, length float64 = 0.0, 0.0, 0.0
	for _, d := range values {
		if !math.IsNaN(d) {
			sum += d
			sumSq += d * d
			length++
		}
	}
	return math.Sqrt((math.Abs(sumSq-sum*sum) /
		((length) * (length - 1))))
}

// Finds the median of a set of values. NaN values are omitted.
//
// @param values The set to find the median of.
// @return The median of those values.
func Median(values []float64) float64 {
	return Percentile(values, 50.0)
}

// Finds the mean of a set of values. NaN values are omitted.
//
// @param values The set of values.
// @return The mean of those values.
func Mean(values []float64) float64 {
	var sum float64 = 0.0
	count := 0
	for _, d := range values {
		if !math.IsNaN(d) {
			sum += d
			count++
		}
	}
	return sum / float64(count)
}

// Finds the sum of a set of values. NaN values are omitted.
//
// @param values The set of values.
// @return The sum of those values.
func Sum(values []float64) float64 {
	var sum float64 = 0.0
	for _, d := range values {
		if !math.IsNaN(d) {
			sum += d
		}
	}
	return sum
}

// Returns the minimum of a set of values. NaN values are omitted.
//
// @param values The set of values to find the minimum for.
// @return The minimum value in the array.
func Min(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	returnValue := math.Inf(1)
	for _, d := range values {
		returnValue = math.Min(returnValue, d)
	}
	return returnValue
}

// Returns the maximum of a set of values. NaN values are omitted.
//
// @param values the set of values to find the maximum for.
// @return The maximum value in the array.
func Max(values []float64) float64 {
	if len(values) == 0 {
		return math.NaN()
	}
	returnValue := math.Inf(-1)
	for _, d := range values {
		returnValue = math.Max(returnValue, d)
	}
	return returnValue
}

// Returns the minimum regular value in a data set. This is the smallest
// value that would not be considered an outlier.
//
// @param values The set of values to find the minimum regular value for.
// @return The minimum regular value.
func MinRegular(values []float64) float64 {
	fq, tq := QuartileRange(values)
	iqr := tq - fq
	return fq - 1.5*iqr
}

// Returns the maximum regular value in a data set. This is the largest
// value that would not be considered an outlier.
//
// @param values The set of values to find the maximum regular value for.
// @return The minimum regular value.
func MaxRegular(values []float64) float64 {
	fq, tq := QuartileRange(values)
	iqr := tq - fq
	return tq + 1.5*iqr
}

// Returns the data value that is the 25th percentile of a set of data.
//
// @param values The set of values to find the first quartile of.
// @return The first quartile data point.
func FirstQuartile(values []float64) float64 {
	return Percentile(values, 25.0)
}

// Returns the data value that is the 75th percentile of a set of data.
//
// @param values The set of values to find the third quartile of.
// @return The third quartile data point.
func ThirdQuartile(values []float64) float64 {
	return Percentile(values, 75.0)
}

func Percentile(values []float64, percentile float64) float64 {
	if percentile < 0 || percentile > 100 {
		fmt.Errorf("Invalid percentile value specified: %f", percentile)
		return math.NaN()
	}
	sorted := sortFloat64s(values)
	if len(sorted) == 0 {
		return math.NaN()
	}
	return percentile_(sorted, percentile)
}

// Finds the range inside which all regular values fall (inclusive).
//
// @param values The set of data to find the regular range for.
// @return A range outside which all other values in the data set would be
//  considered outliers.
func RegularRange(values []float64) (float64, float64) {
	fq, tq := QuartileRange(values)
	iqr := tq - fq
	return fq - 1.5*iqr, tq + 1.5*iqr
}

// Finds the range of the first and third quartiles of a data set, or the
// 25th to 75th percentile of a set of data.
//
// @param values The set of data to find the quartile range for.
// @return A range containing the 25th to 75th percentiles.
func QuartileRange(values []float64) (float64, float64) {
	sorted := sortFloat64s(values)
	return percentile_(sorted, 25.0), percentile_(sorted, 75.0)
}

func percentile_(sorted []float64, percentile float64) float64 {
	index := float64(len(sorted)+1)*(percentile/100) - 1
	ceil := math.Ceil(index)
	floor := math.Floor(index)
	if ceil == floor {
		return sorted[int(index)]
	}
	return (ceil-index)*sorted[int(floor)] +
		(index-floor)*sorted[int(ceil)]
}

func sortFloat64s(values []float64) []float64 {
	length := len(values)
	sorted := make([]float64, 0, length)
	for _, value := range values {
		if !math.IsNaN(value) {
			sorted = append(sorted, value)
		}
	}
	sort.Float64s(sorted)
	return sorted
}
