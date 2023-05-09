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
	aktiv bool				// Gibt an, ob das Objekt aktiv ist
	inhalt string			// enthält möglichen Text-Inhalt
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

func (ob *data) SetzeInhalt(inhalt string) {
	ob.inhalt = inhalt
}

func (ob *data) SetzeErstellung(erstellt int64) {
	ob.erstellt = erstellt
}

func (ob *data) GibErstellung() (int64) {
	return ob.erstellt
}

func (ob *data) Zeichnen() {
	if ob.aktiv {
		switch ob.typ {
// Fadenkreuz-Maus-Zeiger bei Koord x,y
			case 0:	
			Stiftfarbe(0,255,0)
			for i:=uint16(0);i<6;i++ {
				Kreis(ob.x, ob.y, 30-i)
			}
			Vollrechteck(ob.x-3, ob.y-42, 6, 30)
			Vollrechteck(ob.x-3, ob.y+12, 6, 30)
			Vollrechteck(ob.x-42, ob.y-3, 30, 6)
			Vollrechteck(ob.x+12, ob.y-3, 30, 6)

// PAUSE - Anzeige			
			case 1:		
			SetzeFont ("../../Schriftarten/Freshman.ttf", 220 )
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,205)
			Stiftfarbe(230,230,230)
			Vollellipse(600,350,550,200)
			Stiftfarbe(125,0,64)  
			SchreibeFont (205,254,"PAUSE")
			Stiftfarbe(255,0,127)  
			SchreibeFont (200,250,"PAUSE")							// Schreibe mittig Pause
			
// rotes Quadrat ab linker oberer Ecke			// DUMMY
			case 2:		
			Rechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)
			Stiftfarbe(120,160,200)
			Vollrechteck(ob.x,ob.y,ob.qua-1,ob.qua-1)

// Kaffee-Tasse ab linker oberer Ecke			
			case 3:			
			Stiftfarbe(0,0,0)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
			Stiftfarbe(100,100,255)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua*7/15-1)
			
			Stiftfarbe(0,0,0)												// schwarze Umrandung Außen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10, ob.qua*16/100)
			Stiftfarbe(224,224,224)											// weißer Außen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*7/10, ob.qua*15/40, ob.qua*3/20)
			
			Stiftfarbe(0,0,0)												// schwarze Umrandung Innen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*7/10, ob.qua*2/9, ob.qua*2/25)
			Stiftfarbe(155,152,152)											// grauer Innen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*7/10, ob.qua*2/9-ob.qua/60, ob.qua*7/100)
			
			Stiftfarbe(0,0,0)			// schwarze Umrandung Henkel
			Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*12/25,ob.qua*11/100)
			Stiftfarbe(231,62,243)			// pinke Füllung Henkel
			Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*12/25,ob.qua*9/100)
			Stiftfarbe(0,0,0)			// schwaze Innen-Füllung Henkel
			Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*12/25,ob.qua*7/100)
			Stiftfarbe(0,0,255)			// blaues Inneres
			Vollkreis(ob.x+ob.qua*83/100, ob.y+ob.qua*12/25,ob.qua/20)
			
			Stiftfarbe(0,0,0)			// schwarze Umrandung
			Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100,180,0)
			Stiftfarbe(231,62,243)		// pinke Tasse
			Vollkreissektor(ob.x+ob.qua*47/100, ob.y+ob.qua*38/100, ob.qua*35/100,180,0)

			Stiftfarbe(0,0,0)			// schwarze Umrandung Kaffee
			Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*37/100, ob.qua*37/100, ob.qua/10)
			Stiftfarbe(153,76,13)		// brauner Kaffee
			Vollellipse(ob.x+ob.qua*47/100, ob.y+ob.qua*19/50, ob.qua*17/50, ob.qua*7/100)
			
			Stiftfarbe(222,214,214)
			Vollellipse(ob.x+ob.qua*7/25, 		ob.y+ob.qua/5,		ob.qua/60,	ob.qua*2/25)
			Vollellipse(ob.x+ob.qua*29/100+1, 	ob.y+ob.qua*6/25,	ob.qua/60,	ob.qua*2/25)
			Vollellipse(ob.x+ob.qua*7/25, 		ob.y+ob.qua*7/25,	ob.qua/60,	ob.qua*2/25) 

			Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*4/25,	ob.qua/60,	ob.qua*9/100)
			Vollellipse(ob.x+ob.qua*46/100-1, 	ob.y+ob.qua*6/25,	ob.qua/60,	ob.qua*9/100)
			Vollellipse(ob.x+ob.qua*47/100, 	ob.y+ob.qua*8/25,	ob.qua/60,	ob.qua*9/100) 

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
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua/10, ob.x+ob.qua/10,		ob.y+ob.qua/5, 		ob.x+ob.qua*9/10,	ob.y+ob.qua*9/10)
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua/10, ob.x+ob.qua*9/10,	ob.y+ob.qua*9/10, 	ob.x+ob.qua*9/10,	ob.y+ob.qua*4/5)
			
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua*9/10, ob.x+ob.qua/10,	ob.y+ob.qua*4/5, 	ob.x+ob.qua*9/10,	ob.y+ob.qua/10)
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,	ob.y+ob.qua/10, 	ob.x+ob.qua*9/10,	ob.y+ob.qua/5)

// PIZZA
			case 5:																							
			Stiftfarbe(0,0,0)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
			Stiftfarbe(229,212,186)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua*7/15-1)								// Kreis-Hintergrund
			Stiftfarbe(225,141,0)
			Vollkreissektor(ob.x+ob.qua*19/40, ob.y+ob.qua*13/25, ob.qua*2/5, 60, 360)				// brauner Pizzaboden
			Vollkreissektor(ob.x+ob.qua*11/20, ob.y+ob.qua*7/15, ob.qua*2/5, 0, 60)
			Stiftfarbe(255,65,59)
			Vollkreissektor(ob.x+ob.qua*19/40, ob.y+ob.qua*13/25, ob.qua*9/25, 60, 360)				// Tomatensauce
			Vollkreissektor(ob.x+ob.qua*11/20, ob.y+ob.qua*7/15, ob.qua*9/25, 0, 60)
			Stiftfarbe(255,215,33)
			Vollkreissektor(ob.x+ob.qua*19/40, ob.y+ob.qua*13/25, ob.qua*8/25, 60, 360)				// Käse
			Vollkreissektor(ob.x+ob.qua*11/20, ob.y+ob.qua*7/15, ob.qua*8/25, 0, 60)
			Stiftfarbe(169,8,0)
			Vollkreis(ob.x+ob.qua*10/20,ob.y+ob.qua*6/20,	ob.qua/20)								// Salami
			Vollkreis(ob.x+ob.qua*15/20,ob.y+ob.qua*6/20,	ob.qua/20)
			Vollkreis(ob.x+ob.qua*5/20,	ob.y+ob.qua*8/20,	ob.qua/20)
			Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*11/20,	ob.qua/20)
			Vollkreis(ob.x+ob.qua*11/20,ob.y+ob.qua*14/20,	ob.qua/20)
			Vollkreis(ob.x+ob.qua*7/20,	ob.y+ob.qua*15/20,	ob.qua/20)
			
			Vollkreissektor(ob.x+ob.qua*5/8,	ob.y+ob.qua*26/50,	ob.qua/20, 180, 360)
			Vollkreissektor(ob.x+ob.qua*7/10,	ob.y+ob.qua*7/15,	ob.qua/20, 0, 180)
			
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
			Vollkreis(ob.x+ob.qua*10/20,ob.y+ob.qua*6/20,	ob.qua/20)										// Salami
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
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua/10, ob.x+ob.qua/10,		ob.y+ob.qua/5, 		ob.x+ob.qua*9/10,	ob.y+ob.qua*9/10)
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua/10, ob.x+ob.qua*9/10,	ob.y+ob.qua*9/10, 	ob.x+ob.qua*9/10,	ob.y+ob.qua*4/5)
			
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua*9/10, ob.x+ob.qua/10,	ob.y+ob.qua*4/5, 	ob.x+ob.qua*9/10,	ob.y+ob.qua/10)
			Volldreieck(ob.x+ob.qua/10,	ob.y+ob.qua*9/10, ob.x+ob.qua*9/10,	ob.y+ob.qua/10, 	ob.x+ob.qua*9/10,	ob.y+ob.qua/5)
			
// LEVEL 1 - Anzeige			
			case 7:	
			SetzeFont ("../../Schriftarten/Freshman.ttf", 200)
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,200)
			Stiftfarbe(68,225,255)
			Vollellipse(600,350,545,195)
			Stiftfarbe(127,25,64)  
			SchreibeFont (195,254,"LEVEL 1")
			Stiftfarbe(255,0,127)  
			SchreibeFont (190,250,"LEVEL 1")					// Schreibe mittig Level 1

// LEVEL 2 - Anzeige			
			case 8:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 200)
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,200)
			Stiftfarbe(68,225,255)
			Vollellipse(600,350,545,195)
			Stiftfarbe(127,25,64)  
			SchreibeFont (175,254,"LEVEL 2")
			Stiftfarbe(255,0,127)  
			SchreibeFont (170,250,"LEVEL 2")					// Schreibe mittig Level 2
			
// LEVEL 3 - Anzeige			
			case 9:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 200)
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,200)
			Stiftfarbe(68,225,255)
			Vollellipse(600,350,545,195)
			Stiftfarbe(127,25,64)  
			SchreibeFont (175,254,"LEVEL 3")
			Stiftfarbe(255,0,127)  
			SchreibeFont (170,250,"LEVEL 3")					// Schreibe mittig Level 3
			
// LEVEL 4 - Anzeige			
			case 10:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 200)
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,200)
			Stiftfarbe(68,225,255)
			Vollellipse(600,350,545,195)
			Stiftfarbe(127,25,64)  
			SchreibeFont (175,254,"LEVEL 4")
			Stiftfarbe(255,0,127)  
			SchreibeFont (170,250,"LEVEL 4")					// Schreibe mittig Level 4

// LEVEL 5 - Anzeige			
			case 11:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 200)
			Stiftfarbe(125,0,64)
			Vollellipse(600,350,560,200)
			Stiftfarbe(68,225,255)
			Vollellipse(600,350,545,195)
			Stiftfarbe(127,25,64)  
			SchreibeFont (175,254,"LEVEL 5")
			Stiftfarbe(255,0,127)  
			SchreibeFont (170,250,"LEVEL 5")					// Schreibe mittig Level 5
			
// ZIELSCHEIBE			
			case 12:																														
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
			case 13:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 390 )
			Stiftfarbe(186,66,23)
			SchreibeFont (516,215,"3")
			Stiftfarbe(68,215,210)
			SchreibeFont (492,206,"3")

// Countdown - 2			
			case 14:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 390 )
			Stiftfarbe(186,66,23)
			SchreibeFont (516,215,"2")
			Stiftfarbe(68,215,210)
			SchreibeFont (492,206,"2")

// Countdown - 1			
			case 15:																
			SetzeFont ("../../Schriftarten/Freshman.ttf", 390 )
			Stiftfarbe(186,66,23)
			SchreibeFont (528,215,"1")
			Stiftfarbe(68,215,210)
			SchreibeFont (504,206,"1")

// Heidi			
			case 16:																														
			LadeBild (ob.x,ob.y, "../../Bilder/Heidi-2.bmp")

// StEPS-Logo			
			case 17:																														
			LadeBild (ob.x,ob.y, "../../Bilder/StEPS-Logo-2.bmp")

// Kaffee-Tasse 2 ab linker oberer Ecke			
			case 18:			
			Stiftfarbe(0,0,0)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-1)
			Stiftfarbe(100,100,255)
			Vollkreis(ob.x+ob.qua/2-1, ob.y+ob.qua/2-1, ob.qua/2-ob.qua/30-1)
			
			Stiftfarbe(0,0,0)														// schwarze Umrandung Außen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10, ob.qua*16/100)
			Stiftfarbe(224,224,224)													// weißer Außen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*4/10-ob.qua/40, ob.qua*15/100)
			
			Stiftfarbe(0,0,0)														// schwarze Umrandung Innen-Unterteller
			Vollellipse(ob.x+ob.qua*49/100, ob.y+ob.qua*70/100, ob.qua*2/9, ob.qua*8/100)
			Stiftfarbe(155,152,152)													// grauer Innen-Unterteller
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
			Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*2/5, 60, 360)			// brauner Pizzaboden
			Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*2/5, 0, 60)
			Stiftfarbe(255,65,59)
			Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*18/50, 60, 360)		// Tomatensauce
			Vollkreissektor(ob.x+ob.qua/2+ob.qua/20, ob.y+ob.qua/2-ob.qua/30, ob.qua*18/50, 0, 60)
			Stiftfarbe(255,215,33)
			Vollkreissektor(ob.x+ob.qua/2-ob.qua/40, ob.y+ob.qua/2+ob.qua/50, ob.qua*16/50, 60, 360)		// Käse
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
			Stiftfarbe(153,0,153)
			Vollrechteck(480,570,240,80)
			Stiftfarbe(255,0,255)
			Vollrechteck(490,577,220,66)
			SetzeFont ("../../Schriftarten/Freshman.ttf", 56 )
			Stiftfarbe(124,212,255)
			SchreibeFont (554,584,"O K")

// PASST - Objekt			
			case 22:															
			
			Stiftfarbe(173,214,50)
			Vollrechteck(220,460,250,80)
			/*Stiftfarbe(255,0,255)
			Vollrechteck(490,577,220,66)*/
			SetzeFont ("../../Schriftarten/Ubuntu-B.ttf", 50 )
			Stiftfarbe(0,0,0)
			SchreibeFont (252,472,"P a s s t")
			Stiftfarbe(124,212,255)
			SchreibeFont (250,470,"P a s s t")

// PASST-NICHT - Objekt			
			case 23:															
			
			Stiftfarbe(173,214,50)
			Vollrechteck(540,460,450,80)
			/*Stiftfarbe(255,0,255)
			Vollrechteck(490,577,220,66)*/
			SetzeFont ("../../Schriftarten/Ubuntu-B.ttf", 50 )
			Stiftfarbe(0,0,0)
			SchreibeFont (572,472,"P a s s t   n i c h t")
			Stiftfarbe(124,212,255)
			SchreibeFont (570,470,"P a s s t   n i c h t")

// Titel-Objekt			
			case 24:															
			Transparenz(50)
			Stiftfarbe(153,204,0)
			Vollrechteck(150,50,900,80)
			SetzeFont ("../../Schriftarten/Ubuntu-B.ttf", 70 )
			Stiftfarbe(65,96,140)
			SchreibeFont (ob.x,ob.y,ob.inhalt)
			Transparenz(0)
		
// FebWeb - Normal			
			case 25:															
			Transparenz(150)
			LadeBild (ob.x-124,ob.y-137, "../../Bilder/FebWebK.bmp")
			Transparenz(0)
						
// FebWeb - JA			
			case 26:															
			Transparenz(80)
			LadeBild (ob.x-124,ob.y-137, "../../Bilder/FebWebJ.bmp")
			Transparenz(0)
			
// FebWeb - NEIN
			case 27:															
			Transparenz(80)
			LadeBild (ob.x-124,ob.y-137, "../../Bilder/FebWebN.bmp")
			Transparenz(0)
			
// Spielkarte zugedeckt - ab linker oberer Ecke		
			case 31:															
			Stiftfarbe(153,0,153)
			Vollrechteck(ob.x,ob.y,225,150)
			Stiftfarbe(255,0,255)
			Vollrechteck(ob.x+5,ob.y+5,215,140)
			
// Spielkarte aufgedeckt - ab linker oberer Ecke			
			case 32:															
			Stiftfarbe(153,0,153)
			Vollrechteck(ob.x,ob.y,225,150)
			Stiftfarbe(210,250,210)
			Vollrechteck(ob.x+5,ob.y+5,215,140)
			SetzeFont ("../../Schriftarten/Ubuntu-B.ttf", 33 )
			Stiftfarbe(160,100,220)
			SchreibeFont (ob.x+9,ob.y+55,ob.inhalt)

// Spielkarte fertig - ab linker oberer Ecke			
			case 33:															
			//Stiftfarbe(203,0,203)
			
			switch ob.erstellt {
				case 1: Stiftfarbe(160,0,0)
				case 2: Stiftfarbe(160,160,0)
				case 3: Stiftfarbe(0,160,0)
				case 4: Stiftfarbe(0,160,160)
				case 5: Stiftfarbe(0,0,160)
				case 6: Stiftfarbe(160,0,160)
				default:
			}
			Vollrechteck(ob.x,ob.y,225,150)
			//Vollrechteck(ob.x+5,ob.y+5,215,140)
			Stiftfarbe(240,240,240)
			Vollrechteck(ob.x+10,ob.y+10,205,130)
			SetzeFont ("../../Schriftarten/Ubuntu-B.ttf", 33 )
			Stiftfarbe(220,152,255)
			SchreibeFont (ob.x+9,ob.y+55,ob.inhalt)
		}
	}
}

func (ob *data) Getroffen(x,y uint16, opt uint8) (bool,int64) {														// Checkt, ob Hit-Box getroffen
	if ob.aktiv {		
		switch ob.typ {
			case 2:
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				ob.aktiv = false
				SpieleNote("4A",0.1,0)
				return true, time.Now().UnixNano() - ob.erstellt
			} else {
				return false, 0
			}
			case 3:		// Kaffee
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 3 {
					ob.aktiv = false
					SpieleSound("../../Sounds/Trinken.wav")
					return true, time.Now().UnixNano() - ob.erstellt
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
					ob.aktiv = false
					SpieleSound("../../Sounds/Essen.wav")
					return true, time.Now().UnixNano() - ob.erstellt
				}
				ob.SetzeTyp(6)
				SpieleSound("../../Sounds/GameOver.wav")
				return true, 0
			} else {
				return false, 0
			}  
			case 12:
			if math.Hypot( (float64(x)-float64(ob.x)),(float64(y)-float64(ob.y)) ) < float64(ob.qua/2) {
				ob.aktiv = false
				SpieleSound("../../Sounds/Punkt.wav")
				return true, time.Now().UnixNano() - ob.erstellt
			} else {
				return false, 0
			} 
			case 18:		// Kaffee 2
			if ob.x <= x && x < ob.x+ob.qua 	&& 	ob.y <= y && y < ob.y+ob.qua {
				if opt == 3 {
					ob.aktiv = false
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
					ob.aktiv = false
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
			if 480 < x && x < 720 	&& 	570 < y && y < 650 {
				SpieleNote("5A",0.1,0)
				return true, 0
			} else {
				return false, 0
			}
			case 22:		// passt
			if 220 < x && x < 470 && 460 < y && y < 540 {
				return true, 1
			} else {
				return false, 0
			}
			case 23:		// passt nicht
			if 540 < x && x < 990 && 460 < y && y < 540 {
				return true, 0
			} else {
				return false, 0
			}
			case 24:
			case 31:
			if ob.x < x && x < ob.x + ob.qua*3/2 && ob.y < y && y < ob.y+ob.qua {
				return true, ob.erstellt
			} else {
				return false, 0
			}
			case 32:
			case 33:
			return false, 0
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
