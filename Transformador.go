package nomcad

type Transformador struct {
	nombre string

	espirasPrim int
	espirasSec  int

	_vEstatico int
	pot        float32
	fpot       float32

	alimentador *Transformador
	derivados   []*Transformador
}

func NewTransformador(nombre string, espirasPrim, espirasSec int) Transformador {
	t := Transformador{}
	t.Nombre(nombre)
	t.Espiras(espirasPrim, espirasSec)
	return t
}

func NewTransformadorRegulador(nombre string) Transformador {
	return NewTransformador(nombre, 100, 100)
}