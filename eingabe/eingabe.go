// Autor: A. Cyriacus und M. Seiß
// Datum: 12.06.2023
// Zweck: Implementierung des ADO eingabe

package eingabe

import (
	"gfx"
	"../Klassen/vierecke"
	"../darstellung"
	"fmt"
	"../Klassen/spielstaende"
)


// Globale Variablen
// ------------------

var klickbar [][]vierecke.Viereck = make([][]vierecke.Viereck,7)
var spielstand spielstaende.Spielstand



func klickbarElemente() {
		
	var ende, exit, tuer1, tuer2, tuer3, tuer4, tuer5, info vierecke.Viereck
	var fabweb1, wthek1, darth2, jethi2, herk3, wthek3, darth4, amoebi4, wthek4 vierecke.Viereck
	var beenden vierecke.Viereck
	
	ende = vierecke.New(1080,450,1080,615,1180,620,1180,450)
	exit = vierecke.New(1100,565,1190,565,1190,685,1100,685)
	beenden = vierecke.New(350,620,550,620,550,670,350,670)
	
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
	info = vierecke.New(1043,235,1043,365,1153,365,1153,235)			// vergrößert 
	
	klickbar[0] = append(klickbar[0],ende,tuer1,tuer2,tuer3,tuer4,tuer5,info)
	klickbar[1] = append(klickbar[1],exit,fabweb1,wthek1)
	klickbar[2] = append(klickbar[2],exit,darth2,jethi2)
	klickbar[3] = append(klickbar[3],exit,herk3,wthek3)
	klickbar[4] = append(klickbar[4],exit,darth4,amoebi4,wthek4)
	klickbar[5] = append(klickbar[5],exit)								// Raum 5 - Zertifikat
	klickbar[6] = append(klickbar[6],exit,beenden)						// Mülleimer - Ende
	
}

	
func maussteuerung (raumnr int) {
	
	var note float32
	var punkte uint32
	var nschluessel uint16 = gibSchluesselzahl(spielstand)
	var hoSigned bool 													// Hausordnung unterschrieben?

	fmt.Println("Anzahl Schluessel: ",nschluessel)
	
	
A:	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
				
		if taste==1 && status==1 { 										//LINKE Maustaste gerade gedrückt
						
			for index,element := range klickbar[raumnr] { 				//enthält alle klickbaren Objekte im jeweiligen Raum
				
				element.AktiviereKlickbar()
								
				if element.Angeklickt(mausX,mausY) {					//wenn ein Element angeklickt wurde, prüfe wo und welches:
						
						switch raumnr {
							case 0:										//wenn im mainfloor (raumnr 0):
							element.DeaktiviereKlickbar()				//im vorherigen Raum anklickbare Elemente sind deaktiviert
							raumnr = index								//neue raumnr ist index des geklickten Elements (hier der angeklickten Tür, entspricht dem Semester)
							if raumnr == 6 {
								darstellung.InfoDarstellen()
								raumnr = 0
								darstellung.MainfloorDarstellen()
								
							} else if raumnr == 5 {
								if nschluessel == 5	{
									fmt.Println("5. Tür angeklickt")
									darstellung.EndbildschirmDarstellen(spielstand)
								} else {								// Raum zu, kein Schlüssel da
									raumnr = 0
									gfx.SpieleSound("Sounds/Beep.wav")
								}
								
							} else if raumnr == 0 { 					// Spiel beenden
								fmt.Println("Mülleimer angeklickt")
								raumnr = 6
								darstellung.SpielVerlassenDarstellen(spielstand)
								
							} else {
								if nschluessel >= uint16(index)	{
									
									if raumnr == 4 && !hoSigned {										//wenn die Semester-Tür 4 angeklickt wurde,
									
											fmt.Println("Noch nicht unterschrieben!")
											sign, no := darstellung.HeidiDarstellen()					//Heidi mit Bubble und Unterschrift-Aufforderung darstellen

											//Mauslese-Schleife							
											for {
												taste, status, mausX, mausY := gfx.MausLesen1()
												if taste==1 && status==1 {									
													if sign.TesteXYPosInButton(mausX,mausY) {			//wenn der sign-Button angeklickt wurde,
														darstellung.SemesterraumDarstellen(index)
														hoSigned = true
														break											//Mauslese-Schleife verlassen und weiter in den Semesterraum
													} else if no.TesteXYPosInButton(mausX,mausY) {		//wenn der no-Button angeklickt wurde,
														raumnr = 0
														darstellung.MainfloorDarstellen()				//aus Mauslese-Schleife und zurück in den mainfloor
														break													
													}
												}
											} 
									} else  {
											fmt.Println("Raum darstellen: ",index)
											darstellung.SemesterraumDarstellen(index)					//also wird der jeweilige Semesterraum dargestellt
									} 
									
								} else { 																// Raum zu, kein Schlüssel da
									raumnr = 0
									gfx.SpieleSound("Sounds/Beep.wav")
								}
							}

							case 6:										// Spiel verlassen?
							element.DeaktiviereKlickbar()
							if index == 0 {								//Element mit index 0 wurde geklickt, also "exit", d.h. ...
								raumnr = index							//... zurück in den mainfloor (raumnr 0)
								fmt.Println("Zurück zum Mainfloor!")
								darstellung.MainfloorDarstellen()		//deshalb mainfloor darstellen
							} else if index == 1 {						//wenn index 1, beenden, geklickt wurde,
								break A									//deshalb Spiel beendet
							}

							
							default:									//wenn nicht im mainfloor (raumnr != 0):
							element.DeaktiviereKlickbar()
							if index == 0 {								//Element mit index 0 wurde geklickt, also "exit", d.h. ...
								raumnr = index							//... zurück in den mainfloor (raumnr 0)
								darstellung.MainfloorDarstellen()		//deshalb mainfloor darstellen
								
							} else {									//wenn nicht "exit" (index 0) geklickt wurde,
							
								start,no := darstellung.BubbleLaden(raumnr,index)
								
								//Mauslese-Schleife							
								for {
									taste, status, mausX, mausY := gfx.MausLesen1()
									if taste==1 && status==1 {									
										if start.TesteXYPosInButton(mausX,mausY) {					//und der Start-Button angeklickt wurde,
											note,punkte = darstellung.MinigameLaden(raumnr,index)	//dem angeklickten Element (Dozenten) zugehöriges Spiel starten
											speichereMax(note,punkte,raumnr,index)					// Das beste Ergebnis wird gespeichert.
											nschluessel = gibSchluesselzahl(spielstand) 			// Bestimme die neue Schluesselzahl
											break
										} else if no.TesteXYPosInButton(mausX,mausY) {
											break
										}
									}
								}
								darstellung.SemesterraumDarstellen(raumnr)
							}
						}
						
					}
				}	
			}
		}
}


func Eingabe() {

	gfx.Fenster(1200,700)

	spielstand = darstellung.StartFenster()
	darstellung.MainfloorDarstellen()
	
	klickbarElemente()
	maussteuerung(0)
	
}




//  ----------------      Hilfsfunktionen     ---------------------  //


// Erg: Die Anzahl der Schlüssel = Anzahl der geöffneten Räume ist
//		geliefert.
func gibSchluesselzahl(spielstand spielstaende.Spielstand) uint16 {

	var noten []float32 = spielstand.GibNoten()
	var nschluessel uint16 = 1

	if noten[0] < 4.1 && noten[1] < 4.1 && noten[0] > 0.1 && noten[1] > 0.1 {nschluessel++}
	if noten[3] < 4.1 && noten[4] < 4.1 && noten[3] > 0.1 && noten[4] > 0.1 {nschluessel++}
	if noten[6] < 4.1 && noten[7] < 4.1 && noten[6] > 0.1 && noten[7] > 0.1 {nschluessel++}
	if noten[9] < 4.1 && noten[10] < 4.1 && noten[11] < 4.1 &&
		noten[9] > 0.1 && noten[10] > 0.1 && noten[11] > 0.1 {nschluessel++}
	
	return nschluessel

}



// Eff: Die bessere Note ist gespeichert. Bei Notengleichstand ist die
//		bessere Punktzahl gespeichert.
func speichereMax(note1 float32,punkte1 uint32,raumnr,index int)	{

	var noten []float32 = spielstand.GibNoten()
	var punkte []uint32 = spielstand.GibPunkte()
	var ispiel int = (raumnr-1)*3 + (index-1)	// raumnr>0,ispiel>0
	
	fmt.Println(raumnr,index)
	fmt.Println("Erspieltes Ergebnis: ",note1,punkte1)
	fmt.Println("Alter Spielstand: ",spielstand.GibNoten(),spielstand.GibPunkte())
	fmt.Println("Spielnr:", ispiel)
	
	// erstes Spiel? oder bessere Note? oder bessere Punkte?
	if noten[ispiel] == 0 || note1 < noten[ispiel] && note1 > 0.1 || 
			(note1 == noten[ispiel] && punkte1 > punkte[ispiel]) {
		noten[ispiel] = note1
		punkte[ispiel] = punkte1
		spielstand.Speichern(noten,punkte)
	}

	fmt.Println("Neuer Spielstand: ",spielstand.GibNoten(),spielstand.GibPunkte())

}
