// the NET Game - Hauptfunktion und Hilfsfunktionen

// Martin Seiß    24.4.2023


// Level:
// ilevel, netz, punkte, txt, 


package theNETgame


import "fmt"
import "gfx"
import "../../Klassen/buttons"
//import "os"
//import "strconv"
import "time"
import	"../../Klassen/textboxen"
//import	. "../../Klassen/netze"
import  spielfelder "../../Klassen/theNETgameSpielfeld"


const nlevel uint16 = 6				// Anzahl der Level

var ilevel uint16	  				// aktuelle Levelnummer
var nPunkte uint16 = 100			// max. Punkte im Level
var aPunkte uint16					// Punkteanzug durch Kantenpassage			
var punkte uint16 					// aktuelle Punktzahl
var txt textboxen.Textbox
//var text []string = WillkommenText()
var paketid uint32					// id des Pakets
var gewonnen bool					// Level geschafft
var verloren uint16					// 1 = Kante gesperrt, 2 Knoten gesperrt,
									// 3 = Bugget zu Ende
var bestanden bool					// Püfung bestanden
var time0 uint16					// Staŕtzeit in Sekunden

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


func makeKnotenButtonTab(sf spielfelder.Spielfeld) map[uint16]buttons.Button {
	var max uint32 = sf.GibZielID()
	var x,y uint16 			// netz.KnotenKoordinaten(id)
	var ids []uint32 = sf.KnotenID_Liste()
	var buts map[uint16]buttons.Button = make(map[uint16]buttons.Button,0)		
	for _,id := range ids {
		x,y = sf.KnotenKoordinaten(id)
		_,g,_ := sf.Knotenfarbe(id)
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




// Voraus: -
// Eff: Spiel wird gestartet.
// Erg: akteulles Level ilevel+1, Note und Punktestand je Level
//		wird ausgegeben.
func TheNETgame() (float32,uint16) {

	var font string = "Schriftarten/Ubuntu-B.ttf"
	var nachbarn []uint32
	
	gfx.SetzeFont ("Schriftarten/Ubuntu-B.ttf",20)	


	// -----        Erzeuge Buttons zur Spielsteuerung    ----------///
	var weiter,starter,beenden,nochmal buttons.Button
	weiter = buttons.New(1090,650,100,40,255,255,100,false,"  weiter")
	weiter.SetzeFont(font)
	starter = buttons.New(850,650,100,40,255,255,100,true,"  start")
	starter.SetzeFont(font)
	beenden = buttons.New(30,650,100,40,255,255,100,true,"   Ende")
	beenden.SetzeFont(font)
	nochmal = buttons.New(970,650,100,40,255,255,100,false,"nochmal")
	nochmal.SetzeFont(font)

 
 	//  -------------------   baue Spielfeld ------------------------//
	var sf spielfelder.Spielfeld = spielfelder.New(weiter,starter,beenden,nochmal)
	nachbarn = sf.GibNachbarIDs(0)		// Finde die Nachbarn zu ID = 0

	//  --------------------  Router-Buttons ------------------------//
	// erzeuge eine Tabelle von Buttons zu den zugehörigen		     //
	// Netzpunkten - id gibt die Zuordnung							 //
	var sbutton map[uint16]buttons.Button = makeKnotenButtonTab(sf)


	// ----------- starte Grafikausgabe -----------------------------// 
	// ------------und zeichne Spielfeld --------------------------- //

	go sf.Zeichnen()
	
	go sf.Hintergrundmusik()
		
	
	// ----------- Mausabfrage & Spielsteuerung ---------------------//
	
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		if taste==1 && status==1 {
			
			for id,but:= range sbutton {		// Überprüfe Schalter
				if but.TesteXYPosInButton(mausX,mausY) && sf.SpielLaeuft() {
						// !sf.GibGewonnen() && sf.GibVerloren() == 0 {
					fmt.Println("Schalter getroffen: ",id)
					//fmt.Println("alte Nachbarn: ",nachbarn)
					//fmt.Println("Kantenifo:", sf.Kanteninfo(paketid,uint32(id)))
					sf.ErhoeheDistanz(uint16(sf.Kanteninfo(paketid,uint32(id))))
					inaktiviereKnotenButton(sbutton)
					// check Kante verboten?
					r,_,_ := sf.Kantenfarbe(paketid,uint32(id))
					if r == 255 {
						fmt.Println("verbotene Kante betreten: ",paketid,uint32(id))
						sf.SetzeVerloren(1)
						paketid = uint32(id)
						sf.SetzePaketID(paketid)
						nochmal.AktiviereButton()
					}
					// check Knoten verboten?
					r,_,_ = sf.Knotenfarbe(uint32(id))
					if r==255 {
						sf.SetzeVerloren(2)
						paketid = uint32(id)
						sf.SetzePaketID(paketid)
						nochmal.AktiviereButton()
					// check gewonnen?
					} else if id == uint16(sf.GibZielID()) && sf.GibVerloren()==0 {
						fmt.Println("Ziel erreicht!")
						paketid = uint32(id)
						sf.SetzePaketID(paketid)
						sf.SetzeGewonnen(true)
						if ilevel+1 < nlevel {
							weiter.AktiviereButton()
						}
					} else {
						paketid = uint32(id)
						sf.SetzePaketID(paketid)
						nachbarn = sf.GibNachbarIDs(paketid)	
						aktiviereKnotenButton(sbutton,nachbarn)
					}
					//fmt.Println("neue Nachbarn: ",nachbarn)
				}
			}
			
			// check Level gewonnen oder verloren?
			// if sf.GibGewonnen() || sf.GibVerloren()>0 {
			//	nochmal.AktiviereButton()
			// }				

			if weiter.TesteXYPosInButton(mausX,mausY) { // Lade nächtes Level
				weiter.DeaktiviereButton()
				inaktiviereKnotenButton(sbutton)
				sf.NeuesLevel(true)
				ilevel++
				sf.GibNachbarIDs(0)
				//aktiviereKnotenButton(sbutton,sf.GibNachbarIDs(0))
				starter.AktiviereButton()
				// weiter-Button nur wenn Level schon gewonnen
				// und Spiel noch nicht fertig

			}
			
			if nochmal.TesteXYPosInButton(mausX,mausY) { // Level nochmal
				nochmal.DeaktiviereButton()
				inaktiviereKnotenButton(sbutton)
				sf.NeuesLevel(false)
				fmt.Println("Verloren?",sf.GibVerloren())
				sf.GibNachbarIDs(0)
				starter.AktiviereButton()
			}

			if starter.TesteXYPosInButton(mausX,mausY) { // Start Level!
				sf.StartGame()
				aktiviereKnotenButton(sbutton,sf.GibNachbarIDs(0))
				starter.DeaktiviereButton()

			}

			if beenden.TesteXYPosInButton(mausX,mausY) { // Ende des Spiels
				beenden.DeaktiviereButton()
				break
			}

		}
		
		// Bei Verloren, wiederhole Level!
		//if sf.GibVerloren()>0 {
		//		nochmal.AktiviereButton()	
		//}
		// Bei Gewonnen, neues Level!
		//if sf.GibGewonnen() {weiter.AktiviereButton()}

		time.Sleep(1e7)

	}

 
return sf.GibNote(),sf.GibPunktzahl() 
	
}
