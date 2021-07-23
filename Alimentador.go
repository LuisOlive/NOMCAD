package nomcad

func NewAlimentador(nombre string, vLinea int) Transformador {
	a := NewTransformadorRegulador(nombre)
	a.VPrim(vLinea)
	return a
}