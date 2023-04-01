package main
// Autor: B. Schneider
// Datum: 28.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import ( . "gfx"
			"time"
			"fmt"
			"../../Klassen/objekte"
			)

const breite = 800 		// von Gott vorgegeben
const hoehe  = 600  	// von Gott vorgegeben
	
func main () {
	var taste uint8
	var status int8
	var mausX, mausY uint16
	var verz bool								// für die Verzögerung in der Klangwiedergabe
	var attack,decay,sustain,release uint8			// Werte für die Hüllkurve
	
	obj := make([]objekte.Objekt,0,200)			// Array für die Objekte der Welt
	
    maus 		:= objekte.New(0,0,			30		,0)			// Erstellt das Objekt MAUSZEIGER mit Größe (40)
	
	
	Fenster (breite,hoehe)									// Öffnet GFXW-Fenster
	Fenstertitel("Spiel(e) mit Klängen!")								// Gibt Fenster-Titel 
	
	// Objekte werden in der Welt platziert
	erstelleObjekte(&obj, maus)
	
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	//go view_komponente(&obj, maus)
	
	
	// Nebenläufig wird die Kontroll-Komponente 1 für die Maus gestartet.
	go maussteuerung (&obj, maus, &taste, &status, &mausX, &mausY)
	
	// Nebenläufig wird Soundsteuerung gestartet:
	go soundsteuerung(&obj, &taste, &status, &mausX, &mausY, &verz)
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	// In anderen Sprachen würde man hier alle Events abfragen, also
	// ob es eine Tasteneingabe oder Mauseingabe war. Dann gäbe es
	// keine Control-Komponente 1.
	
A:	for {
		taste,gedrueckt,_:= TastaturLesen1 ()
		if gedrueckt == 1 {
			/*
			fmt.Println(taste)
			fmt.Println( GibHuellkurve () )
			* SetzeKlangparameter(rate uint32, aufloesung,kanaele,signal uint8, p float64): Standard: 44100 2 2 1 0.375
			* GibHuellkurve() (float64,float64,float64,float64): 0.002 0.75 0 0.006
			*/
			fmt.Println( GibHuellkurve() )
			switch taste {
				case 'q': 													// mit 'q' wird das Programm beendet!
				break A
				case 273: 												// Oben-Taste
				verz = false
				case 274:												// Unten-Taste
				verz = true
				case 'w':												// s-Taste für die WELLENFORM
				_,_,_,welle,_ := GibKlangparameter()
				switch welle {
					case 1:
					SetzeKlangparameter(44100,2,2,2,0.375)
					Stiftfarbe(255,255,255)
					Vollrechteck(5,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(10,10,"Welle: DREIECK")
					fmt.Println("WELLENform geändert auf DREIECK")
					case 2:
					SetzeKlangparameter(44100,2,2,3,0.375)
					Stiftfarbe(255,255,255)
					Vollrechteck(5,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(10,10,"Welle: SAEGEZAHN")
					fmt.Println("WELLENform geändert auf SÄGEZAHN")
					case 3:
					SetzeKlangparameter(44100,2,2,0,0.375)
					Stiftfarbe(255,255,255)
					Vollrechteck(5,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(10,10,"Welle: SINUS")
					fmt.Println("WELLENform geändert auf SINUS")
					case 0:
					SetzeKlangparameter(44100,2,2,1,0.375)
					Stiftfarbe(255,255,255)
					Vollrechteck(5,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(10,10,"Welle: RECHTECK")
					fmt.Println("WELLENform geändert auf RECHTECK")
				}
				case 'a':
				_,dec,sus,rel := GibHuellkurve()
				switch attack {						// GibHuellkurve() (float64,float64,float64,float64): Standard-> 0.002 0.75 0 0.006
					case 0:
					SetzeHuellkurve(0.01,dec,sus,rel)
					attack = 1
					Stiftfarbe(255,255,255)
					Vollrechteck(200,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(200,10,"Attack: 0,01")
					fmt.Println("ATTACK geändert auf 0,01")
					case 1:
					SetzeHuellkurve(0.1,dec,sus,rel)
					attack = 2
					Stiftfarbe(255,255,255)
					Vollrechteck(200,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(200,10,"Attack: 0,1")
					fmt.Println("ATTACK geändert auf 0,1")
					case 2:
					SetzeHuellkurve(0.5,dec,sus,rel)
					attack = 3
					Stiftfarbe(255,255,255)
					Vollrechteck(200,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(200,10,"Attack: 0,5")
					fmt.Println("ATTACK geändert auf 0,5")
					case 3:
					SetzeHuellkurve(1,dec,sus,rel)
					attack = 4
					Stiftfarbe(255,255,255)
					Vollrechteck(200,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(200,10,"Attack: 1")
					fmt.Println("ATTACK geändert auf 1")
					case 4:
					SetzeHuellkurve(0.002,dec,sus,rel)
					attack = 0
					Stiftfarbe(255,255,255)
					Vollrechteck(200,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(200,10,"Attack: 0,002")
					fmt.Println("ATTACK geändert auf 0,002")
				}
				case 'd':							// GibHuellkurve() (float64,float64,float64,float64): Standard-> 0.002 0.75 0 0.006
				att,_,sus,rel := GibHuellkurve()
				switch decay {						
					case 0:
					SetzeHuellkurve(att,1,sus,rel)
					decay = 1
					Stiftfarbe(255,255,255)
					Vollrechteck(350,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(350,10,"Decay: 1")
					fmt.Println("DECAY geändert auf 1")
					case 1:
					SetzeHuellkurve(att,2.5,sus,rel)
					decay = 2
					Stiftfarbe(255,255,255)
					Vollrechteck(350,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(350,10,"Decay: 2,5")
					fmt.Println("DECAY geändert auf 2,5")
					case 2:
					SetzeHuellkurve(att,5,sus,rel)
					decay = 3
					Stiftfarbe(255,255,255)
					Vollrechteck(350,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(350,10,"Decay: 5")
					fmt.Println("DECAY geändert auf 5")
					case 3:
					SetzeHuellkurve(att,0.2,sus,rel)
					decay = 4
					Stiftfarbe(255,255,255)
					Vollrechteck(350,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(350,10,"Decay: 0,2")
					fmt.Println("DECAY geändert auf 0,2")
					case 4:
					SetzeHuellkurve(att,0.75,sus,rel)
					decay = 0
					Stiftfarbe(255,255,255)
					Vollrechteck(350,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(350,10,"Decay: 0,75")
					fmt.Println("DECAY geändert auf 0,75")
				}
				case 's':							// GibHuellkurve() (float64,float64,float64,float64): Standard-> 0.002 0.75 0 0.006
				att,dec,_,rel := GibHuellkurve()
				switch sustain {						
					case 0:
					SetzeHuellkurve(att,dec,0.25,rel)
					sustain = 1
					Stiftfarbe(255,255,255)
					Vollrechteck(500,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(500,10,"Sutain: 0,25")
					fmt.Println("SUSTAIN geändert auf 0,25")
					case 1:
					SetzeHuellkurve(att,dec,0.5,rel)
					sustain = 2
					Stiftfarbe(255,255,255)
					Vollrechteck(500,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(500,10,"Sutain: 0,5")
					fmt.Println("SUSTAIN geändert auf 0,5")
					case 2:
					SetzeHuellkurve(att,dec,0.7,rel)
					sustain = 3
					Stiftfarbe(255,255,255)
					Vollrechteck(500,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(500,10,"Sutain: 0,75")
					fmt.Println("SUSTAIN geändert auf 0,75")
					case 3:
					SetzeHuellkurve(att,dec,1,rel)
					sustain = 4
					Stiftfarbe(255,255,255)
					Vollrechteck(500,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(500,10,"Sutain: 1")
					fmt.Println("SUSTAIN geändert auf 1")
					case 4:
					SetzeHuellkurve(att,dec,0,rel)
					sustain = 0
					Stiftfarbe(255,255,255)
					Vollrechteck(500,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(500,10,"Sutain: 0")
					fmt.Println("SUSTAIN geändert auf 0")
				}
				case 'r':							// GibHuellkurve() (float64,float64,float64,float64): Standard-> 0.002 0.75 0 0.006
				att,dec,sus,_ := GibHuellkurve()
				switch release {						
					case 0:
					SetzeHuellkurve(att,dec,sus,0.1)
					release = 1
					Stiftfarbe(255,255,255)
					Vollrechteck(650,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(650,10,"Release: 0,1")
					fmt.Println("RELEASE geändert auf 0,1")
					case 1:
					SetzeHuellkurve(att,dec,sus,1)
					release = 2
					Stiftfarbe(255,255,255)
					Vollrechteck(650,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(650,10,"Release: 1")
					fmt.Println("RELEASE geändert auf 1")
					case 2:
					SetzeHuellkurve(att,dec,sus,2.5)
					release = 3
					Stiftfarbe(255,255,255)
					Vollrechteck(650,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(650,10,"Release: 2,5")
					fmt.Println("RELEASE geändert auf 2,5")
					case 3:
					SetzeHuellkurve(att,dec,sus,5)
					release = 4
					Stiftfarbe(255,255,255)
					Vollrechteck(650,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(650,10,"Release: 5")
					fmt.Println("RELEASE geändert auf 5")
					case 4:
					SetzeHuellkurve(att,dec,sus,0.006)
					release = 0
					Stiftfarbe(255,255,255)
					Vollrechteck(650,5,150,19)
					Stiftfarbe(0,0,0)
					Schreibe(650,10,"Release: 0,006")
					fmt.Println("RELEASE geändert auf 0,006")
				}
			}
		} 
	}
}

func erstelleObjekte(obj *[]objekte.Objekt, maus objekte.Objekt) {		// füllt Objekte ins Array
	
	for i:=uint16(0);i<=800;i+=50 {
		for j:=uint16(25)-(i%100)/2;j<=600;j+=50 {
			new := objekte.New(i,j,	90	,2)
			new.SetzeFarbe(uint8(i/4),uint8(j/3),uint8( (i+j)/6 ))
			*obj = append(*obj, new )
		}
	}
	ObjAktualisieren(obj)
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	Transparenz(150)
	for _,ob := range *obj { 					// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Transparenz(0)
	Stiftfarbe(255,255,255)
	Vollrechteck(0,0,800,24)
	Stiftfarbe(0,0,0)
	Schreibe(10,10,"Welle: RECHTECK")
	Schreibe(200,10,"Attack: 0,002")
	Schreibe(350,10,"Decay: 0,75")
	Schreibe(500,10,"Sutain: 0")
	Schreibe(650,10,"Release: 0,006")
	Archivieren()								// Speichert das Hintergrund-Bild
}


// Es folgt die VIEW-Komponente --- Kein Bestandteil der Welt, also unabhängig---------
// Einzig diese Funktion verwendet gfxw-Befehle!!
func view_komponente (obj *[]objekte.Objekt, maus objekte.Objekt) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	for { //Endlos ...
		UpdateAus () 							// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()									// Cleart vollständigen Screen

		Restaurieren(0,0,breite,hoehe)
		//maus.SetzeKoordinaten(<-mauX,<-mauY)		// Aktualisiert Maus-Koordinaten
		// maus.SetzeKoordinaten(*mausX,*mausY)
		maus.Zeichnen()								// Zeichnet Maus
			
		
		if time.Now().UnixNano() - t1 < 1000000000 { //noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() // neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}		//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}		//Frame-Rate :-)
		}
		
		UpdateAn () 									// Nun wird der gezeichnete Frame sichtbar gemacht!
		time.Sleep(time.Duration(verzögerung * 1e5)) 	// Immer ca. 100 FPS !!
	}
}


// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus objekte.Objekt, taste *uint8, status *int8, mausX, mausY *uint16) {
	for {
		*taste, *status, *mausX, *mausY = MausLesen1()
		// fmt.Println (*taste, *status, *mausX,*mausY)
		maus.SetzeKoordinaten(*mausX,*mausY)					// Aktualisiert Maus-Koordinaten
				
	}
}

func soundsteuerung(obj *[]objekte.Objekt, taste *uint8, status *int8, mausX, mausY *uint16, verz *bool) {
	for {
		if *taste==1 && *status==1 { 							//LINKE Maustaste gerade gedrückt
			for *status!=-1 {
				for _,ob := range *obj { 							// Überpüft Objekte
					if ob.Getroffen(*mausX,*mausY) {
						b,h := ob.GibKoordinaten()
						switch b {
							case 0,50:
							SpieleNote(fmt.Sprint((h+80)/80)+"C",0.1,0)
							case 100,150:
							SpieleNote(fmt.Sprint((h+80)/80)+"D",0.1,0)
							case 200,250:
							SpieleNote(fmt.Sprint((h+80)/80)+"E",0.1,0)
							case 300,350:
							SpieleNote(fmt.Sprint((h+80)/80)+"F",0.1,0)
							case 400,450:
							SpieleNote(fmt.Sprint((h+80)/80)+"G",0.1,0)
							case 500,550:
							SpieleNote(fmt.Sprint((h+80)/80)+"A",0.1,0)
							case 600,650:
							SpieleNote(fmt.Sprint((h+80)/80)+"H",0.1,0)
							case 700,750:
							SpieleNote(fmt.Sprint((h+160)/80)+"C",0.1,0)
						 }
					}
				}
				if *verz {
					time.Sleep(time.Duration(3e8) )
				} else {
					time.Sleep(time.Duration(6e7) )
				}
			}
		} else if *taste==3 && *status==1 { 							//RECHTE Maustaste gerade gedrückt
			for *status!=-1 {
				for _,ob := range *obj { 							// Überpüft Objekte
					if ob.Getroffen(*mausX,*mausY) {
						b,h := ob.GibKoordinaten()
						switch b {
							case 0,50:
							SpieleNote(fmt.Sprint((h+80)/80)+"C",0.1,0)
							case 100,150:
							SpieleNote(fmt.Sprint((h+80)/80)+"D",0.1,0)
							case 200,250:
							SpieleNote(fmt.Sprint((h+80)/80)+"E",0.1,0)
							case 300,350:
							SpieleNote(fmt.Sprint((h+80)/80)+"G",0.1,0)
							case 400,450:
							SpieleNote(fmt.Sprint((h+80)/80)+"A",0.1,0)
							case 500,550:
							SpieleNote(fmt.Sprint((h+80)/80)+"H",0.1,0)
							case 600,650:
							SpieleNote(fmt.Sprint((h+80)/80)+"D",0.1,0)
							case 700,750:
							SpieleNote(fmt.Sprint((h+160)/80)+"F",0.1,0)
						 }
					}
				}
				if *verz {
					time.Sleep(time.Duration(3e8) )
				} else {
					time.Sleep(time.Duration(5e7) )
				}
			}
		}
	}
}
