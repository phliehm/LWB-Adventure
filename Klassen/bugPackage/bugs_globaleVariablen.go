package bugPackage


import ("../../Klassen/textboxen")

// Tastatur
var taste uint16
var gedrueckt uint8
var tiefe uint16

// Fenstermaße
const breite uint16 = 1200
const höhe uint16 = 700

const weltHIntro uint16 = 50
const weltB,weltH uint16 = 133,44
var welt [weltH][weltB] uint8 // Welt: Animation 1-3, Zeile, Spalte, Zahl
var weltIntro [weltHIntro][weltB] uint8
const y_offset uint16 = 6 // 5* zH offset damit oben ein schwarzer Balken ist
const zB, zH uint16 = 9,14			// Maße für die Zahlen 0 und 1 (Zellengröße)

var cursor_x, cursor_y uint16  = 0,y_offset*zH

var a uint8 =1 // Bug Animation

var bugArray [10]*bug

var punkteArray [2]uint32 // Punktestand für jedes Level 
var note float32	// Wird an die Main-Funktion des Spiels zurückgegeben

var manual string = "Bewegen :  Pfeiltasten  |  Größere Schritte:  SHIFT + Pfeiltasten  |  Aufgeben:  'q'\n\n" 

var sr,sg,sb uint8 = 0,0,0

var punkteTB textboxen.Textbox 


