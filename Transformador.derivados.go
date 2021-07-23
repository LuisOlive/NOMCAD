package nomcad

func (tr *Transformador) Conectar(args ...*Transformador) {
	tr.derivados = append(tr.derivados, args...)

	for _, c := range args {
		c.alimentador = tr
	}
}