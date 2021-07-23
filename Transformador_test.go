package nomcad

import "testing"

func TestTransformador(t *testing.T) {
	al := NewAlimentador("Generador", 13800)
	
	tr1 := NewTransformador("TR-1", 138, 1320) // 13.8kV -> 132kV
	tr2 := NewTransformador("TR-2", 132, 23) // 132kV -> 23kV
	tr3 := NewTransformador("TR-3", 2300, 44) // 23kV -> 440V
	
	al.Conectar(&tr1)
	tr1.Conectar(&tr2)
	tr2.Conectar(&tr3)
	
	if al.VSec() != 13800 {
		t.Error("Se esperaba salida 13.8kV, se obtuvo", al.VSec())
	}
	
	if tr1.VPrim() != 13800 || tr1.VSec() != 132000 {
		t.Errorf("Se esperaba relación 13.8kV -> 132kV, se obtuvo relación %v/%v", tr1.VPrim(), tr1.VSec())
	}
	
	if tr3.VPrim() != 23000 || tr3.VSec() != 440 {
		t.Errorf("Se esperaba relación 23kV -> 440V, se obtuvo relación %v/%v", tr3.VPrim(), tr3.VSec())
	}
	
	al.VPrim(23000)
	
	if tr3.VSec() != 733.3333 {
		t.Error("Se esperaba tensión de salida 733.3333, Se obtuvo", tr3.VSec())
	}
}