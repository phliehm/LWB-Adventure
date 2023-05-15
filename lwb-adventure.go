package main

// Annalena Cyriacus
// 24.04.2023

import ( . "gfx"
		"fmt"
		"./Klassen/vierecke"
		"./Klassen/objekte"
		"./MiniGames/4_Moorhuhn/moorhuhn"
		)

func maussteuerung (tueren *[]vierecke.Viereck, maus objekte.Objekt) {
	//var taste uint8
	//var status int8
	for {
		taste, status, mausX, mausY := MausLesen1()
		
		maus.SetzeKoordinaten(mausX,mausY)				// Aktualisiert Maus-Koordinaten
		
		if taste==1 && status==1 { 						//LINKE Maustaste gerade gedrückt
			for i,tuer := range *tueren { 				// Zeichnet alleweiteren Objekte ein
				if tuer.Angeklickt(mausX,mausY) {
						//TO DO
						switch i {
							case 0:
							//LadeBild(0,50,"./Bilder/Seminarraum_1.bmp")
							fmt.Println("Du hast ",moorhuhn.Moorhuhn()," Punkte erreicht!")
							case 1:
							LadeBild(0,50,"./Bilder/Seminarraum_2.bmp")
							case 2:
							LadeBild(0,50,"./Bilder/Seminarraum_3.bmp")
							case 3:
							LadeBild(0,50,"./Bilder/Seminarraum_4.bmp")
							case 4:
							LadeBild(0,50,"./Bilder/Tür5.bmp")
						}
				}
			}
		}
		/*
		if taste == 3 && status == 1 { 			//RECHTE Maustaste gerade gedrückt
			//TO DO
		}
		*/
	}
}


func main () {
	Fenster(1200,700)
	LadeBild(0,50,"./Bilder/mainfloor.bmp")
	
	var maus objekte.Objekt
	maus = objekte.New(0,0,0,0)						// Erstellt das Objekt MAUSZEIGER
	tueren := make([]vierecke.Viereck,0)			// Array für die Türen des Mainfloors
	var tuer1, tuer2, tuer3, tuer4, tuer5 vierecke.Viereck
	tuer1 = vierecke.New(45,235,40,595,220,595,220,235)
	tuer2 = vierecke.New(935,290,935,520,1010,545,1020,255)
	tuer3 = vierecke.New(425,330,430,470,465,460,460,350)
	tuer4 = vierecke.New(720,355,710,455,740,460,750,340)
	tuer5 = vierecke.New(570,350,570,435,625,435,625,350)
	
	//---
	/*
	sem4 := make([]vierecke.Viereck,0)			// Array für die Türen des Mainfloors
	var sem4Schm, sem4Amoe vierecke.Viereck
	sem4Schm = vierecke.New(45,45,600,45,600,600,45,600)			// KOORDINATEN ANPASSEN
	sem4Amoe = vierecke.New(935,290,935,520,1010,545,1020,255)
	*/
	//---
		
	Stiftfarbe(0,0,0)
	//tuer1.Zeichnen()
	//tuer2.Zeichnen()
	//tuer3.Zeichnen()
	//tuer4.Zeichnen()
	//tuer5.Zeichnen()
	tueren = append(tueren,tuer1,tuer2,tuer3,tuer4,tuer5)
	
	maussteuerung(&tueren,maus)
	
	
	TastaturLesen1()
}
