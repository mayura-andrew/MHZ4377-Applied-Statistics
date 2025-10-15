package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Sample wheat yield data (30 observations) -- original order as provided
	data := []int{
		145,
		152,
		138,
		167,
		155,
		161,
		143,
		158,
		149,
		172,
		162,
		147,
		154,
		168,
		141,
		159,
		165,
		150,
		163,
		140,
		156,
		169,
		144,
		160,
		153,
		166,
		142,
		157,
		151,
		164,
	}

	// Print the original data in the given order
	fmt.Println("Original data (given order):")
	for _, v := range data {
		fmt.Printf("%d ", v)
	}
	fmt.Println("\n")

	// Raw frequency count for individual values (value -> count)
	rawFreq := make(map[int]int)
	for _, v := range data {
		rawFreq[v]++
	}
	// Print raw frequency table sorted by value
	fmt.Println("Raw frequency (value : count):")
	// collect and sort keys
	keys := make([]int, 0, len(rawFreq))
	for k := range rawFreq {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("%d : %d\n", k, rawFreq[k])
	}
	fmt.Println()

	sort.Ints(data)
	n := len(data)
	min := data[0]
	max := data[n-1]
	rng := max - min

	// Suggested number of classes by square-root rule (for info)
	suggestedClasses := int(math.Round(math.Sqrt(float64(n))))

	// We'll use 6 classes as in your specification
	numClasses := 6

	// Class width: round up range/numClasses to a convenient whole number
	width := int(math.Ceil(float64(rng) / float64(numClasses)))
	if width < 1 {
		width = 1
	}

	// Build intervals (inclusive): start .. end where end = start + width - 1
	intervals := make([][2]int, numClasses)
	for i := 0; i < numClasses; i++ {
		start := min + i*width
		end := start + width - 1
		// Keep the interval endpoints consistent (inclusive). The last
		// interval may extend beyond the max value, which is fine for
		// grouping — it still includes the maximum because max <= end.
		intervals[i][0] = start
		intervals[i][1] = end
	}

	// Tally frequencies
	counts := make([]int, numClasses)
	for _, v := range data {
		for i, intr := range intervals {
			if v >= intr[0] && v <= intr[1] {
				counts[i]++
				break
			}
		}
	}

	// Print the frequency table
	fmt.Println("| Grain Yield (KG) | Frequency |")
	fmt.Println("| :--- | :---: |")
	total := 0
	for i := 0; i < numClasses; i++ {
		fmt.Printf("| %d - %d | %d |\n", intervals[i][0], intervals[i][1], counts[i])
		total += counts[i]
	}
	fmt.Println("| **Total** | **", total, "** |")

	// Print the step-by-step guide summary
	fmt.Println("\n---\n")
	fmt.Println("Step-by-step summary:")
	fmt.Printf("1) Number of observations: %d\n", n)
	fmt.Printf("2) Minimum value: %d\n", min)
	fmt.Printf("3) Maximum value: %d\n", max)
	fmt.Printf("4) Range = Max - Min = %d - %d = %d\n", max, min, rng)
	fmt.Printf("5) Suggested classes by sqrt rule ≈ %d (we use %d)\n", suggestedClasses, numClasses)
	fmt.Printf("6) Class width (rounded up): ceil(%d / %d) = %d\n", rng, numClasses, width)
	fmt.Println("7) Class intervals (inclusive) and frequencies printed above.")

	// Sanity check
	if total != n {
		fmt.Printf("\nWarning: total frequency (%d) != number of observations (%d)\n", total, n)
	} else {
		fmt.Println("\nSanity check: total frequency equals number of observations.")
	}
}
