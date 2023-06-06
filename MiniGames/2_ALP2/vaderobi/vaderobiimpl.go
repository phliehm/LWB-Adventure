package vaderobi
// Autor:(c) Stefan Schmidt, St.Schmidt@online.de, adaptiert von Annalena Cyriacus
// Datum: 09.04.2015 ; letzte Änderung: 18.02.2019 ; zuletzt adaptiert am: 21.03.2023
// Zweck:
/* Das Paket verwaltet eine rechteckige Welt aus schachbrettförmig
 angeordenten Plätzen (Standard: 12 Plätze hoch und 16 Plätze breit)
 und einem Roboter namens Vaderobi. 
 Vaderobi steht immer auf genau einem der Plätze in der Welt (Robis
 Platz) und schaut immer in genau eine Himmelsrichtung (Nord, Ost, Süd,
 West), die als "Vaderobis Richtung" bezeichnet wird.
 In jeder Himmelsrichtung gibt es genau einen Nachbarplatz oder keinen.
 Ein Platz kann durch Vaderobi markiert oder entmarkiert werden.
 Vaderobi hat eine Tasche mit Klötzen, die er einzeln auf seinen Platz
 ablegen oder von dort einzeln aufnehmen kann.
 Vaderobi hat außerdem Zugriff auf Mauersteine und kann Plätze zumauern
 oder entmauern.
 Zu Beginn steht Vaderobi in einer leeren Welt in der linken oberen Ecke
 mit Blick nach Süden und hat 999 Klötze in seiner Tasche.
 Wird eine bereits existierende Welt geladen, so kann Vaderobi an einer
 anderen Position starten. In der Welt befinden sich ebenso insgesamt 
 999 Klötze, allerdings nicht alle unbedingt in seiner Tasche.
 Mit dem Import des Pakets wird automatisch ein Grafikfenster geöffnet,
 indem die Welt von Vaderobi sichtbar gemacht wird. Der Grafikbereich des
 Fensters ist standardmäßig 800 x 600 Pixel. Ein Platz hat dann 50 x 50 Pixel.
 
 In der Grundeinstellung befindet sich Vaderobi im sogenannten Schrittmodus,
 d.h., jede einzelne „Grundaktion“, die von Vaderobi ausgeführt werden soll 
 und die Welt verändert, muss mit einem Druck auf die Entertaste bestätigt
 werden (siehe Funktion Schrittmodus).
 */
import ( "os" ; "runtime" ; "fmt" ; "gfx" ; "dateien" )

var fontpath string = os.Getenv("GOPATH")

var w *welt           // paketinterne Variable - die verwendete Welt

type kachel struct {  // Bestandteil einer Welt; exitiert nur innerhalb einer Welt!
	welt *welt
	ze,sp uint8
	farbe [3]uint8
	farbeM [3]uint8
	markiert bool
	klotzanzahl uint
	zugemauert bool
}

func newKachel (zeile, spalte uint8, r,g,b uint8, r2,g2,b2 uint8, w *welt) *kachel { 
	var f *kachel
	f = new(kachel)
	f.welt = w
	f.ze = zeile
	f.sp = spalte
	f.farbe = [3]uint8{r,g,b}
	f.farbeM = [3]uint8{r2,g2,b2}
	return f
}

func (k *kachel) kodieren () []byte {
	var b []byte = make ([]byte,12)
	b[0] = k.ze
	b[1] = k.sp
	b[2] = k.farbe[0]
	b[3] = k.farbe[1]
	b[4] = k.farbe[2]
	b[5] = k.farbeM[0]
	b[6] = k.farbeM[1]
	b[7] = k.farbeM[2]
	if k.markiert {
		b[8] = 1
	} else {
		b[8] = 0
	}
	b[9] = uint8 (k.klotzanzahl / 256)
	b[10] = uint8 (k.klotzanzahl % 256)
	if k.zugemauert {
		b[11] = 1
	} else {
		b[11] = 0
	}
	return b
}

func (k *kachel) dekodieren (b []byte) {
	k.ze = b[0]
	k.sp = b[1]
	k.farbe[0] = b[2]
	k.farbe[1] = b[3]
	k.farbe[2] = b[4]
	k.farbeM[0] = b[5]
	k.farbeM[1] = b[6]
	k.farbeM[2] = b[7]
	k.markiert = b[8] == 1
	k.klotzanzahl = uint(b[9])*256 + uint(b[10])
	k.zugemauert = b[11] == 1
}

func (k *kachel) zeichnen () {
	feldbreite := k.welt.feldbreite
	if k.markiert {
		gfx.Stiftfarbe(k.farbeM[0],k.farbeM[1],k.farbeM[2])
	}else {
		gfx.Stiftfarbe(k.farbe[0],k.farbe[1],k.farbe[2])
	}
	gfx.Vollrechteck(uint16(k.sp)*feldbreite,uint16(k.ze)*feldbreite,feldbreite,feldbreite)
	if k.klotzanzahl > 0 {
		gfx.Stiftfarbe(240,0,0)
		//gfx.Vollrechteck(uint16(k.sp)*feldbreite,uint16(k.ze)*feldbreite + (4*feldbreite+2)/5,feldbreite/5,feldbreite/5)
		gfx.Kreis(uint16(k.sp)*feldbreite+feldbreite/2,uint16(k.ze)*feldbreite+feldbreite/2,feldbreite/2-1)
		gfx.Kreis(uint16(k.sp)*feldbreite+feldbreite/2,uint16(k.ze)*feldbreite+feldbreite/2,feldbreite/2-2)
		gfx.Kreis(uint16(k.sp)*feldbreite+feldbreite/2,uint16(k.ze)*feldbreite+feldbreite/2,feldbreite/2-3)
		//
		gfx.Schreibe(uint16(k.sp)*feldbreite+2,uint16(k.ze)*feldbreite+2,fmt.Sprint(k.klotzanzahl))
	}
	if k.zugemauert {
		gfx.Stiftfarbe (0,0,0)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+2,uint16(k.ze)*feldbreite+2,feldbreite/2-2,feldbreite/4-2)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite/2+1,uint16(k.ze)*feldbreite+2,feldbreite/2-2,feldbreite/4-2)
		
		gfx.Vollrechteck (uint16(k.sp)*feldbreite,uint16(k.ze)*feldbreite+feldbreite/4+1,feldbreite/4,feldbreite/4-1)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite/4+1,uint16(k.ze)*feldbreite+feldbreite/4+1,feldbreite/2-1,feldbreite/4-1)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite*3/4+1,uint16(k.ze)*feldbreite+feldbreite/4+1,feldbreite/4-1,feldbreite/4-1)
		
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+2,uint16(k.ze)*feldbreite+feldbreite/2+1,feldbreite/2-1,feldbreite/4-2)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite/2+2,uint16(k.ze)*feldbreite+feldbreite/2+1,feldbreite/2-2,feldbreite/4-2)
		
		gfx.Vollrechteck (uint16(k.sp)*feldbreite,uint16(k.ze)*feldbreite+feldbreite*3/4,feldbreite/4,feldbreite/4-1)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite/4+1,uint16(k.ze)*feldbreite+feldbreite*3/4,feldbreite/2-1,feldbreite/4-1)
		gfx.Vollrechteck (uint16(k.sp)*feldbreite+feldbreite*3/4+1,uint16(k.ze)*feldbreite+feldbreite*3/4,feldbreite/4-1,feldbreite/4-1)
	}
	gfx.Stiftfarbe (0,0,0)
	gfx.Rechteck (uint16(k.sp)*feldbreite,uint16(k.ze)*feldbreite,feldbreite,feldbreite)
}

type robi struct {  // Bestandteil einer Welt; existiert nur innerhalb einer Welt!  
	ze, sp uint8
	richtung uint8
	welt *welt
	klotzanzahl uint
}

func newRobi (ze,sp,ri uint8, w *welt) *robi { 
	var r *robi
	r = new(robi)
	r.ze = ze
	r.sp = sp
	r.richtung = ri
	r.welt = w
	r.klotzanzahl = 999
	return r
}

func (r *robi) kodieren() []byte {
	var b []byte = make ([]byte,5)
	b[0] = r.ze
	b[1] = r.sp
	b[2] = r.richtung
	b[3] = uint8 (r.klotzanzahl / 256)
	b[4] = uint8 (r.klotzanzahl % 256)
	return b
}

func (r *robi) dekodieren (b []byte) {
	r.ze = b[0]
	r.sp = b[1]
	r.richtung = b[2]
	r.klotzanzahl = uint(b[3])*256 + uint(b[4])
}
	
func (r *robi) zeichnen () {
	r.welt.feld[r.ze][r.sp].zeichnen()
	xm:=uint16(r.sp)*r.welt.feldbreite+r.welt.feldbreite/2
	ym:=uint16(r.ze)*r.welt.feldbreite+r.welt.feldbreite/2
	b:= r.welt.feldbreite
	switch r.richtung {
		case 2:
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreis(xm,ym-b*7/50,b*15/50)
		gfx.Volldreieck(xm-b*22/50,ym+b*11/50,xm-b*1/5,ym-b*3/10,xm-b*1/5,ym+b*11/50)
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreissektor(xm+b/50,ym-b*7/50,b*15/50,270,90)
		gfx.Volldreieck(xm+b*22/50,ym+b*11/50,xm+b*1/5,ym-b*3/10,xm+b*1/5,ym+b*11/50)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm,ym-b/10,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
		gfx.Volldreieck(xm,ym+b*21/50,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm-(b/50+1),ym+b*1/10,b/50+1,b*8/50)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*4/50,ym+b*7/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm-b*7/50,ym+b*9/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm,ym+b*1/10,b/50+1,b*8/50)									//+1 wegen Division ohne Rest
		gfx.Vollrechteck(xm+b*2/50,ym+b*7/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm+b*5/50,ym+b*9/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm+b*7/50,ym+b*23/50,xm+b*11/50,ym+b*23/50,xm+b*22/50,ym-b*23/50)	
		gfx.Volldreieck(xm+b*7/50,ym+b*23/50,xm+b*20/50,ym-b*23/50,xm+b*22/50,ym-b*23/50)
		case 1:
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreis(xm-b*7/50,ym,b*15/50)
		gfx.Volldreieck(xm+b*11/50,ym-b*22/50,xm-b*3/10,ym-b*1/5,xm+b*11/50,ym-b*1/5)
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreissektor(xm-b*7/50,ym+b/50,b*15/50,180,0)
		gfx.Volldreieck(xm+b*11/50,ym+b*22/50,xm-b*3/10,ym+b*1/5,xm+b*11/50,ym+b*1/5)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm-b/10,ym,xm+b*11/50,ym-b*22/50,xm+b*11/50,ym+b*22/50)
		gfx.Volldreieck(xm+b*21/50,ym,xm+b*11/50,ym-b*22/50,xm+b*11/50,ym+b*22/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm+b*1/10,ym-(b/50+1),b*8/50,b/50+1)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*7/50,ym-b*4/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm+b*9/50,ym-b*7/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm+b*1/10,ym,b*8/50,b/50+1)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*7/50,ym+b*2/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm+b*9/50,ym+b*5/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm+b*23/50,ym-b*7/50,xm+b*23/50,ym-b*11/50,xm-b*23/50,ym-b*22/50)	
		gfx.Volldreieck(xm+b*23/50,ym-b*7/50,xm-b*23/50,ym-b*20/50,xm-b*23/50,ym-b*22/50)
		case 0:
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreis(xm,ym+b*7/50,b*15/50)
		gfx.Volldreieck(xm-b*22/50,ym-b*11/50,xm-b*1/5,ym+b*3/10,xm-b*1/5,ym-b*11/50)
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreissektor(xm+b/50,ym+b*7/50,b*15/50,270,90)
		gfx.Volldreieck(xm+b*22/50,ym-b*11/50,xm+b*1/5,ym+b*3/10,xm+b*1/5,ym-b*11/50)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm,ym+b/10,xm-b*22/50,ym-b*11/50,xm+b*22/50,ym-b*11/50)
		gfx.Volldreieck(xm,ym-b*21/50,xm-b*22/50,ym-b*11/50,xm+b*22/50,ym-b*11/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm-(b/50+1),ym-b*13/50,b/50+1,b*8/50)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*4/50,ym-b*13/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm-b*7/50,ym-b*13/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm,ym-b*13/50,b/50+1,b*8/50)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*2/50,ym-b*13/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm+b*5/50,ym-b*13/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm-b*7/50,ym-b*23/50,xm-b*11/50,ym-b*23/50,xm-b*22/50,ym+b*23/50)	
		gfx.Volldreieck(xm-b*7/50,ym-b*23/50,xm-b*20/50,ym+b*23/50,xm-b*22/50,ym+b*23/50)
		case 3:
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreis(xm+b*7/50,ym,b*15/50)
		gfx.Volldreieck(xm-b*11/50,ym-b*22/50,xm+b*3/10,ym-b*1/5,xm-b*11/50,ym-b*1/5)
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreissektor(xm+b*7/50,ym+b/50,b*15/50,180,0)
		gfx.Volldreieck(xm-b*11/50,ym+b*22/50,xm+b*3/10,ym+b*1/5,xm-b*11/50,ym+b*1/5)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm+b/10,ym,xm-b*11/50,ym-b*22/50,xm-b*11/50,ym+b*22/50)
		gfx.Volldreieck(xm-b*21/50,ym,xm-b*11/50,ym-b*22/50,xm-b*11/50,ym+b*22/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm-b*13/50,ym-(b/50+1),b*8/50,b/50+1)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*13/50,ym-b*4/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm-b*13/50,ym-b*7/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm-b*13/50,ym,b*8/50,b/50+1)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*13/50,ym+b*2/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm-b*13/50,ym+b*5/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm-b*23/50,ym+b*7/50,xm-b*23/50,ym+b*11/50,xm+b*23/50,ym+b*22/50)	
		gfx.Volldreieck(xm-b*23/50,ym+b*7/50,xm+b*23/50,ym+b*20/50,xm+b*23/50,ym+b*22/50)
	}
	if r.welt.feld[r.ze][r.sp].klotzanzahl > 0 {
		gfx.Stiftfarbe(255,0,0)
		//gfx.Vollrechteck(uint16(r.sp)*r.welt.feldbreite,uint16(r.ze)*r.welt.feldbreite + (4*r.welt.feldbreite+2)/5,r.welt.feldbreite/5,r.welt.feldbreite/5)
		gfx.Kreis(uint16(r.sp)*r.welt.feldbreite+r.welt.feldbreite/2,uint16(r.ze)*r.welt.feldbreite+r.welt.feldbreite/2,r.welt.feldbreite/2-1)
		gfx.Kreis(uint16(r.sp)*r.welt.feldbreite+r.welt.feldbreite/2,uint16(r.ze)*r.welt.feldbreite+r.welt.feldbreite/2,r.welt.feldbreite/2-2)
		gfx.Kreis(uint16(r.sp)*r.welt.feldbreite+r.welt.feldbreite/2,uint16(r.ze)*r.welt.feldbreite+r.welt.feldbreite/2,r.welt.feldbreite/2-3)
		//
		gfx.Schreibe(uint16(r.sp)*r.welt.feldbreite+2,uint16(r.ze)*r.welt.feldbreite+2,fmt.Sprint(r.welt.feld[r.ze][r.sp].klotzanzahl))
	}
}

func (r *robi) loeschen () {
	r.welt.feld[r.ze][r.sp].zeichnen()
}

func (r *robi) laufen1 () {
	r.loeschen ()
	switch r.richtung {
		case 0: // nach Norden
		if r.ze > 0 {r.ze--}
		case 1: // nach Osten
		if r.sp < uint8(r.welt.weltbreite/r.welt.feldbreite) - 1 {r.sp++}
		case 2: // nach Süden
		if r.ze < uint8(r.welt.welthoehe/r.welt.feldbreite) - 1 {r.ze++}
		case 3: // nach Westen
		if r.sp > 0 {r.sp--}
	}
	r.zeichnen ()
}

func (r *robi) amRand () bool {
	switch r.richtung {
		case 0: return r.ze == 0
		case 1: return r.sp == uint8(r.welt.weltbreite/r.welt.feldbreite) - 1
		case 2: return r.ze == uint8(r.welt.welthoehe/r.welt.feldbreite) - 1
		case 3: return r.sp == 0
	}
	panic ("Unerwarteter Fehler in -amRand-!")
}

func (r *robi) inLinkerObererEcke () bool {
	return r.ze == 0 && r.sp == 0
}

func (r *robi) linksDrehen () {
	r.richtung = (r.richtung + 3) % 4
	r.zeichnen ()
}

func (r *robi) rechtsDrehen () {
	r.richtung = (r.richtung + 1) % 4
	r.zeichnen ()
}


type welt struct {
	weltbreite uint16
	welthoehe uint16
	feldbreite uint16
	schrittmodus bool
	feld [][]*kachel
	rob *robi
}

/*
func newWelt (weltbreite, welthoehe, feldbreite uint16) *welt {
	var w *welt
	w = new(welt)
	w.weltbreite = weltbreite
	w.welthoehe = welthoehe
	w.feldbreite = feldbreite
	w.schrittmodus = true //Standart: aktiviert
	w.feld = make ([][]*kachel,0)
	if gfx.FensterOffen () {
		if (gfx.Grafikspalten() != weltbreite || gfx.Grafikzeilen() != welthoehe) {
			gfx.FensterAus()
			gfx.Fenster(weltbreite,welthoehe)
		} else {
			gfx.Stiftfarbe(255,255,255)
			gfx.Vollrechteck (0,0,weltbreite,welthoehe)
		}
	} else {
		gfx.Fenster(weltbreite,welthoehe)
	}
	for zeile:=uint8(0);zeile < uint8(welthoehe/feldbreite);zeile++ {
		z:=make([]*kachel,0)
		for spalte:=uint8(0);spalte< uint8(weltbreite/feldbreite);spalte++ {
			f:=newKachel(zeile,spalte,220,220,220,180,250,180,w)
			f.zeichnen()
			z = append(z,f)
		}
		w.feld = append (w.feld,z)
	}
	if welthoehe/feldbreite < 1 || weltbreite/feldbreite < 1 {
		panic ("Die Welt ist zu klein für einen Robi!!")
	}
	w.rob = newRobi(0,0,2,w)												//hier andere Richtung?!
	w.rob.zeichnen()
	return w
}

func (w *welt) kodieren () []byte {
	var b []byte = make ([]byte,8)
	b[0] = uint8(w.weltbreite/256)
	b[1] = uint8(w.weltbreite%256)
	b[2] = uint8(w.welthoehe/256)
	b[3] = uint8(w.welthoehe%256)
	b[4] = uint8(w.feldbreite/256)
	b[5] = uint8(w.feldbreite%256)
	if w.schrittmodus {
		b[6] = 1
	} else {
		b[6] = 0
	} 
	feldstromlaenge:= len(w.feld[0][0].kodieren())
	if feldstromlaenge > 255 {
		panic ("Ein kodiertes Feld darf nicht mehr als 255 Bytes lang sein!")
	} 
	b[7] = uint8(feldstromlaenge)
	ze:=len (w.feld)
	sp:=len (w.feld[0])
	for i:=0; i < ze; i++ {
		for j:=0; j < sp; j++ {
			b = append(b,w.feld[i][j].kodieren()...)
		}
	}
	robistrom:= w.rob.kodieren()
	if len(robistrom) > 255 {
		panic("Ein kodierter Robi darf nicht mehr als 255 Bytes lang sein!")
	}
	b = append (b,uint8(len(robistrom)))
	b = append (b,robistrom...) 
	return b
}
*/

func newWelt (weltbreite, welthoehe, feldbreite uint16) *welt {
	var w *welt
	w = new(welt)
	w.weltbreite = weltbreite
	w.welthoehe = welthoehe
	w.feldbreite = feldbreite
	w.schrittmodus = true //Standart: aktiviert
	w.feld = make ([][]*kachel,0)
	if gfx.FensterOffen () {
		if (gfx.Grafikspalten() < weltbreite || gfx.Grafikzeilen() < welthoehe) {
			panic ("Das offene gfx-Fenster ist zu klein für die neue Welt!")
		} else {
			gfx.Stiftfarbe(255,255,255)
			gfx.Vollrechteck (0,0,weltbreite,welthoehe)
		}
	} else {
		panic ("Kein gfx-Fenster offen!")
	}
	for zeile:=uint8(0);zeile < uint8(welthoehe/feldbreite);zeile++ {
		z:=make([]*kachel,0)
		for spalte:=uint8(0);spalte < uint8(weltbreite/feldbreite);spalte++ {
			//f:=newKachel(zeile,spalte,220,220,220,180,250,180,w)
			f:=newKachel(zeile,spalte,220,220,220,0,255,0,w)
			f.zeichnen()
			z = append(z,f)
		}
		w.feld = append (w.feld,z)
	}
	if welthoehe/feldbreite < 1 || weltbreite/feldbreite < 1 {
		panic ("Die Welt ist zu klein für einen Robi!!")
	}
	w.rob = newRobi(0,0,2,w)												//hier andere Richtung?!
	w.rob.zeichnen()
	return w
}

func (w *welt) kodieren () []byte {
	var b []byte = make ([]byte,8)
	b[0] = uint8(w.weltbreite/256)
	b[1] = uint8(w.weltbreite%256)
	b[2] = uint8(w.welthoehe/256)
	b[3] = uint8(w.welthoehe%256)
	b[4] = uint8(w.feldbreite/256)
	b[5] = uint8(w.feldbreite%256)
	if w.schrittmodus {
		b[6] = 1
	} else {
		b[6] = 0
	} 
	feldstromlaenge:= len(w.feld[0][0].kodieren())
	if feldstromlaenge > 255 {
		panic ("Ein kodiertes Feld darf nicht mehr als 255 Bytes lang sein!")
	} 
	b[7] = uint8(feldstromlaenge)
	ze:=len (w.feld)
	sp:=len (w.feld[0])
	for i:=0; i < ze; i++ {
		for j:=0; j < sp; j++ {
			b = append(b,w.feld[i][j].kodieren()...)
		}
	}
	robistrom:= w.rob.kodieren()
	if len(robistrom) > 255 {
		panic("Ein kodierter Robi darf nicht mehr als 255 Bytes lang sein!")
	}
	b = append (b,uint8(len(robistrom)))
	b = append (b,robistrom...) 
	return b
}

func (w *welt) dekodieren (b []byte) *welt {
	var neueWelt *welt
	weltbreite := uint16(b[0])*256 + uint16(b[1])
	welthoehe  := uint16(b[2])*256 + uint16(b[3])
	feldbreite := uint16(b[4])*256 + uint16(b[5])
	neueWelt = newWelt(weltbreite, welthoehe, feldbreite)
	neueWelt.schrittmodus = b[6] == 1
	feldstromlaenge:= int(b[7])
	aktIndex:=8
	ze:=len (neueWelt.feld)
	sp:=len (neueWelt.feld[0])
	for i:=0; i < ze; i++ {
		for j:=0; j < sp; j++ {
			neueWelt.feld[i][j].dekodieren(b[aktIndex:aktIndex+feldstromlaenge])
			neueWelt.feld[i][j].zeichnen ()
			aktIndex = aktIndex + feldstromlaenge
		} 
	}
	robistromlaenge := int(b[aktIndex])
	aktIndex++
	neueWelt.rob.dekodieren(b[aktIndex:aktIndex+robistromlaenge])
	neueWelt.rob.zeichnen ()
	return neueWelt
}
	
func (w *welt) enter () {
	var taste uint16
	var gedrueckt uint8
	if w.schrittmodus {
		for {
			taste,gedrueckt,_= gfx.TastaturLesen1 ()
			if taste == 13 && gedrueckt == 1 {break}
		}
	}
}

func fehlerUndEnde (text string) {
	//var taste uint16
	//var gedrueckt uint8
	/*
	schrifthoehe := w.feldbreite/3
	ok:= gfx.SetzeFont(fontpath + "LiberationMono-Bold.ttf", int(schrifthoehe-4))
	if ok { // Font gefunden !
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-schrifthoehe,gfx.Grafikspalten(),schrifthoehe)
		gfx.Stiftfarbe (0,255,0)
		gfx.SchreibeFont (10,gfx.Grafikzeilen() - schrifthoehe*19/20,"VOR. VON **"+text+ "** NICHT EINGEHALTEN! --> <ESC>")
	} else { //Font nicht gefunden - Nothilfe: Verwendung des internen Fonts
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-12,gfx.Grafikspalten(),12)
		gfx.Stiftfarbe (0,255,0)
		gfx.Schreibe (10,gfx.Grafikzeilen() - 10," VOR. VON **"+text+ "** NICHT EINGEHALTEN! --> <ESC>")
	}
	*/
	/*
	for {
		taste,gedrueckt,_ = gfx.TastaturLesen1()
		if taste == 27 && gedrueckt == 1 {break}
	}
	os.Exit(1)
	*/
}

// Vor.: Robi hat in seiner Richtung einen betretbaren Nachbarplatz. //
// Eff.: Robi steht in der gleichen Richtung wie vorher auf diesem
// Nachbarplatz.
func Laufen1 () bool {
	w.enter ()
	if AmRand () || VorMauer () {
		fehlerUndEnde ("Laufen1")
		return false
	}
	w.rob.laufen1()
	return true
}

// Vor.: keine //
// Eff.: Robis Richtung ist um 90 Grad nach links gedreht.
func LinksDrehen () {
	w.enter ()
	w.rob.linksDrehen()
}

// Vor.: keine //
// Eff.: Robis Richtung ist um 90 Grad nach rechts gedreht.
func RechtsDrehen () {
	w.enter ()
	w.rob.rechtsDrehen()
}

// Vor.: keine //
// Erg.: -true- ist geliefert, genau dann, wenn Robi in seiner Richtung
// keinen Nachbarplatz hat. 
func AmRand () bool {
	return w.rob.amRand()
}

// Vor.: keine //
// Erg.: -true- ist geliefert, genau dann, wenn Robi in der linken 
// oberen Ecke der Welt (nordwestlichste Ecke) steht.
func InLinkerObererEcke () bool {
	return w.rob.inLinkerObererEcke ()
}

// Vor.: keine //
// Eff.: Wurde Schrittmodus (true) aufgerufen, so muss nun jede einzelne
// Aktion von Robi, die die Welt ändert, mit einem Druck auf die 
// Entertaste aktiviert werden, andernfalls nicht. 
// Standardeinstellung ist -true-.
func Schrittmodus (mode bool) {
	w.schrittmodus = mode
}

// Vor.: -text- ist ein String, der nur aus ASCII-Code-Zeichen besteht.
// Eff.: -text- und -n- sind in der letzten Zeile des Grafikfensters 
// ausgegeben. Der aufrufende Prozess wurde solange angehalten, bis der 
// Benutzer in der Ausgabe mit einem Druck auf <ESC> quittiert hatte.
// Nun ist die Meldung wieder aus dem Grafikfenster entfernt.
func Melden (text string) {
	var taste uint16
	var gedrueckt uint8
	gfx.Archivieren ()
	schrifthoehe := w.feldbreite/3
	ok:= gfx.SetzeFont(fontpath + "LiberationMono-Bold.ttf", int(schrifthoehe-4))
	if ok { // Font gefunden !
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-schrifthoehe ,gfx.Grafikspalten(),schrifthoehe)
		gfx.Stiftfarbe (0,255,0)
		gfx.SchreibeFont (10,gfx.Grafikzeilen() - schrifthoehe*39/40,text + "  --> <ESC>")
	} else { //Font nicht gefunden - Nothilfe: Verwendung des internen Fonts
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-12,gfx.Grafikspalten(),12)
		gfx.Stiftfarbe (0,255,0)
		gfx.Schreibe (10,gfx.Grafikzeilen() - 10,text + "  --> <ESC>")
	}
	for {
		taste,gedrueckt,_ = gfx.TastaturLesen1()
		if taste == 27 && gedrueckt == 1 {break}
	}
	gfx.Restaurieren(0,0,gfx.Grafikspalten(),gfx.Grafikzeilen())
}

// Vor.: -text- ist ein String, der nur aus ASCII-Code-Zeichen besteht.
// Eff.: -text- und -n- sind in der letzten Zeile des Grafikfensters 
// ausgegeben. Der aufrufende Prozess wurde solange angehalten, bis der 
// Benutzer in der Ausgabe mit einem Druck auf <ESC> quittiert hatte.
// Nun ist die Meldung wieder aus dem Grafikfenster entfernt. 
func FehlerMelden (text string) { Melden (text) }

// Vor.: keine //
// Eff.: Das Programm ist mit der Fehlermeldung "Programm beendet! 
// Druecken Sie <ESC> !" angehalten. Nachdem der Benutzer mit einem 
// Druck auf <ESC> quittiert hat, ist das Grafikfenster geschlossen und
// das Programm beendet.
func Fertig () {
	var taste uint16
	var gedrueckt uint8
	schrifthoehe:= w.feldbreite/3
	ok:= gfx.SetzeFont(fontpath + "LiberationMono-Bold.ttf", int(schrifthoehe-4))
	if ok { // Font gefunden !
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-schrifthoehe,gfx.Grafikspalten(),schrifthoehe)
		gfx.Stiftfarbe (0,255,0)
		gfx.SchreibeFont (10,gfx.Grafikzeilen() - schrifthoehe*39/40,"Spiel gemeistert - ALP2 absolviert! :) -->  <ESC> !")
	} else { //Font nicht gefunden - Nothilfe: Verwendung des internen Fonts
		gfx.Stiftfarbe (255,0,0)
		gfx.Vollrechteck (0,gfx.Grafikzeilen()-12,gfx.Grafikspalten(),12)
		gfx.Stiftfarbe (0,255,0)
		gfx.Schreibe (10,gfx.Grafikzeilen() - 10,"Spiel gemeistert - ALP2 absolviert! :) -->  <ESC> !")
	}
	for {
		taste,gedrueckt,_ = gfx.TastaturLesen1()
		if taste == 27 && gedrueckt == 1 {break} 
	}
	os.Exit(0)
}

// Vor.: keine //
// Eff.: Robis Platz ist markiert.
func Markieren () {
	w.enter ()
	w.feld[w.rob.ze][w.rob.sp].markiert = true
	w.feld[w.rob.ze][w.rob.sp].zeichnen ()
	w.rob.zeichnen ()
}

// Vor.: keine //
// Eff.: Robis Platz ist nicht markiert.
func Demarkieren () {
	w.enter ()
	w.feld[w.rob.ze][w.rob.sp].markiert = false
	w.feld[w.rob.ze][w.rob.sp].zeichnen ()
	w.rob.zeichnen ()
}

// Vor.: keine //
// Erg.: -true- ist geliefert, falls Robis Platz markiert ist, 
// andernfalls ist -false- geliefert.
func Markiert () bool {
	return w.feld[w.rob.ze][w.rob.sp].markiert
}

// Vor.: In Robis Richtung gibt es einen Nachbarplatz. //
// Erg.: -true- ist geliefert, falls der Nachbarplatz in Robis Richtung
// markiert ist. Andernfalls ist -false- geliefert.
func NachbarMarkiert () bool {
	if w.rob.amRand() {
		fehlerUndEnde("NachbarMarkiert")
	}
	switch w.rob.richtung {
		case 0: return w.feld[w.rob.ze-1][w.rob.sp].markiert
		case 1: return w.feld[w.rob.ze][w.rob.sp+1].markiert
		case 2: return w.feld[w.rob.ze+1][w.rob.sp].markiert
		case 3: return w.feld[w.rob.ze][w.rob.sp-1].markiert
	}
	panic ("Unerwarteter Fehler in -NachbarMarkiert()-!")
}

// Vor.: keine //
// Erg.: -true- ist geliefert, falls auf Robis Platz kein Klotz liegt,
// anderfalls ist -false- geliefert.
func Leer () bool {
	return w.feld[w.rob.ze][w.rob.sp].klotzanzahl == 0
}

// Vor.: Auf Robis Platz liegt mindestens ein Klotz. //
// Eff.: Auf Robis Platz liegt nun ein Klotz weniger und in Robis Tasche
// ist ein Klotz mehr. In der linken oberen Ecke des Platzes ist die
// aktuelle Anzahl der Klötze auf dem Platz angezeigt. Ist kein Klotz auf 
// dem Platz, so ist keine Anzahl der Klötze angezeigt.
func Leeren1 () bool {
	w.enter ()
	if w.feld[w.rob.ze][w.rob.sp].klotzanzahl > 0 {
		w.feld[w.rob.ze][w.rob.sp].klotzanzahl--
		w.rob.klotzanzahl++
		w.feld[w.rob.ze][w.rob.sp].zeichnen ()
		w.rob.zeichnen()
		return true
	} else {
		fehlerUndEnde("Leeren1")
		return false
	}
}

// Vor.: keine //
// Erg.: -true- ist geliefert, falls Robi in seiner Tasche noch 
// mindestens einen Klotz hat. Andernfalls ist -false- geliefert.
func HatKloetze () bool {
	return w.rob.klotzanzahl > 0
}

// Vor.: Robis Tasche mit Klötzen ist nicht leer. //
// Eff.: In Robis Tasche ist nun ein Klotz weniger und auf seinem Platz 
// ist nun einer mehr. In der linken oberen Ecke des Platzes ist die
// aktuelle Anzahl der Klötze auf dem Platz angezeigt. 
func Legen1 () bool {
	w.enter ()
	if w.rob.klotzanzahl > 0 {
		w.rob.klotzanzahl--
		w.feld[w.rob.ze][w.rob.sp].klotzanzahl++
		w.feld[w.rob.ze][w.rob.sp].zeichnen ()
		w.rob.zeichnen()
		return true
	} else {
		fehlerUndEnde("Legen1")
		return false
	}
} 

// Vor.: keine //
// Erg.: -true- is geliefert, falls in Robis Richtung ein Nachbarfeld
// ist, das zugemauert ist. Andernfalls ist -false- geliefert.
func VorMauer () bool {
	if w.rob.amRand () {
		return false
	}
	switch w.rob.richtung {
		case 0: return w.feld[w.rob.ze-1][w.rob.sp].zugemauert
		case 1: return w.feld[w.rob.ze][w.rob.sp+1].zugemauert
		case 2: return w.feld[w.rob.ze+1][w.rob.sp].zugemauert
		case 3: return w.feld[w.rob.ze][w.rob.sp-1].zugemauert
	}
	panic ("Unerwarteter Fall in -VorMauer-!")
}

// Vor.: In Robis Richtung gibt es ein Nachbarfeld, das nicht 
// zugemauert ist. //
// Eff.: Robi steht jetzt auf diesem Nachbarfeld und sein ursprünglicher
// Standort ist nun zugemauert. Klötze, die sich auf dem zugemauerten
// Feld befunden haben, sind nun in Robis Tasche. Eine auf dem 
// zugemauerten Feld ggf. vorhandene Markierung ist jetzt entfernt.
func Mauern1 () bool {
	w.enter ()
	if !AmRand () && !VorMauer () {
		//plus Markieren()
		w.feld[w.rob.ze][w.rob.sp].markiert = true
		w.feld[w.rob.ze][w.rob.sp].zeichnen ()
		//w.feld[w.rob.ze][w.rob.sp].markiert = false
		w.rob.klotzanzahl = w.rob.klotzanzahl + w.feld[w.rob.ze][w.rob.sp].klotzanzahl
		w.feld[w.rob.ze][w.rob.sp].klotzanzahl=0
		w.feld[w.rob.ze][w.rob.sp].zugemauert = true
		w.rob.laufen1 ()
		return true
	} else {
		fehlerUndEnde("Mauern1")
		return false
	}
}

// Vor.: In Robis Richtung gibt es ein Nachbarfeld, das zugemauert ist. 
// Eff.: Dieses Nachbarfeld ist nun nicht mehr zugemauert und Robi steht
// mit gleicher Richtung auf diesem  entmauerten Platz. 
func Entmauern1 () bool {
	w.enter ()
	if !AmRand() && VorMauer() {
		switch w.rob.richtung {
			case 0: w.feld[w.rob.ze-1][w.rob.sp].zugemauert= false
			case 1: w.feld[w.rob.ze][w.rob.sp+1].zugemauert= false
			case 2: w.feld[w.rob.ze+1][w.rob.sp].zugemauert= false
			case 3: w.feld[w.rob.ze][w.rob.sp-1].zugemauert= false
		}
		w.rob.laufen1 ()
		//einfügen Demarkieren()
		w.feld[w.rob.ze][w.rob.sp].markiert = false
		w.feld[w.rob.ze][w.rob.sp].zeichnen ()
		w.rob.zeichnen ()
		//
		return true
	} else {
		fehlerUndEnde("Entmauern1")
		return false
	}
}

// Vor.: -p- ist ein Wert zwischen 50 und 300 und gibt den Zoomfaktor
//       in Prozent an.
// Eff.: Robis Welt ist entsprechend dem Zoomfaktor auf dem Bildschirm
//       angezeigt. Das Grafikfenster hat seine Größe geändert. Die Feld-
//       breite der Felder von Robis Welt wurde entsprechend angepasst.
func WeltZoomen (p uint16) {
	if p < 50 || p > 300 { return }
	
	zeilenanzahl:= uint8(w.welthoehe/w.feldbreite)   // Anzahl der Zeilen
	spaltenanzahl:= uint8(w.weltbreite/w.feldbreite) // Anzahl der Spalten

	//Jetzt zoomen ...
	w.feldbreite = w.feldbreite * p / 100
	w.weltbreite = w.feldbreite * uint16(spaltenanzahl)
	w.welthoehe  = w.feldbreite * uint16(zeilenanzahl)
		
	if gfx.FensterOffen () {
		if (gfx.Grafikspalten() != w.weltbreite || gfx.Grafikzeilen() != w.welthoehe) {
			gfx.FensterAus()
			gfx.Fenster(w.weltbreite,w.welthoehe)
		} else {
			gfx.Stiftfarbe(255,255,255)
			gfx.Vollrechteck (0,0,w.weltbreite,w.welthoehe)
		}
	} else {
		gfx.Fenster(w.weltbreite,w.welthoehe)
	}
	for z:=uint8(0);z < zeilenanzahl ; z++ {
		for s:=uint8(0);s < spaltenanzahl ; s++ {
			w.feld[z][s].zeichnen() // Felder neu darstellen
		}
	}
	w.rob.zeichnen()  // Robi neu darstellen
}
	
// Vor.: -name- ist ein gültiger Dateiname inkl. Pfad und die zugehörige
// Datei existiert dort. //
// Eff.: Die Welt aus der Datei -name- wurde aktiviert. Dazu gehört auch
// Robis Platz und seine Blickrichtung.
func WeltLaden (name string) {
	var b []byte = make ([]byte,0)
	datei:=dateien.Oeffnen (name,'l')
	stromlaenge:= datei.Groesse ()
	for i:=uint64(0); i < stromlaenge; i++ {
		b = append (b, datei.Lesen())
	}
	datei.Schliessen ()
	w = w.dekodieren (b) 
}

// Vor.: -name- ist ein gültiger Dateiname inkl. Pfad und diese Datei
// existiert bis jetzt noch nicht. //
// Eff.: Diese Datei existiert nun und die aktuelle Welt und Robis
// Eigenschaften sind in der Datei -name- abgespeichert.
func WeltSpeichern (name string) {
	b:=w.kodieren ()
	datei:=dateien.Oeffnen (name,'x')
	groesse:= datei.Groesse ()
	datei.Schliessen ()
	if groesse > 0 {panic ("Datei existierte schon! Programmabbruch!")}
	
	datei = dateien.Oeffnen (name,'s')
	for i:=0;i<len(b);i++ {datei.Schreiben (b[i])}
	datei.Schliessen ()
}

// Vor.: -breite- und -hoehe- sind nicht kleiner als -feldbreite-. //
// Eff.: Im Grafikfenster erscheint eine neue leere Welt mit Robi in der
// linken oberen Ecke mit Blickrichtung nach Süden. Die vorher 
// vorhandene Welt ist verloren.
func NeueWelt (breite, hoehe, feldbreite uint16) {
	w = newWelt(breite, hoehe, feldbreite)
}

// Vor.: keine //
// Eff.: Solange der Benutzer nicht <ESC> drückt ist er im Baumodus
// gefangen und kann Robi mit den Pfeiltasten durch die Welt steuern.
// Außerdem kann er mit <1> Robis Platz markieren bzw. demarkieren,
// mit <2> jeweils einen Klotz auf Robis Platz ablegen, solange Robi
// Klötze in der Tasche hat, und mit <3> jeweils einen Klotz von Robis
// Platz aufnehmen und in die Tasche stecken, solange Klötze auf Robis
// Platz liegen. Mit <4> kann gemauert und mit <5> entmauert werden.
func Baumodus () {
	var taste uint16
	var gedrueckt uint8
	schrittmodus:= w.schrittmodus
	w.schrittmodus = false
	for {
		taste,gedrueckt,_= gfx.TastaturLesen1 ()
		if gedrueckt == 1 {
			switch taste {
				case 27:
				w.schrittmodus = schrittmodus
				return
				case 273:
				if !AmRand() && !VorMauer () {Laufen1()}
				case 276:
				LinksDrehen ()
				case 275:
				RechtsDrehen ()
				case 49:
				if !Markiert () {
					Markieren ()
				} else {
					Demarkieren () 
				}
				case 50:
				if HatKloetze() {Legen1()}
				case 51:
				if !Leer () {Leeren1()}
				case 52:
				if !AmRand() && !VorMauer () {Mauern1()}
				case 53:
				if VorMauer () {Entmauern1()}
			}
		}
	}
}

func WeltOeffnen () {
	switch runtime.GOOS {
		case "linux": 
		if fontpath == "" { //wsl erkennt GOPATH nicht
			fontpath = "/home/lewein/go/"
		} 
		fontpath = fontpath + "/src/gfx/fonts/"
		case "windows":
		fontpath = os.Getenv("GOPATH")+"\\src\\gfx\\fonts\\"
		default:
		fmt.Println ("Betriebssystem nicht erkannt!")
	}
	if !gfx.FensterOffen() {
		gfx.Fenster(1200,700)
	}
	w = newWelt(675,675,75)
}
