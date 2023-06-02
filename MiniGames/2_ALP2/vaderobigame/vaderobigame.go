//Annalena Cyriacus
//Juni 2023
//Minigame "Vaderobi" (ALP2-Game)

package vaderobigame

import (
	"fmt"
	"gfx"
	. "../vaderobi"
	"../../../Klassen/texteditoren"
)

var path string = "" //"../"
var path2 string = "./MiniGames/2_ALP2/"

func rotesKreuz() {						//rotes Kreuz oben links
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

//func Vaderobi() (float32,uint32) {
func Vaderobi() {
	
	var level [7]string
	level = [7] string {"","Welt_FU","Welt_Steps","Welt_robi","Welt_SCM","Welt_CYR","Welt_ALP"}
	var min [7]uint
	min = [7]uint {0,37,31,31,30,40,35}
	var minlaser [7]uint
	minlaser = [7]uint {0,1,3,3,2,1,2}
	var minmark [7]uint
	minmark = [7]uint {0,15,12,13,11,16,14}
	
	//---------------------------------------------------------------------
	for i:=1; i<7; i++ {
		
		gfx.Stiftfarbe(255,255,255)	
		gfx.Cls()
		//WeltLaden("Welt_Steps")
		WeltLaden(level[i])
		rotesKreuz()
		Schrittmodus(false)
		
		//var min uint = 31
		var laseranz uint
		var markanz uint
		var eingaben, fehler, korrekteBefehle uint
		var punkte uint
		var note float32 = 0.0
		
		//Titel
		gfx.Stiftfarbe(0,255,0)
		gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",38)
		gfx.SchreibeFont(700,5,"Super - ALP - Escape")
		
		//Aufgaben-Text
		gfx.Stiftfarbe(0,0,0)
		//gfx.SetzeFont(path + "Schriftarten/ltype.ttf",14)
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
		gfx.SchreibeFont(695,232,"Markiere jeden Deiner Schritte und hinterlasse einen")
		gfx.SchreibeFont(695,253,"Laser-Kreis an jeder Stelle, die Du entmauert hast.")
		gfx.Stiftfarbe(0,255,0)
		gfx.SchreibeFont(695,278,"Wenn Du dabei keine Fehler machst und den kürzesten Weg mit")
		gfx.SchreibeFont(695,299,"den wenigsten Befehlen findest, gibt es die Bestnote!")
		
		//Level
		gfx.Stiftfarbe(0,0,0)
		gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",30)
		gfx.SchreibeFont(695,320,"Level "+fmt.Sprint(i))
		
		//Zähler
		gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
		gfx.SchreibeFont(1020,380,"Eingaben:  "+fmt.Sprint(eingaben))
		gfx.Stiftfarbe(0,255,0)
		gfx.SchreibeFont(1020,440,"Punkte:    "+fmt.Sprint(punkte))
		gfx.Stiftfarbe(255,0,0)
		gfx.SchreibeFont(1020,410,"Fehler:    "+fmt.Sprint(fehler))
		
		var ted texteditoren.Texteditor
		//ted = texteditoren.New(700,350,475,325,20,true)
		ted = texteditoren.New(700,370,300,305,20,true)
		for {
			switch ted.GibString() {
				case "Laufen()":
				switch Laufen1() {
					case false:
					fehler++
					Melden("FEHLER: **Laufen()** NICHT MÖGLICH!")
					case true:
					korrekteBefehle++
					}
				case "LinksDrehen()":
				korrekteBefehle++
				LinksDrehen()
				case "RechtsDrehen()":
				korrekteBefehle++
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
				markanz++
				Markieren()
				case "Demarkieren()":
				korrekteBefehle++
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
					Melden("FEHLER: **Entlasern()** NICHT MÖGLICH!")
					case true:
					korrekteBefehle++
					laseranz--
					}
				case "HatLaserpower()":
				if HatKloetze() {
					Melden("Keine Sorge, ich habe noch Laserpower!")
				} else {
					Melden("Oh nein, ich habe alle Laserpower verschossen!")
				}
				case "Lasern()":
				switch Legen1() {
					case false:
					fehler++
					Melden("FEHLER: **Lasern()** NICHT MÖGLICH!")
					case true:
					korrekteBefehle++
					laseranz++
					}
				case "VorMauer()":
				VorMauer()
				case "Mauern()":
				switch Mauern1() {
					case false:
					fehler++
					Melden("FEHLER: **Mauern()** NICHT MÖGLICH!")
					case true:
					korrekteBefehle++
					}
				case "Entmauern()":
				switch Entmauern1() {
					case false:
					fehler++
					Melden("FEHLER: **Entmauern()** NICHT MÖGLICH!")
					case true:
					korrekteBefehle++
					}
				case "Baumodus()":
				Baumodus()
				default:
				fehler++
				Melden("Eingabefehler! -> Nochmal überlegen und ggf. Syntax prüfen!")
			}
			eingaben++
			punkte = 100+min[i] - eingaben - fehler
			//Zähler
			gfx.UpdateAus()
			gfx.Stiftfarbe(255,255,255)
			gfx.Vollrechteck(1020,380,180,100)
			gfx.Stiftfarbe(0,0,0)
			gfx.SetzeFont(path2 + "terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
			if eingaben > 9 {
				gfx.SchreibeFont(1020,380,"Eingaben: "+fmt.Sprint(eingaben))
			} else {
				gfx.SchreibeFont(1020,380,"Eingaben:  "+fmt.Sprint(eingaben))
			}
			gfx.Stiftfarbe(255,0,0)
			if fehler > 9 {
				gfx.SchreibeFont(1020,410,"Fehler:   "+fmt.Sprint(fehler))
			} else {
				gfx.SchreibeFont(1020,410,"Fehler:    "+fmt.Sprint(fehler))
			}
			gfx.Stiftfarbe(0,0,0)
					
			gfx.UpdateAn()
			
			//Level geschafft
			if InLinkerObererEcke() {
				Legen1()
				if laseranz < minlaser[i] {
					Melden("Du hast vergessen zu lasern oder geschummelt - das gibt Abzug in der B-Note!")
					punkte = punkte -10
				}
				if markanz < minmark[i] {
					Melden("Du hast vergessen zu markieren oder geschummelt - das gibt Abzug in der B-Note!")
					if markanz < minmark[i]/6 {
						punkte = punkte -35
					} else if markanz < minmark[i]/2 {
						punkte = punkte -20
					} else {
						punkte = punkte -15
					}
				}
				gfx.Stiftfarbe(0,255,0)
				if punkte > 99 {
					gfx.SchreibeFont(1020,440,"Punkte:  "+fmt.Sprint(punkte))
				} else if punkte > 9 {
					gfx.SchreibeFont(1020,440,"Punkte:   "+fmt.Sprint(punkte))
				} else {
					gfx.SchreibeFont(1020,440,"Punkte:    "+fmt.Sprint(punkte))
				}
				gfx.Stiftfarbe(0,0,0)
				gfx.Vollrechteck(150,225,375,225)
				gfx.Stiftfarbe(0,255,0)
				gfx.Vollrechteck(160,235,355,205)
				gfx.Stiftfarbe(0,0,0)
				gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",32)
				
				note = notenberechnung(punkte)
				if note == 0.0 {
					gfx.Stiftfarbe(255,0,0)
					gfx.Vollrechteck(160,235,355,205)
					gfx.Stiftfarbe(0,0,0)
					gfx.SchreibeFont(170,255,"Level geschafft!")
					gfx.SchreibeFont(200,300,"Aber leider")
					gfx.SchreibeFont(170,360,"nicht bestanden!!!")
				} else {
					gfx.Stiftfarbe(0,255,0)
					gfx.Vollrechteck(160,235,355,205)
					gfx.Stiftfarbe(0,0,0)
					gfx.SchreibeFont(275,255,"Yeah!!!")
					gfx.SchreibeFont(170,300,"Level geschafft!")
					gfx.SchreibeFont(250,360,"Note: " + fmt.Sprintf("%2.1f",note))
				}
				Melden("Level geschafft! :) Weiter geht's!")
				//Fertig()
				break
			}
			
			//gfx.UpdateAus()
			ted = texteditoren.New(700,370,300,305,20,true)
			//vaderobi.Melden("Neuer Texteditor!",0)
			//gfx.Stiftfarbe(255,255,255)
			//gfx.Vollrechteck(700,675,475,325)
			//gfx.Vollrechteck(700,665,475,4)
			//gfx.UpdateAn()
			
			//Fertig()
		}
	//----------------------------------------------------------------------------
	}
	Fertig()
	//gfx.TastaturLesen1()
	//fmt.Println(ted.GibString())
	
	//fmt.Println(ted.GibPosition())
	//return note, uint32(punkte)
}