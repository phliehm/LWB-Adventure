
// Paket zum Gertränkeautomatenspiel

// Implementation

// 4.6.2023		Martin Seiß



// Folgende Getränkezuteilung wird verwendet:
// ******************************************
//  Dark Schmidther	-	Extraschwarzer Kaffee
//	The Herk		-	Melissentee
//  J.EthI			-	Kaffee mit Milch und 2x Zucker
//  FabFour			-	Cappuccino
//	Amoebi			-	Grüner Tee
//  WtheK			-	Bier


// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 1
// verwendet:
// ******************************************
//  A - Grüner Tee			-	[10,20,50]
//	B - Pfefferminztee		-	[20,50,10]
//  C - Melissentee			-	[50,10,20]
//  D - Schwarzer Tee 		-	[20,10,50]
//	E - Kräutertee			-	[10,50,20]
//  F - Früchtetee			-	[50,20,10]


// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 2
// verwendet:
// ******************************************
//  B - Bier					-	[50,50,10,20]
//	R - Radler					-	[50,50,10,10]
//  E - Extrascharzer Kaffee		[50,20,20,10]
//  L - Latte Macchiato 		-	[50,10,20,20]
//	C - Cappuccino				-	[50,20,10,20]
//  K - Kaffee					-	[50,(10,10,20)]
//		mit Zucker				-   [10,50,(10,20)]
// 		mit 2xZucker			-	[10,10,50,20]	
//		mit Milch				-   [20,50,10,10)]
// 		mit Milch und Zucker	-	[20,10,50,10]
// 		mit Milch und 2xZucker	-	[20,10,50,10]
// in ()-Klammern - Reihenfolge egal


// Folgende Getränkezuteilung mit Name und Nummer wird verwendet:
// ******************************************
//  1 - Extradunkle Schokolade			 	-	[]
//	2 - Melissentee							-	[]
//  3 - Kaffee mit Milch und 2x Zucker		-   []
//  4 - Kaffe mit 1x Zucker					-	[]
//	5 - Grüner Tee							-	[]
//  6 - doppelter Expresso mit 2x Zucker	-   []

// ()-Klammern - Reihenfolge egal


package getraenkeautomat

import "fmt"
import "gfx"
//import "../../Materialordner/gfx"
import "time"
import "../../Klassen/textboxen"
import "../../Klassen/buttons"

var path string = ""				// Pfad für Klassen und Daten
var txt  textboxen.Textbox			// Text für die seitlichen Textbox

// Levelparamter
var getraenkeID rune			// ID des Getränks - zum Automat
var automatNr uint = 1					
var muenzenAnzahl [3]uint		// Anzahl der [10er,20er,50er]
var einwurf []uint 


func Getraenkeautomat() (float32, uint32) {
	
	var note float32 = 6.0
	var punkte uint32
	
	var status uint					// 0 = weitere Münzen einwerfen
									// 1 = richtiges Getränk bestellt
									// 2 = falsches Getränk bestellt
	var font string = path + "Schriftarten/Ubuntu-B.ttf"
	var neueEingabe bool
	var ilevel uint16				// aktuelle Levenummer, beginned bei 0
	var nlevel uint16 = 6			// Anzahl der Level

	// ----------------   initialisiere die Textbox ----------------//
	gfx.SetzeFont (font,10)
	txt = textboxen.New(850,400,300,300)
	txt.SchreibeText("Willkommen zum Getränkeautomatenspiel!\n\n" +
		"Die Dozenten schaffen es nicht sich selber ein Getränk " + 
		"zu kaufen. Obwohl es ganz einfach ist. Man muss nur die " +
		"Münzen in der richtigen Reihenfolge einwerfen. " +
		"Ansonsten hat der Automat nur einen (R)eset-Knopf. Kaufen Sie " +
		"den Dozenten ihr Lieblingsgetränk!\n\n" +
		"Viel Spaß!")

	
	// -----        Erzeuge Buttons zur Spielsteuerung    ----------///
	var weiter,starter,beenden buttons.Button
	weiter = buttons.New(700,600,100,40,255,255,100,false,"  weiter")
	weiter.SetzeFont(font)
	starter = buttons.New(700,600,100,40,255,255,100,true,"   Start")
	starter.SetzeFont(font)
	beenden = buttons.New(700,650,100,40,255,255,100,true,"   Ende")
	beenden.SetzeFont(font)
	
	var zehner,zwanziger,fuenfziger,reset buttons.Button
	zehner = buttons.New(720,70,61,61,255,255,100,false,"")
	zehner.SetzeFont(font)
	zwanziger = buttons.New(710,160,80,80,255,255,100,false,"")
	zwanziger.SetzeFont(font)
	fuenfziger = buttons.New(700,250,100,100,255,255,100,false,"")
	fuenfziger.SetzeFont(font)
	reset = buttons.New(570,48,50,73,255,255,100,false,"")
	reset.SetzeFont(font)


	// ---------------- Zeichne Spielfeld -------------------------- //
	muenzenAnzahl = [3]uint{0,0,0}			// Startwert
	gfx.Stiftfarbe(0,0,0)
	zeichneSpielfeld(1,0,6.,muenzenAnzahl,weiter,starter,beenden)

	// ---------------- starte Musik ------------------------------- //
	go hintergrundmusik(beenden)
	
	// ----------       Eingabesteuerung     ------------------------//
	// ----------- Mausabfrage - Spielsteuerung ---------------------//
	for {

		taste, mstatus, mausX, mausY := gfx.MausLesen1()
		if taste==1 && mstatus==1 {

			if starter.TesteXYPosInButton(mausX,mausY) { // Spiel starten
				starter.DeaktiviereButton()
				aktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset)
				ladeLevel(ilevel)
				//fmt.Println("getraenkeID: ",getraenkeID)				
				zeichneSpielfeld(ilevel+1,punkte,note,muenzenAnzahl,weiter,starter,beenden)
			}

			if weiter.TesteXYPosInButton(mausX,mausY) { // Spiel starten
				weiter.DeaktiviereButton()
				aktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset)
				ladeLevel(ilevel)
				zeichneSpielfeld(ilevel+1,punkte,note,muenzenAnzahl,weiter,starter,beenden)
			}

			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende des Spiels
				beenden.DeaktiviereButton()
				gfx.StoppeAlleSounds()
				break
			}

			if reset.TesteXYPosInButton(mausX,mausY) { // Spiel zurücksetzen
				// weiter.DeaktiviereButton()
				// aktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset)
				gfx.SpieleSound(path + "Sounds/Beep.wav")
				ladeLevel(ilevel)
				zeichneSpielfeld(ilevel+1,punkte,note,muenzenAnzahl,weiter,starter,beenden)
			}

			
			if zehner.TesteXYPosInButton(mausX,mausY) {
				// Münzenanzahl kontrollieren!
				if muenzenAnzahl[0] == 0 {
					gfx.SpieleSound(path + "Sounds/Beep.wav")
				} else {
					gfx.SpieleSound(path + "Sounds/Sparkle.wav")
					muenzenAnzahl[0]--
					einwurf = append(einwurf,10)
					neueEingabe = true
					//fmt.Println("Zehner eingeworfen")
				}
			}

			if zwanziger.TesteXYPosInButton(mausX,mausY) {
				// Münzenanzahl kontrollieren!
				if muenzenAnzahl[1] == 0 {
					gfx.SpieleSound(path + "Sounds/Beep.wav")
				} else {
					gfx.SpieleSound(path + "Sounds/Sparkle.wav")
					muenzenAnzahl[1]--
					einwurf = append(einwurf,20)
					neueEingabe = true
					//fmt.Println("Zwanziger eingeworfen")
				}
			}

			if fuenfziger.TesteXYPosInButton(mausX,mausY) {
				// Münzenanzahl kontrollieren!
				if muenzenAnzahl[2] == 0 {
					gfx.SpieleSound(path + "Sounds/Beep.wav")
				} else {
					gfx.SpieleSound(path + "Sounds/Sparkle.wav")
					muenzenAnzahl[2]--
					einwurf = append(einwurf,50)
					neueEingabe = true
					//fmt.Println("Fünfziger eingeworfen")
				}
			}

			// status: 1 = richtiges Getränk bestellt
			// 2 = falsches Getränk bestellt, 0 = zu wenig Münzen (weiter)
			// 3 = zuviele Münzen
			if neueEingabe {
				status = verarbeitungAutomat(einwurf,getraenkeID,automatNr)
				if status == 1 {			// richtiges Getränk, gewonnen!
					note = note - 1.
					if note < 1. {note = 1.}
					punkte = punkte + 1
					ilevel = ilevel + 1
					txt.SchreibeText("Super, das richtige Getränk ist geliefert. \n\n")
					deaktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset)
					if ilevel == nlevel {
						gfx.SpieleSound(path + "Sounds/Applaus.wav")
						txt.SchreibeText("Super, das richtige Getränk ist geliefert. \n\n" +
							"Sie haben alle Aufgaben bestanden. Danke für das Spielen! \n\n")
					} else {
						gfx.SpieleSound(path + "Sounds/Applaus.wav")
						weiter.AktiviereButton()
					}
				}
				if status == 2 {		// falsches Getränk
					txt.SchreibeText("Schade, das falsche Getränk ist geliefert. \n\n " +
							"Versuchen Sie es noch einmal.")
					gfx.SpieleSound(path + "Sounds/GameOver.wav")
					starter.AktiviereButton()
					//ladeLevel(ilevel)
				}
				if status == 3 {		// zu viele Münzen
					txt.SchreibeText("Schade, es wurden zu viele Münzen" +
							"  eingeworfen. \n\n" +
							" Versuchen Sie es noch einmal.")
					gfx.SpieleSound(path + "Sounds/GameOver.wav")
					starter.AktiviereButton()
				} 					
				if status == 4 {		// Kombination existiert nicht
					txt.SchreibeText("Schade, die Münzen wurden" +
							" in der falschen Reihenfolge eingeworfen. \n\n" +
							" Versuchen Sie es noch einmal.")
					gfx.SpieleSound(path + "Sounds/GameOver.wav")
					starter.AktiviereButton()
				} 				
				neueEingabe = false
				zeichneSpielfeld(ilevel+1,punkte,note,muenzenAnzahl,weiter,starter,beenden)		
			}

		}
	}

	endbildschirm(ilevel+1, 6, note)

	return note,punkte
	
}


// ---------------   Hilfsfunktionen   -------------------------//


// Vor: Ein passendes gfx-Fenster (1200x700) ist geöffnet. Level > 0.
// Eff: Der Endbildschirm für das Spiel ist angezeigt und kann mit einem
//		Mausklick auf das Verlassen-Symbol verlassen werden.
func endbildschirm(level uint16, maxlevel uint16, note float32) {
	
	var path string = ""
	var beenden buttons.Button
	var text textboxen.Textbox
	beenden = buttons.New(1100,570,99,129,255,255,100,true,"")

	// Lade Hintergrund
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.LadeBild(150,80,path + "Bilder/Zertifikat/sprechblase_flipped_400.bmp")
	gfx.LadeBild(20,350,path+"Bilder/Martin/getraenkeautomat/Jethi_klein.bmp")
	gfx.LadeBild(620,80,path + "Bilder/Zertifikat/paper_500.bmp")
	gfx.LadeBild(960,520,path + "Bilder/Zertifikat/certified_100.bmp")
	gfx.LadeBild(1100,570,path + "Bilder/Martin/Zurück-Symbol.bmp")
	
	// Ausgabe der Gesamtnote	
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/ComputerTypewriter.ttf",80)
	gfx.SchreibeFont(25,10,"Getraenkeautomatenspiel")
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	gfx.SchreibeFont(295,120,"Du hast die")
	gfx.SchreibeFont(310,240,"erreicht!")
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	gfx.SchreibeFont(285,150,"Gesamtnote")
	gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	fmt.Println("Final Level: ",level)
	gfx.SchreibeFont(325,175,fmt.Sprintf("%2.1f",note))

	// Schreibe die Punkte pro Level und Gesamtpunkte
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	text = textboxen.New(710,300,350,600)
	text.SetzeSchriftgröße(32)
	text.SchreibeText("Sie haben " + fmt.Sprint(level-1) + " von " +
			fmt.Sprint(maxlevel) + " Level geschafft. \n\n" +
			"Kommen Sie bald wieder!")
	text.Zeichne()		

	// Warte auf Mausklick auf Beenden/Verlassen/Tür-Symbol 
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			if beenden.TesteXYPosInButton(mausX,mausY) {break}
		}
	}

}


// Eff: Weist den Levelparametern die richtigen Werte zu.
func ladeLevel(ilevel uint16) {
	
	if ilevel == 0 {
		automatNr = 1
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{2,3,1}
		txt.SchreibeText("Amoebi trinkt gerne einen grünen Tee.")
		getraenkeID = 'A'				
	} else if ilevel == 1 {
		automatNr = 1
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{2,3,1}
		txt.SchreibeText("Herk trinkt gerne einen Melissentee.")
		getraenkeID = 'C'				
	} else if ilevel == 2 {
		automatNr = 2
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{3,3,3}
		txt.SchreibeText("FabWeb trinkt gerne einen Cappuccino.")
		getraenkeID = 'C'				
	} else if ilevel == 3 {
		automatNr = 2
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{3,3,3}
		txt.SchreibeText("Winnie the K braucht ein Bier zur Stärkung.")
		getraenkeID = 'B'				
	} else if ilevel == 4 {
		automatNr = 2
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{3,3,3}
		txt.SchreibeText("Darth Schmidther trinkt ausschließlich " +
			"extraschwarzen Kaffee.")
		getraenkeID = 'E'	
	} else if ilevel == 5 {
		automatNr = 2
		einwurf = make([]uint,0)
		muenzenAnzahl = [3]uint{3,3,3}
		txt.SchreibeText("Ich mag am liebsten Kaffee mit Milch und 2x Zucker.")
		getraenkeID = 'J'	
	} else {
		fmt.Println("Gewonnen!")
	}
	
}


// Vor: Getränke haben nur Nummern von 1 bis 6 (siehe oben).
// Erg: Der Name des Getränkes ist zur Nummer geliefert.
func getraenkeNr2Name(nr uint) string {
	
	switch nr {
		case 1: return "Extrascharzer Kaffee"
		case 2: return "Melissentee"
		case 3: return "Kaffee mit Milch und 2x Zucker"
		case 4: return "Kaffe mit 1x Zucker"
		case 5: return "Cappuchino"
		case 6: return "Bier"
		default: panic("Falsche Getränkenummer")
	}

} 


// Vor: Getränke haben IDs von A bis ... (siehe oben).
// Erg: Der Name des Getränkes ist zur ID geliefert.
func getraenkeID2Name(id rune, automatNr uint) string {

	if automatNr == 1 {
		switch id {
			case 'A': return "Grüner Tee"
			case 'B': return "Pfefferminztee"
			case 'C': return "Melissentee"
			case 'D': return "Schwarzer Tee"
			case 'E': return "Kräutertee"
			case 'F': return "Früchtetee"
			default: panic("Falsche Getränke ID")
		}
	} else if automatNr == 2 {
		switch id {
			case 'B': return "Bier"
			case 'R': return "Radler"
			case 'E': return "Extraschwarzer Kaffee"
			case 'L': return "Latte Macchiato"
			case 'C': return "Cappuccino"
			case 'K': return "Kaffee"
			default: panic("Falsche Getränke ID")
		}	
	} else {
		panic("Falsche Automatennummer!")
	}
	
} 


// Vor: Eingabe ist Slice der Länge 3.
// Erg: Die passende Rune des Getränks ist zur Eingabe bei Automat 2 
//		geliefert.
//		Wenn das Getränk nicht existiert ist 'X' geliefert.
// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 2
// verwendet:
// ******************************************
//  A - Grüner Tee			-	[10,20,50]
//	B - Pfefferminztee		-	[20,50,10]
//  C - Melissentee			-	[50,10,20]
//  D - Schwarzer Tee 		-	[20,10,50]
//	E - Kräutertee			-	[10,50,20]
//  F - Früchtetee			-	[50,20,10]
func checkEingabe1(einwurf []uint) rune {
	var N int = len(einwurf)
	//var eingabeRune rune = 'X'
	if N == 3 {
		if einwurf[0] == 10 && einwurf[1] == 20 && einwurf[2] == 50 {
			return 'A'
		}
		if einwurf[0] == 20 && einwurf[1] == 50 && einwurf[2] == 10 {
			return 'B'
		}
		if einwurf[0] == 50 && einwurf[1] == 10 && einwurf[2] == 20 {
			return 'C'
		}
		if einwurf[0] == 20 && einwurf[1] == 10 && einwurf[2] == 50 {
			return 'D'
		}
		if einwurf[0] == 10 && einwurf[1] == 50 && einwurf[2] == 20 {
			return 'E'
		}
		if einwurf[0] == 50 && einwurf[1] == 20 && einwurf[2] == 10 {
			return 'F'
		}
	}
	return 'X'
}



// Vor: Getränke haben nur Nummern von 1 bis 6 (siehe oben).
// Erg: Der status ist geliefert: 1 = richtiges Getränk bestellt
// 		2 = falsches Getränk bestellt, 0 = zu wenig Münzen
//		3 = zuviele Münzen
func checkAutomat1(einwurf []uint,getraenkeID rune) uint {

	var N int = len(einwurf)
	var status uint
	var eingabeRune rune
	
	if N < 3 {status = 0}
	if N > 3 {status = 3}

	if N == 3 {		
		eingabeRune = checkEingabe1(einwurf)
		//fmt.Println(eingabeRune,getraenkeID)
		if eingabeRune == getraenkeID {
			status = 1
		} else if eingabeRune == 'X' {
			status = 4			// falsche Kombination
		} else {
			status = 2
		}
	}


return status

}

// Vor: Eingabe ist Slice der Länge 4.
// Erg: Die passende Rune des Getränks ist zur Eingabe bei Automat 1 
//		geliefert.
//		Wenn das Getränk nicht existiert ist 'X' geliefert.
// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 2
// verwendet:
// ******************************************
//  B - Bier					-	[50,50,10,20]
//	R - Radler					-	[50,50,10,10]
//  E - Extrascharzer Kaffee		[50,20,20,10]
//  L - Latte Macchiato 		-	[50,10,20,20]
//	C - Cappuccino				-	[50,20,10,20]
//  K - Kaffee					-	[50,(10,10,20)]
//		mit Zucker				-   [10,50,(10,20)]
// 		mit 2xZucker			-	[10,10,50,20]	
//		mit Milch				-   [20,50,10,10)]
// 		mit Milch und Zucker	-	[20,10,50,10]
// 	J - mit Milch und 2xZucker	-	[20,10,10,50]
// in ()-Klammern - Reihenfolge egal
func checkEingabe2(einwurf []uint) rune {
	//fmt.Println(einwurf)
	var N int = len(einwurf)
	if N == 4 {
		if einwurf[0] == 50 && einwurf[1] == 50 && einwurf[2] == 10 &&  einwurf[3] == 20 {
			return 'B'
		}
		if einwurf[0] == 50 && einwurf[1] == 50 && einwurf[2] == 10 &&  einwurf[3] == 10 {
			return 'R'
		}
		if einwurf[0] == 50 && einwurf[1] == 20 && einwurf[2] == 20 &&  einwurf[3] == 10 {
			return 'E'
		}
		if einwurf[0] == 50 && einwurf[1] == 10 && einwurf[2] == 20 &&  einwurf[3] == 20 {
			return 'L'
		}
		if einwurf[0] == 50 && einwurf[1] == 20 && einwurf[2] == 10 &&  einwurf[3] == 20 {
			return 'C'
		}
		if einwurf[0] == 20 && einwurf[1] == 10 && einwurf[2] == 10 &&  einwurf[3] == 50 {
			return 'J'
		}	
	}
	return 'X'
}


// Vor: Getränke mit IDs (rune) von B,R,E,L,C,J können richtig sein
//		(siehe oben).
// Erg: Der status ist geliefert: 1 = richtiges Getränk bestellt
// 		2 = falsches Getränk bestellt, 0 = zu wenig Münzen
//		3 = zuviele Münzen
func checkAutomat2(einwurf []uint,getraenkeID rune) uint {

	var N int = len(einwurf)
	var status uint
	var eingabeRune rune
	
	if N < 4 {status = 0}		// noch nicht genug Münzen
	if N > 4 {status = 3}		// zu viele Münzen

	if N == 4 {		// richtige Anzahl Münzen
		eingabeRune = checkEingabe2(einwurf)
		//fmt.Println(eingabeRune,getraenkeID)
		if eingabeRune == getraenkeID {
			status = 1			// richtiges Getränk
		} else if eingabeRune == 'X' {
			status = 4			// falsche Kombination
		} else {
			status = 2			// falsches Getränk
		}
	}


return status

}



// Vor: Getränke haben nur Nummern von 1 bis 6 (siehe oben).
//		Es gibt die Automaten 1 und 2.
// Erg: Der status ist geliefert: 1 = richtiges Getränk bestellt
// 		2 = falsches Getränk bestellt, 0 = zu wenig Münzen
//		4 = zuviele Münzen
func verarbeitungAutomat(einwurf []uint,getraenkeID rune,automatNr uint) uint {
	var erg uint
	if automatNr == 1 {
		erg = checkAutomat1(einwurf,getraenkeID)
	} else if automatNr == 2 {
		erg = checkAutomat2(einwurf,getraenkeID)
	} else{
		panic("Automat nicht bekannt!")
	}
	return erg
}



// Eff: Die Buttons sind aktiviert.
func aktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset buttons.Button) {
	zehner.AktiviereButton()
	zwanziger.AktiviereButton()
	fuenfziger.AktiviereButton()
	reset.AktiviereButton()	
}


// Eff: Die Buttons sind deaktiviert.
func deaktiviereSpielbutton(zehner,zwanziger,fuenfziger,reset buttons.Button) {
	zehner.DeaktiviereButton()
	zwanziger.DeaktiviereButton()
	fuenfziger.DeaktiviereButton()
	reset.DeaktiviereButton()
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Schriebt Note und Level ins Spielfeld.
func schreibeSpielstand(level uint16,punkte uint32, note float32) {
	gfx.SchreibeFont(850,10,"Level: " + fmt.Sprint(level))
	//gfx.SchreibeFont(150,15,"Punkte: " + fmt.Sprint(punkte))
	gfx.SchreibeFont(1100,10,"Note: " + fmt.Sprint(note))
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Die Münzen sind gezeichnet und die Anzahlen sind geschreiben.
func zeichneMuenzen(muenzenzahl [3]uint) {

	var x,y,r uint16		// Parameter für Kreis

	// zeichne Zehner
	x = 750
	y = 100
	r = 30
	gfx.Stiftfarbe(255,255,0)
	gfx.Vollkreis(x,y,r)
	gfx.Stiftfarbe(0,0,0)
	gfx.Kreis(x,y,r)
	gfx.Kreis(x,y,r-4)
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",30)
	gfx.SchreibeFont(x-17,y-17,"10")
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-65,y-12,fmt.Sprint(muenzenzahl[0])+" x")
	
	// zeichne Zwanziger
	x = 750
	y = 200
	r = 35
	gfx.Stiftfarbe(255,255,0)
	gfx.Vollkreis(x,y,r)
	gfx.Stiftfarbe(0,0,0)
	gfx.Kreis(x,y,r)
	gfx.Kreis(x,y,r-4)
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",35)
	gfx.SchreibeFont(x-19,y-20,"20")
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-70,y-12,fmt.Sprint(muenzenzahl[1])+" x")

	// zeichne Fünfziger
	x = 750
	y = 300
	r = 40
	gfx.Stiftfarbe(255,255,0)
	gfx.Vollkreis(x,y,r)
	gfx.Stiftfarbe(0,0,0)
	gfx.Kreis(x,y,r)
	gfx.Kreis(x,y,r-4)
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",40)
	gfx.SchreibeFont(x-21,y-23,"50")
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-75,y-12,fmt.Sprint(muenzenzahl[2])+" x")

}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Der Automat ist gezeichnet.
func zeichneAutomat() {

	// Automatenfarbe
	gfx.Stiftfarbe(139,69,19)		
	gfx.Vollrechteck(10,10,650,680)
	gfx.Stiftfarbe(0,0,0)		
	gfx.Rechteck(10,10,650,680)
	
	// Display
	gfx.LadeBild(550,30,"Bilder/Martin/getraenkeautomat/Display.bmp")
	gfx.Rechteck(550,30,91,400)

	// Reset-Button
	gfx.Stiftfarbe(255,255,0)		
	gfx.Vollrechteck(570,48,50,73)
	gfx.Stiftfarbe(0,0,0)		
	gfx.Rechteck(570,48,50,73)
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",50)
	gfx.SchreibeFont(577,55,"R")

	// Graph
	gfx.Stiftfarbe(255,255,255)		
	gfx.Vollrechteck(20,20,500,500)
	if automatNr == 1 {
		gfx.LadeBildMitColorKey (40,60,path+"Bilder/Martin/getraenkeautomat/Automat1.bmp",0,0,0)
	} else {
		gfx.LadeBildMitColorKey (30,100,path+"Bilder/Martin/getraenkeautomat/Automat2.bmp",0,0,0)
	}
	
	// Ausgabe
	gfx.LadeBild(70,530,"Bilder/Martin/getraenkeautomat/Ausschank.bmp")

	// Getränkeliste
	gfx.Stiftfarbe(255,211,155)		
	gfx.Vollrechteck(350,530,280,150)
	gfx.Stiftfarbe(0,0,0)		
	gfx.Rechteck(350,530,280,150)
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",20)
	gfx.Stiftfarbe(0,0,0)			
	var x,y uint16 = 360,535
	if automatNr == 1 {
		gfx.SchreibeFont(x,y,"A - " + getraenkeID2Name('A',1))
		gfx.SchreibeFont(x,y+24,"B - " + getraenkeID2Name('B',1))
		gfx.SchreibeFont(x,y+24*2,"C - " + getraenkeID2Name('C',1))
		gfx.SchreibeFont(x,y+24*3,"D - " + getraenkeID2Name('D',1))
		gfx.SchreibeFont(x,y+24*4,"E - " + getraenkeID2Name('E',1))
		gfx.SchreibeFont(x,y+24*5,"F - " + getraenkeID2Name('F',1))
	} else {
		gfx.SchreibeFont(x,y,"B - " + getraenkeID2Name('B',2))
		gfx.SchreibeFont(x,y+24,"R - " + getraenkeID2Name('R',2))
		gfx.SchreibeFont(x,y+24*2,"E - " + getraenkeID2Name('E',2))
		gfx.SchreibeFont(x,y+24*3,"L - " + getraenkeID2Name('L',2))
		gfx.SchreibeFont(x,y+24*4,"C - " + getraenkeID2Name('C',2))
		gfx.SchreibeFont(x,y+24*5,"K - " + getraenkeID2Name('K',2))		
	}	
	
}

// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Die Buttons sind gezeichnet.
func zeichneButtons(weiter,starter,beenden buttons.Button) {
	if weiter.GibAktivitaetButton() {
			weiter.ZeichneButton()
	}
	if starter.GibAktivitaetButton() {
			starter.ZeichneButton()
	}
	if beenden.GibAktivitaetButton() {
			beenden.ZeichneButton()
	}
}


// Vor: Ein passendes gfx-Fenster ist geöffnet.
// Eff: Das Spielfeld mit Automat, Münzen, Text, J.EthI, Buttons und
//	 	Spielstand ist gezeichnet.
func zeichneSpielfeld(ilevel uint16, punkte uint32, note float32,
		muenzenzahl [3]uint, weiter,starter,beenden buttons.Button) {

	var fontsize int = 20

	gfx.UpdateAus()

	// Säubere Hintergrund
	gfx.Stiftfarbe(255,255,255)		
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		

	// Dozentenbild
	gfx.LadeBild(840,10,path + "Bilder/Martin/getraenkeautomat/Jethi_klein.bmp")
	gfx.Linie(830,0,830,700-1)
	gfx.Linie(830,380,1200-1,380)

	zeichneAutomat()
	zeichneMuenzen(muenzenzahl)
	zeichneButtons(weiter,starter,beenden)

	// Spielstand
	gfx.SetzeFont (path + "Schriftarten/Ubuntu-B.ttf",fontsize)
	schreibeSpielstand(ilevel,punkte,note)

	// schreibe den Text in die Box
	txt.Zeichne()		
	
	gfx.UpdateAn()

}


// Vor: Ein gfx-Grafikfenster ist geöffnet.
// Eff: Hintergrundmusik ist gestartet. (Als go-Routine ausführen
//		damit das Spiel weitergeht.)
func hintergrundmusik(beenden buttons.Button) {
	var soundstr string = "Sounds/Music/getraenkeautomat.wav"
	for beenden.GibAktivitaetButton() {
		gfx.SpieleSound(soundstr)
		time.Sleep (time.Duration(19197e6))
	}
	//fmt.Println("Sound gestoppt")
}





