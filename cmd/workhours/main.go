package main

import (
	"fmt"
	"sort"
)

func main() {
	// Employee work hours data (30 employees)
	workHours := []int{
		38, 42, 35, 40, 44, 37, 41, 39, 45, 36,
		43, 38, 40, 42, 35, 44, 39, 41, 37, 43,
		36, 45, 38, 40, 42, 39, 41, 37, 44, 40,
	}

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("    ANALYZING WEEKLY WORK HOURS - 30 Employees")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()

	// Display the data
	fmt.Println("Original data (Employee ID 1-30):")
	for i, hours := range workHours {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", hours)
	}
	fmt.Println()
	fmt.Println()

	// Calculate and display Mean, Median, and Mode
	calculateMean(workHours)
	calculateMedian(workHours)
	calculateMode(workHours)

	// Summary
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    SUMMARY")
	fmt.Println("═══════════════════════════════════════════════════════════")

	mean := getMean(workHours)
	median := getMedian(workHours)
	modes := getModes(workHours)

	fmt.Printf("Mean:   %.2f hours\n", mean)
	fmt.Printf("Median: %.1f hours\n", median)
	fmt.Printf("Mode:   ")
	for i, mode := range modes {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d hours", mode)
	}
	fmt.Println()
	fmt.Println("═══════════════════════════════════════════════════════════")
}

func calculateMean(data []int) {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    1. MEAN (Average)")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("The mean is the average of all values.")
	fmt.Println("Formula: Mean = (Sum of all values) / (Number of values)")
	fmt.Println()

	n := len(data)
	var sum int

	fmt.Println("Step 1: Add all the work hours together")
	fmt.Print("        Sum = ")
	for i, hours := range data {
		sum += hours
		if i < len(data)-1 {
			fmt.Printf("%d + ", hours)
		} else {
			fmt.Printf("%d", hours)
		}
		if (i+1)%10 == 0 && i < len(data)-1 {
			fmt.Println()
			fmt.Print("              ")
		}
	}
	fmt.Println()
	fmt.Printf("        Sum = %d hours\n", sum)
	fmt.Println()

	mean := float64(sum) / float64(n)

	fmt.Printf("Step 2: Divide the sum by the number of employees\n")
	fmt.Printf("        Mean = %d / %d = %.4f hours\n", sum, n, mean)
	fmt.Println()
	fmt.Printf("✓ Mean = %.2f hours per week\n", mean)
	fmt.Println()
	fmt.Printf("Interpretation: On average, employees work %.2f hours per week.\n", mean)
	fmt.Println()
}

func calculateMedian(data []int) {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    2. MEDIAN (Middle Value)")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("The median is the middle value when data is arranged in order.")
	fmt.Println("If there's an even number of values, it's the average of the two middle values.")
	fmt.Println()

	n := len(data)
	sorted := make([]int, n)
	copy(sorted, data)
	sort.Ints(sorted)

	fmt.Println("Step 1: Arrange all values in ascending order")
	fmt.Println("Sorted work hours:")
	for i, hours := range sorted {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%d ", hours)
	}
	fmt.Println()
	fmt.Println()

	var median float64

	fmt.Printf("Step 2: Find the middle position(s)\n")
	fmt.Printf("        Number of values (n) = %d (even number)\n", n)
	fmt.Printf("        Middle positions: %d and %d\n", n/2, (n/2)+1)
	fmt.Println()

	if n%2 == 0 {
		// Even number of values
		mid1Idx := (n / 2) - 1
		mid2Idx := n / 2
		mid1 := sorted[mid1Idx]
		mid2 := sorted[mid2Idx]
		median = float64(mid1+mid2) / 2.0

		fmt.Printf("Step 3: Since we have an even number of values, average the two middle values\n")
		fmt.Printf("        Value at position %d: %d hours\n", mid1Idx+1, mid1)
		fmt.Printf("        Value at position %d: %d hours\n", mid2Idx+1, mid2)
		fmt.Printf("        Median = (%d + %d) / 2 = %.1f hours\n", mid1, mid2, median)
	} else {
		// Odd number of values
		midIdx := n / 2
		median = float64(sorted[midIdx])
		fmt.Printf("Step 3: The middle value is at position %d\n", midIdx+1)
		fmt.Printf("        Median = %.1f hours\n", median)
	}

	fmt.Println()
	fmt.Printf("✓ Median = %.1f hours per week\n", median)
	fmt.Println()
	fmt.Println("Interpretation: Half of the employees work less than or equal to")
	fmt.Printf("%.1f hours, and half work more than or equal to %.1f hours.\n", median, median)
	fmt.Println()
}

func calculateMode(data []int) {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    3. MODE (Most Frequent)")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("The mode is the value(s) that appear most frequently in the dataset.")
	fmt.Println()

	// Create frequency map
	frequency := make(map[int]int)
	for _, hours := range data {
		frequency[hours]++
	}

	// Sort keys for organized display
	keys := make([]int, 0, len(frequency))
	for k := range frequency {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	fmt.Println("Step 1: Count how many times each value appears")
	fmt.Println()
	fmt.Println("        Work Hours | Frequency (Count)")
	fmt.Println("        -----------|------------------")

	maxFreq := 0
	for _, hours := range keys {
		count := frequency[hours]
		fmt.Printf("             %2d    |        %d\n", hours, count)
		if count > maxFreq {
			maxFreq = count
		}
	}
	fmt.Println()

	fmt.Printf("Step 2: Identify the highest frequency\n")
	fmt.Printf("        Highest frequency = %d\n", maxFreq)
	fmt.Println()

	// Find all modes (values with highest frequency)
	modes := []int{}
	for _, hours := range keys {
		if frequency[hours] == maxFreq {
			modes = append(modes, hours)
		}
	}

	fmt.Printf("Step 3: Find all values with the highest frequency\n")

	if maxFreq == 1 {
		fmt.Println("        All values appear exactly once.")
		fmt.Println("        There is NO MODE (no value repeats).")
		fmt.Println()
		fmt.Println("✓ Mode: None (all values are unique)")
	} else if len(modes) == 1 {
		fmt.Printf("        The value %d appears %d times (most frequent)\n", modes[0], maxFreq)
		fmt.Println()
		fmt.Printf("✓ Mode = %d hours per week\n", modes[0])
		fmt.Println()
		fmt.Printf("Interpretation: %d hours is the most common weekly work time,\n", modes[0])
		fmt.Printf("appearing %d times in the dataset.\n", maxFreq)
	} else {
		fmt.Printf("        Multiple values appear %d times (tied for most frequent):\n", maxFreq)
		fmt.Print("        ")
		for i, mode := range modes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d hours", mode)
		}
		fmt.Println()
		fmt.Println()
		fmt.Print("✓ Modes (multimodal) = ")
		for i, mode := range modes {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%d hours", mode)
		}
		fmt.Println()
		fmt.Println()
		fmt.Printf("Interpretation: The dataset has multiple modes. These values\n")
		fmt.Printf("each appear %d times, which is the highest frequency.\n", maxFreq)
	}
	fmt.Println()
}

// Helper functions for summary
func getMean(data []int) float64 {
	var sum int
	for _, v := range data {
		sum += v
	}
	return float64(sum) / float64(len(data))
}

func getMedian(data []int) float64 {
	n := len(data)
	sorted := make([]int, n)
	copy(sorted, data)
	sort.Ints(sorted)

	if n%2 == 0 {
		return float64(sorted[n/2-1]+sorted[n/2]) / 2.0
	}
	return float64(sorted[n/2])
}

func getModes(data []int) []int {
	frequency := make(map[int]int)
	for _, v := range data {
		frequency[v]++
	}

	maxFreq := 0
	for _, count := range frequency {
		if count > maxFreq {
			maxFreq = count
		}
	}

	if maxFreq == 1 {
		return []int{} // No mode
	}

	modes := []int{}
	for value, count := range frequency {
		if count == maxFreq {
			modes = append(modes, value)
		}
	}
	sort.Ints(modes)
	return modes
}
