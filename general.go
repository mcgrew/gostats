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

// Returns the standard deviation of a set of values.
// 
// @param values The set to find the standard deviation of.
// @return The standard deviation as a double.
func StandardDeviation(values []float64) float64 {
  sum, sumSq := float64(0.0), float64(0.0)
  length := float64(0.0);
  for _,d := range values {
    if !math.IsNaN( d ) {
      sum += d;
      sumSq += d*d;
      length++;
    }
  }
  return math.Sqrt(( math.Abs( sumSq - sum * sum ) /
                  (( length ) * ( length - 1 ))));
}

//// Finds the median of a set of values. NaN values are omitted.
//// 
//// @param values The set to find the median of.
//// @return The median of those values.
//func Median(values []float64) float64 {
//  int length = values.length;
//  for ( double value : values ) {
//    if ( Double.isNaN( value ))
//      length--;
//  }
//  if ( length == 0 )
//    return Double.NaN;
//  double [] sorted = Arrays.copyOf( values, values.length );
//  Arrays.sort( sorted );
//  if (( length & 1 ) == 0 ) {
//    return ( sorted[ length / 2 ] + sorted[ length / 2 + 1 ]) / 2;
//  } else {
//    return sorted[ length / 2 ];
//  }
//}
//
//// Finds the mean of a set of values. NaN values are omitted.
//// 
//// @param values The set of values.
//// @return The mean of those values.
//func Mean(values []float64) float64 {
//  double sum = 0.0;
//  int count = 0;
//  for ( double d : values ) {
//    if ( !Double.isNaN( d )) {
//      sum += d;
//      count++;
//    }
//  }
//  return sum / count;
//}
//
//// Finds the sum of a set of values. NaN values are omitted.
//// 
//// @param values The set of values.
//// @return The sum of those values.
//func Sum(values []float64) float64 {
//  double sum = 0.0;
//  for ( double d : values ) {
//    if ( !Double.isNaN( d ))
//      sum += d;
//  }
//  return sum;
//}
//
//// Returns the minimum of a set of values. NaN values are omitted.
//// 
//// @param values The set of values to find the minimum for.
//// @return The minimum value in the array.
//func Min(values []float64) float64 {
//  if ( values.length == 0 )
//    return Double.NaN;
//  double returnValue = Double.POSITIVE_INFINITY;;
//  for ( double d : values ) {
//    returnValue = Math.min( returnValue, d );
//  }
//  return returnValue;
//}
//
//// Returns the maximum of a set of values. NaN values are omitted.
//// 
//// @param values the set of values to find the maximum for.
//// @return The maximum value in the array.
//func Max(values []float64) float64 {
//  if ( values.length == 0 )
//    return Double.NaN;
//  double returnValue = Double.NEGATIVE_INFINITY;
//  for ( double d : values ) {
//    if ( !Double.isNaN( d )) {
//      returnValue = Math.max( returnValue, d );
//    }
//  }
//  return returnValue;
//}
//
//// Returns the minimum regular value in a data set. This is the smallest
//// value that would not be considered an outlier.
//// 
//// @param values The set of values to find the minimum regular value for.
//// @return The minimum regular value.
//func MinRegular(values []float64) float64 {
//  double [] sorted = Arrays.copyOf( values, values.length );
//  Arrays.sort( sorted );
//  double q1 = firstQuartile( sorted );
//  for ( double d : sorted ) {
//    if ( Double.compare( d, q1 ) >= 0 )
//      return d;
//  }
//  return Double.NaN;
//}
//
//// Returns the maximum regular value in a data set. This is the largest
//// value that would not be considered an outlier.
//// 
//// @param values The set of values to find the maximum regular value for.
//// @return The minimum regular value.
//func MaxRegular(values []float64) float64 {
//  double [] sorted = Arrays.copyOf( values, values.length );
//  Arrays.sort( sorted );
//  double q3 = thirdQuartile( sorted );
//  for ( int i = sorted.length; i >= 0; i-- ) {
//    if ( Double.compare( sorted[ i ], q3 ) <= 0 )
//      return sorted[ i ];
//  }
//  return Double.NaN;
//}
//
//// Returns the data value that is the 25th percentile of a set of data.
//// 
//// @param values The set of values to find the first quartile of.
//// @return The first quartile data point.
//func FirstQuartile(values []float64) float64 {
//}
//
//// Returns the data value that is the 75th percentile of a set of data.
//// 
//// @param values The set of values to find the third quartile of.
//// @return The third quartile data point.
//func ThirdQuartile(values []float64) float64 {
//}
//
//// Finds the outliers in a data set.
//// 
//// @param values The set of data to find the outliers of.
//// @return All values that fall outside of the regular value range.
//public static double[] outliers( double[] values ) {
//  NumberList outliers = new NumberList( );
//  Range regularRange = regularRange( values );
//  for ( double value : values ) {
//    if ( !regularRange.contains( value ))
//      outliers.add( value );
//  }
//  return outliers.toDoubleArray( );
//}
//
//// Finds the range inside which all regular values fall (inclusive).
//// 
//// @param values The set of data to find the regular range for.
//// @return A range outside which all other values in the data set would be
////  considered outliers.
//func RegularRange(values []float64) float64,float64 {
//  double [] sorted = Arrays.copyOf( values, values.length );
//  Arrays.sort( sorted );
//  Range range = new Range( firstQuartile( sorted ), thirdQuartile( sorted ))
//    .scale( 4.0 );
//  double low = Double.NaN;
//  double high = Double.NaN;
//  for ( double d : sorted ) {
//    if ( Double.isNaN( low ) && range.contains( d ))
//      low = d;
//    if ( Double.isNaN( high ) || range.contains( d ))
//      high = d;
//  }
//  return low, high
//}
//
//// Finds the range of the first and third quartiles of a data set, or the
//// 25th to 75th percentile of a set of data.
//// 
//// @param values The set of data to find the quartile range for.
//// @return A range containing the 25th to 75th percentiles.
//func QuartileRange(values []float64) float64,float64 {
//}
//
//
