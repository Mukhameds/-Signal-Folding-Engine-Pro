package optimize

import (
	"signal_folding_pro/core"
	"signal_folding_pro/signal"
)

type FoldingEngine struct {
	Protein   *core.Protein
	Field     *signal.Field
	Phantoms  *signal.PhantomSystem
	Iterations int
	EnergyLog  []float64
}

func NewFoldingEngine(p *core.Protein, f *signal.Field, ps *signal.PhantomSystem) *FoldingEngine {
	return &FoldingEngine{
		Protein:   p,
		Field:     f,
		Phantoms:  ps,
		Iterations: 1000,
		EnergyLog:  []float64{},
	}
}

func (fe *FoldingEngine) Run() {
	bestEnergy := fe.Protein.TotalEnergy()
	for step := 0; step < fe.Iterations; step++ {
		fe.Field.Propagate()
		idx := fe.Field.MaxSignalIndex()
		fe.Phantoms.Fluctuate(idx)

		e := fe.Protein.TotalEnergy()
		fe.EnergyLog = append(fe.EnergyLog, e)

		if e < bestEnergy {
			bestEnergy = e
		}
	}
}
