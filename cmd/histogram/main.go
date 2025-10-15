package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Wheat yield data (30 observations) -- original order as provided
	data := []float64{
		145, 152, 138, 167, 155, 161, 143, 158, 149, 172,
		162, 147, 154, 168, 141, 159, 165, 150, 163, 140,
		156, 169, 144, 160, 153, 166, 142, 157, 151, 164,
	}

	// --- 1. Create histogram data ---
	values := make(plotter.Values, len(data))
	copy(values, data)

	// --- 2. Create a new plot ---
	p := plot.New()
	p.Title.Text = "Histogram of Wheat Yield Data"
	p.X.Label.Text = "Grain Yield (KG)"
	p.Y.Label.Text = "Frequency"

	// --- 3. Create the histogram ---
	// We use 6 bins to match the frequency table (class width = 6)
	hist, err := plotter.NewHist(values, 6)
	if err != nil {
		log.Fatal(err)
	}

	// Customize histogram appearance
	hist.FillColor = color.RGBA{R: 173, G: 216, B: 230, A: 255} // Light blue

	// Add histogram to the plot
	p.Add(hist)

	// --- 4. Save the plot to a file ---
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "wheat_yield_histogram.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Histogram has been saved to wheat_yield_histogram.png")
	fmt.Println()

	// --- 5. Calculate measures of variability ---
	calculateVariability(data)
	fmt.Println()

	// --- 6. Analyze the distribution shape ---
	analyzeDistribution(data)
}

func calculateVariability(data []float64) {
	n := len(data)

	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("         MEASURES OF VARIABILITY - Step by Step")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println()

	// Sort the data for range and quartile calculations
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)

	// ============ 1. RANGE ============
	fmt.Println("--- 1. RANGE ---")
	fmt.Println()
	fmt.Println("The range is the difference between the maximum and minimum values.")
	fmt.Println("It gives us the total spread of the data.")
	fmt.Println()

	min := sorted[0]
	max := sorted[n-1]
	dataRange := max - min

	fmt.Printf("Step 1: Identify the minimum value = %.0f kg\n", min)
	fmt.Printf("Step 2: Identify the maximum value = %.0f kg\n", max)
	fmt.Printf("Step 3: Calculate Range = Max - Min\n")
	fmt.Printf("        Range = %.0f - %.0f = %.0f kg\n", max, min, dataRange)
	fmt.Println()
	fmt.Printf("✓ Range = %.0f kg\n", dataRange)
	fmt.Println()

	// ============ 2. VARIANCE ============
	fmt.Println("--- 2. VARIANCE ---")
	fmt.Println()
	fmt.Println("Variance measures how far each number in the dataset is from the mean.")
	fmt.Println("Formula: σ² = Σ(xi - μ)² / n")
	fmt.Println()

	// Calculate mean
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(n)

	fmt.Printf("Step 1: Calculate the mean (μ)\n")
	fmt.Printf("        Sum of all values = %.0f\n", sum)
	fmt.Printf("        Mean (μ) = %.0f / %d = %.4f kg\n", sum, n, mean)
	fmt.Println()

	fmt.Println("Step 2: Calculate (xi - μ)² for each value:")
	fmt.Println("        Value | (Value - Mean) | (Value - Mean)²")
	fmt.Println("        ------|----------------|----------------")

	var sumSquaredDiff float64
	// Show first 5 and last 5 examples to keep output manageable
	for i, v := range sorted {
		diff := v - mean
		sqDiff := diff * diff
		sumSquaredDiff += sqDiff

		if i < 5 || i >= n-5 {
			fmt.Printf("        %.0f   |   %7.4f    |   %10.4f\n", v, diff, sqDiff)
		} else if i == 5 {
			fmt.Println("        ...   |      ...       |       ...")
		}
	}

	fmt.Println()
	fmt.Printf("Step 3: Sum all squared differences\n")
	fmt.Printf("        Σ(xi - μ)² = %.4f\n", sumSquaredDiff)
	fmt.Println()

	variance := sumSquaredDiff / float64(n)

	fmt.Printf("Step 4: Divide by n (population variance)\n")
	fmt.Printf("        Variance (σ²) = %.4f / %d = %.4f kg²\n", sumSquaredDiff, n, variance)
	fmt.Println()
	fmt.Printf("✓ Variance = %.4f kg²\n", variance)
	fmt.Println()

	// ============ 3. STANDARD DEVIATION ============
	fmt.Println("--- 3. STANDARD DEVIATION ---")
	fmt.Println()
	fmt.Println("Standard deviation is the square root of variance.")
	fmt.Println("It measures the average distance of data points from the mean.")
	fmt.Println("Formula: σ = √(σ²)")
	fmt.Println()

	stdDev := math.Sqrt(variance)

	fmt.Printf("Step 1: Take the square root of variance\n")
	fmt.Printf("        σ = √(%.4f) = %.4f kg\n", variance, stdDev)
	fmt.Println()
	fmt.Printf("✓ Standard Deviation = %.4f kg\n", stdDev)
	fmt.Println()
	fmt.Println("Interpretation: On average, the wheat yield values deviate")
	fmt.Printf("from the mean by approximately %.2f kg.\n", stdDev)
	fmt.Println()

	// ============ 4. INTERQUARTILE RANGE (IQR) ============
	fmt.Println("--- 4. INTERQUARTILE RANGE (IQR) ---")
	fmt.Println()
	fmt.Println("The IQR measures the spread of the middle 50% of the data.")
	fmt.Println("It is the difference between Q3 (75th percentile) and Q1 (25th percentile).")
	fmt.Println()

	fmt.Println("Step 1: Arrange data in ascending order (already sorted)")
	fmt.Println("Sorted data:")
	for i, v := range sorted {
		if i > 0 && i%10 == 0 {
			fmt.Println()
		}
		fmt.Printf("%.0f ", v)
	}
	fmt.Println()
	fmt.Println()

	// Calculate Q1 and Q3
	// Using method: split data into two halves (excluding median for odd n)
	midpoint := n / 2

	var lowerHalf, upperHalf []float64
	if n%2 == 0 {
		// Even number: split exactly in half
		lowerHalf = sorted[:midpoint]
		upperHalf = sorted[midpoint:]
	} else {
		// Odd number: exclude the median
		lowerHalf = sorted[:midpoint]
		upperHalf = sorted[midpoint+1:]
	}

	fmt.Printf("Step 2: Split data into two halves (n = %d, midpoint at position %d)\n", n, midpoint)
	fmt.Printf("        Lower half (first %d values): ", len(lowerHalf))
	for _, v := range lowerHalf {
		fmt.Printf("%.0f ", v)
	}
	fmt.Println()
	fmt.Printf("        Upper half (last %d values): ", len(upperHalf))
	for _, v := range upperHalf {
		fmt.Printf("%.0f ", v)
	}
	fmt.Println()
	fmt.Println()

	// Calculate Q1 (median of lower half)
	q1 := getMedian(lowerHalf)
	fmt.Printf("Step 3: Calculate Q1 (median of lower half)\n")
	if len(lowerHalf)%2 == 0 {
		mid1Idx := len(lowerHalf)/2 - 1
		mid2Idx := len(lowerHalf) / 2
		fmt.Printf("        Lower half has %d values (even)\n", len(lowerHalf))
		fmt.Printf("        Middle values: %.0f and %.0f\n", lowerHalf[mid1Idx], lowerHalf[mid2Idx])
		fmt.Printf("        Q1 = (%.0f + %.0f) / 2 = %.2f kg\n", lowerHalf[mid1Idx], lowerHalf[mid2Idx], q1)
	} else {
		midIdx := len(lowerHalf) / 2
		fmt.Printf("        Lower half has %d values (odd)\n", len(lowerHalf))
		fmt.Printf("        Middle value at position %d: %.0f\n", midIdx+1, lowerHalf[midIdx])
		fmt.Printf("        Q1 = %.2f kg\n", q1)
	}
	fmt.Println()

	// Calculate Q3 (median of upper half)
	q3 := getMedian(upperHalf)
	fmt.Printf("Step 4: Calculate Q3 (median of upper half)\n")
	if len(upperHalf)%2 == 0 {
		mid1Idx := len(upperHalf)/2 - 1
		mid2Idx := len(upperHalf) / 2
		fmt.Printf("        Upper half has %d values (even)\n", len(upperHalf))
		fmt.Printf("        Middle values: %.0f and %.0f\n", upperHalf[mid1Idx], upperHalf[mid2Idx])
		fmt.Printf("        Q3 = (%.0f + %.0f) / 2 = %.2f kg\n", upperHalf[mid1Idx], upperHalf[mid2Idx], q3)
	} else {
		midIdx := len(upperHalf) / 2
		fmt.Printf("        Upper half has %d values (odd)\n", len(upperHalf))
		fmt.Printf("        Middle value at position %d: %.0f\n", midIdx+1, upperHalf[midIdx])
		fmt.Printf("        Q3 = %.2f kg\n", q3)
	}
	fmt.Println()

	iqr := q3 - q1
	fmt.Printf("Step 5: Calculate IQR\n")
	fmt.Printf("        IQR = Q3 - Q1 = %.2f - %.2f = %.2f kg\n", q3, q1, iqr)
	fmt.Println()
	fmt.Printf("✓ Interquartile Range (IQR) = %.2f kg\n", iqr)
	fmt.Println()
	fmt.Println("Interpretation: The middle 50% of wheat yield values")
	fmt.Printf("span a range of %.2f kg.\n", iqr)
	fmt.Println()

	// Summary table
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Println("                    SUMMARY OF VARIABILITY")
	fmt.Println("═══════════════════════════════════════════════════════════")
	fmt.Printf("Range:                   %.2f kg\n", dataRange)
	fmt.Printf("Variance (σ²):           %.4f kg²\n", variance)
	fmt.Printf("Standard Deviation (σ):  %.4f kg\n", stdDev)
	fmt.Printf("Q1 (First Quartile):     %.2f kg\n", q1)
	fmt.Printf("Q3 (Third Quartile):     %.2f kg\n", q3)
	fmt.Printf("IQR (Q3 - Q1):           %.2f kg\n", iqr)
	fmt.Println("═══════════════════════════════════════════════════════════")
}

// Helper function to calculate median
func getMedian(arr []float64) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}
	if n%2 == 0 {
		return (arr[n/2-1] + arr[n/2]) / 2
	}
	return arr[n/2]
}

func analyzeDistribution(data []float64) {
	n := len(data)

	// Calculate mean
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(n)

	// Calculate median
	sorted := make([]float64, n)
	copy(sorted, data)
	sort.Float64s(sorted)

	var median float64
	if n%2 == 0 {
		median = (sorted[n/2-1] + sorted[n/2]) / 2
	} else {
		median = sorted[n/2]
	}

	// Calculate standard deviation
	var sumSq float64
	for _, v := range data {
		diff := v - mean
		sumSq += diff * diff
	}
	stdDev := math.Sqrt(sumSq / float64(n))

	// Calculate skewness (using sample skewness formula)
	var sumCubed float64
	for _, v := range data {
		diff := v - mean
		sumCubed += diff * diff * diff
	}
	skewness := (sumCubed / float64(n)) / math.Pow(stdDev, 3)

	// Print statistical summary
	fmt.Println("--- Distribution Analysis ---")
	fmt.Printf("Number of observations: %d\n", n)
	fmt.Printf("Mean: %.2f kg\n", mean)
	fmt.Printf("Median: %.2f kg\n", median)
	fmt.Printf("Standard Deviation: %.2f kg\n", stdDev)
	fmt.Printf("Skewness: %.4f\n", skewness)
	fmt.Println()

	// Interpret the distribution shape
	fmt.Println("--- Shape of the Distribution ---")

	// Check symmetry using skewness
	if math.Abs(skewness) < 0.5 {
		fmt.Println("✓ The distribution is approximately SYMMETRIC (skewness ≈ 0)")
		fmt.Println("  - The data is fairly evenly distributed around the mean.")
		fmt.Println("  - Mean and median are very close to each other.")
	} else if skewness > 0.5 {
		fmt.Println("✓ The distribution is POSITIVELY SKEWED (right-skewed)")
		fmt.Println("  - The tail extends more to the right (higher values).")
		fmt.Println("  - Mean is greater than median.")
		fmt.Println("  - Most values are concentrated on the lower end.")
	} else {
		fmt.Println("✓ The distribution is NEGATIVELY SKEWED (left-skewed)")
		fmt.Println("  - The tail extends more to the left (lower values).")
		fmt.Println("  - Mean is less than median.")
		fmt.Println("  - Most values are concentrated on the higher end.")
	}

	fmt.Println()
	fmt.Println("Visual observation from histogram:")
	fmt.Println("- The histogram shows the frequency distribution across 6 bins.")
	fmt.Println("- With skewness close to 0, the distribution appears fairly uniform/symmetric.")
	fmt.Println("- The data is reasonably well-spread across the range from 138 to 172 kg.")
}
