package main

import (
	"fmt"
	"image/color"
	"math"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {
	// Data
	type Rec struct {
		ID          int
		Weight      float64
		Crunchiness string
		Sweetness   float64
		Ripeness    int
		Quality     string
	}

	data := []Rec{
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

	// Scatter plot: Weight vs Sweetness, color by Quality
	p := plot.New()
	p.Title.Text = "Weight vs Sweetness (colored by Quality)"
	p.X.Label.Text = "Weight"
	p.Y.Label.Text = "Sweetness"

	// Separate points by quality
	goodPts := make(plotter.XYs, 0)
	badPts := make(plotter.XYs, 0)
	var xsum, ysum float64
	for _, r := range data {
		xsum += r.Weight
		ysum += r.Sweetness
		if r.Quality == "good" {
			goodPts = append(goodPts, plotter.XY{X: r.Weight, Y: r.Sweetness})
		} else {
			badPts = append(badPts, plotter.XY{X: r.Weight, Y: r.Sweetness})
		}
	}

	// scatterters
	sg, _ := plotter.NewScatter(goodPts)
	sg.GlyphStyle.Color = color.RGBA{G: 180, B: 60, A: 255} // greenish
	sg.GlyphStyle.Radius = vg.Points(3)
	sb, _ := plotter.NewScatter(badPts)
	sb.GlyphStyle.Color = color.RGBA{R: 200, A: 255} // red
	sb.GlyphStyle.Radius = vg.Points(3)
	p.Add(sg, sb)
	p.Legend.Add("good", sg)
	p.Legend.Add("bad", sb)

	// Compute Pearson correlation and linear regression
	n := float64(len(data))
	xMean := xsum / n
	yMean := ysum / n
	var cov, varx, vary float64
	for _, r := range data {
		dx := r.Weight - xMean
		dy := r.Sweetness - yMean
		cov += dx * dy
		varx += dx * dx
		vary += dy * dy
	}
	cov /= n
	varx /= n
	vary /= n
	pearson := cov / (math.Sqrt(varx) * math.Sqrt(vary))

	// regression slope and intercept (OLS using means)
	slope := cov / varx
	intercept := yMean - slope*xMean

	// draw regression line across x range
	xmin := data[0].Weight
	xmax := data[0].Weight
	for _, r := range data {
		if r.Weight < xmin {
			xmin = r.Weight
		}
		if r.Weight > xmax {
			xmax = r.Weight
		}
	}
	linePts := plotter.XYs{{X: xmin, Y: intercept + slope*xmin}, {X: xmax, Y: intercept + slope*xmax}}
	line, _ := plotter.NewLine(linePts)
	line.Color = color.RGBA{R: 0, G: 0, B: 139, A: 255}
	line.LineStyle.Width = vg.Points(1)
	p.Add(line)
	p.Legend.Add(fmt.Sprintf("regression (r=%.3f)", pearson), line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, "weight_vs_sweetness.png"); err != nil {
		panic(err)
	}

	fmt.Printf("Saved weight_vs_sweetness.png (Pearson r = %.4f, slope = %.4f, intercept = %.4f)\n", pearson, slope, intercept)

	// Composition: sweetness bins vs quality (good/bad)
	// Define bins: <=1, (1,2], (2,3], >3
	bins := []struct {
		label    string
		min, max float64
	}{
		{"<=1", math.Inf(-1), 1.0},
		{"1-2", 1.0, 2.0},
		{"2-3", 2.0, 3.0},
		{">3", 3.0, math.Inf(1)},
	}

	// counts per bin per quality
	goodCounts := make([]float64, len(bins))
	badCounts := make([]float64, len(bins))
	for _, r := range data {
		for i, b := range bins {
			if r.Sweetness > b.min && r.Sweetness <= b.max {
				if r.Quality == "good" {
					goodCounts[i]++
				} else {
					badCounts[i]++
				}
				break
			}
		}
	}

	// Build grouped bar chart: for each bin, show bars for good and bad side-by-side
	p2 := plot.New()
	p2.Title.Text = "Composition: Sweetness bins vs Quality"
	p2.Y.Label.Text = "Count"
	p2.X.Label.Text = "Sweetness bins"

	labels := make([]string, len(bins))
	for i, b := range bins {
		labels[i] = b.label
	}

	valsGood := make(plotter.Values, len(goodCounts))
	valsBad := make(plotter.Values, len(badCounts))
	for i := range valsGood {
		valsGood[i] = goodCounts[i]
		valsBad[i] = badCounts[i]
	}

	barw := vg.Points(20)
	bg, _ := plotter.NewBarChart(valsGood, barw)
	bg.Color = color.RGBA{R: 46, G: 139, B: 87, A: 255}
	bg.Offset = -vg.Points(11)
	bb, _ := plotter.NewBarChart(valsBad, barw)
	bb.Color = color.RGBA{R: 178, G: 34, B: 34, A: 255}
	bb.Offset = vg.Points(11)

	p2.Add(bg, bb)
	p2.NominalX(labels...)
	p2.Legend.Add("good", bg)
	p2.Legend.Add("bad", bb)
	p2.Add(plotter.NewGrid())

	if err := p2.Save(6*vg.Inch, 3*vg.Inch, "sweetness_quality_composition.png"); err != nil {
		panic(err)
	}

	// Print contingency table and percentages
	fmt.Println("\nSweetness bin vs Quality (counts):")
	for i, b := range bins {
		fmt.Printf("%s: good=%d, bad=%d\n", b.label, int(goodCounts[i]), int(badCounts[i]))
	}

	fmt.Println("Saved sweetness_quality_composition.png")

	// Interpretation
	fmt.Println("\nInterpretation:")
	fmt.Println("1) Weight vs Sweetness:")
	fmt.Printf("   - Pearson correlation r = %.3f (close to 0 indicates weak linear association).\n", pearson)
	if math.Abs(pearson) < 0.3 {
		fmt.Println("   - There is little to no linear relationship between weight and sweetness in this sample.")
	} else if pearson > 0 {
		fmt.Println("   - Positive correlation: heavier apples tend to be sweeter.")
	} else {
		fmt.Println("   - Negative correlation: heavier apples tend to be less sweet.")
	}
	fmt.Printf("   - Regression line: sweetness = %.3f * weight + %.3f\n", slope, intercept)

	fmt.Println("\n2) Composition by sweetness bin and quality:")
	fmt.Println("   - The bar chart shows counts of 'good' vs 'bad' within each sweetness bin.")
	fmt.Println("   - Use the plot 'sweetness_quality_composition.png' for a quick view of how quality distributes across sweetness levels.")
}
