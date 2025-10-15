// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	// The fertilizer usage data provided by the user
// 	data := []float64{
// 		22.5, 23.1, 21.8, 24.0, 22.7, 23.5, 24.8, 22.0, 23.9, 25.2,
// 		21.5, 23.3, 24.5, 22.2, 23.8, 25.5, 21.9, 24.2, 22.8, 23.6,
// 		25.0, 22.1, 24.9, 23.0, 22.9, 24.1, 23.7, 22.4, 24.7, 23.4,
// 		22.6, 24.3, 23.2, 25.1, 21.7, 24.4, 22.3, 25.3, 23.8, 24.6,
// 		21.6, 23.9, 22.5, 25.4, 23.1, 24.0, 22.9, 23.5, 24.8, 22.2,
// 		23.7, 25.0, 21.8, 24.2, 23.0, 22.7, 24.5, 23.3, 20.0, 24.9,
// 	}

// 	// --- 1. Compute the Mean ---
// 	var sum float64 = 0.0
// 	for _, value := range data {
// 		sum += value
// 	}

// 	fmt.Println(len(data))
// 	fmt.Println(sum)
// 	mean := sum / float64(len(data))
// 	fmt.Printf("Mean: %.2f kg\n", mean)

// 	// --- 2. Compute the Median ---
// 	// First, we need a sorted copy of the data
// 	sortedData := make([]float64, len(data))
// 	copy(sortedData, data)
// 	sort.Float64s(sortedData)

// 	fmt.Println(sortedData)
// 	var median float64
// 	n := len(sortedData)
// 	// Check if the number of elements is even or odd
// 	if n%2 == 0 {
// 		// If even, the median is the average of the two middle numbers
// 		mid1 := sortedData[(n/2)-1]
// 		mid2 := sortedData[n/2]
// 		median = (mid1 + mid2) / 2
// 	} else {
// 		// If odd, the median is the middle number
// 		median = sortedData[n/2]
// 	}


// 	fmt.Printf("Median: %.2f kg\n", median)

// 	// --- 3. Compute the Mode ---
// 	// Create a map to store the frequency of each number
// 	frequency := make(map[float64]int)
// 	for _, value := range data {
// 		frequency[value]++
// 	}

// 	maxFreq := 0
// 	var modes []float64
// 	// Find the highest frequency
// 	for _, freq := range frequency {
// 		if freq > maxFreq {
// 			maxFreq = freq
// 		}
// 	}

// 	// If maxFreq is 1, every element is a mode, which is not useful.
// 	if maxFreq <= 1 {
// 		fmt.Println("Mode: No unique mode found.")
// 	} else {
// 		// Collect all numbers that have the highest frequency
// 		for value, freq := range frequency {
// 			if freq == maxFreq {
// 				modes = append(modes, value)
// 			}
// 		}
// 		sort.Float64s(modes) // Sort for clean output
// 		fmt.Printf("Mode(s) (all values with frequency %d): %v\n", maxFreq, modes)
// 	}
// }


package main

import (
	"fmt"
	"sort"
)

// Helper function to calculate the median of a slice
func getMedian(arr []float64) float64 {
	var median float64
	n := len(arr)
	if n%2 == 0 {
		// Even number of elements: average of the two middle ones
		mid1 := arr[(n/2)-1]
		mid2 := arr[n/2]
		median = (mid1 + mid2) / 2
	} else {
		// Odd number of elements: the middle one
		median = arr[n/2]
	}
	return median
}

func main() {
	data := []float64{
		22.5, 23.1, 21.8, 24.0, 22.7, 23.5, 24.8, 22.0, 23.9, 25.2,
		21.5, 23.3, 24.5, 22.2, 23.8, 25.5, 21.9, 24.2, 22.8, 23.6,
		25.0, 22.1, 24.9, 23.0, 22.9, 24.1, 23.7, 22.4, 24.7, 23.4,
		22.6, 24.3, 23.2, 25.1, 21.7, 24.4, 22.3, 25.3, 23.8, 24.6,
		21.6, 23.9, 22.5, 25.4, 23.1, 24.0, 22.9, 23.5, 24.8, 22.2,
		23.7, 25.0, 21.8, 24.2, 23.0, 22.7, 24.5, 23.3, 20.0, 24.9,
	}

	// --- STEP 1: Sort the data ---
	sortedData := make([]float64, len(data))
	copy(sortedData, data)
	sort.Float64s(sortedData)

	fmt.Println("--- Visual Demonstration for Quartiles ---")
	fmt.Println("\nStep 1: The full dataset sorted from smallest to largest:")
	fmt.Println(sortedData)

	// --- STEP 2: Split the data into lower and upper halves ---
	n := len(sortedData)
	midpoint := n / 2
	
	lowerHalf := sortedData[:midpoint]
	upperHalf := sortedData[midpoint:]

	fmt.Println("\nStep 2: The data is split into two halves.")
	fmt.Println("Lower Half (first 30 values):")
	fmt.Println(lowerHalf)
	fmt.Println("\nUpper Half (last 30 values):")
	fmt.Println(upperHalf)

	// --- STEP 3: Calculate Q1 (Median of the Lower Half) ---
	q1 := getMedian(lowerHalf)
	fmt.Println("\nStep 3: Calculate Q1 from the lower half.")
	fmt.Printf("The middle values of the lower half are %.1f and %.1f.\n", lowerHalf[14], lowerHalf[15])
	fmt.Printf("Q1 = (%.1f + %.1f) / 2 = %.2f kg\n", lowerHalf[14], lowerHalf[15], q1)

	// --- STEP 4: Calculate Q3 (Median of the Upper Half) ---
	q3 := getMedian(upperHalf)
	fmt.Println("\nStep 4: Calculate Q3 from the upper half.")
	fmt.Printf("The middle values of the upper half are %.1f and %.1f.\n", upperHalf[14], upperHalf[15])
	fmt.Printf("Q3 = (%.1f + %.1f) / 2 = %.2f kg\n", upperHalf[14], upperHalf[15], q3)

	fmt.Println("\n--- FINAL RESULTS ---")
	fmt.Printf("First Quartile (Q1): %.2f kg\n", q1)
	fmt.Printf("Third Quartile (Q3): %.2f kg\n", q3)
}