package eingabe

import (
	"gfx"
	"../Klassen/vierecke"
	"../darstellung"
)

// Globale Variablen
// ------------------

//var klickbar [][]vierecke.Viereck = make([][]vierecke.Viereck,0)
var klickbar [][]vierecke.Viereck = make([][]vierecke.Viereck,6)
//var klickElemente0 []vierecke.Viereck = make([]vierecke.Viereck,0)


func klickbarElemente() {
	
	var ende, exit, tuer1, tuer2, tuer3, tuer4, tuer5 vierecke.Viereck
	var fabweb1, wthek1, darth2, jethi2, herk3, wthek3, darth4, amoebi4, wthek4 vierecke.Viereck
	
	ende = vierecke.New(1080,495,1075,615,1130,620,1135,500)
	exit = vierecke.New(1100,565,1190,565,1190,685,1100,685)
	
	fabweb1 = vierecke.New(900,265,900,565,1045,565,1045,265)
	wthek1 = vierecke.New(85,265,85,560,255,560,255,265)
	darth2 = vierecke.New(130,405,130,680,280,680,280,405)
	jethi2 = vierecke.New(800,435,785,495,880,500,900,430)
	herk3 = vierecke.New(100,270,70,415,175,385,160,270)
	wthek3 = vierecke.New(560,245,560,495,705,495,705,245)
	darth4 = vierecke.New(205,140,210,490,385,490,385,140)
	amoebi4 = vierecke.New(1040,315,1010,405,1155,405,1080,295)
	wthek4 = vierecke.New(655,195,655,515,845,515,845,195)
	
	tuer1 = vierecke.New(45,235,40,595,220,595,220,235)
	tuer2 = vierecke.New(935,290,935,520,1010,545,1020,255)
	tuer3 = vierecke.New(425,330,430,470,465,460,460,350)
	tuer4 = vierecke.New(720,355,710,455,740,460,750,340)
	tuer5 = vierecke.New(570,350,570,435,625,435,625,350)
	
	
	klickbar[0] = append(klickbar[0],ende,tuer1,tuer2,tuer3,tuer4,tuer5)
	klickbar[1] = append(klickbar[1],exit,fabweb1,wthek1)
	klickbar[2] = append(klickbar[2],exit,darth2,jethi2)
	klickbar[3] = append(klickbar[3],exit,herk3,wthek3)
	klickbar[4] = append(klickbar[4],exit,darth4,amoebi4,wthek4)
	klickbar[5] = append(klickbar[5],exit)
}

	
func maussteuerung (raumnr int) {
	//var taste uint8
	//var status int8
	
	/*for _,el := range klickbar[raumnr] {
		el.SetzeFarbe(0,0,0)
		el.Zeichnen()
	}*/
	
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
				
		if taste==1 && status==1 { 										//LINKE Maustaste gerade gedrückt
						
			for index,element := range klickbar[raumnr] { 				//enthält alle klickbaren Objekte im jeweiligen Raum
				
				element.AktiviereKlickbar()
								
				if element.Angeklickt(mausX,mausY) {					//wenn ein Element angeklickt wurde, prüfe wo und welches:
						
						switch raumnr {
							case 0:										//wenn im mainfloor (raumnr 0):
							element.DeaktiviereKlickbar()
							raumnr = index								//neue raumnr ist index des geklickten Elements (hier der angeklickten Tür, entspricht dem Semester)
							darstellung.SemesterraumDarstellen(index)	//also wird der jeweilige Semesterraum dargestellt
							
							/*for _,el := range klickbar[raumnr] {
								el.SetzeFarbe(0,0,0)
								el.Zeichnen()
							}*/
							
							default:									//wenn nicht im mainfloor (raumnr != 0):
							element.DeaktiviereKlickbar()
							if index == 0 {								//Element mit index 0 wurde geklickt, also "exit", d.h. ...
								raumnr = index							//... zurück in den mainfloor (raumnr 0)
								darstellung.MainfloorDarstellen()		//deshalb mainfloor darstellen
								
								/*for _,el := range klickbar[raumnr] {
									el.SetzeFarbe(0,0,0)
									el.Zeichnen()
								}*/
								
							} else {									//wenn nicht "exit" (index 0) geklickt wurde,
								darstellung.MinigameLaden(raumnr,index)	//dem angeklickten Element (Dozenten) zugehöriges Spiel starten
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
