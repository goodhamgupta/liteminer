/*
 *  Brown University, CS138, Spring 2020
 *
 *  Purpose: contains all interval related logic.
 */

package pkg

// Interval represents [Lower, Upper)
type Interval struct {
	Lower uint64 // Inclusive
	Upper uint64 // Exclusive
}

// GenerateIntervals divides the range [0, upperBound] into numIntervals intervals.
func GenerateIntervals(upperBound uint64, numIntervals int) (intervals []Interval) {
	// TODO: Students should implement this.
	var intervalSize uint64 = upperBound / uint64(numIntervals)
	for i := 1; i <= numIntervals; i++ {
		intervals = append(intervals, Interval{Lower: uint64(i - 1), Upper: uint64(i) * intervalSize})
	}
    // Clamp to upperBound
    if intervals[numIntervals-1].Upper > upperBound {
        intervals[numIntervals-1].Upper = upperBound
    }
	return
}
