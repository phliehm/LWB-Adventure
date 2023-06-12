//Annalena Cyriacus
//Juni 2023
//Minigame "Vaderobi" (ALP2-Game)

package vaderobigame

import (
	"fmt"
	"gfx"
	. "../vaderobi"
	"../../../Klassen/texteditoren"
	"../../../Klassen/vierecke"
)

var path string = "" //"../"
var path2 string = "./MiniGames/2_ALP2/"

//-------------------------Hilfsfunktionen------------------------------
//----------------------------------------------------------------------

//Darstellung rotes Kreuz oben links in vaderobis Welt
func rotesKreuz() {
	gfx.Stiftfarbe(255,0,0)
	gfx.Volldreieck(3,3,71,67,71,71)
	gfx.Volldreieck(3,3,7,3,71,71)
	gfx.Volldreieck(3,3,7,3,71,67)
	gfx.Volldreieck(3,3,67,71,71,71)
	gfx.Volldreieck(3,3,3,7,71,71)
	gfx.Volldreieck(3,3,3,7,67,71)
	
	gfx.Volldreieck(71,3,67,3,3,71)
	gfx.Volldreieck(71,3,3,67,3,71)
	gfx.Volldreieck(71,3,67,3,3,67)
	gfx.Volldreieck(71,3,71,7,3,71)
	gfx.Volldreieck(71,3,71,7,7,71)
	gfx.Volldreieck(71,3,3,71,7,71)
}

//Berechnung der Note aus Punktzahl 
func notenberechnung(punkte uint) float32 {
	if punkte >= 100 { return 1.0
	} else if punkte >= 90 { 
		return 1.3
	} else if punkte >= 80 { 
		return 1.7
	} else if punkte >= 75 {
		return 2.0
	} else if punkte >= 70 {
		return 2.3
	} else if punkte >= 65 {
		return 2.7
	} else if punkte >= 60 {
		return 3.0
	} else if punkte >= 55 {
		return 3.3
	} else if punkte >= 50 {
		return 4.0
	} else { return 0.0 }
}

// Darstellung der Aufgabentexte der verschiedenen Level
func aufgabentexte(n int) {																		//TODO
		
	switch n {
		case 0:
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,70,"Kleine Vorübung: Probiere aus, was VADEROBI alles kann! ")
		gfx.SchreibeFont(694,91,"Hier gibt es noch keine Punkte, also lass es krachen!")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-4.49.2.ttf",14)
		gfx.SchreibeFont(695,122,"Folgende Befehle kannst Du einzeln per Tastatur eingeben")
		gfx.SchreibeFont(695,141,"und den VADEROBI mit der ENTER-Taste ausführen lassen:")
		gfx.SchreibeFont(695,165,"Laufen()     LinksDrehen()     RechtsDrehen()")
		gfx.SchreibeFont(695,184,"Lasern()     Entlasern()       Markieren()")
		gfx.SchreibeFont(695,203,"Mauern()     Entmauern()       Demarkieren()")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,232,"Du kannst auch andere ALP2-robi-Befehle ausprobieren und")
		gfx.SchreibeFont(695,253,"versuchen Befehle mit kombinieren, z.B. LaufenUndLasern()!")
		gfx.Stiftfarbe(0,255,0)
		gfx.SchreibeFont(695,278,"Wenn Du Dich bereit fühlst mit Level 1 zu starten,")
		gfx.SchreibeFont(695,299,"gehe zum Ziel in der linken oberen Ecke des Spielfelds!")
		case 1:
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,70,"Lotse den VADEROBI mit möglichst wenigen Befehlen")
		gfx.SchreibeFont(694,91,"zum Ziel in der linken oberen Ecke des Spielfelds!")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-4.49.2.ttf",14)
		gfx.SchreibeFont(695,122,"Folgende Befehle kannst Du dafür nacheinander eingeben und")
		gfx.SchreibeFont(695,141,"den VADEROBI mit der ENTER-Taste ausführen lassen:")
		gfx.SchreibeFont(695,165,"Laufen()     LinksDrehen()     RechtsDrehen()")
		gfx.SchreibeFont(695,184,"Lasern()     Entlasern()       Markieren()")
		gfx.SchreibeFont(695,203,"Mauern()     Entmauern()       Demarkieren()")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,232,"Im ersten Level gibt es noch keine Zusatz-Aufgaben,")
		gfx.SchreibeFont(695,253,"aber auch hier geht es schon um Effizienz:")
		case 2:
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,70,"Lotse den VADEROBI mit möglichst wenigen Befehlen")
		gfx.SchreibeFont(694,91,"zum Ziel in der linken oberen Ecke des Spielfelds!")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-4.49.2.ttf",14)
		gfx.SchreibeFont(695,122,"Folgende Befehle kannst Du dafür nacheinander eingeben und")
		gfx.SchreibeFont(695,141,"den VADEROBI mit der ENTER-Taste ausführen lassen:")
		gfx.SchreibeFont(695,165,"Laufen()     LinksDrehen()     RechtsDrehen()")
		gfx.SchreibeFont(695,184,"Lasern()     Entlasern()       Markieren()")
		gfx.SchreibeFont(695,203,"Mauern()     Entmauern()       Demarkieren()")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,232,"Markiere auf dem Weg zum Ziel jeden Deiner Schritte!")
		gfx.SchreibeFont(695,253,"und denke an die maximale Effizienz (Kombis möglich!):")
		default:
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,70,"Lotse den VADEROBI mit möglichst wenigen Befehlen")
		gfx.SchreibeFont(694,91,"zum Ziel in der linken oberen Ecke des Spielfelds!")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-4.49.2.ttf",14)
		gfx.SchreibeFont(695,122,"Folgende Befehle kannst Du dafür nacheinander eingeben und")
		gfx.SchreibeFont(695,141,"den VADEROBI mit der ENTER-Taste ausführen lassen:")
		gfx.SchreibeFont(695,165,"Laufen()     LinksDrehen()     RechtsDrehen()")
		gfx.SchreibeFont(695,184,"Lasern()     Entlasern()       Markieren()")
		gfx.SchreibeFont(695,203,"Mauern()     Entmauern()       Demarkieren()")
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",16)
		gfx.SchreibeFont(695,232,"Markiere jeden Deiner Schritte UND hinterlasse einen")
		gfx.SchreibeFont(695,253,"Laser-Kreis an jeder Stelle, die Du entmauert hast.")
		//case 3:
		//case 4:
		//case 5:
		//case 6:
		//default:
	}
	
	if n != 0 {
		gfx.Stiftfarbe(0,255,0)
		gfx.SchreibeFont(695,278,"Wenn Du keine Fehler machst und den kürzesten Weg mit")
		gfx.SchreibeFont(695,299,"den wenigsten Befehlen findest, gibt es die Bestnote!")
	}
	
}


// eigentliche (exportierbare) Spiel-Funktion "Vaderobi()"
//----------------------------------------------------------------------
// Vor.: -
// Eff.: Das vaderobigame (Spiel "SUPER-ALP2-Escape") ist gestartet.
// Erg.: Die erspielte Gesamtpunktzahl (uint32) und die daraus berechnete Note (float32) ist geliefert.
func Vaderobi() (float32,uint32) {
	
	//Initialisierung der Slices für Level-Welten und Mindestanzahlen (für Befehle, Laser- und Markierungen)
	var level []string
	level = [] string {"Start-Welt","Welt_FU","Welt_Steps","Welt_robi","Welt_SCM","Welt_CYR","Welt_ALP"}
	var min []uint
	min = []uint {0,21,28,31,28,24,29}
	var minlaser []uint
	minlaser = []uint {0,0,0,3,2,1,2}
	var minmark []uint
	minmark = []uint {0,0,12,13,11,16,14}
	
	//Initialisierung der Variablen und Speicher-Slices für Punkte und Noten
	var punktespeicher []uint
	punktespeicher = make([]uint,len(level))
	var notenspeicher []float32
	notenspeicher = make([]float32,len(level))
	var gesamtpunkte uint32
	var gesamtnote float32
	//var maxpunktzahl uint (optional, ggf. TODO)
	
	//Initialisierung des klickbaren Beenden-Symbols (für End-Bildschirm)
	var exit vierecke.Viereck = vierecke.New(1080,30,1080,145,1170,145,1170,30)
	
	//---------------------------------------------------------------------
	
	//Initialisiere Darstellung der vaderobi-Welt
	WeltOeffnen()
	
	//Spiel-Schleife
	for {
		
		//Level-Schleife	
		for i:=0; i<len(level); i++ {
						
			//Initialisierung verschiedener Steuer- und Speicher-Variablen
			var laseranz, markanz, eingaben, fehler, korrekteBefehle uint
			var punkte, punktabzug uint
			var note float32 = 0.0
			
			
			//Zeichne Spielfeld
			//---------------------------------------------------------------------------
			
			gfx.UpdateAus()
			// erneuere Spiel-Fenster
			gfx.Stiftfarbe(255,255,255)
			gfx.Cls()
			
			//aktuelle Level-Welt laden
			WeltLaden(level[i])
			rotesKreuz()
			Schrittmodus(false)
			
			//Titel
			gfx.Stiftfarbe(0,255,0)
			gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",38)
			gfx.SchreibeFont(700,5,"Super - ALP - Escape")
			
			//Aufgaben-Text
			aufgabentexte(i)
			
			//Beenden-Hinweis
			gfx.Stiftfarbe(0,0,0)
			gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-4.49.2.ttf",14)
			gfx.SchreibeFont(1020,640,"Beenden des Spiels:")
			gfx.SchreibeFont(1020,660,"Tippe exit und ENTER!")
			
			
			//Level
			gfx.Stiftfarbe(0,0,0)
			gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",30)
			gfx.SchreibeFont(695,320,"Level "+fmt.Sprint(i))
			
			//Zähler
			gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
			gfx.SchreibeFont(1020,380,"Eingaben:   "+fmt.Sprint(eingaben))
			//gfx.Stiftfarbe(0,255,0)
			//gfx.SchreibeFont(1020,440,"Punkte:     "+fmt.Sprint(punkte))
			gfx.Stiftfarbe(255,0,0)
			gfx.SchreibeFont(1020,410,"Fehler:     "+fmt.Sprint(fehler))
			
			//Sound
			if i == 0 {
				gfx.SpieleSound("./Sounds/lordvaderrise.wav")
			} else {
				gfx.SpieleSound("./Sounds/imperial_march.wav")
			}
						
			gfx.UpdateAn()
			
			//Texteditor
			var ted texteditoren.Texteditor = texteditoren.New(700,370,300,305,20,true)
			
			//Eingabe-Verarbeitung
			//--------------------------------------------------------------------------
			for {
					
				//vorzeitiges Beenden	
				if ted.GibString() == "exit" {
					i = len(level)
					gfx.StoppeAlleSounds()
					break
				}
				
				//Fallunterscheidung Eingabe		
				switch ted.GibString() {
					case "Laufen()":
					switch Laufen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Laufen()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_impact11.wav")
						korrekteBefehle++
						}
					case "LaufenUndMarkieren()":
					switch Laufen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **LaufenUndMarkieren()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_impact11.wav")
						gfx.SpieleSound("./Sounds/vader_breathing.wav")
						markanz++
						Markieren()
						korrekteBefehle++
						}
					case "MarkierenUndLaufen()":
					gfx.SpieleSound("./Sounds/vader_breathing.wav")
					markanz++
					Markieren()
					switch Laufen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Laufen...()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_impact11.wav")
						korrekteBefehle++
						}
					case "LaufenUndLasern()":
					switch Laufen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Laufen...()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_impact11.wav")
						gfx.SpieleSound("./Sounds/vader_breathing.wav")
						laseranz++
						Legen1()
						korrekteBefehle++
						}
					case "LasernUndLaufen()":
					gfx.SpieleSound("./Sounds/vader_breathing.wav")
					laseranz++
					Legen1()
					switch Laufen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Laufen...()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_impact11.wav")
						korrekteBefehle++
						}
					case "LinksDrehen()":
					korrekteBefehle++
					gfx.SpieleSound("./Sounds/sfx_sounds_impact2.wav")
					LinksDrehen()
					case "RechtsDrehen()":
					korrekteBefehle++
					gfx.SpieleSound("./Sounds/sfx_sounds_impact2.wav")
					RechtsDrehen()
					case "AmRand()":
					if AmRand() {
						Melden("Ich stehe am Rand - siehst Du doch!")
					} else {
						Melden("Ich stehe NICHT am Rand - siehst Du doch!")
					}
					case "InLinkerObererEcke()":
					if InLinkerObererEcke() {
						Melden("Ich stehe in der linken oberen Ecke - siehst Du doch!")
					} else {
						Melden("Ich stehe NICHT in der linken oberen Ecke - siehst Du doch!")
					}
					case "Fertig()":
					Fertig()
					case "Markieren()":
					korrekteBefehle++
					gfx.SpieleSound("./Sounds/vader_breathing.wav")
					markanz++
					Markieren()
					case "Demarkieren()":
					korrekteBefehle++
					gfx.SpieleSound("./Sounds/vader_breathing.wav")
					markanz--
					Demarkieren()
					case "Markiert()":
					if Markiert() {
						Melden("Mein Feld ist markiert - siehst Du doch!")
					} else {
						Melden("Mein Feld ist NICHT markiert - siehst Du doch!")
					}
					case "NachbarMarkiert()":
					if NachbarMarkiert() {
						Melden("Das Feld vor mir ist markiert - siehst Du doch!")
					} else {
						Melden("Das Feld vor mir ist NICHT markiert - siehst Du doch!")
					}
					case "Leer()":
					if Leer() {
						Melden("Mein Feld ist leer - siehst Du doch!")
					} else {
						Melden("Mein Feld ist NICHT leer - siehst Du doch!")
					}
					case "Entlasern()":
					switch Leeren1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Entlasern()** NICHT MÖGLICH!")
						case true:
						korrekteBefehle++
						gfx.SpieleSound("./Sounds/light-saber-off.wav")
						laseranz--
						}
					case "HatKloetze()":
					if HatKloetze() {
						Melden("Keine Sorge, ich habe noch Laserpower!")
					} else {
						Melden("Oh nein, ich habe alle Laserpower verschossen!")
					}
					case "Lasern()":
					switch Legen1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Lasern()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/light-saber-on.wav")
						korrekteBefehle++
						laseranz++
						}
					case "VorMauer()":
					VorMauer()
					case "Mauern()":
					switch Mauern1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Mauern()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_movement_portal1.wav")
						korrekteBefehle++
						}
					case "Entmauern()":
					switch Entmauern1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Entmauern()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_falling6.wav")
						korrekteBefehle++
						}
					case "EntmauernUndLasern()":
					switch Entmauern1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **EntmauernUndLasern()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_falling6.wav")
						Legen1()
						gfx.SpieleSound("./Sounds/sfx_movement_portal1.wav")
						laseranz++
						korrekteBefehle++
						}
					case "LasernUndEntmauern()":
					Legen1()
					gfx.SpieleSound("./Sounds/sfx_movement_portal1.wav")
					laseranz++
					switch Entmauern1() {
						case false:
						fehler++
						gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
						Melden("FEHLER: **Entmauern...()** NICHT MÖGLICH!")
						case true:
						gfx.SpieleSound("./Sounds/sfx_sounds_falling6.wav")
						korrekteBefehle++
						}
					case "Baumodus()":														//Cheat-Modus
					gfx.SpieleSound("./Sounds/ooh_sw_luke_dontdothat.wav")
					Baumodus()
					default:
					fehler++
					gfx.SpieleSound("./Sounds/sw_luke_dontdothat.wav")
					Melden("Eingabefehler! -> Nochmal überlegen und ggf. Syntax prüfen!")
				}
				
				if i > 0 {
					//Punkte-Berechnung
					eingaben++
					punkte = 100+min[i]
					if punkte >= eingaben {
						punkte = punkte - eingaben
					} else {
						punkte = 0
					}
					if punkte >= fehler {
						punkte = punkte - fehler
					} else {
						punkte = 0
					}
					
					//Zeichne Zähler
					gfx.UpdateAus()
					gfx.Stiftfarbe(255,255,255)
					gfx.Vollrechteck(1020,380,180,100)
					gfx.Stiftfarbe(0,0,0)
					gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
					if eingaben > 9 {
						gfx.SchreibeFont(1020,380,"Eingaben:  "+fmt.Sprint(eingaben))
					} else {
						gfx.SchreibeFont(1020,380,"Eingaben:   "+fmt.Sprint(eingaben))
					}
					gfx.Stiftfarbe(255,0,0)
					if fehler > 9 {
						gfx.SchreibeFont(1020,410,"Fehler:    "+fmt.Sprint(fehler))
					} else {
						gfx.SchreibeFont(1020,410,"Fehler:     "+fmt.Sprint(fehler))
					}
					gfx.Stiftfarbe(0,0,0)
							
					gfx.UpdateAn()
				}
				
				//prüfe, ob Level geschafft
				if InLinkerObererEcke() {
					Legen1()
					gfx.SpieleSound("./Sounds/laser_all2easy.wav")
					Melden("Ziel erreicht!")
					
					//Prüfung auf Lasern, wo entmauert wurde
					if laseranz < minlaser[i] {
						gfx.SpieleSound("./Sounds/ooh.wav")
						Melden("Du hast vergessen zu lasern oder geschummelt - das gibt Abzug in der B-Note!")
						if punkte >= 10 {
							punkte = punkte -10
							punktabzug = punktabzug + 10
						} else {
							punkte = 0
							punktabzug = punktabzug + punkte
						}
						
					}
					
					//Prüfung auf Markieren
					if markanz < minmark[i] {
						gfx.SpieleSound("./Sounds/ooh2.wav")
						Melden("Du hast vergessen zu markieren oder geschummelt - das gibt Abzug in der B-Note!")
						if markanz < minmark[i]/6 && punkte >= 35 {
							punkte = punkte -35
							punktabzug = punktabzug + 35
						} else if markanz < minmark[i]/2 && punkte >= 20 {
							punkte = punkte -20
							punktabzug = punktabzug + 20
						} else if punkte >= 15 {
							punkte = punkte -15
							punktabzug = punktabzug + 15
						} else {
							punktabzug = punktabzug + punkte
							punkte = 0
						}
					}
					
					//Schreibe Punkte und Abzüge
					gfx.Stiftfarbe(255,0,0)
					gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
					if punktabzug > 99 {
						gfx.SchreibeFont(1020,440,"Abzüge:   "+fmt.Sprint(punktabzug))
					} else if punktabzug > 9 {
						gfx.SchreibeFont(1020,440,"Abzüge:    "+fmt.Sprint(punktabzug))
					} else {
						gfx.SchreibeFont(1020,440,"Abzüge:     "+fmt.Sprint(punktabzug))
					}
					
					gfx.Stiftfarbe(0,255,0)
					if punkte > 99 {
						gfx.SchreibeFont(1020,470,"Punkte:   "+fmt.Sprint(punkte))
					} else if punkte > 9 {
						gfx.SchreibeFont(1020,470,"Punkte:    "+fmt.Sprint(punkte))
					} else {
						gfx.SchreibeFont(1020,470,"Punkte:     "+fmt.Sprint(punkte))
					}
					
					if i > 0 {
						//Level-Punkte in Punktespeicher schreiben und Level-Note berechnen
						punktespeicher[i] = punkte
						note = notenberechnung(punkte)
						notenspeicher[i] = note
						gesamtpunkte = gesamtpunkte + uint32(punkte)
						gesamtnote = notenberechnung(uint(gesamtpunkte)/uint(len(notenspeicher)-1))
						fmt.Println(gesamtnote)
						
						//Zeichne Level-Abschluss-Meldung
						gfx.Stiftfarbe(0,0,0)
						gfx.Vollrechteck(150,225,375,225)
						gfx.Stiftfarbe(0,255,0)
						gfx.Vollrechteck(160,235,355,205)
						gfx.Stiftfarbe(0,0,0)
						gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",32)
					
						if note == 0.0 {												//TODO: Level wiederholen!!!
							gfx.Stiftfarbe(255,0,0)
							gfx.Vollrechteck(160,235,355,205)
							gfx.Stiftfarbe(0,0,0)
							gfx.SchreibeFont(170,255,"Level geschafft!")
							gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",30)
							gfx.SchreibeFont(230,315,"Aber leider")
							gfx.SchreibeFont(172,360,"nicht bestanden!!!")
							gfx.SpieleSound("./Sounds/vader_breathing.wav")
						} else {
							gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",32)
							gfx.Stiftfarbe(0,255,0)
							gfx.Vollrechteck(160,235,355,205)
							gfx.Stiftfarbe(0,0,0)
							gfx.SchreibeFont(275,255,"Yeah!!!")
							gfx.SchreibeFont(170,300,"Level geschafft!")
							gfx.SchreibeFont(250,360,"Note: " + fmt.Sprintf("%2.1f",note))
							gfx.SpieleSound("./Sounds/swsidious_youhavebeenwelltrained.wav")
						}
						Melden("Level geschafft! :) Weiter geht's!")
					}
					break
				}
				
				//neuer Texteditor für nächste Eingabe
				ted = texteditoren.New(700,370,300,305,20,true)
				
			}
			gfx.StoppeAlleSounds()
		//----------------------------------------------------------------------------
		}
			
		//----------------- Endbildschirm --------------------------------------------
		gfx.Stiftfarbe(255,255,255)										// Lösche Inhalt des gfx-Fensters
		gfx.Cls()
		
		gfx.SpieleSound(path + "Sounds/the_force.wav")
		
		//Grafik (Bilder)
		gfx.LadeBild(150,100,path + "Bilder/vaderobiGame/sprechblase_flipped_400.bmp")
		gfx.LadeBildMitColorKey(100,350,path + "Bilder/vaderobiGame/Darth_200.bmp",255,255,255)
		gfx.LadeBild(620,80,path + "Bilder/vaderobiGame/paper_500.bmp")
		gfx.LadeBild(960,520,path + "Bilder/vaderobiGame/certified_100.bmp")
		gfx.LadeBild(1080,30,path + "Bilder/vaderobiGame/Zurück-Symbol.bmp")
		exit.AktiviereKlickbar()
		fmt.Println("exit aktiviert")
		
		//Überschrift	
		gfx.Stiftfarbe(0,255,0)
		gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
		gfx.SchreibeFont(330,10,"Super - ALP - Escape")
		
		//Sprechblase Darth Schmidter
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
		
		if gesamtnote == 0.0 {
			gfx.SchreibeFont(295,140,"Du hast die")
			gfx.SchreibeFont(285,170,"Prüfung leider")
			gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",32)
			gfx.Stiftfarbe(255,0,0)
			gfx.SchreibeFont(310,195,"Nicht")
			gfx.SchreibeFont(260,235,"Bestanden!")
		} else {		
			gfx.SchreibeFont(295,140,"Du hast die")
			gfx.SchreibeFont(310,260,"erreicht!")
			gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
			gfx.SchreibeFont(285,170,"Gesamtnote")
			gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
			gfx.Stiftfarbe(0,255,0)
			gfx.SchreibeFont(325,195,fmt.Sprintf("%2.1f",gesamtnote))
		}
		
		//Inhalt Abschluss-Zertifikat für Mini-Game
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
		
		for i:=1; i<len(level); i++ {
			gfx.SchreibeFont(710,150+uint16((i-1)*68), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(punktespeicher[i]) + " Punkte")
			gfx.SchreibeFont(710,175+uint16((i-1)*68),"           Note " + fmt.Sprintf("%2.1f",notenspeicher[i]))
		}
		gfx.SchreibeFont(700,130+uint16(6*70),"----------------------")
		gfx.SchreibeFont(710,160+uint16(6*70),"Gesamt:    " + fmt.Sprint(gesamtpunkte) + " Punkte")	
		
		//Mauslese-Schleife für Klick auf Exit-Symbol
		for {
			taste, status, mausX, mausY := gfx.MausLesen1()
			if taste==1 && status==1 {
				if exit.Angeklickt(mausX,mausY) { 							// Ende des Spiels
					//fmt.Println("exit geklickt")
					gfx.StoppeAlleSounds()
					break
				}
			}
		}
		break
	}		
	
	//Rückgabe von Note und Gesamtpunktzahl
	return gesamtnote, gesamtpunkte

}
