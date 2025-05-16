package core

import (
	"math"
	props "signal_folding_pro/core/properties"
)

type Residue struct {
	Type        string
	Index       int
	Phi, Psi, Omega float64 // Углы вращения
	X_N, Y_N, Z_N     float64
	X_CA, Y_CA, Z_CA  float64
	X_C, Y_C, Z_C     float64
	Charge      float64
	Hydrophobic float64
	Volume      float64
}

func NewResidue(resType string, index int) Residue {
	prop := props.GetProperty(resType)
	return Residue{
		Type:        resType,
		Index:       index,
		Phi:         -60.0,
		Psi:         -45.0,
		Omega:       180.0, // транс-ориентация
		Charge:      prop.Charge,
		Hydrophobic: prop.Hydrophobic,
		Volume:      prop.Volume,
	}
}

// Возвращает центр масс аминокислоты (в среднем по 3 точкам)
func (r *Residue) Center() (float64, float64, float64) {
	x := (r.X_N + r.X_CA + r.X_C) / 3
	y := (r.Y_N + r.Y_CA + r.Y_C) / 3
	z := (r.Z_N + r.Z_CA + r.Z_C) / 3
	return x, y, z
}

// Эвклидово расстояние между центрами двух остатков
func Distance(a, b *Residue) float64 {
	x1, y1, z1 := a.Center()
	x2, y2, z2 := b.Center()
	dx := x1 - x2
	dy := y1 - y2
	dz := z1 - z2
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}