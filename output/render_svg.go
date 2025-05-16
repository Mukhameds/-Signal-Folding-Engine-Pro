package output

import (
	"fmt"
	"os"
	"signal_folding_pro/core"
)

func RenderSVG(p *core.Protein, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("[ERROR] creating SVG:", err)
		return
	}
	defer f.Close()

	f.WriteString(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="800" height="800" viewBox="-100 -100 200 200">
<style>circle { stroke: black; stroke-width: 0.3; }</style>
`)

	for _, r := range p.Chain {
		cx := fmt.Sprintf("%.2f", r.X_CA)
		cy := fmt.Sprintf("%.2f", r.Z_CA) // Z как вертикаль
		color := "blue"
		if r.Hydrophobic > 0 {
			color = "red"
		}
		f.WriteString(fmt.Sprintf("<circle cx='%s' cy='%s' r='2' fill='%s' />\n", cx, cy, color))
	}

	f.WriteString("</svg>\n")
}
