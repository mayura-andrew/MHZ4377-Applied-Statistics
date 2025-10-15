package main

import (
	"fmt"
	"image/color"
	"math"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Marks dataset
	marks := []float64{10, 43, 25, 34, 31, 9, 25, 30, 28, 12, 26, 19, 11, 8, 35, 41, 28, 19, 8, 21, 20, 47, 32, 28, 21}
	n := len(marks)
	sort.Float64s(marks)

	fmt.Printf("Sorted marks (%d values): %v\n\n", n, marks)

	// 1) Number of classes by sqrt(n)
	numClasses := int(math.Round(math.Sqrt(float64(n))))
	if numClasses < 1 {
		numClasses = 1
	}
	fmt.Printf("Number of classes (sqrt(n) rounded): %d\n", numClasses)

	// range and class width
	min := marks[0]
	max := marks[n-1]
	rangeVal := max - min
	width := math.Ceil(rangeVal / float64(numClasses))
	if width < 1 {
		width = 1
	}
	fmt.Printf("Range = %.0f - %.0f = %.0f\n", max, min, rangeVal)
	fmt.Printf("Class width (rounded up) = %.0f\n\n", width)

	// Build class intervals (inclusive lower, exclusive upper except last)
	types := make([][2]float64, numClasses)
	start := min
	for i := 0; i < numClasses; i++ {
		end := start + width
		// last class include max
		if i == numClasses-1 {
			end = max + 0.0001
		}
		types[i][0] = start
		types[i][1] = end
		start = end
	}

	// Tally frequencies
	freq := make([]int, numClasses)
	for _, v := range marks {
		for i, intr := range types {
			if v >= intr[0] && v < intr[1] {
				freq[i]++
				break
			}
		}
	}

	// Print frequency table
	fmt.Println("Frequency table:")
	fmt.Println("Class interval\tFrequency")
	total := 0
	for i := 0; i < numClasses; i++ {
		low := types[i][0]
		high := types[i][1]
		fmt.Printf("%.0f - %.0f\t\t%d\n", low, math.Floor(high-0.0001), freq[i])
		total += freq[i]
	}
	fmt.Printf("Total\t\t%d\n\n", total)

	// 2) Draw histogram with numClasses bins
	vals := plotter.Values(marks)
	p := plot.New()
	p.Title.Text = "Histogram of Student Marks"
	p.X.Label.Text = "Marks"
	p.Y.Label.Text = "Frequency"

	hist, err := plotter.NewHist(vals, numClasses)
	if err != nil {
		panic(err)
	}
	hist.FillColor = color.RGBA{R: 100, G: 149, B: 237, A: 255}
	p.Add(hist)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "marks_histogram.png"); err != nil {
		panic(err)
	}
	fmt.Println("Saved marks_histogram.png")

	// 3) Measures of variability and central tendency
	mean := mean(marks)
	median := median(marks)
	modes := mode(marks)
	varPop := variancePopulation(marks)
	varSample := varianceSample(marks)
	stdPop := math.Sqrt(varPop)
	stdSample := math.Sqrt(varSample)
	q1 := quartile(marks, 0.25)
	q3 := quartile(marks, 0.75)
	iqr := q3 - q1

	fmt.Println()
	fmt.Println("Central tendency:")
	if len(modes) == 0 {
		fmt.Println("Mode: None (all values unique)")
	} else {
		fmt.Printf("Mode(s): %v\n", modes)
	}
	fmt.Printf("Mean: %.4f\n", mean)
	fmt.Printf("Median: %.4f\n", median)

	fmt.Println()
	fmt.Println("Measures of variability:")
	fmt.Printf("Range: %.4f\n", rangeVal)
	fmt.Printf("Population Variance: %.4f\n", varPop)
	fmt.Printf("Sample Variance: %.4f\n", varSample)
	fmt.Printf("Population Std Dev: %.4f\n", stdPop)
	fmt.Printf("Sample Std Dev: %.4f\n", stdSample)
	fmt.Printf("Q1: %.4f\n", q1)
	fmt.Printf("Q3: %.4f\n", q3)
	fmt.Printf("IQR: %.4f\n", iqr)

	// 4) Suitable measure for variability: choose based on skewness
	skew := skewness(marks)
	fmt.Println()
	fmt.Printf("Skewness: %.4f\n", skew)
	if math.Abs(skew) < 0.5 {
		fmt.Println("Distribution roughly symmetric -> standard deviation (sample) is a suitable measure of variability.")
	} else {
		fmt.Println("Distribution skewed -> IQR is a more robust measure of variability.")
	}

	// Comment on shape
	if skew > 0.5 {
		fmt.Println("Histogram shape: Positively skewed (right-skewed).")
	} else if skew < -0.5 {
		fmt.Println("Histogram shape: Negatively skewed (left-skewed).")
	} else {
		fmt.Println("Histogram shape: Approximately symmetric.")
	}
}

func mean(a []float64) float64 {
	var s float64
	for _, v := range a {
		s += v
	}
	return s / float64(len(a))
}

func median(a []float64) float64 {
	n := len(a)
	if n%2 == 0 {
		return (a[n/2-1] + a[n/2]) / 2
	}
	return a[n/2]
}

func mode(a []float64) []float64 {
	freq := make(map[float64]int)
	for _, v := range a {
		freq[v]++
	}
	maxf := 0
	for _, f := range freq {
		if f > maxf {
			maxf = f
		}
	}
	if maxf == 1 {
		return []float64{}
	}
	m := []float64{}
	for v, f := range freq {
		if f == maxf {
			m = append(m, v)
		}
	}
	sort.Float64s(m)
	return m
}

func variancePopulation(a []float64) float64 {
	m := mean(a)
	var s float64
	for _, v := range a {
		d := v - m
		s += d * d
	}
	return s / float64(len(a))
}

func varianceSample(a []float64) float64 {
	m := mean(a)
	var s float64
	for _, v := range a {
		d := v - m
		s += d * d
	}
	return s / float64(len(a)-1)
}

func quartile(a []float64, p float64) float64 {
	// using linear interpolation method
	n := len(a)
	pos := p * float64(n+1)
	if pos <= 1 {
		return a[0]
	}
	if pos >= float64(n) {
		return a[n-1]
	}
	lower := int(math.Floor(pos)) - 1
	upper := lower + 1
	frac := pos - math.Floor(pos)
	return a[lower] + frac*(a[upper]-a[lower])
}

func skewness(a []float64) float64 {
	m := mean(a)
	n := float64(len(a))
	var s2, s3 float64
	for _, v := range a {
		d := v - m
		s2 += d * d
		s3 += d * d * d
	}
	s2 /= n
	s3 /= n
	sd := math.Sqrt(s2)
	if sd == 0 {
		return 0
	}
	return s3 / (sd * sd * sd)
}
