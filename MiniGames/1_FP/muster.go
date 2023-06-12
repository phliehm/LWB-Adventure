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
	
func Muster() (note float32, punktExp uint32) {
	var mutex sync.Mutex					// erstellt Mutex
	var wg sync.WaitGroup					// erstellt Waitgroup
	var gedrueckt uint8
	var taste,tiefe uint16
	
	var eingabe string						// zur "Editor"-Eingabe
	var punkte int16 = 0					// Spiel-Punktzahl
	var punkteArr [4]int16					// Slice für Punktzahlen
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
	go view_komponente(&obj, maus, okayObjekt, &signal, &stop, &akt, &ende, &punkte, &diff, &mutex, &eingabe, &wg)
	
	// Objekte werden nach und nach in der Welt platziert
	go spielablauf(&obj, maus, random, &mutex, &akt, &tastatur, &stop, &signal, &ende, &eingabe, &wert, &punkte, &punkteArr, kanal, &wg)
	
	// Nebenläufig wird die Kontroll-Komponente für die Maus gestartet.
	go maussteuerung(&obj, maus, okayObjekt, &signal, &stop, &akt, &ende, &punkte, &diff, &wert, kanal, &wg)
	
	go musikhintergrund(&ende)
	
	// Die Kontroll-Komponente 2 ist die 'Mainloop' im Hauptprogramm	
	// Wir fragen hier nur die Tastatur ab.
	
	wg.Add(3)						// Wait-Group erhält Counter 3 zum Warten auf das Ende der nebenläufigen Prozesse
	
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 28 )
	
	
A:	for {
		taste, gedrueckt, tiefe = TastaturLesen1()
		
		if tastatur {
			if gedrueckt == 1  { 						// Beim Drücken der Taste, nicht beim Loslassen!
				switch {
					case taste == 27:  									// ESC-Taste
					break A
					case taste==13 || taste==271:  						// Enter-Taste(n)
					signal = true
					case taste == 32:  									// Leer-Taste
					eingabe += " "
					case taste ==  8:  									// Backspace-Taste
					if eingabe != "" {
						eingabe = eingabe [:len(eingabe)-1]
					}
					case taste ==  276:  								// LINKS-Taste
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
			if gedrueckt == 1  { 					// Beim Drücken der Taste (1), nicht beim Loslassen!
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
	
	stop = false
	ende = true 
	if !signal {
		signal = true
	} else {
		kanal <- false
	}
	
	wg.Wait()
	
	punkte = punkteArr[0] + punkteArr[1] + punkteArr[2] + punkteArr[3]
	if punkte>0 {
		punktExp = uint32(punkte)
	} else {
		punktExp = 0
	}
	switch {
		case punktExp == 0:		note = 5.0
		case punktExp < 135:	note = 4.7
		case punktExp < 270:	note = 4.3
		case punktExp < 405:	note = 4.0
		case punktExp < 540:	note = 3.7
		case punktExp < 675:	note = 3.3
		case punktExp < 810:	note = 3.0
		case punktExp < 945:	note = 2.7
		case punktExp < 1080:	note = 2.3
		case punktExp < 1215:	note = 2.0
		case punktExp < 1350:	note = 1.7
		case punktExp < 1485:	note = 1.3
		case punktExp > 1615:	note = 1.0
	}
		
	endbildschirm(punktExp,punkteArr,note)
	
	fmt.Println("Du hast ",punktExp," Punkte erreicht!")
	fmt.Println("Damit erreichst du die Note ",note," .\nHerzlichen Glückwunsch!")
	time.Sleep( time.Duration(1e9) )
	
	return
}

func musikhintergrund(ende *bool) {
	for !*ende {
		SpieleSound("./Sounds/Muster.wav")
		time.Sleep( time.Duration(44e8) )
	}
}

func endbildschirm(EndP uint32, punkteArr [4]int16, note float32) {
	
	Stiftfarbe(255,255,255)
	Cls()
	
	LadeBild(150,100, "./Bilder/Zertifikat/sprechblase_flipped_400.bmp")
	LadeBild(230,390, "./Bilder/FP/FabWeb_fullBody_gespiegelt.bmp")
	
	LadeBild(620,80, "./Bilder/Zertifikat/paper_500.bmp")
	LadeBild(960,520, "./Bilder/Zertifikat/certified_100.bmp")
	LadeBild(720,550, "./Bilder/FP/fu-logo.bmp")
	
		
	Stiftfarbe(0,255,0)
	SetzeFont( "./Schriftarten/ComputerTypewriter.ttf",70)
	SchreibeFont(40,10,"Funktionale  Programmierung")
	Stiftfarbe(0,0,0)
	SetzeFont( "./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	SchreibeFont(295,140,"Du hast die")
	SchreibeFont(310,260,"erreicht!")
	SetzeFont( "./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	SchreibeFont(285,170,"Gesamtnote")
	SetzeFont( "./Schriftarten/Starjedi.ttf",42)
	
	SchreibeFont(325,195,fmt.Sprintf("%2.1f",note))
	
	SetzeFont( "./Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	
	SchreibeFont(730,160+uint16(0*35),"Muster-Erkennung: ")
	SchreibeFont(850,160+uint16(1*35), fmt.Sprint(punkteArr[0]) + " Punkte")
	SchreibeFont(730,160+uint16(2*35),"Muster-Memory Lvl. 1: ")
	SchreibeFont(850,160+uint16(3*35), fmt.Sprint(punkteArr[1]) + " Punkte")
	SchreibeFont(730,160+uint16(4*35),"Muster-Memory Lvl. 2: ")
	SchreibeFont(850,160+uint16(5*35), fmt.Sprint(punkteArr[2]) + " Punkte")
	SchreibeFont(730,160+uint16(6*35),"Muster-Memory Lvl. 3: ")
	SchreibeFont(850,160+uint16(7*35), fmt.Sprint(punkteArr[3]) + " Punkte")
	SchreibeFont(720,150+uint16(8*35),"----------------------")
	SchreibeFont(730,180+uint16(8*35),"Gesamt:    " + fmt.Sprint(EndP) + " Punkte")

	TastaturLesen1()
}

	
func spielablauf(obj *[]objekte.Objekt, maus objekte.Objekt, random *rand.Rand, mutex *sync.Mutex, akt, tastatur, stop, signal, ende *bool, 
			eingabe *string, wert *uint8, punkte *int16, punkteArr *[4]int16, kanal chan bool, wg *sync.WaitGroup) {
	var neuerZustand bool
	
	defer wg.Done()
	time.Sleep( time.Duration(1e8) )
	zwischentext(&texte.MusterEinl, mutex, stop)			// Einleitungs-Text
	
	if *ende {return}
	
	zwischentext(&texte.MusterErk, mutex, stop)
	
	musterSpiel(obj, maus, akt, signal, tastatur, ende, random, mutex, eingabe, wert, punkte, kanal)
	if *ende { 
		fmt.Println("Beende spielablauf")	
		return 
	}
	
	punkteArr[0] = *punkte
	*punkte = 0
	
	*obj = make([]objekte.Objekt,0)
	time.Sleep( time.Duration(1e9) )
	
	zwischentext(&texte.MusterEins, mutex, stop)
	memorySpiel(obj, akt, 2, random)					// auf Level 1
	
	
	neuerZustand = <- kanal
	punkteArr[1] = *punkte						// speichert Punkte für Lvl.1
	*punkte = 0	
	if !neuerZustand { return }
										
	time.Sleep( time.Duration(2e9) )
	*obj = make([]objekte.Objekt,0)				// ---- leere und warte
	time.Sleep( time.Duration(1e8) )
	
	zwischentext(&texte.MusterZwei, mutex, stop)
	memorySpiel(obj, akt, 3, random)					// auf Level 2
	
	neuerZustand = <- kanal
	punkteArr[2] = *punkte						// speichert Punkte für Lvl.2
	*punkte = 0	
	if !neuerZustand { return }
	
	
	time.Sleep( time.Duration(2e9) )
	*obj = make([]objekte.Objekt,0)				// ---- leere und warte
	time.Sleep( time.Duration(1e8) )
	
	zwischentext(&texte.MusterDrei, mutex, stop)
	memorySpiel(obj, akt, 5, random)					// auf Level 3
		
	neuerZustand = <- kanal
	punkteArr[3] = *punkte						// speichert Punkte für Lvl.3
	if !neuerZustand { return }
	
	fmt.Println("Beende Spielablauf")
	*signal = false
}

func zwischentext(textArr *[]string, mutex *sync.Mutex, stop *bool) {
	mutex.Lock()
	LadeBild (0,0, "./Bilder/FP/Funktionale.bmp")		// Hintergrund des Muster-Raumes wird gezeichnet
	Transparenz(120)
	Stiftfarbe(76,0,153)														
	Vollrechteck(100,50,1000,600)
	Transparenz(0)
	SetzeFont ("./Schriftarten/Ubuntu-B.ttf", 50 )
	Stiftfarbe(20,20,20)									// schreibt den Schatten
	for ind,str := range *textArr {
		SchreibeFont (207, uint16(69+ind*55) ,str )
	}
	Stiftfarbe(124,212,255)									// schreibt den eigentlichen Text
	for ind,str := range *textArr {
		SchreibeFont (211, uint16(71+ind*55) ,str )
	}
	Archivieren()
	mutex.Unlock()
	
	*stop = true
	for *stop { time.Sleep( time.Duration(1e8) ) }
}

func musterSpiel(obj *[]objekte.Objekt, maus objekte.Objekt, akt, signal, tastatur, ende *bool, rand *rand.Rand, 
							mutex *sync.Mutex, eingabe *string, wert *uint8, punkte *int16, kanal chan bool) { // gibt Muster zur Abfrage
	
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
		} else if *ende {
			fmt.Println("Beende Muster-Spiel.")
			return
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
		SchreibeFont (130, 340 , texte.MusterV[auswahl] )								// Muster-Vorgabe
		
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
		if *ende { 
			fmt.Println("Beende Muster-Spiel.")
			return
		}
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
			*tastatur = true							// aktiviert die Tastatur-Eingabe
			
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
			if *ende { 
				fmt.Println("Beende Muster-Spiel.")
				return
			}
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
			if *ende { 
				fmt.Println("Beende Muster-Spiel.")
				return
			}
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
		time.Sleep( time.Duration(2e8) )				// vor neuem Muster warte 0,2 Sekunden
	}
	*tastatur = false
	*eingabe = ""
	// kanal <- true
}

func musterabfrage(i int) {
	LadeBild (0,0, "./Bilder/FP/Funktionale.bmp")			// Hintergrund des Muster-Raumes wird gezeichnet
	
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
	LadeBild (0,0, "./Bilder/FP/Funktionale.bmp")			// Hintergrund des Muster-Raumes wird gezeichnet
	
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
		case 1:	musterListe = texte.MusterListe1
		case 2:	musterListe = texte.MusterListe2
		case 3:	musterListe = texte.MusterListe3
	}
	typListe := [12]int64{1,1,2,2,3,3,4,4,5,5,6,6}	// zugehörige Typen zu den Karten-Beschriftungen
	
	rand.Shuffle(12, func(i,j int) {
		typListe[i],	typListe[j] 	= typListe[j],	 typListe[i] 
		musterListe[i],	musterListe[j] 	= musterListe[j],musterListe[i] } )		// mischt die beiden Listen randomisiert (identisch) durch
	
	for i:=uint16(75);i<1000;i+=275 {
		for j:=uint16(150);j<600;j+=190 {										// erstellt Karten mit je einem Muster & Typ an verschiedenen Koordinaten (i,j) 
			switch level {
				case 1:
				neu = objekte.New(i,j,150,31)
				neu.SetzeInhalt( musterListe[counter] )
				case 2:
				neu = objekte.New(i,j,150,32)
				neu.SetzeInhalt( musterListe[counter] )
				case 3:
				neu = objekte.New(i,j,150,31)
				neu.SetzeInhalt( musterListe[counter] )
				case 5:
				neu = objekte.New(i,j,150,35)
			}
			neu.SetzeErstellung( typListe[counter] )
			*obj = append(*obj, neu)											// fügt die neu erstellte Karte dem Objekte-Array hinzu
			counter++
		}
	}
	*akt = true	
}

// Es folgt die VIEW-Komponente
func view_komponente (obj *[]objekte.Objekt, maus,okayObjekt objekte.Objekt, signal, stop ,akt, ende *bool, 
													punkte, diff *int16, mutex *sync.Mutex, eingabe *string, wg *sync.WaitGroup) {   	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	defer wg.Done()
	
	for { //Endlos ...
		mutex.Lock()
		UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		
		Stiftfarbe(255,255,255)
		Cls()												// Cleart vollständigen Screen
		
		if *akt { 
			objAktualisieren(obj)
			*akt = false
		} else {
			Restaurieren(0,0,1200,700)						// Restauriert das alte Hintergrundbild
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
		
		if *ende { 
			fmt.Println("View-Komponente beendet!")
			return 
		}
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
	}
}

func objAktualisieren(obj *[]objekte.Objekt) {
	LadeBild (0,0, "./Bilder/FP/Funktionale.bmp")		// Hintergrund des Muster-Raumes wird gezeichnet
	
	for _,ob := range *obj { 								// Zeichnet alleweiteren Objekte ein
		ob.Zeichnen()
	}
	Archivieren()											// Speichert das Hintergrund-Bild
}

// Es folgt die Maus-Komponente 1 --- Kein Bestandteil der Welt, also unabhängig -----
func maussteuerung (obj *[]objekte.Objekt, maus,okayObjekt objekte.Objekt, signal, stop, akt, ende *bool, 
					punkte, diff *int16, wert *uint8, kanal chan bool, wg *sync.WaitGroup) {
	
	var aufgedeckt,warten bool = false,false			// gibt an, ob eine Karte aufgedeckt wurde, auf das Zudecken gewartet wird
	var objektSpeicher,objektSpeicher2 objekte.Objekt	// Speichert aufgedeckte Karten
	var zaehler uint8									// überprüft, wie viele Paare aufgedeckt wurden
	
	defer wg.Done()
	
	for {
		_, status, mausX, mausY := MausLesen1()
		maus.SetzeKoordinaten(mausX,mausY)								// Aktualisiert Maus-Koordinaten
		
		if *ende {
			fmt.Println("Maussteuerung beendet")
			return
		} else if *stop {												// bei Zwischenanzeigen wird auf "OK" gewartet
			if status==1 { 												// Maustaste wird gedrückt
				if ja,_ := okayObjekt.Getroffen(mausX,mausY,1); ja {
					*stop = false
				}
			}
		} else {
			if status==1 { 									// Maustaste im normalen Modus gedrückt
				if warten {
					if objektSpeicher.GibTyp() == 34 {
						objektSpeicher2.SetzeTyp(32)
						objektSpeicher.SetzeTyp(32)
					} else if objektSpeicher.GibTyp() == 32 {
						objektSpeicher2.SetzeTyp(31)
						objektSpeicher.SetzeTyp(31)
					} else if objektSpeicher.GibTyp() == 36 {
						objektSpeicher2.SetzeTyp(35)
						objektSpeicher.SetzeTyp(35)
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
										//*signal = true
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
										//*signal = true
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
								*akt = true

							} else if !aufgedeckt && ob.GibTyp() == 35 {			// Fälle für Level 3 SOUND
								ob.SetzeTyp(36)
								go spieleKlang(lang)
								objektSpeicher = ob
								aufgedeckt = true
								*akt = true
							} else if aufgedeckt && ob!=objektSpeicher && ob.GibTyp()==35 {
								ob.SetzeTyp(36)
								go spieleKlang(lang)
								*akt = true
								if lang == objektSpeicher.GibErstellung() {
									ob.SetzeTyp(33)
									objektSpeicher.SetzeTyp(33)
									*punkte += 50
									maus.SetzeTyp(26)
									zaehler++
									if zaehler == 6 {
										zaehler = 0
										fmt.Println("Sende Signal in maussteuerung")
										kanal <- true
										//*signal = true
									}
									go setzeMaus(maus)
								} else {
									*punkte -= 10
									objektSpeicher2 = ob
									warten = true
									maus.SetzeTyp(27)
									go setzeMaus(maus)
								}
								aufgedeckt = false
								*akt = true 
						
							} else {
								*wert = uint8(lang)
								*signal = true
							}
							
						}
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

func spieleKlang(lang int64) {
	switch lang {
		case 1: SpieleSound("./Sounds/1Air.wav")
		case 2: SpieleSound("./Sounds/2Bergkoenig.wav")
		case 3: SpieleSound("./Sounds/3Koenigin.wav")
		case 4: SpieleSound("./Sounds/4Pachelbel.wav")
		case 5: SpieleSound("./Sounds/5Zuckerfee.wav")
		case 6: SpieleSound("./Sounds/6Jahreszeiten.wav")
	}
}	
