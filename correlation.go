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
	"math"
)

// These funcions are ported from java code in bbc-util:
//   https://github.com/mcgrew/bbc-util/blob/master/src/edu/purdue/bbc/util/Statistics.java

// Returns the Pearson correlation coefficient of the 2 sets of values.
// Each array should be the same length and contain at least 3 values.
//
// <PRE>
//    x = sample values in the first set
//    y = sample values in the second set
//    n = number of values in each set
//    Sx = standard deviation of x
//    Sy = standard deviation of y
//
//                                               _              _
//                     sum( i=0 to n-1, ( x[i] - x ) * ( y[i] - y ))
// correlationValue = -----------------------------------------------
//                                ( n - 1 ) Sx * Sy
//
// References:
// <CITE>Wikipedia, the free encyclopedia (2010, April 13).
//     <I>Correlation and Dependence</I>.
//     Retrieved April 14, 2010, from http://en.wikipedia.org/wiki/Correlation#Pearson.27s_product-moment_coefficient</CITE>
// </PRE>
//
// Parameters:
//   x: The first set of values to use for the calculation.
//   y: The second set of values to use for the calculation.
//
// Return value:
//   float64: The Pearson correlation value.
func PearsonCorrelation(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n < 3 {
		return math.NaN()
	}

	var Sx, Sy, meanX, meanY float64 = 0.0, 0.0, 0.0, 0.0
	var numerator float64 = 0.0
	var sumX, sumY, sumXSq, sumYSq float64 = 0.0, 0.0, 0.0, 0.0

	for _, currentX := range x {
		meanX += currentX
	}
	for _, currentY := range y {
		meanY += currentY
	}
	meanX /= float64(n)
	meanY /= float64(n)
	for i := 0; i < n; i++ {
		thisX := x[i]
		thisY := y[i]
		numerator += (thisX - meanX) * (thisY - meanY)
		sumX += thisX
		sumY += thisY
		sumXSq += thisX * thisX
		sumYSq += thisY * thisY
	}
	Sx = math.Sqrt((sumXSq - sumX*sumX/float64(n)) / float64(n-1))
	Sy = math.Sqrt((sumYSq - sumY*sumY/float64(n)) / float64(n-1))

	return numerator / (float64(n-1) * Sx * Sy)

}

// Returns the Spearman rank correlation coefficient of the 2 sets of data.
// Each array should be the same length and contain at least 3 values.
//
// <PRE>
//    x = values in the first set
//    y = values in the second set
//    n = number of values in each set
//    Rx = rank array of x, ie. the new locations of each element of x if x were sorted
//    Ry = rank array of y, ie. the new locations of each element of y if y were sorted
//
//                         6 * sum( i=0 to n-1, ( Rx[i] - Ry[i] )^2 )
// correlationValue = 1 - --------------------------------------------
//                                       n * ( n^2 - 1 )
//
// References:
// <CITE>Wikipedia, the free encyclopedia (2010, April 1).
//     <I>Spearman's rank correlation coefficient</I>.
//     Retrieved April 14, 2010, from http://en.wikipedia.org/wiki/Spearman's_rank_correlation_coefficient</CITE>
// <CITE>Spiegel, Murray R. (1961). Statistics, 2nd Edition (pp 376,391-393). New York.</CITE>
// </PRE>
//
// @param x The first set of data to use for the calculation.
// @param y The second set of data to use for the calculation.
// @return The Spearman correlation value.
// Parameters:
//   x: The first set of values to use for the calculation.
//   y: The second set of values to use for the calculation.
//
// Return value:
//   float64: The Pearson correlation value.
func SpearmanCorrelation(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n < 3 {
		return math.NaN()
	}

	Rx := getRank(x)
	Ry := getRank(y)

	var numerator float64 = 0.0
	for i := 0; i < n; i++ {
		numerator += (Rx[i] - Ry[i]) * (Rx[i] - Ry[i])
	}

	return 1.0 - (6.0*numerator)/float64(n*(n*n-1))
}

// Returns the Kendall tau rank correlation coefficient of the 2 sets of data.
// Each array should be the same length and contain at least 3 values.
//
// <PRE>
//    x = values in the first set
//    y = values in the second set
//    n = number of values in each set
//    Rx = the rank of the values in the first set
//    Ry = the rank of the values in the second set
//    Tx = the number of ties in x.
//    Ty = the number of ties in y.
//    concordant = The number of concordant pairs
//    discordant = The number of discordant pairs
//
//                                 concordant - discordant
// correlationValue = ---------------------------------------------
//                              n*(n-1)            n*(n-1)
//                     sqrt( ( ------- - Tx ) * (  ------- - Ty ) )
//                                2                   2
//
// References:
// <CITE>Wikipedia, the free encyclopedia (2010, March 25).
//     <I>Kendall tau rank correlation coefficient</I>.
//     Retrieved April 14, 2010, from http://en.wikipedia.org/wiki/Kendall_tau_rank_correlation_coefficient</CITE>
// </PRE>
//
// Parameters:
//   x: The first set of values to use for the calculation.
//   y: The second set of values to use for the calculation.
//
// Return value:
//   float64: The Kendall Tau correlation value.
func KendallCorrelation(x, y []float64) float64 {
	n := len(x)
	if n != len(y) || n < 3 {
		return math.NaN()
	}

	Rx := getRank(x)
	Ry := getRank(y)

	// Count the number of concordant and discordant pairs, and ties.
	var concordant, discordant, xTies, yTies float64 = 0.0, 0.0, 0.0, 0.0

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			xRel := Rx[j] - Rx[i]
			yRel := Ry[j] - Ry[i]

			if xRel == 0.0 || yRel == 0.0 { // check for ties first
				if xRel == 0.0 {
					xTies++
				}
				if yRel == 0.0 {
					yTies++
				}
			} else if (xRel > 0 && yRel > 0) || (xRel < 0 && yRel < 0) {
				concordant++
			} else if (xRel > 0 && yRel < 0) || (xRel < 0 && yRel > 0) {
				discordant++
			}
		}
	}
	denominator := float64(n*(n-1)) / 2.0
	return float64(concordant-discordant) /
		math.Sqrt((denominator-float64(xTies))*(denominator-float64(yTies)))
}

func getRank(x []float64) []float64 {
	n := len(x)
	order := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		order = append(order, float64(i))
	}
	TandemSort(x, order)
	rankSlice := make([]float64, n, n)
	returnvalue := make([]float64, n, n)
	lastDiff := 0
	for i := 0; i < n; i++ {
		rankSlice[i] = float64(i)
		if x[i] > x[lastDiff] {
			if i-lastDiff > 1 {
				newRank := float64(lastDiff) + float64(((i-1)-lastDiff))/2.0
				for j := lastDiff; j < i; j++ {
					rankSlice[j] = newRank
				}
			}
			lastDiff = i
		}
	}
	for i := 0; i < n; i++ {
		returnvalue[int(order[i])] = rankSlice[i]
	}
	return returnvalue
}
