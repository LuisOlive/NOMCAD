package nomcad

// // asigna o devuelve una tensión de linea al alimentador
// func (a *Alimentador) VLinea_(arg ...int) float32 {
// 	if len(arg) == 1 {
// 		a.vLinea = arg[0]
// 	}
// 	return float32(a.vLinea)
// }

// //devuelve la tensión reglamentaria del 95% de acuerdo con 210-19. a) 1) Nota 4
// func (t *Transformador) VMin() float32 {
// 	return t.VLinea() * .95
// }

// func (t *Transformador) VLinea() float32 {
// 	t.vLinea = int(t.VSup() * t.Relacion())
// 	return float32(t.vLinea)
// }

func (tr *Transformador) VPrim(arg ...int) float32 {
	if tr.alimentador != nil {
		tr._vEstatico = -1
		return tr.alimentador.VSec()
	}
	
	if len(arg) == 1 && arg[0] >= 0{
		tr._vEstatico = arg[0]
	}
	
	return float32(tr._vEstatico)
}

func (tr *Transformador) VSec() float32 {
	return tr.VPrim() * tr.Relacion()
}