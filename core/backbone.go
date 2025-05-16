package core

import (
	"math"
)

type Protein struct {
	Sequence []string
	Chain    []*Residue
}

func NewBackboneProtein(seq string) *Protein {
	s := []string{}
	for _, r := range seq {
		s = append(s, string(r))
	}
	p := &Protein{Sequence: s}
	for i, code := range s {
		res := NewResidue(code, i)
		p.Chain = append(p.Chain, &res)
	}
	return p
}

// Основная геометрия по φ, ψ, ω
func (p *Protein) InitializeBackboneGeometry() {
	bondLength := 3.8 // Å между Cα–Cα условно
	angleStep := 120.0 // поворот в градусах (упрощённо)

	x, y, z := 0.0, 0.0, 0.0
	angle := 0.0

	for _, res := range p.Chain {

		rad := angle * math.Pi / 180.0
		x += bondLength * math.Cos(rad)
		y += bondLength * math.Sin(rad)
		z += 0

		res.X_CA = x
		res.Y_CA = y
		res.Z_CA = z

		res.X_N = x - 1.2 * math.Cos(rad)
		res.Y_N = y - 1.2 * math.Sin(rad)
		res.Z_N = z - 0.5

		res.X_C = x + 1.2 * math.Cos(rad)
		res.Y_C = y + 1.2 * math.Sin(rad)
		res.Z_C = z + 0.5

		angle += angleStep + res.Phi*0.02 - res.Psi*0.01 // псевдо-вращение
	}
}

func (p *Protein) TotalEnergy() float64 {
	total := 0.0
	for i := 0; i < len(p.Chain); i++ {
		for j := i + 2; j < len(p.Chain); j++ {
			d := Distance(p.Chain[i], p.Chain[j])
			if d < 0.1 {
				continue
			}
			vol := (p.Chain[i].Volume + p.Chain[j].Volume) / 2.0
			epsilon := -0.01 * p.Chain[i].Hydrophobic * p.Chain[j].Hydrophobic
			lj := 4 * epsilon * (math.Pow(vol/d, 12) - math.Pow(vol/d, 6))
			total += lj
		}
	}
	return total
}
