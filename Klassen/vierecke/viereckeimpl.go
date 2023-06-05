package vierecke

// Annalena Cyriacus
// 24.04.2023

import ( "gfx" ; "fmt" ; "unsafe")

// Datenstruktur eines Vierecks
// ----------------------------

type data struct {
	xA, yA uint16
	xB, yB uint16 
	xC, yC uint16
	xD, yD uint16
	r, g, b uint8 //optional
}

// Methoden eines Dreiecks
// -----------------------

// Vor.: -
// Erg.: Ein Viereck mit den Koordinaten (xA,yA), (xB,yB), (xC,yC) und
//       (xD,yD) ist geliefert. A ist dabei die obere linke Ecke.
func New (xA,yA,xB,yB,xC,yC,xD,yD uint16) *data { // *data erfüllt das Interface Viereck
	var v *data
	v = new(data)
	v.xA = xA
	v.yA = yA
	v.xB = xB
	v.yB = yB
	v.xC = xC
	v.yC = yC
	v.xD = xD
	v.yD = yD
	return v
}


// Vor.: -
// Erg.: Alle Koordinaten der vier Ecken des Vierecks sind geliefert.
func (v *data) GetKoordinaten () (xA,yA,xB,yB,xC,yC,xD,yD uint16) {
	return v.xA,v.yA,v.xB,v.yB,v.xC,v.yC,v.xD,v.yD
}

// Vor.: -
// Erg.: Das Viereck hat die eingegebenen Koordinaten.
func (v *data) SetKoordinaten (xA,yA,xB,yB,xC,yC,xD,yD uint16) () {
	v.xA = xA
	v.yA = yA
	v.xB = xB
	v.yB = yB
	v.xC = xC
	v.yC = yC
	v.xD = xD
	v.yD = yD
}

func (v *data) GibKoordA() (uint16,uint16) {
	return v.xA, v.yA
}

func (v *data) GibKoordB() (uint16,uint16) {
	return v.xB, v.yB
}

func (v *data) GibKoordC() (uint16,uint16) {
	return v.xC, v.yC
}

func (v *data) GibKoordD() (uint16,uint16) {
	return v.xD, v.yD
}

func (v *data) SetzeKoordA(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xA = x
		v.yA = y
	}
}
		
func (v *data) SetzeKoordB(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xB = x
		v.yB = y
	}
}

func (v *data) SetzeKoordC(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xC = x
		v.yC = y
	}
}

func (v *data) SetzeKoordD(x,y uint16) {
	if x < 1200 && y < 700 {
		v.xD = x
		v.yD = y
	}
}

/*		
func (v *data) GibFarbe() (uint8,uint8,uint8) {
	return v.r, v.g, v.b
}

func (v *data) SetzeFarbe (r,g,b uint8) {
	v.r = r
	v.g = g
	v.b = b
}

*/

// Vor.: -
// Erg.: Eine Repräsentation des Vierecks in Form eines Strings ist geliefert.
func (v *data) String () string {
	return "Viereck mit den Koordinaten A("+fmt.Sprint(v.xA)+";"+fmt.Sprint(v.yA)+"), B("+fmt.Sprint(v.xB)+";"+fmt.Sprint(v.yB)+"), C("+fmt.Sprint(v.xC)+";"+fmt.Sprint(v.yC)+") und D("+fmt.Sprint(v.xD)+";"+fmt.Sprint(v.yD)+")"
}
	
	
// Vor.: -
// Erg.: Ein stets gleichlanger Bytestrom, der die serialisierte Form
//       des Vierecks darstellt, ist geliefert.
func (v *data) Kodieren () (b []byte) {
	b = (*(*[unsafe.Sizeof(*v)]byte)(unsafe.Pointer(&(*v))))[:] // Es wird rüberkopiert!
	return
}
	

// Vor.: -
// Eff.: Das Viereck hat jetzt die Eigenschaften, die im Bytestrom
//       enthalten waren. Seine vorher vorhandenen Eigenschaften
//       gibt es nicht mehr.
func (v *data) Dekodieren (b []byte) {
	(*v) = *(*data)(unsafe.Pointer(&b[0])) 
}
	
// Vor.: -
// Erg.: Eine tiefe Kopie des Vierecks v in Form einer Variablen des
//       Typs interface {} ist geliefert. Die Objekte v und 
//       v.Kopie().(Viereck) sind unterschiedliche Objekte, haben aber
//       komplett identische Eigenschaften.
func (v * data) Kopie () interface {} {
	var kopie *data = new(data) // 1. Erzeuge eine Kopie-Instanz!
	*kopie = *v // Der Inhalt des Structs von kopie sei bitte mit den Werten des Inhalts des Structs von v belegt.
	return kopie
}

// Vor.: -
// Erg.: Das Viereck ist in der aktuellen Stiftfarbe gezeichnet.
func (v *data) Zeichnen() {
	gfx.Stiftfarbe(v.r, v.g, v.b)
	gfx.Linie(v.xA, v.yA, v.xB, v.yB)
	gfx.Linie(v.xB, v.yB, v.xC, v.yC)
	gfx.Linie(v.xC, v.yC, v.xD, v.yD)
	gfx.Linie(v.xD, v.yD, v.xA, v.yA)
}

func (v *data) Angeklickt(x,y uint16) bool {						// Checkt, ob angeklickt
	
	var min, max func(a,b uint16) uint16
	min = func (a,b uint16) uint16 {
		if a < b {return a} else {return b}
	}
	max = func (a,b uint16) uint16 {
		if a > b {return a} else {return b}
	}
	
	return min(v.xA,v.xB) <= x && x <= max(v.xC,v.xD) && min(v.yA,v.yD) <= y && y <= max(v.yB,v.yC)
}
