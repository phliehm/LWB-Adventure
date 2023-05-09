package graphen
// Autor: St. Schmidt
// Datum: 17.01.2023
// Zweck: effiziente Implementierung des ADT Graph - ALLE Methoden mit O(1) bzw. O(k) !!!!
//        --> Damit ist auch das von Herrn Vater aufgeworfene Problem beseitigt!!
// ---------------------------------------------------------------------------------------

import ( "gfx" ; "fmt" ; "math" )

type data struct {
	gerichtet bool					// true: gerichteter Graph ; false: ungerichteter Graph
	knotentab map[uint32]*knoten 	// Tabelle aller Knoten des Graphen (Schlüssel sind die Knotennummern)
	kantenanzahl uint32				// Anzahl aller Kanten des Graphen
}

type knoten struct {
	index uint32	            	// seine ID	
	x,y uint16 	 	            	// seine Grafikkordinaten
	r,g,b uint8	     				// seine Farbe
	info uint32                   	// ggf. zusätzliche Information des Knotens
	kantentabAB map[uint32]*kante 	// Tabelle aller Kanten von diesem Knoten zum Knoten mit der angegebenen ID 
	kantentabAN map[uint32]*kante	// Tabelle aller Kanten zu  diesem Knoten vom Knoten mit der angegebenen ID 
	}
	
type kante struct {
	nachIndex uint32	// ZielknotenID
	r,g,b uint8         // ihre Farbe
	info uint32      	// ggf. zusätzliche Informationen der Kante
}

func New (gerichtet bool) *data {
	var g *data = new(data)
	g.gerichtet = gerichtet
	g.knotentab = make (map[uint32]*knoten,0)
	g.kantenanzahl = 0
	return g
}

func (g *data) Gerichtet () bool {
	return g.gerichtet
}

func (g *data) Leer () bool {
	return len(g.knotentab) == 0
}

func (g *data) Knotenanzahl () uint32 {
	return uint32(len(g.knotentab))
}

func (g *data) Kantenanzahl () uint32 {
	return g.kantenanzahl
}

func (g *data) KnotenEinfuegen (id uint32, x,y uint16, knoteninfo uint32) {
	if id > 99 { panic ("Voraussetzung von 'KnotenEinfuegen' nicht eingehalten!") }
	_,ok:=g.knotentab[id]
	if ok { panic ("ERROR: Knoten mit dieser ID existiert schon!") }
	var k knoten = knoten{id,x,y,200,200,200,knoteninfo,make(map[uint32]*kante,0),make(map[uint32]*kante,0)}
	g.knotentab[id]=&k
}

func (g *data) KanteEinfuegen (index1, index2 uint32, kanteninfo uint32) {
	if k1,ok:=g.knotentab[index1];ok { // Es gibt den Knoten mit der ID index1 ...
		if k2,ok:=g.knotentab[index2]; ok { // ... und es gibt den Knoten mit der ID index2 ...
			_,ok:=k1.kantentabAB[index2]
			if ok { panic ("ERROR: Einzufügende Kante existiert schon!") }
			neueKante:= kante{index2,0,0,0,kanteninfo}
			k1.kantentabAB[index2]= &neueKante
			k2.kantentabAN[index1]= &neueKante
			g.kantenanzahl++ // Eine Kante mehr im Graphen
			if !g.gerichtet { // Bei ungerichteten Graphen erfolgt die Eintragung auch in der Gegenrichtung ohne Erhöhung der Kantenzahl!
				neueKante2:= kante{index1,0,0,0,kanteninfo}
				k2.kantentabAB[index1]=&neueKante2
				k1.kantentabAN[index2]=&neueKante2
			}
		}
	}
}

func (g *data) Knotenfarbe (id uint32) (r,gr,b uint8) {
	if k,ok:=g.knotentab[id]; ok {
		r,gr,b = k.r,k.g,k.b
	}
	return
}

func (g *data) KnotenKoordinaten (id uint32) (x,y uint16) {
	if k,ok:=g.knotentab[id]; ok {
		x,y = k.x,k.y
	}
	return
}

func (g *data) Knoteninfo (id uint32) (info uint32) {
	if k,ok:=g.knotentab[id]; ok {
		info = k.info
	}
	return 
}

func (g *data) Kantenfarbe (vonID, nachID uint32) (r,gr,b uint8) {
	if k,ok:=g.knotentab[vonID]; ok {
		if kan,ok:=k.kantentabAB[nachID]; ok {
			r,gr,b=kan.r,kan.g,kan.b
		}
	}
	return
}

func (g *data) Kanteninfo (vonID,nachID uint32) (erg uint32) {
	if k,ok:= g.knotentab[vonID]; ok {
		if kan,ok:=k.kantentabAB[nachID]; ok {
			erg = kan.info
		}
	}
	return
}

func (g *data) KnotenFaerben (id uint32, r,gr,b uint8) {
	if k,ok:=g.knotentab[id]; ok {
		k.r,k.g,k.b=r,gr,b
	}
}

func (g *data) KanteFaerben (vonID, nachID uint32, r,gr,b uint8) {
	if k,ok:=g.knotentab[vonID] ; ok {
		if kan,ok:= k.kantentabAB[nachID]; ok {
			kan.r,kan.g,kan.b = r,gr,b
		}
	}
	if !g.gerichtet { //auch Gegenrichtung faerben
		if k,ok:=g.knotentab[nachID] ; ok {
			if kan,ok:= k.kantentabAB[vonID]; ok {
				kan.r,kan.g,kan.b = r,gr,b
			}
		}
	}	
}

func (g *data) KnoteninfoSetzen (id uint32, info uint32) {
	if k,ok:=g.knotentab[id]; ok {
		k.info = info
	}
}

func (g *data) KanteninfoSetzen (vonID, nachID uint32, info uint32) {
	if k,ok:=g.knotentab[vonID] ; ok {
		if kan,ok:= k.kantentabAB[nachID]; ok {
			kan.info = info
		}
	}
	if !g.gerichtet { //auch in Gegenrichtung die neue Info setzen
		if k,ok:=g.knotentab[nachID] ; ok {
			if kan,ok:= k.kantentabAB[vonID]; ok {
				kan.info = info
			}
		}
	}
}

func (g *data) KnotenID_Liste () (erg []uint32) {
	for id,_:= range g.knotentab {
		erg = append(erg,id)
	}
	return
}

func (g *data) Kantenliste () (erg [][2]uint32) {
	for _,k:= range g.knotentab {
		for nachID,_ := range k.kantentabAB {
			erg = append(erg,[2]uint32{k.index,nachID})
		}
	}
	return
}
			
func (g *data) Benachbart (index1, index2 uint32) (erg bool) {
	if k,ok:= g.knotentab[index1] ; ok {
		if _,ok:= k.kantentabAB[index2]; ok {
			erg = true
		}
	}
	return
}

func (g *data) Enthalten (id uint32) (erg bool) {
	if _,ok:=g.knotentab[id]; ok {
		erg =true
	}
	return
}

func (g *data) KnotenEntfernen (id uint32) {
	if k,ok:= g.knotentab[id]; ok { // Wenn es den Knoten mit der ID id gibt ...
		//1. In allen Knoten, die eine Kante ZUM Knoten mit -id- haben, muss der Verweis auf diese Kante dort gelöscht werden ...
		for id2,_ := range k.kantentabAN {
			delete (g.knotentab[id2].kantentabAB,id)
			g.kantenanzahl--
			if !g.gerichtet { // Bei ungerichteten Graphen wird dort auch die eingetragene Gegenrichtung ohne Verringerung der Kantananzahl gelöscht!
				delete (g.knotentab[id2].kantentabAN,id)
			}
		}
		//2. Danach löscht man den Knoten (inklusive seiner Kantentabellen). 
		//   Damit löscht die Garbage-Collection alle nicht mehr adressierbaren Kanten!!
		//   Bei gerichteten Graphen muss noch die Kantenanzahl um die vom Knoten id abgehenden Kanten verringert werden,
		//   bei ungerichteten wurden die Kanten schon oben abgezogen!
		if g.gerichtet {
			g.kantenanzahl = g.kantenanzahl - uint32(len (k.kantentabAB))
		}
		delete (g.knotentab, id) 
	}
}

func (g *data) KanteEntfernen (index1, index2 uint32) {
	if k1,ok:=g.knotentab[index1]; ok { // Es gibt den Knoten mit der ID index1 ...
		if k2,ok:= g.knotentab[index2]; ok { // ... und es gibt den Knoten mit der ID index2 ...
			if _,ok:= k1.kantentabAB[index2]; ok { // und es gibt die Kante wirklich ...
				delete(k1.kantentabAB,index2)
				delete(k2.kantentabAN,index1)
				g.kantenanzahl--  // Eine Kante weniger im Graphen!
				if !g.gerichtet { // Bei ungerichteten Graphen erfolgt die Löschung der Einträge auch in Gegenrichtung!
					delete (k2.kantentabAB,index1)
					delete (k1.kantentabAN,index2)
				}
			}
			// Damit gibt es keinen Verweis mehr auf die Kante und die Garbage-Collection löscht die Kante!
		}
	}
}

func (g *data) Darstellen () {
	if !gfx.FensterOffen () {gfx.Fenster (640,480)}
	gfx.Stiftfarbe (255,255,255)
	//gfx.Cls ()
	//Kanten
	for _,k:= range g.knotentab {
		x1:= k.x
		y1:= k.y
		for _,kan:= range k.kantentabAB {
			gfx.Stiftfarbe (kan.r, kan.g, kan.b)
			nachID:=kan.nachIndex
			k2:=g.knotentab[nachID]
			x2:= k2.x
			y2:= k2.y
			if (x1 !=x2) || (y1 != y2) {
				if !g.gerichtet {
					gfx.Linie (x1,y1,x2,y2)
				} else {			
					l:= math.Sqrt((float64(x1)-float64(x2))*(float64(x1)-float64(x2)) + (float64(y1)-float64(y2))*(float64(y1)-float64(y2)))
					xa:= uint16((l-20.0)/l * (float64(x2)-float64(x1)) + float64(x1) + 0.5)
					ya:= uint16((l-20.0)/l * (float64(y2)-float64(y1)) + float64(y1) + 0.5)
					xb:= uint16((l-10.0)/l * (float64(x2)-float64(x1)) + float64(x1) + 0.5)
					yb:= uint16((l-10.0)/l * (float64(y2)-float64(y1)) + float64(y1) + 0.5)
					xc:= uint16(3.0 * (float64(y1)-float64(y2))/l + float64(xa) + 0.5)
					yc:= uint16(3.0 * (float64(x2)-float64(x1))/l + float64(ya) + 0.5)
					//xd:= uint16(-3.0 * (float64(y1)-float64(y2))/l + float64(xa) + 0.5)
					//yd:= uint16(-3.0 * (float64(x2)-float64(x1))/l + float64(ya) + 0.5)
					//xe:= uint16(20.0/l * (float64(x2)-float64(x1)) + float64(x1) + 0.5)
					//ye:= uint16(20.0/l * (float64(y2)-float64(y1)) + float64(y1) + 0.5)
					xf:= uint16(6.0 * (float64(y1)-float64(y2))/l + float64(xa) + 0.5)
					yf:= uint16(6.0 * (float64(x2)-float64(x1))/l + float64(ya) + 0.5)
					xg:= uint16(3.0 * (float64(y1)-float64(y2))/l + float64(xb) + 0.5)
					yg:= uint16(3.0 * (float64(x2)-float64(x1))/l + float64(yb) + 0.5)
					gfx.Stiftfarbe (kan.r, kan.g, kan.b)
					gfx.Linie (x1,y1,xc,yc)
					gfx.Volldreieck (xf,yf,xa,ya,xg,yg)
				}
				xm, ym:=(x1+x2)/2, (y1+y2)/2
				l:= math.Sqrt((float64(x1)-float64(x2))*(float64(x1)-float64(x2)) + (float64(y1)-float64(y2))*(float64(y1)-float64(y2)))
				xneu:= uint16(12.0 * (float64(y1)-float64(y2))/l + float64(xm) + 0.5)
				yneu:= uint16(12.0 * (float64(x2)-float64(x1))/l + float64(ym) + 0.5)
				//gfx.Stiftfarbe (255,255,255)
				//gfx.Vollkreis (xm,ym,8)
				gfx.Stiftfarbe (0,0,0)
				if g.gerichtet || nachID > k.index {
					gfx.Schreibe (xneu-8,yneu-4,fmt.Sprint(kan.info))
				}		
			} else {
				gfx.Kreis (x1+30,y1-15,30)
				if g.gerichtet {
					gfx.Volldreieck(x1,y1-10,x1+3,y1-20,x1-3,y1-20)
				}
			}	
		}
	}
	//Knoten
	for _,k:= range g.knotentab {
		gfx.Stiftfarbe (k.r, k.g, k.b)
		gfx.Vollkreis (k.x, k.y, 10)
		gfx.Stiftfarbe (0,0,0)
		if k.index < 10 {
			gfx.Schreibe (k.x-3,k.y-3,fmt.Sprint(k.index))
		} else {
			gfx.Schreibe(k.x-8, k.y-3,fmt.Sprint(k.index))
		}
	}
}
