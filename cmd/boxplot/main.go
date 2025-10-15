package main

import (
	"fmt"
	"image/color"
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// The fertilizer usage data
	data := plotter.Values{
		22.5, 23.1, 21.8, 24.0, 22.7, 23.5, 24.8, 22.0, 23.9, 25.2,
		21.5, 23.3, 24.5, 22.2, 23.8, 25.5, 21.9, 24.2, 22.8, 23.6,
		25.0, 22.1, 24.9, 23.0, 22.9, 24.1, 23.7, 22.4, 24.7, 23.4,
		22.6, 24.3, 23.2, 25.1, 21.7, 24.4, 22.3, 25.3, 23.8, 24.6,
		21.6, 23.9, 22.5, 25.4, 23.1, 24.0, 22.9, 23.5, 24.8, 22.2,
		23.7, 25.0, 21.8, 24.2, 23.0, 22.7, 24.5, 23.3, 20.0, 24.9,
	}

	// --- 1. Create a new plot ---
	p := plot.New()

	// --- 2. Set the title and labels ---
	p.Title.Text = "Box Plot of Fertilizer Usage"
	p.Y.Label.Text = "Fertilizer Usage (KG)"

	// --- 3. Create the box plot ---
	// The width parameter controls the width of the box.
	// We pass the data and it automatically calculates quartiles, median, etc.
	box, err := plotter.NewBoxPlot(vg.Points(50), 0, data)
	if err != nil {
		log.Fatal(err)
	}

	// --- 4. Customize the plot appearance (optional) ---
	box.FillColor = color.RGBA{R: 173, G: 216, B: 230, A: 255} // Light blue
	// Only FillColor is set here; BoxPlot does not expose a direct Color/LineStyle field

	// Add the box plot to the plot
	p.Add(box)

	// --- 5. Save the plot to a file ---
	// The dimensions are in standard points (e.g., inches * 72).
	if err := p.Save(8*vg.Inch, 6*vg.Inch, "fertilizer_boxplot.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Box plot has been saved to fertilizer_boxplot.png")
}
