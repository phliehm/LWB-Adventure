// the NET Game - Hauptfunktion und Hilfsfunktionen

// Martin Seiß    24.4.2023


package theNETgame


import "fmt"
import "gfx"
import "../../Klassen/buttons"
//import "os"
//import "strconv"
import "time"
import . "../../Klassen/graphen"
import "math/rand" 

const nlevel uint16 = 9				// Anzahl der Level

var ilevel uint16	  				// aktuelle Levelnummer
var nPunkte uint16 = 100			// max. Punkte im Level
var aPunkte uint16					// Punkteanzug durch Kantenpassage			
var punkte uint16 					// aktuelle Punktzahl
var text []string = WillkommenText()
var paketid uint32					// id des Pakets
var gewonnen bool					// Level geschafft
var verloren uint16					// 1 = Kante gesperrt, 2 Knoten gesperrt,
									// 3 = Bugget zu Ende
var bestanden bool					// Püfung bestanden
var time0 uint16				// Staŕtzeit in Sekunden


//var gPunkte	uint16				// Gesamtpunkte erreicht
//	var maxPunkte uint16			// maximale erreichbare Geamtpunktzahl
//	var happy bool = true			// Winnie sieht happy aus
//	var levelNeuLaden bool			// Level neu laden
//	var neuZeichnen bool 			// Schaltkreis neu zeichnen
//	var bestanden bool				// Prüfung bestanden?
//	var gewonnen bool				// Level geschafft
								




func erzeugeKnotenButton(x,y uint16, g uint8) buttons.Button {
	if g == 255 {
		return buttons.New(x-20,y-25,40,50,0,0,0,true, "")
	} else {
		return buttons.New(x-20,y-25,40,50,0,0,0,false, "")
	}	
}


func erzeugeStartStopButton(x,y uint16) buttons.Button {
	return buttons.New(x-30,y-15,60,30,0,0,0,false, "")
}


func makeKnotenButtonTab(netz Graph) map[uint16]buttons.Button {
	var max uint32 = maxID(netz)	
	var x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = netz.KnotenID_Liste()
	var buts map[uint16]buttons.Button = make(map[uint16]buttons.Button,0)		
	for _,id := range ids {
		x,y = netz.KnotenKoordinaten(id)
		_,g,_ := netz.Knotenfarbe(id)
		if id == 0 && id == max {
			buts[uint16(id)] = erzeugeStartStopButton(x,y)
		} else {
			buts[uint16(id)] = erzeugeKnotenButton(x,y,g)
		}
	}
	return buts
}


func inaktiviereKnotenButton(buts map[uint16]buttons.Button) {
	for _,button := range buts {
			button.DeaktiviereButton()
	}
}


func aktiviereKnotenButton(buts map[uint16]buttons.Button, idlist []uint32) {
	for _,id := range idlist {
			buts[uint16(id)].AktiviereButton()
	}
}


func WillkommenText() []string {
	var erg []string = make([]string,0)
	erg = append(erg,"Willkommen zum NET-Game!")
	erg = append(erg,"")	
	erg = append(erg,"")
	erg = append(erg,"Bewegen Sie das Datenpaket")
	erg = append(erg,"entlang der Verbindungen")
	erg = append(erg,"zum Ziel, indem Sie auf die")
	erg = append(erg,"grünen Nachbarrouter klicken.")
	erg = append(erg,"Nutzen Sie die kostengünstigste")
	erg = append(erg,"Verbindung.")
	erg = append(erg,"Beachten Sie Kosten der")
	erg = append(erg,"Verbindungen und die Zeit.")
	erg = append(erg,"")
	erg = append(erg,"Viel Spaß!")
	return erg
}


func schreibeGewonnen(diff uint16, weiter buttons.Button) []string {
	var erg []string = make([]string,0)
	
	erg = append(erg,"Glückwunsch Sie haben die")
	erg = append(erg,"Aufgabe geschafft!")
	erg = append(erg,"")
	// Wie gewonnen?
	if ilevel+1 == nlevel {		// Letztes Level? => Spiel zu Ende
		weiter.DeaktiviereButton()
		erg = append(erg,"Sie haben alle Level")
		erg = append(erg,"geschafft!")
		erg = append(erg,"")
		gfx.SpieleSound("../Sounds/Sparkle.wav")
		time.Sleep (time.Duration(4e8))
		gfx.SpieleSound("../Sounds/Sparkle.wav")							
		time.Sleep (time.Duration(4e8))
		gfx.SpieleSound("../Sounds/Sparkle.wav")
	} else if ilevel == 2 && !bestanden {	// oder Prüfung bestanden?
		weiter.AktiviereButton()
		bestanden = true							
		erg = append(erg,"Sie haben damit auch die")
		erg = append(erg,"Prüfung bestanden!")
		erg = append(erg,"")
		time.Sleep (time.Duration(4e8))
		gfx.SpieleSound("../Sounds/Sparkle.wav")
		time.Sleep (time.Duration(4e8))
		gfx.SpieleSound("../Sounds/Sparkle.wav")							
		time.Sleep (time.Duration(4e8))
		gfx.SpieleSound("../Sounds/Sparkle.wav")
	} else { // oder nur Level gewonnen
		erg = append(erg,"Sie die ideale Route um")
		erg = append(erg,fmt.Sprint(diff)+" Punkte verfehlt.")
		erg = append(erg,"")
		weiter.AktiviereButton()
		gfx.SpieleSound("../Sounds/Sparkle.wav")
	}			
	erg = append(erg,"")	
	erg = append(erg,"Auf zur nächsten Aufgabe oder")
	erg = append(erg,"versuchen Sie es noch einmal.")
	return erg
}



func schreibeVerloren(verloren uint16) []string {
	var erg []string = make([]string,0)

	erg = append(erg,"Es tut mit Leid, aber die")
	erg = append(erg,"Kosten waren zu groß. Sie")	
	erg = append(erg,"haben leider verloren!")
	erg = append(erg,"")			
	erg = append(erg,"Aber versuchen Sie es noch")
	erg = append(erg,"einmal.")
	return erg
}



func berechneNote(punkte, maxPunkte uint16) string {
	var note string
	var prozente float32 = float32(punkte)/float32(maxPunkte)*100.
	if prozente > 90 {
		note = "1.0"
	} else if prozente > 85 {
		note = "1.3"
	} else if prozente > 80 {
		note = "1.7"
	} else if prozente > 75 {
		note = "2.0"
	} else if prozente > 70 {
		note = "2.3"
	} else if prozente > 65 {
		note = "2.7"
	} else if prozente > 60 {
		note = "3.0"
	} else if prozente > 55 {
		note = "3.3"
	} else if prozente > 50 {
		note = "3.7"
	} else if prozente > 45 {
		note = "4.0"
	} else {
		note = "n.B."
	}
	return note
}


func schreibeSpielstand(level,punkte, maxPunkte uint16) {
	var note string =  berechneNote(punkte, maxPunkte)
	gfx.SchreibeFont(20,15,"Level: " + fmt.Sprint(level))
	gfx.SchreibeFont(150,15,"Punkte: " + fmt.Sprint(punkte))
	gfx.SchreibeFont(320,15,"Note: " + fmt.Sprint(note))
}


func zeichneButtons(weiter,zurueck,beenden,nochmal buttons.Button) {
	if weiter.GibAktivitaetButton() {
			weiter.ZeichneButton()
	}
	if zurueck.GibAktivitaetButton() {
			zurueck.ZeichneButton()
	}
	if beenden.GibAktivitaetButton() {
			beenden.ZeichneButton()
	}
	if nochmal.GibAktivitaetButton() {
			nochmal.ZeichneButton()
	}
}



func zeichneStart(netz Graph) {
	var x,y uint16 = netz.KnotenKoordinaten(0)
	//var ids []uint32 = netz.KnotenID_Liste()
	r,g,_ := netz.Knotenfarbe(0)
	if r == 255 && g == 255 {
		gfx.Stiftfarbe(r,g,0)
	} else {
		gfx.Stiftfarbe(255,255,255)
	}
	gfx.Vollrechteck(x-30,y-15,60,30)
	gfx.Stiftfarbe(0,0,0)
	gfx.Rechteck(x-30,y-15,60,30)
	gfx.SetzeFont ("../Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-24,y-12,"Start")
}


func zeichneZiel(netz Graph) {
	var max uint32 = maxID(netz)
	var x,y uint16 = netz.KnotenKoordinaten(max)
	//var ids []uint32 = netz.KnotenID_Liste()
	r,g,_ := netz.Knotenfarbe(max)
	if r == 255 && g == 255 {
		gfx.Stiftfarbe(r,g,0)
	} else {
		gfx.Stiftfarbe(255,255,255)
	}
	gfx.Vollrechteck(x-30,y-15,60,30)
	gfx.Stiftfarbe(0,0,0)
	gfx.Rechteck(x-30,y-15,60,30)
	gfx.SetzeFont ("../Schriftarten/Ubuntu-B.ttf",20)
	gfx.SchreibeFont(x-17,y-12,"Ziel")
}


func zeichneComputer(netz Graph) {
	var max uint32 = maxID(netz)	
	var x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = netz.KnotenID_Liste()
	for i:=0; i<len(ids); i++ {
		x,y = netz.KnotenKoordinaten(ids[i])
		r,g,_ := netz.Knotenfarbe(ids[i])
		if ids[i] != 0 && ids[i] != max {
			if r == 255 {
				gfx.LadeBild(x-20,y-25,"../Bilder/Computer_klein_rot.bmp")
				gfx.LadeBild(x-20,y-40,"../Bilder/Feuer.bmp")
			} else if g == 255 {
				gfx.LadeBild(x-20,y-25,"../Bilder/Computer_klein_gruen.bmp")
			} else {
				gfx.LadeBild(x-20,y-25,"../Bilder/Computer_klein.bmp")
			}
		}
	}

}



func zeichneDarthSchmidtar(netz Graph) {
	//var max uint32 = maxID(netz)	
	var x0,y0,x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = netz.KnotenID_Liste()
	fmt.Println("Neue Liste")
	for i:=0; i<len(ids); i++ {
		x0,y0 = netz.KnotenKoordinaten(ids[i])
		for j:=0; j<len(ids); j++ {
			//fmt.Println(i,j)		
//			if i == 0 && j == 1 { 
//				gfx.LadeBild(100,100,"../Bilder/DarthSchmidtarExtraTiny.bmp")
//			}
			if netz.Benachbart(ids[i],ids[j]) { //&&  ids[i]<ids[j] {
				r,_,_ := netz.Kantenfarbe(ids[i],ids[j])
				if r == 255 {
					x,y = netz.KnotenKoordinaten(ids[j])
					//fmt.Println("Knotenkoord: ",ids[i],x0,y0)
					fmt.Println("Kantenkoord: ",ids[i],ids[j],x,y)
					x = x0 + (x-x0)/2
					y = y0 + (y-y0)/2
					//fmt.Println("Darth Pos: ",x,y)
					gfx.LadeBild(x-16,y-25,"../Bilder/DarthSchmidtarExtraTiny.bmp")
				}
			}
		}
	}

}



func zeichnePaket(id uint32,netz Graph) {
	var x,y uint16 = netz.KnotenKoordinaten(id)
//	gfx.LadeBildMitColorKey(x-25,y-25,"../Bilder/paket_klein.bmp",255,255,255)
	gfx.LadeBild(x-25,y-25,"../Bilder/paket_klein.bmp")
}



func zeichneSpielfeld(ilevel, punkte, maxPunkte uint16, text []string,netz Graph) {

	var fontsize int = 20

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		

//	if happy {
		gfx.LadeBild(840,10,"../Bilder/WtheK_black.bmp")
//	} else {
//		gfx.LadeBild(840,10,"../Bilder/WtheK_black_sad.bmp")
//	}
	gfx.Linie(830,0,830,700-1)
	gfx.Linie(830,380,1200-1,380)

	gfx.SetzeFont ("../Schriftarten/Ubuntu-B.ttf",fontsize)
	schreibeSpielstand(ilevel+1,punkte,maxPunkte)
	for i:=0; i<len(text); i++ {
		gfx.SchreibeFont(850,400+20*uint16(i),text[i])
	}
	
	netz.Darstellen()
	zeichneStart(netz)
	zeichneZiel(netz)
	zeichneComputer(netz)
//	zeichneDarthSchmidtar(netz)

}



func baueGraph() Graph {

	var g Graph = New(false)		// initialisiert einen leeren Graphen
	var m,n uint32 = 8,11			// Anzahl der Konten horizontal und verikal
	var dm,dn uint16 = 100,50		// Abstand zwischen Knoten 
	var k, kmax uint32 = 1,10		// Kosten und maximale Kosten					
	var id, id2 uint32				// ID des Knoten, und des 2. Kantenknoten

	// Zeichne Knoten
	for i:=uint32(0);i<m;i++ {
		for j:=uint32(0);j<n;j++ {
			id = i+j*m-j/2
			if j % 2 == 0 {
				g.KnotenEinfuegen(id,uint16(i)*dm+50,uint16(j)*dn+100,0)
			} else {
				if i < m-1 {
					g.KnotenEinfuegen(id,uint16(i)*dm+50+dm/2,uint16(j)*dn+100,0)
				}		
			}
		}
	}
  
	// Zeichne Verbindungen
	// gerade Zeilen
	for i:=uint32(0);i<m;i++ {
		for j:=uint32(0);j<n;j=j+2 {
			id = i+j*m-j/2
			k = zufallszahl(1,kmax)
			if j < n-1 {
				if i == 0 {
					id2 = id+m				
					if g.Enthalten(id2) {
						if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
					}
				} else if i == m-1 {
					id2 = id-1+m		
					if g.Enthalten(id2) {
						if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
					}
				} else {			
					id2 = id+m
					if g.Enthalten(id2) {
						if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
					} 
					id2 = id-1+m
					if g.Enthalten(id2) {
						if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
					}
				}			
			}
			if i < m-1 {
				id2 = id+1
				if g.Enthalten(id2) {
					if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
				}
			}
			if j < n-1 {
				id2 = id+2*m-1
				if g.Enthalten(id2) {
					if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
				}

			}
		}
	}

	// ungearde Zeilen
	for i:=uint32(0);i<m;i++ {
		for j:=uint32(0)+1;j<n;j=j+2 {
			id = i+j*m-j/2
			k = zufallszahl(1,kmax)
			if i != m-1 {
				id2 = id+m-1
				if g.Enthalten(id2) {
					if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
				}
				id2 = id+m
				if g.Enthalten(id2) {
					if falseORtrue() {g.KanteEinfuegen(id,id2,k)}
				}			
			}
		}
	}

	return g

}



func DijkstraAlgorithmus(g Graph) (bool,uint32) {
 
	// Algorithmus von Dijkstra
	var minknoten uint32
	var IDs []uint32 = g.KnotenID_Liste()
	var gelbOK bool					// Gelber Knoten gefunden?
	var mininfo uint32
	var ok bool						// Endknoten erreicht?

	
	var startknoten uint32 = 0					// Startknoten wählen
	g.KnotenFaerben (startknoten,255,255,0)		// Färbe Startknoten gelb


	
	// Wiederhole solange bis es keine gelben Knoten mehr gibt, dann 
	// break
	for {
		// suche gelben Knoten - mit kürzester Distanz
		gelbOK = false
		for _,index := range IDs {
			r,gr,b := g.Knotenfarbe(index) 
			info := g.Knoteninfo(index)
			if r == 255 && gr == 255 && b == 0 {
				if gelbOK {
					if info < mininfo {
						minknoten = index
						mininfo = info
					}
				} else {
					gelbOK = true
					minknoten = index
					mininfo = info
				}  
			}
		}
		if !gelbOK {break} // wenn keinen gelben Knoten gefunden dann beende den Algorithmus
		
		// Färbe minimalen gelben Knoten grün
		g.KnotenFaerben (minknoten,0,255,0)
		// Überprüfe alle Nachbarknoten
		for _,ID := range IDs {
			if g.Benachbart(minknoten,ID) {
				r,gr,b := g.Knotenfarbe(uint32(ID))
				if gr != 255 { // wenn Knoten nicht gelb oder grün (noch nicht besucht)
					g.KanteFaerben(minknoten,ID,255,0,0)		// Kantefarbe rot
					g.KnotenFaerben(ID,255,255,0)					// Knoten gelb
					g.KnoteninfoSetzen(ID,mininfo+g.Kanteninfo(minknoten,ID)) // Distanz setzen
				} else if r == 255 && gr == 255 && b == 0  { // Knoten gelb? (besucht)
					if mininfo+g.Kanteninfo(minknoten,ID) < g.Knoteninfo(ID) {
						g.KanteFaerben(minknoten,ID,255,0,0)	// Kantefarbe rot
						g.KnoteninfoSetzen(ID,mininfo+g.Kanteninfo(minknoten,ID)) //Distanz setzen
						// alte rote Kante gelb färben
							for _,ID2 := range IDs {
								if g.Benachbart(ID2,ID) {
								r1,gr1,_ := g.Kantenfarbe(ID2,ID)	
								if r1 == 255 && gr1 == 0 && ID2 != minknoten {
									g.KanteFaerben(ID2,ID,255,255,0)
								}
							}
						}
					} else { 				// Nachbarknoten gelb, aber Dist. nicht minimal
						g.KanteFaerben(ID,minknoten,255,255,0)	// => Kantenfarbe gelb
					}
				} else { // Nachbarknoten schon abgearbeitet -> grün => Kante gelb
					// g.KanteFaerben(minknoten,ID,255,255,0)	// Kantefarbe gelb
					// !!!! Weglassen da gerichteter Graph und sonst rote Kanten nicht sichtbar
				}
			}
		}
	}

	// Kantenfarbe zurücksetzen
	for _,ID1 := range IDs {
		for _,ID2 := range IDs {
			g.KanteFaerben(ID1,ID2,0,0,0)			// Kantefarbe schwarz
		}	
	}		
	
	max := maxID(g)
	r,gr,_ := g.Knotenfarbe(max)
	// Endknoten erreichbar?
	if r==0 && gr==255 {
		ok = true
	}

	return ok,g.Knoteninfo(max)

}


func zufallszahl(m0,m1 uint32) uint32 {
	var delta float64 = float64(m1-m0)
	return uint32(rand.Float64() * delta)+1
}


func falseORtrue() bool {
	return rand.Float64() > 0.5
}


func maxID(netz Graph) uint32 {
	var ids []uint32 = netz.KnotenID_Liste()
	var max uint32
	for i:=0; i<len(ids); i++ {
		if max <= ids[i] {
			max = ids[i]
		}
	}
	return max
}


// Vor: Es gibt ein Netz von Knoten und Kanten.
// Eff: Die Nachbarn werden grün gesetzt.
// Erg: Eine Liste der Nachbarn wird zurückgegeben.
func findeNachbarn(id uint32, netz Graph) []uint32 {
	var nachbarn []uint32 = make([]uint32,0)
	var ids []uint32 = netz.KnotenID_Liste()
	for _,index2:= range ids {
		r,_,b := netz.Knotenfarbe(index2) 
		if netz.Benachbart (id, index2) {
			nachbarn = append(nachbarn,index2)
			netz.KnotenFaerben(index2,r,255,b)
		} else {
			netz.KnotenFaerben(index2,r,0,b)
		}
	}
	return nachbarn
}


func hindernisse(netz Graph, pKnotensperre,pKantensperre float64) {
	var ids []uint32 = netz.KnotenID_Liste()
	var max uint32 = maxID(netz)
	// Knoten sperren = rot
	for {
		if !gewonnen && verloren == 0 {
			for _,index:= range ids {
				_,g,b := netz.Knotenfarbe(index)
				if index != 0 && index != max {
					if rand.Float64() > pKnotensperre {
						netz.KnotenFaerben(index,0,g,b)
					} else {
						netz.KnotenFaerben(index,255,g,b)
					}
				}
				for _,index2:= range ids {
					_,g,b := netz.Kantenfarbe(index,index2) 
					if rand.Float64() > pKantensperre && index < index2 {
						netz.KanteFaerben(index,index2,0,g,b)
						netz.KanteFaerben(index2,index,0,g,b)
					} else {
						netz.KanteFaerben(index,index2,255,g,b)
						netz.KanteFaerben(index2,index,255,g,b)
					}
				}
			}
		}
		time.Sleep (time.Duration(2e9))
	}
}


func zeichnen(weiter,zurueck,beenden,nochmal buttons.Button, netz Graph) {

		var dtime uint16				// Zeit seit Start

		for {
			if !gewonnen && verloren==0 { 
				dtime = uint16(float64(time.Now().UnixNano())/1e9) - time0
			}
			punkte = nPunkte - aPunkte - dtime
			if punkte == 0 || punkte > nPunkte {
				verloren = 3
				punkte = 0
			}
			fmt.Println("Punkte:", punkte)
			gfx.UpdateAus()
			zeichneSpielfeld(ilevel,punkte,0,text,netz)
			zeichnePaket(paketid,netz)
			zeichneButtons(weiter,zurueck,beenden,nochmal)
			gfx.UpdateAn()
			time.Sleep (time.Duration(50e7))
		}
}



// Voraus: -
// Eff: Spiel wird gestartet.
// Erg: akteulles Level ilevel+1, Note und Punktestand je Level
//		wird ausgegeben.
func TheNETgame(ilevel uint16,ePunkte []uint16) (uint16,string,[]uint16) {

//	var ilevelGeschafft	uint16		// höchstes geschafftes Level
//	var nlevel uint16 = 9				// Anzahl der Level
//	var nPunkte uint16 = 200		// neue Punkte im Level
//	var aPunkte uint16				// Punkteanzug durch Kantenpassage			
//	var punkte uint16 				// aktuelle Punktzahl
//	var ePunkte [nlevel]uint16		// Punkte erreicht im Level
	var gPunkte	uint16				// Gesamtpunkte erreicht
	var maxPunkte uint16			// maximale erreichbare Geamtpunktzahl
//	var happy bool = true			// Winnie sieht happy aus
//	var levelNeuLaden bool			// Level neu laden
//	var neuZeichnen bool 			// Schaltkreis neu zeichnen
//	var bestanden bool				// Prüfung bestanden?
									
//	var text []string = WillkommenText()
	var font string = "../Schriftarten/Ubuntu-B.ttf"
//	var sound string = ""
//	var soundAn bool				// soll Sound gespielt werden?
	var mindist uint32				// minimale Distanz zum Ziel
	var ok bool						// alles OK?
	var zielID uint32				// id des Ziels
	
	
	ePunkte = make([]uint16,nlevel)
	time0 = uint16(float64(time.Now().UnixNano())/1e9)


	//  --------------------   baue Netz ----------------------------//
	
	rand.Seed(time.Now().UnixNano())		// setzt Saat der Zufallszahlen
	var netz Graph = baueGraph()			// Netz mit Knoten und Kanten
	for i:=0; i<100; i++ {					// Check ob Graph zum Ende führt
		ok,mindist = DijkstraAlgorithmus(netz)
		if ok {break}
		if i==9 {panic("Kein zusammenhängenden Graphen gefunden!")}
		netz = baueGraph()		// erzeuge neues Netz!
	}
	fmt.Println(mindist)

	// Liste der Nachbarn zum aktuellen Knoten
	var nachbarn []uint32 = findeNachbarn(0,netz)	
	zielID = maxID(netz)					// Zielknoten ID
	var ids []uint32 = netz.KnotenID_Liste()


	//  --------------------   Buttons ------------------------------//
	
	// erzeuge eine Tabelle von Buttons zu den zugehörigen Netzpunkten //
	// id gibt die Zuordnung
	var sbutton map[uint16]buttons.Button = makeKnotenButtonTab(netz)

	// Erzeuge Buttons zur Spielsteuerung
	var weiter,zurueck,beenden,nochmal buttons.Button
	weiter = buttons.New(1090,650,100,40,255,255,100,false,"  weiter")
	weiter.SetzeFont(font)
	zurueck = buttons.New(850,650,100,40,255,255,100,false,"  zurück")
	zurueck.SetzeFont(font)
	beenden = buttons.New(30,650,100,40,255,255,100,true,"   Ende")
	beenden.SetzeFont(font)
	nochmal = buttons.New(970,650,100,40,255,255,100,false,"nochmal")
	nochmal.SetzeFont(font)
 

	// ----------- starte Grafikausgabe ----------------------------// 
	// ------------und zeichne Spielfeld -------------------------- //

//	gfx.Fenster(1200,700)
/*	time0 = uint16(float64(time.Now().UnixNano())/1e9)
	dtime = uint16(float64(time.Now().UnixNano())/1e9) - time0
	punkte = nPunkte	
	gfx.SetzeFont ("../Schriftarten/Ubuntu-B.ttf",20)
	gfx.UpdateAus()
	zeichneSpielfeld(happy,ilevel,punkte,maxPunkte,text,netz)
	zeichnePaket(0,netz)
	zeichneButtons(weiter,zurueck,beenden,nochmal)
	gfx.UpdateAn()
*/

	var pKnotensperre,pKantensperre float64 = 0.1,0.025
	go hindernisse(netz, pKnotensperre,pKantensperre) 

	go zeichnen(weiter,zurueck,beenden,nochmal,netz)


	// ----------- Mausabfrage & Spielsteuerung ---------------------//
	
	//taste, status, mausX, mausY := gfx.MausLesen1()
	
	
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			for id,but:= range sbutton {		// Überprüfe Schalter
				if but.TesteXYPosInButton(mausX,mausY) {
					fmt.Println("Schalter getroffen: ",id)
					fmt.Println(netz.Kanteninfo(paketid,uint32(id)))
					aPunkte = aPunkte + uint16(netz.Kanteninfo(paketid,uint32(id)))
					inaktiviereKnotenButton(sbutton)
					paketid = uint32(id)
					fmt.Println("Nachbarn: ",nachbarn)
					for _,index2:= range ids {
						r,_,_ := netz.Kantenfarbe(paketid,index2)
						if r == 255 {
							verloren = 2
							fmt.Println("verbotene Kante betreten: ",id,index2)
							break
						}
					}
					r,_,_ := netz.Knotenfarbe(paketid)
					if r==255 {
						verloren = 1
					} else if id == uint16(zielID) && verloren==0 {
						fmt.Println("Ziel erreicht!")
						gewonnen = true 	
					} else {
						nachbarn = findeNachbarn(paketid,netz)		
						aktiviereKnotenButton(sbutton,nachbarn)
					}
					//neuZeichnen = true
				}
			}
			// check Level gewonnen oder verloren?
			if gewonnen {			// check: Level gewonnen?
				text = schreibeGewonnen(aPunkte-uint16(mindist),weiter)
				// Merke die Punkte im Level
				if punkte > ePunkte[ilevel] {ePunkte[ilevel] = punkte} // Verbesserung?
				gPunkte = 0					// Berechne Gesamtpunktzahl
				for i:=uint16(0); i<nlevel; i++ {
					gPunkte = gPunkte + ePunkte[i]
				}
				nochmal.AktiviereButton()
			} else if verloren > 0 {
				text = schreibeVerloren(verloren)
				nochmal.AktiviereButton()
			}
/*				// Merke die Punkte im Level

			} else if nPunkte == 0 {	   	// wenn zu viele Versuche!!!
				inaktiviereSchalter(sbutton)
				text = schreibeVerloren()
				happy = false
				nochmal.AktiviereButton()
				if !soundAn {			// Spiele Sound nur einmal
					gfx.SpieleSound("../Sounds/GameOver.wav")
					soundAn = true	
				}
			}
			if weiter.TesteXYPosInButton(mausX,mausY) { // nächstes Level
				// Lade nächtes Level
				ilevel++
				zurueck.AktiviereButton()
				// weiter-Button nur wenn Level schon gewonnen
				// und Spiel noch nicht fertig
				if ilevel == ilevelGeschafft || ilevel+1 == nlevel {
					weiter.DeaktiviereButton()
				}
				levelNeuLaden = true
				neuZeichnen = true		
			}
			*/
			if nochmal.TesteXYPosInButton(mausX,mausY) { // Level nochmal
				aPunkte = 0
				inaktiviereKnotenButton(sbutton)
				paketid = 0	
				nachbarn = findeNachbarn(paketid,netz)		
				aktiviereKnotenButton(sbutton,nachbarn)
				gewonnen = false
				verloren = 0
//				levelNeuLaden = true
//				neuZeichnen = true
				time0 = uint16(float64(time.Now().UnixNano())/1e9)

			}
			/*
			if zurueck.TesteXYPosInButton(mausX,mausY) { // Level zurück
				// Lade vorheriges Level 
				ilevel--
				if ilevel == 0 {zurueck.DeaktiviereButton()}
				if ilevel < ilevelGeschafft {weiter.AktiviereButton()}
				levelNeuLaden = true
				neuZeichnen = true
			}
*/
			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende des Spiels
				break
			}
/*
			if levelNeuLaden {
//				lev  = level.New()		// Veränderungen rückgängig machen
				sk = lev.GibSchaltkreis(ilevel)
				sbutton = makeSchalterbuttonTab(sk,xSize)
				nPunkte = lev.GibMaxPunktzahl(ilevel) + lev.GibMinSchalter(ilevel)
				xSize = lev.GibXSize(ilevel)
				text = lev.GibText(ilevel)
				nochmal.DeaktiviereButton()
				levelNeuLaden = false
				soundAn = false
			}
			
			if neuZeichnen {
					gfx.UpdateAus()
					zeichneSpielfeld(happy,ilevel,punkte,gPunkte,text,netz)
					zeichnePaket(paketid,netz)
					zeichneButtons(weiter,zurueck,beenden,nochmal)
					gfx.UpdateAn()
					neuZeichnen = false
			}
*/
		}
		//neuZeichnen = false
		//time.Sleep (time.Duration(1e7))
	}

 
return ilevel,berechneNote(gPunkte,maxPunkte),ePunkte
	
}
