//Annalena Cyriacus
//Juni 2023
//Klasse Texteditoren
//basierend auf dem Quelltext von Stefan Schmidt (LWB Informatik, ALP 3)

package texteditoren

import (
	"gfx"
	"./editor"
	"./editor1line"
)

type data struct {
	x,y,breite,höhe uint16 	// Koordinaten und breite und Höhe
	//const spaltenanzahl, zeilenanzahl = 64, 21 entspricht breite und höhe?!
	//font string
	schriftgr int
	text string
	einzeilig bool
}

func New(posx,posy,breite,höhe uint16, schriftgr int, einzeilig bool) *data {
	var ted *data
	ted = new(data)
	ted.x = posx
	ted.y = posy
	ted.höhe = höhe
	ted.breite = breite
	ted.einzeilig = einzeilig
	//ted.schriftgr = 20
	//ted.font = gfx.GibFont()
	if !gfx.FensterOffen () { gfx.Fenster(1200,700) }
	gfx.Stiftfarbe(0,0,0)
	gfx.Rechteck(posx,posy,breite,höhe)
	if einzeilig {
		ted.text = editor1line.Editor(posx,posy,breite,höhe,schriftgr)
	} else {
		ted.text = editor.Editor(posx,posy,breite,höhe,schriftgr)
	}
	return ted
}

/*
func darstellen(posx,posy,breite,höhe uint16) {
	if !gfx.FensterOffen () { gfx.Fenster(1200,700) }
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(posx,posy,breite,höhe)
	editor.Editor(posx,posy) 
}
*/

//-------------------------- Methoden ---------------------------- 

func (ted *data) GibPosition() (x,y uint16) {
	return ted.x, ted.y
}

func (ted *data) SetzePosition(xneu,yneu uint16) {
	ted.x = xneu
	ted.y = yneu
}

/*
func (ted *data) GibFont() string {
	return ted.font
}

func (ted *data) SetzeFont(fontneu string) {
	ted.font = 
}
*/

func (ted *data) GibSchriftgröße() (schriftgr int) {
	return ted.schriftgr
}

func (ted *data) SetzeSchriftgröße(schriftgrneu int) {
	ted.schriftgr = schriftgrneu
}

func (ted *data) GibHöhe() (höhe uint16) {
	return ted.höhe
}

func (ted *data) SetzeHöhe(höheneu uint16) {
	ted.höhe = höheneu
}

func (ted *data) GibBreite() (breite uint16) {
	return ted.breite
}

func (ted *data) SetzeBreite(breiteneu uint16) {
	ted.breite = breiteneu
}

func (ted *data) IstEinzeilig() bool {
	return ted.einzeilig
}

func (ted *data) NeuerTexteditor(posx,posy,breite,höhe uint16, schriftgr int, einzeilig bool) {
	ted = New(posx,posy,breite,höhe,schriftgr,einzeilig)
}

func (ted *data) GibString() string {
	return ted.text
}
