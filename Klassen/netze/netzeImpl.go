// Klasse Netz für theNETgame wird von Graph abgeleitet.

// Implementation

// Martin Seiß			12.5.2023	(Start)

//  !!!!!! Möglicher Bug im Dijkstra-Algorithmus. Erspielte Distanz manchmal
//		geringer als über den Algorthmus bestimmte, was nicht sein 
//		darf.

package netze

import "../graphen"
import "math/rand"
import "time"
import "fmt"


type data struct {
	graphen.Graph							// Graph für das Computernetzwerk
	pKnotensperre,pKantensperre float64		// W'keit für Knoten und Kantensperre
	mindist uint32							// Minimale Distanz (Kosten)
											// zwischen Start und Ziel
	zielID	uint32							// ID des Ziels
	gewonnen bool							// Spiel gewonnen
	verloren uint16							// Spiel verloren?
											// 0 = nicht verloren
											// 1 = verbotene Kante betreten
											// 2 = defekten Router betreten
											// 3 = Kredit verbraucht
}



func New(pKnotensperre,pKantensperre float64) *data {
	var n *data = new(data)
	n.pKnotensperre = pKnotensperre
	n.pKantensperre = pKantensperre
	n.Graph = graphen.New(false)
	n.baueNetzgraph()
	n.gewonnen = false
	n.verloren = 0
	fmt.Println("W'keiten:",n.pKnotensperre,n.pKantensperre)
	return n
}



func (netz *data) GibNachbarIDs(id uint32) []uint32 {
	var nachbarn []uint32 = make([]uint32,0)
	var ids []uint32 = netz.KnotenID_Liste()
	for _,index2:= range ids {
		r,_,b := netz.Knotenfarbe(index2) 
		if netz.Benachbart (id, index2) {
			nachbarn = append(nachbarn,index2)
			netz.KnotenFaerben(index2,r,255,b)
		} else if netz.Benachbart (index2,id) {
			nachbarn = append(nachbarn,index2)
			netz.KnotenFaerben(index2,r,255,b)			
		} else {
			netz.KnotenFaerben(index2,r,0,b)
		}
	}
	return nachbarn
}



func (netz *data) GibZielID() uint32 {
	return netz.zielID
}



func (netz *data) GibMinDist() uint32 {
	return netz.mindist
}



func (n *data) SetzeWkeitHindernisse(pKnotensperre,pKantensperre float64) {
	n.pKnotensperre = pKnotensperre
	n.pKantensperre = pKantensperre

}



func (netz *data) Hindernisse() {

	var ids []uint32 = netz.KnotenID_Liste()
	var max uint32 = netz.zielID
	
	for {
		if !netz.gewonnen && netz.verloren == 0 {
			//fmt.Println("verloren: ",verloren)
			for _,index:= range ids {
				_,g,b := netz.Knotenfarbe(index)
				if index != 0 && index != max {
					if rand.Float64() > netz.pKnotensperre {
						netz.KnotenFaerben(index,0,g,b)
					} else {
						netz.KnotenFaerben(index,255,g,b)
						//fmt.Println("Knoten gesperrt: ",index)
					}
				}
				for _,index2:= range ids {
					_,g,b := netz.Kantenfarbe(index,index2)
					if netz.Benachbart(index,index2) && index < index2 {
						if rand.Float64() > netz.pKantensperre  {
							netz.KanteFaerben(index,index2,0,g,b)
						} else {
							netz.KanteFaerben(index,index2,255,g,b)
							//fmt.Println("Kante gesperrt: ",index,index2)
						}
					}
				}
			}
		}
		time.Sleep (time.Duration(2e9))
	}
	fmt.Println("Beende Hindernisse")
}



func (netz *data) SetzeVerloren(verloren uint16) {
	netz.verloren = verloren
}


func (netz *data) SetzeGewonnen(gewonnen bool) {
	netz.gewonnen = gewonnen
}


func (netz *data) GibVerloren() uint16 {
	return netz.verloren
}


func (netz *data) GibGewonnen() bool {
	return netz.gewonnen
}

////////////////////////////////////////////////////////////////

////////      Hilfsfunktionen      /////////////////////////////

////////////////////////////////////////////////////////////////


// Vor: Leerer Graph ist vorhanden!
// Eff: Ein für das Spiel geeigneter Graph mit Verbindung von Start
//		und Ziel ist hinzugefügt.
func (netz *data) baueNetzgraph() {
	
	var mindist uint32
	var ok bool
	
	rand.Seed(time.Now().UnixNano())		// setzt Saat der Zufallszahlen
	netz.baueGraph()						// erzeuge Netz mit Knoten und Kanten
	netz.zielID = netz.berechneZielID()		// bestimme Ziel ID
	for i:=0; i<1000; i++ {					// Check ob Graph zum Ende führt
		ok,mindist = netz.dijkstraAlgorithmus()
		fmt.Println("Graph OK? ",ok)
		if ok {break}
		if i % 100 == 99 {panic("Probleme zusammenhängenden Graphen zu finden!")}
		netz.Graph = graphen.New(false)
		netz.baueGraph()					// erzeuge neues Netz!
	}
	
	var gzahl uint
	for _,id:= range netz.KnotenID_Liste() {
		_,g,_ :=  netz.Knotenfarbe(id) 
		if g == 255 {gzahl++}
	}
	
	netz.mindist = mindist


}



// Vor: -
// Erg: Ein Graph der ein Computernetz repräsentieren könnte ist geleifert.
func (netz *data) baueGraph() {

	var m,n uint32 = 8,11			// Anzahl der Konten horizontal und verikal
	var dm,dn uint16 = 100,50		// Abstand zwischen Knoten 
	var k, kmax uint32 = 1,10		// Kosten und maximale Kosten					
	var id, id2 uint32				// ID des Knoten, und des 2. Kantenknoten

	// Zeichne Knoten
	for i:=uint32(0);i<m;i++ {
		for j:=uint32(0);j<n;j++ {
			id = i+j*m-j/2
			if j % 2 == 0 {
				netz.KnotenEinfuegen(id,uint16(i)*dm+50,uint16(j)*dn+100,0)
			} else {
				if i < m-1 {
					netz.KnotenEinfuegen(id,uint16(i)*dm+50+dm/2,uint16(j)*dn+100,0)
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
					if netz.Enthalten(id2) {
						if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
					}
				} else if i == m-1 {
					id2 = id-1+m		
					if netz.Enthalten(id2) {
						if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
					}
				} else {			
					id2 = id+m
					if netz.Enthalten(id2) {
						if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
					} 
					id2 = id-1+m
					if netz.Enthalten(id2) {
						if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
					}
				}			
			}
			if i < m-1 {
				id2 = id+1
				if netz.Enthalten(id2) {
					if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
				}
			}
			if j < n-1 {
				id2 = id+2*m-1
				if netz.Enthalten(id2) {
					if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
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
				if netz.Enthalten(id2) {
					if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
				}
				id2 = id+m
				if netz.Enthalten(id2) {
					if falseORtrue() {netz.KanteEinfuegen(id,id2,k)}
				}			
			}
		}
	}

}


// Vor: -
// Erg: True ist geliefert, wenn erster (Start) und letzter (Ziel) 
//		miteinander verbunden sind. Wenn True, dann ist die minimale Disatnz
//		von Start zum Ziel geliefert, sonst 0.
// 		Grundlage ist der Dijkstra-Algorithmus. 
//  !!!!!! Möglicher Bug im Algorithmus. Erspielte Distanz manchmal
//		geringer als über den Algorthmus bestimmte, was nicht sein 
//		dürfte.
func (g *data) dijkstraAlgorithmus() (bool,uint32) {
 
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
		
	r,gr,_ := g.Knotenfarbe(g.zielID)
	// Endknoten erreichbar?
	if r==0 && gr==255 {
		ok = true
	}

	return ok,g.Knoteninfo(g.zielID)

}


// Erg: Die größte ID also die ID des Ziels ist geliefert. 
func (netz *data) berechneZielID() uint32 {
	var ids []uint32 = netz.KnotenID_Liste()
	var max uint32
	for i:=0; i<len(ids); i++ {
		if max <= ids[i] {
			max = ids[i]
		}
	}
	return max
}



// Eine zufallszahl im Intervall m0 bis m1 ist geliefert.
func zufallszahl(m0,m1 uint32) uint32 {
	var delta float64 = float64(m1-m0)
	return uint32(rand.Float64() * delta)+1
}



// Erg: True ist mit einer Wahrscheinlichkeit von 0.5 geliefert, sonst
//		false.
func falseORtrue() bool {
	return rand.Float64() > 0.5
}




