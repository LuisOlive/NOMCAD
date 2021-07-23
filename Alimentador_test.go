package nomcad

import "testing"

func TestAlimentador(t *testing.T) {
	// // instanciamos un alimentador
	// a1 := NewAlimentador(100)
	
	// // instanciamos 4 circuitos derivados
	// c1 := NewCircuito("Circuito 1", 53000, .88)
	// c2 := NewCircuito("Circuito 2", 2000, .95)
	// // c3 := NewCircuito("Circuito 3", 400, 1)
	// // c4 := NewCircuito("Circuito 4", 150000, .7)
	
	// // probamos circuito 1 sin conectar
	// if c1.VSup() != 0 {
	// 	t.Error("se esperaban 0V por que no hay fuente, se obtuvieron", c1.VSup())
	// }
	
	// // conectamos con referencia al alimentador
	// a1.Conectar(&c1, &c2)
	
	// // ahora esperamos que sí tenga alimentacion de 100v
	// if c1.VSup() != 100 || c2.VSup() != 100 {
	// 	t.Error("se esperaban 100V por que están conectadas, se obtuvieron", c1.VSup())
	// }
	
	// // probando que la tensión es escuchada por el circuito 1
	// a1.VLinea(220)
	
	// if c1.VSup() != 220 || c2.VSup() != 220 {
	// 	t.Error("se esperaba que cambiara a 220, se obtuvieron", c1.VSup())
	// }
}