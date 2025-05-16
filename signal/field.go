package signal

import (
	"math"
	"signal_folding_pro/core"
)

type Field struct {
	Protein *core.Protein
	Signals []float64
}

func NewField(p *core.Protein) *Field {
	return &Field{
		Protein: p,
		Signals: make([]float64, len(p.Chain)),
	}
}

// Обновляет поле сигналов на основе расстояний и гидрофобности
func (f *Field) Propagate() {
	for i := range f.Protein.Chain {
		signal := 0.0
		source := f.Protein.Chain[i]
		for j := range f.Protein.Chain {
			if i == j {
				continue
			}
			target := f.Protein.Chain[j]
			d := core.Distance(source, target)
			if d < 8.0 {
				impact := target.Hydrophobic / (d + 1e-6)
				signal += impact
			}
		}
		f.Signals[i] = signal
	}
}

func (f *Field) MaxSignalIndex() int {
	maxIdx := 0
	maxVal := -math.MaxFloat64
	for i, val := range f.Signals {
		if val > maxVal {
			maxVal = val
			maxIdx = i
		}
	}
	return maxIdx
}
