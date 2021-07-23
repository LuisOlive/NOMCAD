package nomcad

func (tr *Transformador) Nombre(arg ...string) string {
	if len(arg) == 1 {
		tr.nombre = arg[0]
	}
	return tr.nombre
}