package textboxen

import (
	"gfx"
	"strings"
	//"fmt"
	"unicode/utf8"
)

type data struct {
	x,y,breite,höhe uint16 	// Koordinaten und breite und Höhe
	font string 
	zeilenAbstand uint16
	text string
	schriftgr int
	r,g,b uint8
	linksbündig bool
	zentriert bool
	
}

func New(posx,posy,breite,höhe uint16) *data{
	var tb *data
	tb = new(data)
	tb.x = posx
	tb.y = posy
	tb.höhe = höhe
	tb.breite = breite
	tb.schriftgr = 20
	tb.font =  "../Schriftarten/Ubuntu-B.ttf"
	tb.linksbündig = true
	return tb
}
// Vor.: --
// Eff.: Die Position der Textbox ist gesetzt
func (tb *data)	SetzePosition(x,y uint16) () {
	tb.x = x
	tb.y = y
} 

// Vor.: --
// Erg.: Position der Textbox ist geliefert
func (tb *data) GibPosition() (uint16,uint16) {
	return tb.x, tb.y
}

// Vor.: 
// Erg.: Die Größe (Breite und Höhe) der Textbox ist geliefert
func (tb *data)	GibGröße()(uint16,uint16) {
	return tb.breite, tb.höhe
}
// Vor.: --
// Eff.: Der Text wird der Textbox hinzugefügt 
func (tb *data)	SchreibeText(text string) () {
	tb.text = text
}
	
// Vor.: --
// Eff.: Der Zeilenabstand wurde geändert
func (tb *data)	SetzeZeilenAbstand(za uint16) () {
	tb.zeilenAbstand = za
}
	
// Vor.: --
// Erg.: Der Zeilenabstand ist geliefert
func (tb *data)	GibZeilenAbstand() (uint16) {
	return tb.zeilenAbstand
}

// Vor.: Es existiert eine Font-Datei im ttf-Format im angegebenen Ordner
// Eff.: Der Font des Textes in der Textbox ist gesetzt. 
func (tb *data)	SetzeFont(font string) {
	tb.font = font
}

// Vor.: --
// Erg.: Der Font der Textbox ist geliefert
func (tb *data)	GibFont()string {
	return tb.font
}

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geändert
func (tb *data)	SetzeSchriftgröße (gr int) {
	tb.schriftgr = gr
}

// Vor.: --
// Erg.: Die Schriftgröße der Textbox ist geliefert
func (tb *data)	GibSchriftgröße () int {
	return tb.schriftgr
}


// Vor.: --
// Erg.: Die Farbe des Textes und des Rahmens (gleich) ist gesetzt
func (tb *data) SetzeFarbe(r,g,b uint8) {
	tb.r,tb.g,tb.b  = r,g,b
}

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() linksbündig ausgegeben (default)
func (tb *data) SetzeLinksbündig() {
	tb.linksbündig = true
	tb.zentriert = false
}

// Vor.: --
// Erg.: Text wird beim Aufruf von Zeichne() zentriert ausgeben
func (tb *data) SetzeZentriert() {
	tb.linksbündig = false
	tb.zentriert = true
}

// Vor.: Ein gfx-Fenster ist offen
// Eff.: Zeichnet die Textbox in das gfx-Fenster
func (tb *data)	Zeichne() {
	gfx.SetzeFont(tb.font,tb.schriftgr)
	gfx.Stiftfarbe(tb.r, tb.g, tb.b)
	//gfx.Rechteck(tb.x,tb.y,tb.breite,tb.höhe)
	
	// Zeilenumbrüche generieren
	//var textTemp string = tb.text
	var zeilenHöhe int		// y-Position der Zeile
	var zeilenAnzahl uint16	
	var zeile string	// Zeileninhalt
	var zeichenbreite int
	var zeilenLänge, zeilenLänge_temp uint16
	//var linkerAbstand uint16 // Abstand zur linken Boxgrenze zur Darstellung von zentriertem Text
	zeichenbreite = tb.schriftgr/2
	if zeichenbreite == 0 {zeichenbreite = 1}	// Vermeide eine Zeichenbreite von 0
	
	// Wörter sollen ganz gelassen werden, Zeilenumbrüche also nur bei Leerzeichen
	
	worte := strings.Fields(tb.text)	// Slice mit durch Leerzeichen getrennte Worte
	
	for _,w:= range worte {
		// utf8.RuneCountInString(w) gibt die korrekte Anzahl an Zeichen in einem UTF8-Wort, also die Länge des Wortes
		zeilenLänge = uint16(len(zeile)*zeichenbreite + utf8.RuneCountInString(w)*zeichenbreite)
		if zeilenLänge <= tb.breite {		// Prüfe ob Zeilenlänge überschritten wird
			//l += utf8.RuneCountInString(w)	*zeichenbreite						// Füge Länge des Wortes zur Zeilenlänge hinzu
			zeile += w + " "					// Wort zur Zeile, inklusive Leerzeichen danach
			//fmt.Println(w, len(w),utf8.RuneCountInString(w))
			zeilenLänge_temp = zeilenLänge			// temp wird gebraucht um später den Text zentriert darstellen zu können
			

		} else {								// Zeile zu lang
			// Schreibe den Text entweder linksbündig oder zentriert
			if tb.linksbündig {
				gfx.SchreibeFont(tb.x,tb.y+uint16(zeilenHöhe)+zeilenAnzahl*tb.zeilenAbstand,zeile)
			} else {
				gfx.SchreibeFont(tb.x+(tb.breite-zeilenLänge_temp)/2,tb.y+uint16(zeilenHöhe)+zeilenAnzahl*tb.zeilenAbstand,zeile)	// Schreibe Zeile an richtige Stelle
			}
			
			zeilenHöhe+=tb.schriftgr					// erhöhe Zeilenanzahl
			zeilenLänge=uint16(utf8.RuneCountInString(w)*zeichenbreite)		// setze Zeilenlänge neu (sonst wäre sie jetzt zu lang
			zeilenLänge_temp = zeilenLänge									// Setze auch temp neu, sonst wird die falsche länge bei nur einem Wort pro Zeile genommen
			zeile = w + " "					// Beginne eine neue Zeile
			zeilenAnzahl+=1							// Erhöhe die Zeilenanzahl
			//fmt.Println(w, len(w),utf8.RuneCountInString(w))
		}
	}
	if tb.linksbündig {
		gfx.SchreibeFont(tb.x,tb.y+uint16(zeilenHöhe)+zeilenAnzahl*tb.zeilenAbstand,zeile)
	} else {
		gfx.SchreibeFont(tb.x+(tb.breite-zeilenLänge_temp)/2,tb.y+uint16(zeilenHöhe)+zeilenAnzahl*tb.zeilenAbstand,zeile)	// Schreibe Zeile an richtige Stelle
	}
}


