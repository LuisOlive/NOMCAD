package nomcad

// devuelve las espiras en el lado primario del transformador
func (t *Transformador) EspirasPrim() int {
	return t.espirasPrim
}

// devuelve las espiras en el lado secundario del transformador
func (t *Transformador) EspirasSec() int {
	return t.espirasSec
}

// configura las espiras para el primario y el secundario
func (t *Transformador) Espiras(prim, sec int) {
	if prim <= 0 || sec <= 0 { return }
	t.espirasPrim = prim
	t.espirasSec = sec
}

// devuelve la relación del transformador dados los números de vueltas en primario y secundario
func (t *Transformador) Relacion() float32 {
	return float32(t.espirasSec) / float32(t.espirasPrim)
}