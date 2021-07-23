package nomcad

type Circuito struct {
	Transformador
}

func NewCircuito(nombre string, _, __ float32) Circuito {
	c := Circuito{}
	c.Transformador = NewTransformadorRegulador(nombre)
	return c
}