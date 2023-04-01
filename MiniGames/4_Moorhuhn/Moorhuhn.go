package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import ( . "gfx" ;  "time" ; "fmt"; "../../Klassen/objekte"; "../../Klassen/raeume" )

const breite = 800 		// von Gott vorgegeben
const hoehe  = 600  	// von Gott vorgegeben
	
func main () {
	var punkte uint16 = 10					// Spiel-Punktzahl
	var pause bool = false
	var akt	bool = true						// Prüft, ob Grafik aktualisiert werden muss
	obj := make([]objekte.Objekt,0)			// Array für die Objekte der Welt
	
    
	pauseObjekt := objekte.New(breite,hoehe,hoehe/4	,1)			// Erstellt das Objekt PAUSE 
	maus 		:= objekte.New(0,0,			30		,0)			// Erstellt das Objekt MAUSZEIGER mit Größe (40)
	
	/*
	mauX := make(chan uint16)									// Erstellt Channel zur Maussteuerung
	mauY := make(chan uint16)
	*/
	
	Fenster (breite,hoehe)									// Öffnet GFXW-Fenster
	Fenstertitel("StEPS-Wars")								// Gibt Fenster-Titel 
	// SetzeFont ("collegeb.ttf", hoehe/4)					// Setzt Schriftart
	// SetzeFont ("Prisma.ttf", hoehe/4)
	SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/20 ) 
	
	// Objekte werden nach und nach in der Welt platziert
	go erstelleObjekte(&obj, &pause, &akt)
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, pauseObjekt, &pause, &akt, &punkte)
	
	
	// Nebenläufig wird die Kontroll-Komponente 1 für die Maus gestartet.
	go maussteuerung (&obj, maus, &pause, &akt, &punkte)
	
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	// In anderen Sprachen würde man hier alle Events abfragen, also
	// ob es eine Tasteneingabe oder Mauseingabe war. Dann gäbe es
	// keine Control-Komponente 1.
	
A:	for {
		taste,gedrueckt,_:= TastaturLesen1 ()
		if gedrueckt == 1 {
			switch taste {
				case 'q': 													// mit 'q' wird das Programm beendet!
				break A
				case 'p': 
				pause = !pause												// Pause-Modus !!
				if pause {
					pause = false
					
				} else {
					pause = true
					SetzeFont ("Freshman.ttf", hoehe/4 )
				}												
				pause = !pause
				case ' ':
				case 'y':
				case 'x':
				
			}
		}
	}
	// Mit dem Ende des Hauptprogramms (Control-Komponente 2)
	// werden auch die anderen Komponenten, die hier gestartet wurden,
	// beendet!
	

}

func erstelleObjekte(obj *[]objekte.Objekt, pause,akt *bool) {		// füllt Objekte ins Array
	
	time.Sleep( time.Duration(2e9) )
	
	*obj = append(*obj, objekte.New(0,150,		350	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(200,0,		300	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(400,300,	250	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(600,200,	200	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(400,100,	150	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(650,80,	100	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(600,50,	50	,3) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	for *pause { time.Sleep( time.Duration(2e8) ) }
	*obj = append(*obj, objekte.New(500,40,	50	,5) )
	*akt = true
	time.Sleep( time.Duration(2e9) )
	
	test := objekte.New(0,350,		100	,2)
	test.SetzeFarbe(255,0,0)
	*obj = append(*obj, test )
	*akt = true
	
	for i:=uint16(0);i<300;i++ {
		test.SetzeKoordinaten(3*i,3*i)
		*akt=true
		time.Sleep( time.Duration(1e8) )
	}
	
	
	
}

// Es folgt die VIEW-Komponente --- Kein Bestandteil der Welt, also unabhängig---------
// Einzig diese Funktion verwendet gfxw-Befehle!!
func view_komponente (obj *[]objekte.Objekt, maus,pauseObjekt objekte.Objekt, pause,akt *bool, punkte *uint16) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	for { //Endlos ...
		UpdateAus () 							// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()									// Cleart vollständigen Screen
		
		if *akt { 
			ObjAktualisieren(obj)
			*akt = false
		} else {
			Restaurieren(0,0,breite,hoehe)		// Restauriert das alte Hintergrundbild
		}
	
		
		//maus.SetzeKoordinaten(<-mauX,<-mauY)		// Aktualisiert Maus-Koordinaten
		// maus.SetzeKoordinaten(*mausX,*mausY)
		maus.Zeichnen()								// Zeichnet Maus
		
		SetzeFont ("../../Schriftarten/Freshman.ttf", hoehe/20 )
		Stiftfarbe(76,0,153)  
		SchreibeFont (breite*14/20,0,"Punkte : "+fmt.Sprint (*punkte))	// Schreibe rechts oben Punkte
		Stiftfarbe(100,10,155)
		Schreibe (0,0,"FPS:"+fmt.Sprint (anzahl))					// Schreibe links oben FPS
		if *pause { pauseObjekt.Zeichnen() }
			
		
		if time.Now().UnixNano() - t1 < 1000000000 { //noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() // neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}		//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}		//Frame-Rate :-)
		}
		
		UpdateAn () // Nun wird der gezeichnete Frame sichtbar gemacht!
		time.Sleep(time.Duration(verzögerung * 1e5)) // Immer ca. 100 FPS !!
		
	}
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	raeume.Moorhuhn(breite)						// Hintergrund des Moorhuhn-Raumes wird gezeichnet
	// raeume.Hauptflur(breite)
	
	for _,ob := range *obj { 					// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()							// Speichert das Hintergrund-Bild
}

// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus objekte.Objekt, pause,akt *bool, punkte *uint16) {
	/*var taste uint8
	var status int8 */
	for {
		taste, status, mausX, mausY := MausLesen1()
		
		maus.SetzeKoordinaten(mausX,mausY)				// Aktualisiert Maus-Koordinaten
		/*
		mauX <- mausX									// Channel-Weiterleitung der Maus-Koordinaten
		mauY <- mausY
		*/
		
		if *pause==false && taste==1 && status==1 { 				//LINKE Maustaste gerade gedrückt
			
			for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
				if ob.Getroffen(mausX,mausY) {
					switch ob.GibTyp() {
						case 2:
						SpieleNote("4C",0.1,0)
						case 3:
						ob.SetzeTyp(4)
						SpieleNote("4A",0.1,0)
						case 4:
						SpieleNote("4A",0.1,0)
						case 5:
						ob.SetzeTyp(6)
						SpieleNote("4E",0.1,0)
						case 6:
						SpieleNote("4E",0.1,0)
					}
					*punkte++
					*akt = true
				}
			}
			fmt.Println ("links: ", taste, status,mausX,mausY) 	// printet Koordinaten des Klicks
		}
		
		if *pause==false && taste == 3 && status == 1 { 			//RECHTE Maustaste gerade gedrückt
			
			for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
				if ob.Getroffen(mausX,mausY) {
					switch ob.GibTyp() {
						case 4:
						ob.SetzeTyp(3)
						case 6:
						ob.SetzeTyp(5)
					}
					SpieleSound("../../Sounds/GameOver.wav")
					*punkte--
					*akt = true
				}
			}
			fmt.Println ("rechts: ", taste, status,mausX,mausY)	// printet Koordinaten des Klicks
		}
	}
}

