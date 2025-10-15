package main

import (
	"fmt"
	"image/color"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Work hours data (30 employees)
	workHours := []int{
		38, 42, 35, 40, 44, 37, 41, 39, 45, 36,
		43, 38, 40, 42, 35, 44, 39, 41, 37, 43,
		36, 45, 38, 40, 42, 39, 41, 37, 44, 40,
	}

	// Compute frequency per hour
	freq := make(map[int]int)
	for _, h := range workHours {
		freq[h]++
	}

	// Collect sorted hour values
	keys := make([]int, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	values := make(plotter.Values, len(keys))
	labels := make([]string, len(keys))
	for i, k := range keys {
		values[i] = float64(freq[k])
		labels[i] = fmt.Sprintf("%d", k)
	}

	p := plot.New()
	p.Title.Text = "Distribution of Weekly Work Hours"
	p.Y.Label.Text = "Frequency (Number of Employees)"
	p.X.Label.Text = "Work Hours"

	// Create bar chart
	bar, err := plotter.NewBarChart(values, vg.Points(20))
	if err != nil {
		panic(err)
	}
	bar.Color = color.RGBA{R: 100, G: 149, B: 237, A: 255} // cornflower blue
	bar.LineStyle.Width = vg.Points(0.5)

	p.Add(bar)

	// Adjust X axis to show labels centered under bars
	p.NominalX(labels...)

	// Optional: improve layout
	p.X.Tick.Label.Rotation = 0
	p.X.Tick.Label.YAlign = -0.5
	p.X.Tick.Label.XAlign = 0.5

	// Add a light grid for readability
	grid := plotter.NewGrid()
	p.Add(grid)

	// Compute maximum bar height for line drawing
	maxY := 0.0
	for i := range values {
		if values[i] > maxY {
			maxY = values[i]
		}
	}

	// Compute mean and median and map them to bar-index coordinates
	var sumH float64
	for _, h := range workHours {
		sumH += float64(h)
	}
	mean := sumH / float64(len(workHours))

	// Build full sorted unique hour keys list (the keys slice is sorted hour values)
	// Map a numeric hour to x-position (index) using linear interpolation between keys
	hourToPos := func(x float64) float64 {
		if len(keys) == 0 {
			return 0
		}
		if x <= float64(keys[0]) {
			return 0
		}
		last := float64(keys[len(keys)-1])
		if x >= last {
			return float64(len(keys) - 1)
		}
		for j := 0; j < len(keys)-1; j++ {
			a := float64(keys[j])
			b := float64(keys[j+1])
			if x >= a && x <= b {
				// interpolate between j and j+1
				frac := (x - a) / (b - a)
				return float64(j) + frac
			}
		}
		return 0
	}

	// median calculation
	sortedVals := make([]int, len(workHours))
	copy(sortedVals, workHours)
	sort.Ints(sortedVals)
	var median float64
	nn := len(sortedVals)
	if nn%2 == 0 {
		median = float64(sortedVals[nn/2-1]+sortedVals[nn/2]) / 2.0
	} else {
		median = float64(sortedVals[nn/2])
	}

	meanPos := hourToPos(mean)
	medianPos := hourToPos(median)

	// Create vertical lines for mean and median
	lineHeight := maxY + 1.0
	meanLinePts := plotter.XYs{{X: meanPos, Y: 0}, {X: meanPos, Y: lineHeight}}
	medianLinePts := plotter.XYs{{X: medianPos, Y: 0}, {X: medianPos, Y: lineHeight}}
	meanLine, _ := plotter.NewLine(meanLinePts)
	medianLine, _ := plotter.NewLine(medianLinePts)
	meanLine.Color = color.RGBA{R: 220, G: 20, B: 60, A: 255} // crimson
	meanLine.LineStyle.Width = vg.Points(1.5)
	medianLine.Color = color.RGBA{R: 34, G: 139, B: 34, A: 255} // forest green
	medianLine.LineStyle.Width = vg.Points(1.5)
	medianLine.LineStyle.Dashes = []vg.Length{vg.Points(2), vg.Points(2)}

	p.Add(meanLine)
	p.Add(medianLine)
	p.Legend.Add("Mean", meanLine)
	p.Legend.Add("Median", medianLine)

	// Save annotated plot
	if err := p.Save(8*vg.Inch, 4*vg.Inch, "work_hours_distribution_annotated.png"); err != nil {
		panic(err)
	}

	fmt.Println("Saved work_hours_distribution_annotated.png")
}
