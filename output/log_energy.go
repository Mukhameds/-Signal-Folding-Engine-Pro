package output

import (
	"fmt"
	"os"
	"signal_folding_pro/optimize"
)

func LogEnergyPlot(engine *optimize.FoldingEngine, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("[ERROR] creating energy log:", err)
		return
	}
	defer f.Close()

	f.WriteString("iteration,energy\n")
	for i, e := range engine.EnergyLog {
		line := fmt.Sprintf("%d,%.6f\n", i, e)
		f.WriteString(line)
	}
}