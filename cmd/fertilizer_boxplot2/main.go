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
	data := plotter.Values{
		65, 64, 80, 66, 62, 67, 75, 54, 50, 74, 68, 65,
		67, 55, 73, 71, 74, 61, 64, 52, 64, 60, 72,
	}

	p := plot.New()
	p.Title.Text = "Boxplot of Fertilizer Usage (grams)"
	p.Y.Label.Text = "Grams"

	box, err := plotter.NewBoxPlot(vg.Points(60), 0, data)
	if err != nil {
		log.Fatal(err)
	}
	box.FillColor = color.RGBA{R: 173, G: 216, B: 230, A: 255}
	p.Add(box)

	if err := p.Save(4*vg.Inch, 6*vg.Inch, "fertilizer_boxplot2.png"); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Saved fertilizer_boxplot2.png")
}
