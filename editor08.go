package main
// Autor: St. Schmidt
// Datum: 13.09.2022
// Zweck: "Anwendung" des ADT Folge (polymorph) - Der Mehrzeilen-Editor!!
// ---> Mit Verbesserungen: 1.) Es gehen keine Tastendrücke mehr verloren!
// --->                     2.) Kein Flackern des Editorfensters mehr!
//-----------------------------------------------------------------------------------------------
import ("gfx"; "fmt" ; . "./folgen")

const spaltenanzahl, zeilenanzahl = 64, 21
	
var f Folge = New (rune(0)) //neue Folge mit Elementen vom Typ rune
var zeilen Folge = New (f)  //neue Folge mit Elementen vom Typ 'Folge von rune-Elementen'
	
var gedrueckt uint8
var taste,tiefe uint16
	
var startspalte,startzeile, index uint = 0,0,0

func String (f Folge, startposition uint) (erg string) {
	var index uint = f.AktuellerIndex ()
	f.Positionieren (startposition)
	for i:=uint(0); i < f.Laenge (); i++ { 
		e,_:=f.AktuellesElement()
		erg = erg + string(e.(rune)) //Folgenelemente sind vom Typ Element !!
		f.Vor ()
	}
	f.Positionieren (index)
	return
}

func updateFenster () {
	gfx.UpdateAus()
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls ()
	gfx.Stiftfarbe(0,0,0)
	z,_:= zeilen.AktuellesElement ()
	f = z.(Folge)
	for f.AktuellerIndex () >= startspalte + spaltenanzahl {startspalte++}
	for f.AktuellerIndex () < startspalte {startspalte--}
	index = zeilen.AktuellerIndex () //merke aktuelle Zeile
	for index > startzeile + zeilenanzahl - 2 {startzeile++}
	for index < startzeile {startzeile--}
	for i:= startzeile; i < startzeile + zeilenanzahl - 1; i++ {
		zeilen.Positionieren (i)
		if a,ok:= zeilen.AktuellesElement (); ok {
			if a.(Folge).Laenge () > startspalte {
				gfx.SchreibeFont (0,uint16((i-startzeile)*28),String(a.(Folge),startspalte) + "  ")
			}
		}
	}
	zeilen.Positionieren (index) //aktiviere wieder aktuelle Zeile		
	gfx.SchreibeFont (0,572,"Sp:       ") 
	gfx.SchreibeFont (70,572,fmt.Sprint(f.AktuellerIndex ()+1))
	gfx.SchreibeFont (120,572,"Zeile:        ")
	gfx.SchreibeFont  (230,572,fmt.Sprint(zeilen.AktuellerIndex()+1))
	gfx.SchreibeFont (uint16((f.AktuellerIndex()-startspalte)*14),28*uint16(index-startzeile),"_") //Kursor
	gfx.UpdateAn()
}
	
func main () {
	zeilen.Einfuegen (New (rune(0)))  //Die erste Zeile ist immer da.
	zeilen.Positionieren (0) //Die erste Zeile ist am Anfang aktuell.
	gfx.Fenster (900,600)
	gfx.Fenstertitel ("Unser erster eigener Editor!")
	gfx.SetzeFont ("./terminus-font/TerminusTTF-4.49.2.ttf",28)
	gfx.TastaturpufferAn()
	updateFenster () //Damit man der Cursor in der oberen Ecke sieht.
A:	for {
		taste, gedrueckt, tiefe = gfx.TastaturpufferLesen1()
		//fmt.Println (taste,gedrueckt,tiefe, gfx.Tastaturzeichen(taste, tiefe), string (gfx.Tastaturzeichen(taste,tiefe)))
		if gedrueckt == 1  { // Beim Drücken der Taste, nicht beim Loslassen!
			fmt.Println(taste)
			switch {
				case taste == 27:  // ESC-Taste
				break A
				case taste == 276:  // Pfeil nach links
				f.Zurueck ()		
				case taste == 275:  // Pfeil nach rechts
				f.Vor ()
				case taste == 278:  // POS1-Taste
				f.Positionieren (0)
				case taste == 279: // Ende - Taste
				f.Positionieren (f.Laenge ())
				case taste == 127: // Entf-Taste
				f.Loeschen ()
				case taste ==  8:  //Backspace-Taste
				if f.AktuellerIndex () > 0 {
					f.Zurueck ()
					f.Loeschen ()
				}
				case taste == 13:  //ENTER-Taste
				zeilen.Vor ()
				zeilen.Einfuegen (New (rune(0)))
				zeilen.Zurueck ()
				case taste == 273:  // Pfeil nach oben
				zeilen.Zurueck ()
				case taste == 274:  // Pfeil nach unten
				if zeilen.AktuellerIndex() < zeilen.Laenge () - 1 {
					zeilen.Vor ()
				}
				case taste >= 32 && taste <= 270:  
				z:= gfx.Tastaturzeichen(taste, tiefe)
				if z != 0 { f.Einfuegen (z) }
			}
			updateFenster()
		}
	}
}
