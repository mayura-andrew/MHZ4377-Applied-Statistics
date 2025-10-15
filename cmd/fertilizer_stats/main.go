package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []float64{65, 64, 80, 66, 62, 67, 75, 54, 50, 74, 68, 65, 67, 55, 73, 71, 74, 61, 64, 52, 64, 60, 72}
	sort.Float64s(data)
	fmt.Println("Sorted data:", data)

	n := len(data)
	// Mean
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(n)

	// Median
	var median float64
	if n%2 == 0 {
		median = (data[n/2-1] + data[n/2]) / 2
	} else {
		median = data[n/2]
	}

	// Mode (may be multiple)
	freq := make(map[float64]int)
	for _, v := range data {
		freq[v]++
	}
	maxf := 0
	modes := []float64{}
	for _, f := range freq {
		if f > maxf {
			maxf = f
		}
	}
	if maxf == 1 {
		modes = []float64{} // no mode
	} else {
		for v, f := range freq {
			if f == maxf {
				modes = append(modes, v)
			}
		}
		sort.Float64s(modes)
	}

	// Quartiles Q1 Q3 (method: median of halves)
	mid := n / 2
	var lower, upper []float64
	if n%2 == 0 {
		lower = data[:mid]
		upper = data[mid:]
	} else {
		lower = data[:mid]
		upper = data[mid+1:]
	}
	q1 := medianOfSlice(lower)
	q3 := medianOfSlice(upper)
	iqr := q3 - q1

	// Variability note: range
	rangeVal := data[n-1] - data[0]

	// Print results
	fmt.Println()
	fmt.Println("--- Central Tendency ---")
	if len(modes) == 0 {
		fmt.Printf("Mode: None (all values unique)\n")
	} else {
		fmt.Printf("Mode(s): %v (frequency %d)\n", modes, maxf)
	}
	fmt.Printf("Mean: %.4f\n", mean)
	fmt.Printf("Median: %.4f\n", median)

	fmt.Println()
	fmt.Println("--- Quartiles and IQR ---")
	fmt.Printf("Q1: %.4f\n", q1)
	fmt.Printf("Q3: %.4f\n", q3)
	fmt.Printf("IQR: %.4f\n", iqr)
	fmt.Printf("Range: %.4f\n", rangeVal)

	// Boxplot interpretation
	fmt.Println()
	fmt.Println("Interpretation (boxplot):")
	fmt.Printf("- Median = %.4f indicates center.\n", median)
	fmt.Printf("- IQR = %.4f shows spread of middle 50%% of data.\n", iqr)
	fmt.Printf("- Values outside [Q1 - 1.5*IQR, Q3 + 1.5*IQR] are potential outliers.\n")
	lowFence := q1 - 1.5*iqr
	highFence := q3 + 1.5*iqr
	fmt.Printf("- Lower fence = %.4f, Upper fence = %.4f\n", lowFence, highFence)

	// Identify outliers
	outliers := []float64{}
	for _, v := range data {
		if v < lowFence || v > highFence {
			outliers = append(outliers, v)
		}
	}
	if len(outliers) == 0 {
		fmt.Println("- No outliers detected by 1.5*IQR rule.")
	} else {
		fmt.Printf("- Outliers: %v\n", outliers)
	}
}

func medianOfSlice(s []float64) float64 {
	m := len(s)
	if m == 0 {
		return 0
	}
	if m%2 == 0 {
		return (s[m/2-1] + s[m/2]) / 2
	}
	return s[m/2]
}
