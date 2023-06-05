package eingabe

import (
	"gfx"
	"../Klassen/vierecke"
	"../darstellung"
)

// Globale Variablen
// ------------------

var klickbar [][]vierecke.Viereck = make([][]vierecke.Viereck,0)
//var klickElemente0 []vierecke.Viereck = make([]vierecke.Viereck,0)


func klickbarElemente() {
	
	var tuer1, tuer2, tuer3, tuer4, tuer5, exit, ende vierecke.Viereck
	tuer1 = vierecke.New(45,235,40,595,220,595,220,235)
	tuer2 = vierecke.New(935,290,935,520,1010,545,1020,255)
	tuer3 = vierecke.New(425,330,430,470,465,460,460,350)
	tuer4 = vierecke.New(720,355,710,455,740,460,750,340)
	tuer5 = vierecke.New(570,350,570,435,625,435,625,350)
	exit = vierecke.New(1100,565,1190,565,1190,685,1100,685)
	ende = vierecke.New(1080,445,1075,555,1130,570,1135,450)
	
	klickbar[0] = append(klickbar[0],ende,tuer1,tuer2,tuer3,tuer4,tuer5)
	klickbar[1] = append(klickbar[1],exit)
	klickbar[2] = append(klickbar[2],exit)
	klickbar[3] = append(klickbar[3],exit)
	klickbar[4] = append(klickbar[4],exit)
	klickbar[5] = append(klickbar[5],exit)
}

	
func maussteuerung (raumnr int) {
	//var taste uint8
	//var status int8
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
				
		if taste==1 && status==1 { 										//LINKE Maustaste gerade gedrückt
						
			for index,element := range klickbar[raumnr] { 				//enthält alle klickbaren Objekte
				if element.Angeklickt(mausX,mausY) {
						
						switch raumnr {
							case 0:
							raumnr = index
							darstellung.SemesterraumDarstellen(index)
							default:
							if index == 0 {
								darstellung.MainfloorDarstellen()
							} else {
								//darstellung.MinigameLaden(index)
							}
						}
						
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


func Eingabe() {
	
	klickbarElemente()
	maussteuerung(0)
}
