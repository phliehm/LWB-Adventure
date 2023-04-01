package objekte 
// Autor: B. Schneider
// Datum: 21.03.2023
// Zweck: Objekte für das SWP

import (. "gfx" ; "sync")

var m sync.Mutex
	
type data struct {                                      // Zugriff
	x,y uint16          	// Koordinaten der linken oberen Ecke
	typ uint8            	// Objekt-Typ
	qua	uint16				// Quadratgröße = Breite des Objekts
	r,g,b uint8				// Farbe des Objekts (optional)
}

/*
 * OBJEKTE:
 * 0: rotes Quadrat
 * 1: grüne Kugel
 * 
 * Semaphor s / Mutex m im Sync-Paket: s.P() & s.V() / m.Lock() & m.Unlock() 
*/


func New (x,y, qua uint16, typ uint8) *data {
	var ob *data
	ob = new(data)
	ob.x = x
	ob.y = y
	ob.qua = qua
	ob.typ = typ
	return ob
}

func (ob *data) SetzeKoordinaten(x,y uint16) {
	//m.Lock()
	ob.x = x
	ob.y = y
	//m.Unlock()
}

func (ob *data) GibKoordinaten() (uint16,uint16) {
	return ob.x, ob.y
}

func (ob *data) SetzeTyp(t uint8) {
	ob.typ = t
}

func (ob *data) GibTyp() (uint8) {
	return ob.typ
}

func (ob *data) SetzeFarbe(r,g,b uint8) {
	ob.r = r
	ob.g = g
	ob.b = b
}

func (ob *data) GibFarbe() (uint8, uint8, uint8) {
	return ob.r,ob.g,ob.b
}

func (ob *data) Zeichnen() {
	switch ob.typ {
		case 0:																// Fadenkreuz-Maus-Zeiger bei Koord x,y
		m.Lock()
		Stiftfarbe(255,255,255)
		for i:=uint16(0);i<ob.qua/10;i++ {
			Kreis(ob.x, ob.y, ob.qua/2-i)
		}
		Vollrechteck(ob.x-ob.qua/20, 	ob.y-ob.qua*7/10,	ob.qua/10,ob.qua/2)
		Vollrechteck(ob.x-ob.qua/20, 	ob.y+ob.qua*2/10,	ob.qua/10,ob.qua/2)
		Vollrechteck(ob.x-ob.qua*7/10, 	ob.y-ob.qua/20,		ob.qua/2,ob.qua/10)
		Vollrechteck(ob.x+ob.qua*2/10, 	ob.y-ob.qua/20,		ob.qua/2,ob.qua/10)
		m.Unlock()
		
		case 1:																// PAUSE - Anzeige
		m.Lock()
		SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua) )
		Stiftfarbe(230,230,230)
		Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
		Stiftfarbe(255,0,127)  
		SchreibeFont (ob.x/6,ob.y/2-ob.qua*5/12,"PAUSE")					// Schreibe mittig Pause
		m.Unlock()
		
		case 2:																// rotes Quadrat ab linker oberer Ecke
		m.Lock()
		Rechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)
		Stiftfarbe(ob.r,ob.g,ob.b)
		Vollrechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)
		m.Unlock()
		
		case 3:																// Kaffee-Tasse ab linker oberer Ecke
		m.Lock()
		Stiftfarbe(0,0,0)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
		Stiftfarbe(100,100,255)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-ob.qua/30-1)
		
		Stiftfarbe(0,0,0)												// schwarze Umrandung Außen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10, ob.qua*16/100)
		Stiftfarbe(224,224,224)											// weißer Außen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10-ob.qua/40, ob.qua*15/100)
		
		Stiftfarbe(0,0,0)												// schwarze Umrandung Innen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*2/9, ob.qua*8/100)
		Stiftfarbe(155,152,152)											// grauer Innen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*2/9-ob.qua/60, ob.qua*7/100)
		
		Stiftfarbe(0,0,0)			// schwarze Umrandung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*11/100)
		Stiftfarbe(231,62,243)			// pinke Füllung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*9/100)
		Stiftfarbe(0,0,0)			// schwaze Innen-Füllung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*7/100)
		Stiftfarbe(0,0,255)			// blaues Inneres
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*5/100)
		
		Stiftfarbe(0,0,0)			// schwarze Umrandung
		Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100,180,0)
		Stiftfarbe(231,62,243)		// pinke Tasse
		Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*38/100, ob.qua*35/100,180,0)

		Stiftfarbe(0,0,0)			// schwarze Umrandung Kaffee
		Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100, ob.qua*10/100)
		Stiftfarbe(153,76,13)		// brauner Kaffee
		// Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*34/100, ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*38/100, ob.qua*34/100, ob.qua*7/100)
		
		Stiftfarbe(222,214,214)
		Vollellipse(ob.x+ob.qua*28/100, 	ob.y+ob.qua*20/100,	ob.qua/60,	ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*29/100+1, 	ob.y+ob.qua*24/100,	ob.qua/60,	ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*28/100, 	ob.y+ob.qua*28/100,	ob.qua/60,	ob.qua*8/100) 

		Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*8/50,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*46/100-1, 	ob.y+ob.qua*12/50,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*16/50,	ob.qua/60,	ob.qua*9/100) 

		Vollellipse(ob.x+ob.qua*63/100, 	ob.y+ob.qua*19/100,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*64/100+1, 	ob.y+ob.qua*23/100,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*63/100, 	ob.y+ob.qua*27/100,	ob.qua/60,	ob.qua*9/100) 
		m.Unlock()
		
		case 4:																// Kaffee-Tasse ab linker oberer Ecke - GESTRICHEN
		m.Lock()
		Stiftfarbe(0,0,0)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
		Stiftfarbe(100,100,255)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-ob.qua/30-1)
		
		Stiftfarbe(0,0,0)												// schwarze Umrandung Außen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10, ob.qua*16/100)
		Stiftfarbe(224,224,224)											// weißer Außen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10-ob.qua/40, ob.qua*15/100)
		
		Stiftfarbe(0,0,0)												// schwarze Umrandung Innen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*2/9, ob.qua*8/100)
		Stiftfarbe(155,152,152)											// grauer Innen-Unterteller
		Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*2/9-ob.qua/60, ob.qua*7/100)
		
		Stiftfarbe(0,0,0)			// schwarze Umrandung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*11/100)
		Stiftfarbe(231,62,243)			// pinke Füllung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*9/100)
		Stiftfarbe(0,0,0)			// schwaze Innen-Füllung Henkel
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*7/100)
		Stiftfarbe(0,0,255)			// blaues Inneres
		Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*48/100,ob.qua*5/100)
		
		Stiftfarbe(0,0,0)			// schwarze Umrandung
		Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100,180,0)
		Stiftfarbe(231,62,243)		// pinke Tasse
		Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*38/100, ob.qua*35/100,180,0)

		Stiftfarbe(0,0,0)			// schwarze Umrandung Kaffee
		Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100, ob.qua*10/100)
		Stiftfarbe(153,76,13)		// brauner Kaffee
		// Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*34/100, ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*38/100, ob.qua*34/100, ob.qua*7/100)
		
		Stiftfarbe(222,214,214)
		Vollellipse(ob.x+ob.qua*28/100, 	ob.y+ob.qua*20/100,	ob.qua/60,	ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*29/100+1, 	ob.y+ob.qua*24/100,	ob.qua/60,	ob.qua*8/100)
		Vollellipse(ob.x+ob.qua*28/100, 	ob.y+ob.qua*28/100,	ob.qua/60,	ob.qua*8/100) 

		Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*8/50,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*46/100-1, 	ob.y+ob.qua*12/50,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*16/50,	ob.qua/60,	ob.qua*9/100) 

		Vollellipse(ob.x+ob.qua*63/100, 	ob.y+ob.qua*19/100,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*64/100+1, 	ob.y+ob.qua*23/100,	ob.qua/60,	ob.qua*9/100)
		Vollellipse(ob.x+ob.qua*63/100, 	ob.y+ob.qua*27/100,	ob.qua/60,	ob.qua*9/100) 
			
		Stiftfarbe(255,0,0)																			// Durchstreichung
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua/10, ob.x+ob.qua/10,ob.y+ob.qua/5, ob.x+ob.qua*9/10,ob.y+ob.qua*9/10)
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua/10, ob.x+ob.qua*9/10,ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,ob.y+ob.qua*8/10)
		
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua*9/10, ob.x+ob.qua/10,ob.y+ob.qua*8/10, ob.x+ob.qua*9/10,ob.y+ob.qua/10)
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,ob.y+ob.qua/10, ob.x+ob.qua*9/10,ob.y+ob.qua/5)
		m.Unlock()

		case 5:																							// PIZZA
		Stiftfarbe(0,0,0)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
		Stiftfarbe(229,212,186)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-ob.qua/30-1)								// Kreis-Hintergrund
		Stiftfarbe(225,141,0)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 60, 360)		// brauner Pizzaboden
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*2/5, 0, 60)
		Stiftfarbe(255,65,59)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*18/50, 60, 360)	// Tomatensauce
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*18/50, 0, 60)
		Stiftfarbe(255,215,33)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*16/50, 60, 360)	// Käse
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*16/50, 0, 60)
		Stiftfarbe(169,8,0)
		Vollkreis(ob.x+ob.qua*10/20,	ob.y+ob.qua*6/20,	ob.qua/20)										// Salami
		Vollkreis(ob.x+ob.qua*15/20,ob.y+ob.qua*6/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*5/20,	ob.y+ob.qua*8/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*11/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*11/20,ob.y+ob.qua*14/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*15/20,	ob.qua/20)
		
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/8,	ob.y+ob.qua/2+ob.qua/50,	ob.qua/20, 180, 360)
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/5,	ob.y+ob.qua/2-ob.qua/30,	ob.qua/20, 0, 180)
		
		Stiftfarbe(0,0,0)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 60, 119)			// Pizzastücke-Umrandung
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 120, 179)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 180, 239)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 240, 299)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 300, 360)
		Kreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*2/5, 0, 60)
		
		case 6:																							// PIZZA - gestrichen
		Stiftfarbe(0,0,0)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
		Stiftfarbe(229,212,186)
		Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-ob.qua/30-1)								// Kreis-Hintergrund
		Stiftfarbe(225,141,0)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 60, 360)		// brauner Pizzaboden
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*2/5, 0, 60)
		Stiftfarbe(255,65,59)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*18/50, 60, 360)	// Tomatensauce
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*18/50, 0, 60)
		Stiftfarbe(255,215,33)
		Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*16/50, 60, 360)	// Käse
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*16/50, 0, 60)
		Stiftfarbe(169,8,0)
		Vollkreis(ob.x+ob.qua*10/20,	ob.y+ob.qua*6/20,	ob.qua/20)										// Salami
		Vollkreis(ob.x+ob.qua*15/20,ob.y+ob.qua*6/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*5/20,	ob.y+ob.qua*8/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*11/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*11/20,ob.y+ob.qua*14/20,	ob.qua/20)
		Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*15/20,	ob.qua/20)
		
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/8,	ob.y+ob.qua/2+ob.qua/50,	ob.qua/20, 180, 360)
		Vollkreissektor(ob.x+ob.qua/2+ob.qua/5,	ob.y+ob.qua/2-ob.qua/30,	ob.qua/20, 0, 180)
		
		Stiftfarbe(0,0,0)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 60, 119)			// Pizzastücke-Umrandung
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 120, 179)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 180, 239)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 240, 299)
		Kreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 300, 360)
		Kreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*2/5, 0, 60)
		
		Stiftfarbe(255,0,0)																			// Durchstreichung
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua/10, ob.x+ob.qua/10,ob.y+ob.qua/5, ob.x+ob.qua*9/10,ob.y+ob.qua*9/10)
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua/10, ob.x+ob.qua*9/10,ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,ob.y+ob.qua*8/10)
		
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua*9/10, ob.x+ob.qua/10,ob.y+ob.qua*8/10, ob.x+ob.qua*9/10,ob.y+ob.qua/10)
		Volldreieck(ob.x+ob.qua/10,ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,ob.y+ob.qua/10, ob.x+ob.qua*9/10,ob.y+ob.qua/5)

	}
}

func (ob *data) Getroffen(x,y uint16) bool {														// Checkt quasi, ob Hit-Box getroffen
	switch ob.typ {
		case 2:
		if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
			return true
		} else {
			return false
		}  

		default:
		if ob.x+ob.qua/10 < x && x < ob.x+ob.qua*9/10 	&& 	ob.y+ob.qua/10 < y && y < ob.y+ob.qua*9/10 {
			return true
		} else {
			return false
		}
	}
}
