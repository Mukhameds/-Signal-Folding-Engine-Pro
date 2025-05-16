package main

import (
	"fmt"
	"os"

	"signal_folding_pro/core"
	props "signal_folding_pro/core/properties"
	"signal_folding_pro/optimize"
	"signal_folding_pro/output"
	"signal_folding_pro/signal"
)

func main() {
	// Загрузка базы аминокислот
	err := props.LoadAAProperties("datasets/amino_acids.json")
	if err != nil {
		fmt.Println("[ERROR] Failed to load amino acid properties:", err)
		os.Exit(1)
	}

	// Последовательность белка (пример)
	sequence := "ACDEFGHIKLMNPQRSTVWY"
	protein := core.NewBackboneProtein(sequence)
	protein.InitializeBackboneGeometry()

	// Сигнальное поле и фантомы
	field := signal.NewField(protein)
	phantoms := signal.NewPhantomSystem(protein)

	// Оптимизация структуры
	optimizer := optimize.NewFoldingEngine(protein, field, phantoms)
	optimizer.Run()

	// Сохранение результатов
	output.ExportPDB(protein, "output/final_structure.pdb")
	output.RenderSVG(protein, "output/final_structure.svg")
	output.LogEnergyPlot(optimizer, "output/energy_log.csv")

	fmt.Println("✅ Folding complete. Final energy:", protein.TotalEnergy())
}
