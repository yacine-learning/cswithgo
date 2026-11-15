package main

// Adds counts the additions and subtractions, to compare with recomputing.
var Adds int

// MovingAverage returns the average of every window of k consecutive values.
func MovingAverage(xs []float64, k int) []float64 {
	if k <= 0 || k > len(xs) {
		return nil
	}
	sum := 0.0
	for _, x := range xs[:k] { // the first window
		sum += x
		Adds++
	}
	out := []float64{sum / float64(k)}
	for right := k; right < len(xs); right++ {
		sum += xs[right] - xs[right-k] // one value enters, one leaves
		Adds += 2
		out = append(out, sum/float64(k))
	}
	return out
}
