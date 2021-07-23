package nomcad

import "github.com/chewxy/math32"

// Devuelve la potencia activa de una Transformador
func (tr *Transformador) Pot(arg ...float32) float32 {
	if len(arg) == 1 {
		tr.pot = arg[0]
	}
	
	return tr.pot
}

// Devuelve el factor de potencia de una Transformador
func (tr *Transformador) FPot(arg ...float32) float32 {
	if len(arg) == 1 {
		tr.fpot = arg[0]
	}
	
	return tr.fpot
}

// Devuelve la potencia real de una Transformador
func (tr *Transformador) PotS() float32 {
	return tr.Pot() / tr.FPot()
}

// Devuelve el ángulo de desfasamiento de onda de una Transformador
func (tr *Transformador) Angulo() float32 {
	return math32.Acos(tr.FPot())
}

func sgn(x float32) float32 {
	if x >= 0 {
		return 1
	}
	return -1
}

// Devuelve la potencia reactiva de una Transformador
func (tr *Transformador) PotQ() float32 {
	return tr.PotS() * math32.Sin(tr.Angulo()) * sgn(tr.FPot())
}

