package objekte 
// Autor: B. Schneider
// Datum: 21.03.2023
// Zweck: Objekte für das SWP

import (. "gfx" ; "sync"; "time"; "math")

var m sync.Mutex
	
type data struct {                                      // Zugriff
	x,y uint16          	// Koordinaten der linken oberen Ecke
	typ uint8            	// Objekt-Typ
	qua	uint16				// Quadratgröße = Breite des Objekts
	r,g,b uint8				// Farbe des Objekts (optional)
	aktiv bool				// Gibt an, ob das Objekt aktiv ist
	erstellt int64			// Zeit der Erstellung
}


func New (x,y, qua uint16, typ uint8) *data {
	var ob *data
	ob = new(data)
	ob.x = x
	ob.y = y
	ob.qua = qua
	ob.typ = typ
	ob.aktiv = true
	ob.erstellt = time.Now().UnixNano()
	return ob
}

func (ob *data) SetzeKoordinaten(x,y uint16) {
	ob.x = x
	ob.y = y
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

func (ob *data) SetzeAkt(akt bool) {
	ob.aktiv = akt
}

func (ob *data) GibAkt() (bool) {
	return ob.aktiv
}

func (ob *data) GibErstellung() (int64) {
	return ob.erstellt
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
	if ob.aktiv {
		switch ob.typ {
// Fadenkreuz-Maus-Zeiger bei Koord x,y
			case 0:	
			Stiftfarbe(240,240,240)
			for i:=uint16(0);i<ob.qua/10;i++ {
				Kreis(ob.x, ob.y, ob.qua/2-i)
			}
			Vollrechteck(ob.x-ob.qua/20, 	ob.y-ob.qua*7/10,	ob.qua/10,ob.qua/2)
			Vollrechteck(ob.x-ob.qua/20, 	ob.y+ob.qua*2/10,	ob.qua/10,ob.qua/2)
			Vollrechteck(ob.x-ob.qua*7/10, 	ob.y-ob.qua/20,		ob.qua/2,ob.qua/10)
			Vollrechteck(ob.x+ob.qua*2/10, 	ob.y-ob.qua/20,		ob.qua/2,ob.qua/10)

// PAUSE - Anzeige			
			case 1:		
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua) )
			Stiftfarbe(230,230,230)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y/2-ob.qua*5/12,"PAUSE")					// Schreibe mittig Pause
			
// rotes Quadrat ab linker oberer Ecke			
			case 2:		
			Rechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)
			Stiftfarbe(ob.r,ob.g,ob.b)
			Vollrechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)

// Kaffee-Tasse ab linker oberer Ecke			
			case 3:			
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

// Kaffee-Tasse ab linker oberer Ecke - GESTRICHEN			
			case 4:			
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

// PIZZA
			case 5:																							
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

// PIZZA - gestrichen			
			case 6:																							
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
			
// LEVEL 1 - Anzeige			
			case 7:	
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*5/24 )
			Stiftfarbe(144,99,31)
			Vollellipse(ob.x/2,ob.y/2,ob.x*26/60,ob.y*13/60)
			Stiftfarbe(68,225,255)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(127,25,64)  
			SchreibeFont (ob.x/6,ob.y*20/48,"LEVEL 1")
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y*19/48,"LEVEL 1")					// Schreibe mittig Level 1

// LEVEL 2 - Anzeige			
			case 8:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*5/24 )
			Stiftfarbe(144,99,31)
			Vollellipse(ob.x/2,ob.y/2,ob.x*26/60,ob.y*13/60)
			Stiftfarbe(68,225,255)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(127,25,64)  
			SchreibeFont (ob.x/6,ob.y*20/48,"LEVEL 2")
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y*19/48,"LEVEL 2")					// Schreibe mittig Level 2
			
// LEVEL 3 - Anzeige			
			case 9:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*5/24 )
			Stiftfarbe(144,99,31)
			Vollellipse(ob.x/2,ob.y/2,ob.x*26/60,ob.y*13/60)
			Stiftfarbe(68,225,255)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(127,25,64)  
			SchreibeFont (ob.x/6,ob.y*20/48,"LEVEL 3")
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y*19/48,"LEVEL 3")					// Schreibe mittig Level 3
			
// LEVEL 4 - Anzeige			
			case 10:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*5/24 )
			Stiftfarbe(144,99,31)
			Vollellipse(ob.x/2,ob.y/2,ob.x*26/60,ob.y*13/60)
			Stiftfarbe(68,225,255)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(127,25,64)  
			SchreibeFont (ob.x/6,ob.y*20/48,"LEVEL 4")
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y*19/48,"LEVEL 4")					// Schreibe mittig Level 4

// ZIELSCHEIBE			
			case 11:																														
			Stiftfarbe(230,0,0)
			Vollkreis(ob.x,ob.y,ob.qua/2)
			Stiftfarbe(0,0,0)
			Kreis(ob.x,ob.y,ob.qua/2)	
			Stiftfarbe(255,255,255)  
			Vollkreis(ob.x,ob.y,ob.qua*4/10)
			Stiftfarbe(230,0,0)
			Vollkreis(ob.x,ob.y,ob.qua*3/10)
			Stiftfarbe(255,255,255)  
			Vollkreis(ob.x,ob.y,ob.qua/5)
			Stiftfarbe(230,0,0)
			Vollkreis(ob.x,ob.y,ob.qua/10)

// Countdown - 3			
			case 12:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*13/30 )
			/*Stiftfarbe(249,195,141)
			Vollkreis(ob.x/2,ob.y/2,ob.qua/4)
			Stiftfarbe(72,226,180)
			Vollkreis(ob.x/2,ob.y/2,ob.qua*9/40)*/
			Stiftfarbe(186,66,23)
			SchreibeFont (ob.x*43/100,ob.y*35/100,"3")
			Stiftfarbe(68,215,210)
			SchreibeFont (ob.x*41/100,ob.y*34/100,"3")

// Countdown - 2			
			case 13:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*13/30 )
			Stiftfarbe(186,66,23)
			SchreibeFont (ob.x*43/100,ob.y*35/100,"2")
			Stiftfarbe(68,215,210)
			SchreibeFont (ob.x*41/100,ob.y*34/100,"2")

// Countdown - 1			
			case 14:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*13/30 )
			Stiftfarbe(186,66,23)
			SchreibeFont (ob.x*44/100,ob.y*35/100,"1")
			Stiftfarbe(68,215,210)
			SchreibeFont (ob.x*42/100,ob.y*34/100,"1")

// LEVEL 5 - Anzeige			
			case 15:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)*5/24 )
			Stiftfarbe(144,99,31)
			Vollellipse(ob.x/2,ob.y/2,ob.x*26/60,ob.y*13/60)
			Stiftfarbe(68,225,255)
			Vollellipse(ob.x/2,ob.y/2,ob.x*5/12,ob.y/5)
			Stiftfarbe(127,25,64)  
			SchreibeFont (ob.x/6,ob.y*20/48,"LEVEL 5")
			Stiftfarbe(255,0,127)  
			SchreibeFont (ob.x/6,ob.y*19/48,"LEVEL 5")					// Schreibe mittig Level 5

// Heidi			
			case 16:																														
			LadeBild (ob.x,ob.y, "../../Bilder/Heidi-2.bmp")

// Kaffee-Tasse 2 ab linker oberer Ecke			
			case 18:			
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

// PIZZA 2
			case 19:																							
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

// OK - Objekt			
			case 20:															
			Stiftfarbe(142,36,24)
			Vollrechteck(ob.x*4/10,ob.y*8/10,ob.x/5,ob.y/10)
			Stiftfarbe(255,124,128)
			Vollrechteck(ob.x*41/100,ob.y*81/100,ob.x*9/50,ob.y*8/100)
			SetzeFont ("../../Schriftarten/Freshman.ttf", int(ob.qua)/15 )
			Stiftfarbe(124,212,255)
			SchreibeFont (ob.x*46/100,ob.y*82/100,"O K")
			
		}
	}
}

func (ob *data) Getroffen(x,y uint16, opt uint8) (bool,int64) {														// Checkt, ob Hit-Box getroffen
	if ob.aktiv {		
		switch ob.typ {
			case 2:
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				ob.SetzeAkt(false)
				SpieleNote("4A",0.1,0)
				return true, time.Now().UnixNano() - ob.GibErstellung()
			} else {
				return false, 0
			}
			case 3:		// Kaffee
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 3 {
					ob.SetzeAkt(false)
					SpieleSound("../../Sounds/Trinken.wav")
					return true, time.Now().UnixNano() - ob.GibErstellung()
				}
				ob.SetzeTyp(4)
				SpieleSound("../../Sounds/GameOver.wav")
				return true, 0
			} else {
				return false, 0
			}  
			case 5:		// Pizza
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 1 {
					ob.SetzeAkt(false)
					SpieleSound("../../Sounds/Essen.wav")
					return true, time.Now().UnixNano() - ob.GibErstellung()
				}
				ob.SetzeTyp(6)
				SpieleSound("../../Sounds/GameOver.wav")
				return true, 0
			} else {
				return false, 0
			}  
			case 11:
			if math.Hypot( (float64(x)-float64(ob.x)),(float64(y)-float64(ob.y)) ) < float64(ob.qua/2) {
				ob.SetzeAkt(false)
				SpieleSound("../../Sounds/Punkt.wav")
				return true, time.Now().UnixNano() - ob.GibErstellung()
			} else {
				return false, 0
			}
			/*
			case 16: 		// HEIDI
			if ob.x <= x && x < ob.x+85 	&& 	ob.y <= y && y < ob.y+88 {
				SpieleSound("../../Sounds/Roar.wav")
				return true, 0
			} else {
				return false, 0
			} 
			*/ 
			case 18:		// Kaffee 2
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 3 {
					ob.SetzeAkt(false)
					SpieleSound("../../Sounds/Trinken.wav")
					return true, 6e8
				}
				ob.SetzeTyp(4)
				SpieleSound("../../Sounds/GameOver.wav")
				return true, 0
			} else {
				return false, 0
			}  
			case 19:		// Pizza 2
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 1 {
					ob.SetzeAkt(false)
					SpieleSound("../../Sounds/Essen.wav")
					return true, 6e8
				}
				ob.SetzeTyp(6)
				SpieleSound("../../Sounds/GameOver.wav")
				return true, 0
			} else {
				return false, 0
			}  
			case 20:
			if ob.x*4/10 < x && x < ob.x*6/10 	&& 	ob.y*8/10 < y && y < ob.y*9/10 {
				SpieleNote("5A",0.1,0)
				return true, 0
			} else {
				return false, 0
			}
			
			default:
			if ob.x+ob.qua/10 < x && x < ob.x+ob.qua*9/10 	&& 	ob.y+ob.qua/10 < y && y < ob.y+ob.qua*9/10 {
				ob.SetzeAkt(false)
				SpieleNote("3A",0.1,0)
				return true, time.Now().UnixNano() - ob.GibErstellung()
			} else {
				return false, 0
			}
		}
	}
	return false, 0
}
