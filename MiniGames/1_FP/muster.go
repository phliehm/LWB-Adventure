package muster
// Autor: B. Schneider
// Datum: 25.04.2023
// Zweck: SWP - Minigame Mustererkennung
//--------------------------------------------------------------------

import ( 	. "gfx"
			"time"
			"fmt"
			"sync"
			"../../Klassen/objekte"
			"../../Klassen/texte"
			"math/rand"
			)

	
func Muster() int16 {
	var mutex sync.Mutex					// erstellt Mutex
	var gedrueckt uint8
	var taste,tiefe uint16
	
	var eingabe string						// zur "Editor"-Eingabe
	var punkte int16 = 0					// Spiel-Punktzahl
	var diff int16 = 0						// Punkte-Veränderung
	var wert uint8
	
	var stop bool = false					// für OK-Objekt
	var signal bool = false					// falls Signal
	var ende bool = false					// falls Ende (Spiel soll beendet werden)
	var tastatur = false					// falls Tastatur-Eingabe aktiv ist
	
	var kanal chan bool = make(chan bool)
	
	var akt	bool = true						// Prüft, ob Grafik aktualisiert werden muss
	obj := make([]objekte.Objekt,0)			// Array für die Objekte der Welt
	
	
	random := rand.New( rand.NewSource( time.Now().UnixNano() ) )	// Initialisiere Random-Objekt mit der Systemzeit
	
	maus 		:= objekte.New(0, 0, 0, 25)			// Erstellt das Objekt MAUSZEIGER
	okayObjekt 	:= objekte.New(0, 0, 0, 20)			// OK-Objekt
	
	Fenstertitel("Muster, Muster, nichts als Muster - und dazwischen Muster")								// Gibt Fenster-Titel 
	
	// Das Hauptprogramm startet die View-Komponente als nebenläufigen Prozess!
	go view_komponente(&obj, maus, okayObjekt, &signal, &stop, &akt, &ende, &punkte, &diff, &mutex, &eingabe)
	
	// Objekte werden nach und nach in der Welt platziert
	go spielablauf(&obj, maus, random, &mutex, &akt, &tastatur, &stop, &signal, &eingabe, &wert, &punkte, kanal)
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung(&obj, maus, okayObjekt, &signal, &stop, &akt, &ende, &punkte, &diff, &wert, kanal)
	
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	
	
	
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 28 )
	
	
A:	for {
		taste, gedrueckt, tiefe = TastaturLesen1()
		
		fmt.Println (taste,gedrueckt,tiefe)
		if tastatur {
			if gedrueckt == 1  { // Beim Drücken der Taste, nicht beim Loslassen!
				switch {
					case taste == 27:  									// ESC-Taste
					break A
					case taste==13 || taste==271:  						// Enter-Taste(n)
					signal = true
					case taste == 32:  									// Leer-Taste
					eingabe += " "
					case taste ==  8:  									//Backspace-Taste
					if eingabe != "" {
						eingabe = eingabe [:len(eingabe)-1]
					}
					case taste >= 48 && taste < 58 && tiefe == 0:  		// Zahlen
					eingabe += string(taste)
					case taste == 44:
					eingabe += ","
					case taste == 46:
					eingabe += "."
					case taste == 55 && tiefe > 0:  					// 7
					eingabe += "["
					case taste == 56 && tiefe > 0:  					// 8
					eingabe += "("
					case taste == 57 && tiefe > 0:    					// 9		
					eingabe += ")"
					case taste == 48 && tiefe > 0:  		  			// 0
					eingabe += "]"
					case taste == 46 && tiefe > 0:  		
					eingabe += ":"
					case taste == 49 && tiefe > 0:  		
					eingabe += ":"
					case taste == 50 && tiefe > 0:  		
					eingabe += "\""
					case taste == 51 && tiefe > 0:  		
					eingabe += "'"
					case taste == 92 && tiefe > 0:  		
					eingabe += "'"
					case taste >= 97 && taste < 123 && tiefe == 0:  	// Kleinbuchstaben
					eingabe += string(taste)
					case taste >= 97 && taste < 123 && tiefe > 0:		// Großbuchstaben
					eingabe += string(taste-32)
					default:
				}
			}
		} else {
			if gedrueckt == 1  { // Beim Drücken der Taste, nicht beim Loslassen!
				switch {
					case taste == 27:  									// ESC-Taste
					break A
					case taste == 13:  									// Enter-Taste
					case taste == 271: 									// 2. Enter-Taste
					case taste == 32:  									// Leer-Taste
				}
			}
		}
	}
	fmt.Println("Vielen Dank für's Spielen!")
	time.Sleep( time.Duration(2e8) )
	
	return punkte
}

func spielablauf(obj *[]objekte.Objekt, maus objekte.Objekt, random *rand.Rand, mutex *sync.Mutex, akt, tastatur, stop, signal *bool, 
			eingabe *string, wert *uint8, punkte *int16, kanal chan bool) {
	var neuerZustand bool
	
	
	zwischentext(&texte.MusterEinl, mutex, stop)		// Einleitungs-Text
	
	//musterSpiel(obj, maus, akt, signal, tastatur, random, mutex, eingabe, wert, punkte)
	
	//time.Sleep( time.Duration(3e9) )
	
	*obj = make([]objekte.Objekt,0)
	*akt = true
	
	memorySpiel(obj, akt, 2, random)					// auf Level 1
	
	neuerZustand = <- kanal
	
	memorySpiel(obj, akt, 1, random)					// auf Level 2
	
	neuerZustand = <- kanal
	
	fmt.Println(neuerZustand)
	//for !*signal { time.Sleep( time.Duration(2e9) ) }
	
}

func zwischentext(textArr *[]string, mutex *sync.Mutex, stop *bool) {
	mutex.Lock()
	LadeBild (0,0, "./Bilder/Funktionale.bmp")		// Hintergrund des Muster-Raumes wird gezeichnet
	Transparenz(120)
	Stiftfarbe(76,0,153)														
	Vollrechteck(100,50,1000,600)
	Transparenz(0)
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 50 )
	Stiftfarbe(124,212,255)
	for ind,str := range *textArr {
		SchreibeFont (210, uint16(70+ind*55) ,str )
	}
	Archivieren()
	mutex.Unlock()
	
	*stop = true
	for *stop { time.Sleep( time.Duration(1e8) ) }
}

func musterSpiel(obj *[]objekte.Objekt, maus objekte.Objekt, akt, signal, tastatur *bool, rand *rand.Rand, 
							mutex *sync.Mutex, eingabe *string, wert *uint8, punkte *int16) { // gibt Muster zur Abfrage
	
	var zufallSpalte int
	var versuch uint8									// zählt, ob es schon einen Lösungsversuch gab
	var musterwahl int									
	var musterspeicher []int						// speichert genutzte Muster ab
	
	titel := objekte.New(230,50,150,24)
	titel.SetzeInhalt("MUSTER-ERKENNUNG")
	
	passt 		:= objekte.New(0, 0, 0, 22)				// Passt-Objekt
	passtNicht 	:= objekte.New(0, 0, 0, 23)				// Passt-Nicht-Objekt
	
	*obj = append(*obj, passt, passtNicht)
	
	for i:=1;i<11;i++ {					// Schleife für die Abfrage von 10 Mustern
		
		if *tastatur {
			passt.SetzeAkt(true)
			passtNicht.SetzeAkt(true)
			*tastatur = false
		}
		
Neu0:			
		auswahl := rand.Intn(6)								// wählt zufällig eine der Muster-Zeilen 0 bis 5 aus 
		
		wahrOderFalsch := rand.Intn(2)						// durch Zufall wird wahre 1 oder falsche 0 Antwort gewählt
		if wahrOderFalsch == 0 {
			zufallSpalte = rand.Intn( len(texte.MusterN[auswahl]) )
		} else {
			zufallSpalte = rand.Intn( len(texte.MusterJ[ auswahl ]) )
		}
		
		musterwahl = 10*auswahl + zufallSpalte					// ausgewähltes Muster
		for i:=0;i<len(musterspeicher);i++ {
			if musterspeicher[i] == musterwahl {
				goto Neu0
			}
		}
		musterspeicher = append(musterspeicher, auswahl)		// fügt neue Musterkombination hinzu
		
		// -------------------------------------- Zeichnet den Hitnergrund zur Musterabfrage "Passt (nicht)"
		mutex.Lock()
		musterabfrage(i)
		
		titel.Zeichnen()
		
		SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
		Stiftfarbe(180,50,35)
		SchreibeFont (130, 340 , texte.MusterV[auswahl] )							// Muster-Vorgabe
		
		if wahrOderFalsch == 0 {
			SchreibeFont (480, 340 , texte.MusterN[ auswahl ][ zufallSpalte ] )			// falsches Muster
		} else if wahrOderFalsch == 1 {
			SchreibeFont (480, 340 , texte.MusterJ[ auswahl ][ zufallSpalte ] )			// richtiges Muster
		}
		passt.Zeichnen()
		passtNicht.Zeichnen()
		
		Archivieren()
		mutex.Unlock()
		// -----
		
		for !*signal { time.Sleep( time.Duration(2e8) ) }
		*signal = false
		if *wert == uint8(wahrOderFalsch) {
			SpieleSound("./Sounds/Sparkle.wav")
			*punkte += 50
			maus.SetzeTyp(26)
			go setzeMaus(maus)
		} else {
			SpieleSound("./Sounds/Beep.wav")
			*punkte -= 10
			maus.SetzeTyp(27)
			go setzeMaus(maus)
			time.Sleep( time.Duration(1e9) )
		}
		
		if wahrOderFalsch == 1 {
			passt.SetzeAkt(false)
			passtNicht.SetzeAkt(false)
			*tastatur = true													// aktiviert die Tastatur-Eingabe
			
			// -------------------------------------- Zeichnet den Hintergrund zur Mustereingabe "f="
			mutex.Lock()
			mustereingabe(i,0)
			
			titel.Zeichnen()	
			
			SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
			Stiftfarbe(180,50,35)
			SchreibeFont (130, 340 , texte.MusterV[auswahl] )							// Muster-Vorgabe
			
			SchreibeFont (480, 340 , texte.MusterJ[ auswahl ][ zufallSpalte ] )			// richtiges Muster
			
			Archivieren()
			mutex.Unlock()
			
Neu1:		// -----
			for !*signal { time.Sleep( time.Duration(2e8) ) }
			*signal = false
																							
			if *eingabe == texte.MusterL[ auswahl ][ zufallSpalte ][0] {		// ABFRAGE: richtiges Muster: Lösung 1
				SpieleSound("./Sounds/Sparkle.wav")
				*eingabe = ""
				*punkte += 50
				maus.SetzeTyp(26)
				go setzeMaus(maus)
			} else if versuch==0 {
				SpieleSound("./Sounds/Beep.wav")
				*punkte -= 10
				maus.SetzeTyp(27)
				go setzeMaus(maus)
				versuch++
				goto Neu1
			} else {
				versuch = 0
			}
			
			// -------------------------------------- Zeichnet den Hintergrund zur Mustereingabe "g="
			mutex.Lock()
			mustereingabe(i,1)
			
			titel.Zeichnen()	
			
			SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
			Stiftfarbe(180,50,35)
			SchreibeFont (130, 340 , texte.MusterV[auswahl] )							// Muster-Vorgabe
			
			SchreibeFont (480, 340 , texte.MusterJ[ auswahl ][ zufallSpalte ] )			// richtiges Muster
			
			Archivieren()
			mutex.Unlock()
			// -----
Neu2:			
			for !*signal { time.Sleep( time.Duration(2e8) ) }
			*signal = false
																							
			if *eingabe == texte.MusterL[ auswahl ][ zufallSpalte ][1] {		// ABFRAGE: richtiges Muster: Lösung 2
				SpieleSound("./Sounds/Sparkle.wav")
				*eingabe = ""
				*punkte += 50
				maus.SetzeTyp(26)
				go setzeMaus(maus)
			} else if versuch==0 {
				SpieleSound("./Sounds/Beep.wav")
				*punkte -= 10
				maus.SetzeTyp(27)
				go setzeMaus(maus)
				versuch++
				goto Neu2
			} else {
				versuch = 0
			}
		}
		time.Sleep( time.Duration(2e9) )				// vor neuem Muster warte 2 Sekunden
	}
	*tastatur = false
	*eingabe = ""

}

func musterabfrage(i int) {
	LadeBild (0,0, "./Bilder/Funktionale.bmp")			// Hintergrund des Muster-Raumes wird gezeichnet
	
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
	
	Stiftfarbe(220,220,220)														
	Vollrechteck(360,150,490,80)
	
	Transparenz(40)
	Stiftfarbe(76,0,153)														
	Vollrechteck(100,250,300,180)												
	Vollrechteck(450,250,650,180)
	Transparenz(0)
	
	Stiftfarbe(100,180,255)	
	SchreibeFont (400, 150 , "Muster Nr. " + fmt.Sprint(i) )
	Stiftfarbe(30,30,30)
	SchreibeFont (120, 250 , "Muster:      Argument:" )	
}

func mustereingabe(i,opt int) {
	LadeBild (0,0, "./Bilder/Funktionale.bmp")			// Hintergrund des Muster-Raumes wird gezeichnet
	
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
	
	Stiftfarbe(220,220,220)														
	Vollrechteck(360,150,490,80)
	Stiftfarbe(100,180,255)	
	SchreibeFont (400, 150 , "Muster Nr. " + fmt.Sprint(i) )
	
	Transparenz(40)
	Stiftfarbe(76,0,153)														
	Vollrechteck(100,250,300,180)												
	Vollrechteck(450,250,650,180)
	Vollrechteck(150,575,900,100)
	
	Stiftfarbe(153,204,0)
	Vollrechteck(100,450,1000,90)
	Transparenz(0)
		
	Stiftfarbe(30,30,30)
	SchreibeFont (120, 250 , "Muster:      Argument:" )
	if opt == 0 {
		SchreibeFont (195,580,"Bindung:   f =")
		Stiftfarbe(65,96,140)
		SchreibeFont (140,455,"Tippe die Bindung für f ein!")
	} else {
		SchreibeFont (195,580,"Bindung:  w =")
		Stiftfarbe(65,96,140)
		SchreibeFont (135,455,"Tippe die Bindung für w ein!")
	}
}

func memorySpiel(obj *[]objekte.Objekt, akt *bool, level uint8, rand *rand.Rand) {		// füllt Karten ins Array
	
	titel := objekte.New(300,50,150,24)
	titel.SetzeInhalt("MUSTER-MEMORY")
	*obj = append(*obj, titel)
	*akt = true
	
	var musterListe [12]string		// wird mit Mustern gefüllt
	var counter uint8				// zählt die neu erstellten Karten (bis 12)
	var neu objekte.Objekt
	
	switch level {					// Karten-Beschriftung als Liste
		case 1:
		musterListe = [12]string{	"     [ (x:[]) ]","       [ ['a'] ]","      ( 'a':y )","        \"aa\"","     ( x:\"b\" )", "        \"bb\"",
								"     [ x , \"b\" ]","  [ \"b\" , \"b\" ]","      ( x , [] )","     ( 'a' , [] )","     ( x:'b':y )","       \"bba\""}
		case 2:
		musterListe = [12]string{	"      [(ix:y)]","        [\"a\"]","        ('x':y)","         \"xy\"","       (b:\"a\")", "         \"aa\"",
								"       [a,\"b\"]","     [\"b\",\"b\"]","        (u,v)","(\"Not\",False)","      (x:'b':y)","      \"oben\""}
		case 3:
		musterListe = [12]string{	"    [(wa:nn)]","     [[2023]]","    ('L':iebe)","      \"LWB\"","      (bl:\"a\")", "        \"Ja\"",
								"  [wer,\"MP\"]"," [\"AB\",\"MP\"]","   (can,find)","(True,\"Love\")","    (o:'d':er)"," \"Adventure\""}
	}
	// zugehörige Typen zu den Beschriftungen:
	typListe := [12]int64{1,1,2,2,3,3,4,4,5,5,6,6}
	
	rand.Shuffle(12, func(i,j int) {
		typListe[i],	typListe[j] 	= typListe[j],	 typListe[i] 
		musterListe[i],	musterListe[j] 	= musterListe[j],musterListe[i] } )		// mischt die beiden Listen randomisiert (identisch) durch
	
	for i:=uint16(75);i<1000;i+=275 {
		for j:=uint16(150);j<600;j+=190 {										// erstellt Karten mit je einem Muster & Typ an verschiedenen Koordinaten (i,j) 
			if level == 2 {
				neu = objekte.New(i,j,150,32)
			} else {
				neu = objekte.New(i,j,150,31)
			}
			neu.SetzeInhalt( musterListe[counter] )
			neu.SetzeErstellung( typListe[counter] )
			*obj = append(*obj, neu)											// fügt die neu erstellte Karte dem Objekte-Array hinzu
			counter++
		}
	}
	*akt = true	
}

// Es folgt die VIEW-Komponente
func view_komponente (obj *[]objekte.Objekt, maus,okayObjekt objekte.Objekt, signal, stop ,akt, ende *bool, 
													punkte, diff *int16, mutex *sync.Mutex, eingabe *string) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	for { //Endlos ...
		mutex.Lock()
		UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()																// Cleart vollständigen Screen
		
		if *akt { 
			ObjAktualisieren(obj)
			*akt = false
		} else {
			Restaurieren(0,0,1200,700)										// Restauriert das alte Hintergrundbild
		}
		
		if *stop {
			okayObjekt.Zeichnen()
		}
		SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 70 )
		Stiftfarbe(120,180,120)
		SchreibeFont (645,580,*eingabe)
		
		SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 35 )
		Stiftfarbe(76,0,153)  
		SchreibeFont (500,12,"Punkte : "+fmt.Sprint (*punkte))				// Schreibe rechts oben Punkte
		Stiftfarbe(100,10,155)
		Schreibe (2,2,"FPS:"+fmt.Sprint (anzahl))							// Schreibe links oben FPS
		if *signal {  }
			
		maus.Zeichnen()														// Zeichnet Maus
			
		
		if time.Now().UnixNano() - t1 < 1000000000 { //noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() 				// neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}				//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}				//Frame-Rate :-)
		}
		
		UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
		mutex.Unlock()
		
		if *ende { return }
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
		
	}
}

func ObjAktualisieren(obj *[]objekte.Objekt) {
	LadeBild (0,0, "./Bilder/Funktionale.bmp")		// Hintergrund des Muster-Raumes wird gezeichnet
	
	for _,ob := range *obj { 								// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()											// Speichert das Hintergrund-Bild
}

// Es folgt die CONTROL-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus,okayObjekt objekte.Objekt, signal, stop, akt, ende *bool, punkte, diff *int16, wert *uint8, kanal chan bool) {
	//var taste uint8
	var aufgedeckt,warten bool = false,false			// gibt an, ob eine Karte aufgedeckt wurde, auf das Zudecken gewartet wird
	var objektSpeicher,objektSpeicher2 objekte.Objekt	// Speichert aufgedeckte Karten
	var zaehler uint8									// überprüft, wie viele Paare aufgedeckt wurden
	
	for {
		taste, status, mausX, mausY := MausLesen1()
		// fmt.Println(taste, status, mausX, mausY)
		maus.SetzeKoordinaten(mausX,mausY)					// Aktualisiert Maus-Koordinaten
		
		if *stop {
			if status==1 { 									// Maustaste wird gedrückt
				if ja,_ := okayObjekt.Getroffen(mausX,mausY,1); ja {
					*stop = false
				}
			}
		} else if *ende {
			return
		} else {
			if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
				if warten {
					if objektSpeicher.GibTyp() == 34 {
						objektSpeicher2.SetzeTyp(32)
						objektSpeicher.SetzeTyp(32)
					} else {
						objektSpeicher2.SetzeTyp(31)
						objektSpeicher.SetzeTyp(31)
					}
					warten = false
					*akt = true
				} else {
					for _,ob := range *obj { 									// überprüft Objekte im Array
						if get,lang := ob.Getroffen(mausX,mausY,1); get {
							if !aufgedeckt && ob.GibTyp() == 32 {					// Fälle für Level 1
								SpieleNote("5A",0.1,0)
								ob.SetzeTyp(34)
								objektSpeicher = ob
								aufgedeckt = true
								*akt = true
							} else if aufgedeckt && ob!=objektSpeicher && ob.GibTyp()==32 {
								*akt = true
								if lang == objektSpeicher.GibErstellung() {
									ob.SetzeTyp(33)
									objektSpeicher.SetzeTyp(33)
									SpieleSound("./Sounds/Sparkle.wav")
									*punkte += 50
									maus.SetzeTyp(26)
									zaehler++
									if zaehler == 6 {
										zaehler = 0
										kanal <- true
										*signal = true
									}
									go setzeMaus(maus)
								} else {
									ob.SetzeTyp(34)
									SpieleSound("./Sounds/Beep.wav")
									*punkte -= 10
									objektSpeicher2 = ob
									warten = true
									maus.SetzeTyp(27)
									go setzeMaus(maus)
								}
								aufgedeckt = false
								*akt = true
							} else if !aufgedeckt && ob.GibTyp() == 31 {			// Fälle für Level 2 und 3
								SpieleNote("5A",0.1,0)
								ob.SetzeTyp(32)
								objektSpeicher = ob
								aufgedeckt = true
								*akt = true
							} else if aufgedeckt && ob!=objektSpeicher && ob.GibTyp()==31 {
								*akt = true
								if lang == objektSpeicher.GibErstellung() {
									ob.SetzeTyp(33)
									objektSpeicher.SetzeTyp(33)
									SpieleSound("./Sounds/Sparkle.wav")
									*punkte += 50
									maus.SetzeTyp(26)
									zaehler++
									if zaehler == 6 {
										zaehler = 0
										kanal <- true
										*signal = true
									}
									go setzeMaus(maus)
								} else {
									ob.SetzeTyp(32)
									SpieleSound("./Sounds/Beep.wav")
									*punkte -= 10
									objektSpeicher2 = ob
									warten = true
									maus.SetzeTyp(27)
									go setzeMaus(maus)
								}
								aufgedeckt = false
								*akt = true //
							} else {
								*wert = uint8(lang)
								*signal = true
								//SpieleSound("./Sounds/Beep.wav")
							}
							
						}
					}
				}
			}
			if taste == 3 && status == 1 { 			//RECHTE Maustaste gerade gedrückt
				for _,ob := range *obj { 							// Zeichnet alleweiteren Objekte ein
					if get,lang :=  ob.Getroffen(mausX,mausY,3); get {
						if lang == 0 {
							*punkte -= 5
							
						} else if lang < 3.5e8 {
							 *punkte += 20
							 *diff = 20
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 20 Punkte")
						} else if lang < 4.2e8 {
							 *punkte += 15
							 *diff = 15
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 15 Punkte")
						} else if lang < 5.2e8 {
							 *punkte += 10
							 *diff = 10
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 10 Punkte")
						} else if lang < 7e8 {
							 *punkte += 5
							 *diff = 5
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 5 Punkte")
						} else if lang < 1e9 {
							 *punkte += 2
							 *diff = 2
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 2 Punkte")
						} else {
							 *punkte += 1
							 *diff = 1
							 fmt.Println("Reaktionszeit: ",lang/1e6, " 1 Punkt")
						}		
						*akt = true
					}
				}
			}
		}
	}
}

func setzeMaus (maus objekte.Objekt) {
	time.Sleep(2e9)
	maus.SetzeTyp(25)
}
	
