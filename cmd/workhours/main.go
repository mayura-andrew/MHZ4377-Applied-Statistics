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

	// Calculate and display Variance and Standard Deviation
	calculateVarianceAndStdDev(workHours)

	// Summary
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    SUMMARY")
	fmt.Println("═══════════════════════════════════════════════════════════")

	mean := getMean(workHours)
	median := getMedian(workHours)
	modes := getModes(workHours)
	variance := getVariance(workHours)
	stdDev := getStdDev(workHours)

	fmt.Printf("Mean:               %.2f hours\n", mean)
	fmt.Printf("Median:             %.1f hours\n", median)
	fmt.Printf("Mode:               ")
	for i, mode := range modes {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d hours", mode)
	}
	fmt.Println()
	fmt.Printf("Variance:           %.4f hours²\n", variance)
	fmt.Printf("Standard Deviation: %.4f hours\n", stdDev)
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

func calculateVarianceAndStdDev(data []int) {
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("     4. VARIANCE AND STANDARD DEVIATION (Variability)")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()
	fmt.Println("Variance and standard deviation measure how spread out the data is.")
	fmt.Println("They tell us how much the work hours vary from the average.")
	fmt.Println()

	n := len(data)

	// Calculate mean first
	var sum int
	for _, hours := range data {
		sum += hours
	}
	mean := float64(sum) / float64(n)

	// --- VARIANCE ---
	fmt.Println("--- VARIANCE (σ²) ---")
	fmt.Println()
	fmt.Println("Variance measures the average squared distance from the mean.")
	fmt.Println("Formula: σ² = Σ(xi - μ)² / n")
	fmt.Println("         where xi = each value, μ = mean, n = number of values")
	fmt.Println()

	fmt.Printf("Step 1: We already know the mean (μ) = %.4f hours\n", mean)
	fmt.Println()

	fmt.Println("Step 2: Calculate (xi - μ)² for each employee's work hours:")
	fmt.Println()
	fmt.Println("        Employee | Hours | (Hours - Mean) | (Hours - Mean)²")
	fmt.Println("        ---------|-------|----------------|----------------")

	var sumSquaredDiff float64
	for i, hours := range data {
		diff := float64(hours) - mean
		sqDiff := diff * diff
		sumSquaredDiff += sqDiff

		// Show first 10 and last 5 to keep output manageable
		if i < 10 || i >= n-5 {
			fmt.Printf("           %2d    |  %2d   |   %8.4f     |   %10.4f\n", i+1, hours, diff, sqDiff)
		} else if i == 10 {
			fmt.Println("          ...    |  ...  |      ...       |       ...")
		}
	}

	fmt.Println()
	fmt.Printf("Step 3: Sum all the squared differences\n")
	fmt.Printf("        Σ(xi - μ)² = %.4f\n", sumSquaredDiff)
	fmt.Println()

	variance := sumSquaredDiff / float64(n)

	fmt.Printf("Step 4: Divide by n to get the variance\n")
	fmt.Printf("        Variance (σ²) = %.4f / %d = %.4f hours²\n", sumSquaredDiff, n, variance)
	fmt.Println()
	fmt.Printf("✓ Variance = %.4f hours²\n", variance)
	fmt.Println()

	// --- STANDARD DEVIATION ---
	fmt.Println("--- STANDARD DEVIATION (σ) ---")
	fmt.Println()
	fmt.Println("Standard deviation is the square root of variance.")
	fmt.Println("It's easier to interpret because it's in the same units as the data (hours).")
	fmt.Println("Formula: σ = √(σ²)")
	fmt.Println()

	stdDev := 0.0
	// Manual square root calculation visualization
	fmt.Printf("Step 1: Take the square root of the variance\n")
	fmt.Printf("        σ = √(%.4f)\n", variance)
	fmt.Println()

	// Show the calculation
	stdDev = 1.0
	for i := 0; i < 10; i++ {
		stdDev = (stdDev + variance/stdDev) / 2.0
	}

	// Use actual square root for accuracy
	import_math := func(x float64) float64 {
		if x <= 0 {
			return 0
		}
		// Newton's method for square root
		guess := x / 2.0
		for i := 0; i < 20; i++ {
			guess = (guess + x/guess) / 2.0
		}
		return guess
	}
	stdDev = import_math(variance)

	fmt.Printf("        σ = %.4f hours\n", stdDev)
	fmt.Println()
	fmt.Printf("✓ Standard Deviation = %.4f hours\n", stdDev)
	fmt.Println()

	// Interpretation
	fmt.Println("--- INTERPRETATION ---")
	fmt.Println()
	fmt.Printf("Variance (σ²) = %.4f hours²\n", variance)
	fmt.Printf("  - This measures the average squared deviation from the mean.\n")
	fmt.Printf("  - The squared units make it less intuitive to interpret directly.\n")
	fmt.Println()
	fmt.Printf("Standard Deviation (σ) = %.4f hours\n", stdDev)
	fmt.Printf("  - On average, work hours deviate from the mean by about %.2f hours.\n", stdDev)
	fmt.Printf("  - This tells us the 'typical' distance from the average of %.2f hours.\n", mean)
	fmt.Println()

	// Context
	fmt.Println("What does this mean?")
	if stdDev < 2.0 {
		fmt.Println("  - LOW variability: Work hours are very consistent across employees.")
		fmt.Println("  - Most employees work very close to the average.")
	} else if stdDev < 4.0 {
		fmt.Println("  - MODERATE variability: There is some variation in work hours.")
		fmt.Println("  - Most employees work within a reasonable range of the average.")
	} else {
		fmt.Println("  - HIGH variability: Work hours vary significantly across employees.")
		fmt.Println("  - Employees' work schedules differ substantially.")
	}
	fmt.Println()

	// Range interpretation
	fmt.Printf("Approximately 68%% of employees work between %.2f and %.2f hours\n",
		mean-stdDev, mean+stdDev)
	fmt.Printf("(within 1 standard deviation of the mean).\n")
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

func getVariance(data []int) float64 {
	n := len(data)
	if n == 0 {
		return 0
	}

	mean := getMean(data)
	var sumSquaredDiff float64

	for _, v := range data {
		diff := float64(v) - mean
		sumSquaredDiff += diff * diff
	}

	return sumSquaredDiff / float64(n)
}

func getStdDev(data []int) float64 {
	variance := getVariance(data)

	// Newton's method for square root
	if variance <= 0 {
		return 0
	}

	guess := variance / 2.0
	for i := 0; i < 20; i++ {
		guess = (guess + variance/guess) / 2.0
	}
	return guess
}
