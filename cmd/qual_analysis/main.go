package main

import (
	"fmt"
	"image/color"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Dataset (20 observations)
	type Record struct {
		ID          int
		Weight      float64
		Crunchiness string // high/medium/low
		Sweetness   float64
		Ripeness    int    // 1-4 (ordinal)
		Quality     string // good/bad
	}

	data := []Record{
		{1, 70, "high", 0.5, 1, "good"},
		{2, 90, "medium", 1, 2, "bad"},
		{3, 83, "high", 1, 1, "good"},
		{4, 85, "low", 3, 3, "good"},
		{5, 90, "low", 3, 3, "good"},
		{6, 78, "high", 1.5, 2, "good"},
		{7, 78, "high", 0.5, 4, "bad"},
		{8, 93, "medium", 1, 4, "good"},
		{9, 85, "high", 3.5, 3, "good"},
		{10, 88, "low", 2, 3, "bad"},
		{11, 86, "high", 2, 2, "good"},
		{12, 92, "low", 2, 2, "good"},
		{13, 95, "medium", 1.5, 3, "bad"},
		{14, 100, "high", 4, 3, "good"},
		{15, 94, "medium", 2.5, 2, "good"},
		{16, 96, "low", 2, 4, "good"},
		{17, 70, "low", 1, 1, "good"},
		{18, 82, "medium", 3, 1, "bad"},
		{19, 90, "high", 3, 2, "bad"},
		{20, 78, "high", 0.5, 2, "bad"},
	}

	// Qualitative variables to plot: Crunchiness, Quality, Ripeness (ordinal)

	// 1) Crunchiness (high/medium/low) - choose order high, medium, low
	crunchFreq := make(map[string]int)
	for _, r := range data {
		crunchFreq[r.Crunchiness]++
	}
	crunchKeys := []string{"high", "medium", "low"}
	crunchVals := make(plotter.Values, len(crunchKeys))
	for i, k := range crunchKeys {
		crunchVals[i] = float64(crunchFreq[k])
	}
	plotBar(crunchKeys, crunchVals, "crunchiness_distribution.png", "Crunchiness", "Count", color.RGBA{R: 70, G: 130, B: 180, A: 255})

	// 2) Quality (good/bad) - order good, bad
	qualityFreq := make(map[string]int)
	for _, r := range data {
		qualityFreq[r.Quality]++
	}
	qualityKeys := []string{"good", "bad"}
	qualityVals := make(plotter.Values, len(qualityKeys))
	for i, k := range qualityKeys {
		qualityVals[i] = float64(qualityFreq[k])
	}
	plotBar(qualityKeys, qualityVals, "quality_distribution.png", "Quality", "Count", color.RGBA{R: 46, G: 139, B: 87, A: 255})

	// 3) Ripeness (1-4) treat as ordinal categories
	ripenessFreq := make(map[int]int)
	for _, r := range data {
		ripenessFreq[r.Ripeness]++
	}
	ripenessKeys := []int{1, 2, 3, 4}
	ripenessVals := make(plotter.Values, len(ripenessKeys))
	labelsRip := make([]string, len(ripenessKeys))
	for i, k := range ripenessKeys {
		ripenessVals[i] = float64(ripenessFreq[k])
		labelsRip[i] = fmt.Sprintf("%d", k)
	}
	plotBar(labelsRip, ripenessVals, "ripeness_distribution.png", "Ripeness (1=low → 4=high)", "Count", color.RGBA{R: 255, G: 165, B: 0, A: 255})

	// Interpretations printed to console
	fmt.Println("--- Interpretations for qualitative variables ---")

	fmt.Println("1) Crunchiness (high / medium / low):")
	for _, k := range crunchKeys {
		fmt.Printf("   - %s: %d observations\n", k, crunchFreq[k])
	}
	fmt.Println("   Interpretation: Most items are 'high' or 'medium' crunchiness; 'low' is less common.")

	fmt.Println()
	fmt.Println("2) Quality (good / bad):")
	for _, k := range qualityKeys {
		fmt.Printf("   - %s: %d observations\n", k, qualityFreq[k])
	}
	fmt.Println("   Interpretation: Majority are labeled 'good' — quality appears generally positive in this sample.")

	fmt.Println()
	fmt.Println("3) Ripeness (1 to 4):")
	for i, k := range ripenessKeys {
		fmt.Printf("   - %d: %d observations\n", k, int(ripenessVals[i]))
	}
	fmt.Println("   Interpretation: Ripeness levels are distributed across the scale; identify if certain ripeness levels coincide with 'bad' quality for downstream analysis.")

	fmt.Println("\nSaved PNGs: crunchiness_distribution.png, quality_distribution.png, ripeness_distribution.png")
}

// plotBar draws and saves a bar chart. keys are the x labels (strings), vals are counts.
func plotBar(keys []string, vals plotter.Values, filename, xlabel, ylabel string, col color.RGBA) {
	p := plot.New()
	p.Title.Text = filename
	p.X.Label.Text = xlabel
	p.Y.Label.Text = ylabel

	bar, err := plotter.NewBarChart(vals, vg.Points(20))
	if err != nil {
		panic(err)
	}
	bar.Color = col
	bar.LineStyle.Width = vg.Points(0.5)

	p.Add(bar)
	p.NominalX(keys...)
	p.Add(plotter.NewGrid())

	if err := p.Save(6*vg.Inch, 3*vg.Inch, filename); err != nil {
		panic(err)
	}
}
