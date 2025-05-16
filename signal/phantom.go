package signal

import (
	"math/rand"
	"time"
	"signal_folding_pro/core"
)

type PhantomSystem struct {
	Protein *core.Protein
}

func NewPhantomSystem(p *core.Protein) *PhantomSystem {
	rand.Seed(time.Now().UnixNano())
	return &PhantomSystem{Protein: p}
}

// Возбуждает фантомную флуктуацию (случайный сдвиг φ/ψ)
func (ps *PhantomSystem) Fluctuate(index int) {
	res := ps.Protein.Chain[index]
	dPhi := (rand.Float64() - 0.5) * 30.0 // ±15°
	dPsi := (rand.Float64() - 0.5) * 30.0

	res.Phi += dPhi
	res.Psi += dPsi
}

// Может использоваться в каскаде возбуждений или от внешнего сигнала
func (ps *PhantomSystem) ApplyCascade(indices []int) {
	for _, idx := range indices {
		ps.Fluctuate(idx)
	}
}
