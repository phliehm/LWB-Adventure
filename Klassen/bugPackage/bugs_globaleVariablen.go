package bugPackage


import (
		"../../Klassen/textboxen"
		"sync")

// Tastatur
var taste uint16
var gedrueckt uint8
var tiefe uint16

// Fenstermaße
const breite uint16 = 1200
const höhe uint16 = 700

const weltHIntro uint16 = 50
const weltH,weltB uint16 = 44,133
var welt [weltH][weltB] uint8 // Welt: Animation 1-3, Zeile, Spalte, Zahl
var weltIntro [weltHIntro][weltB] uint8
const y_offset uint16 = 6 // 6* zH offset damit oben ein schwarzer Balken ist
const zB, zH uint16 = 9,14			// Maße für die Zahlen 0 und 1 (Zellengröße)

var cursor_x, cursor_y uint16  = 0,y_offset*zH

var a uint8 =1 // Bug Animation

var bugArray [20]*bug
var bugArraySchloss sync.Mutex

var manual string = "Bewegen :  Pfeiltasten  |  Größere Schritte:  SHIFT + Pfeiltasten  |  Aufgeben:  'q'\n\n" 

var sr,sg,sb uint8 = 0,0,0

var punkteTB textboxen.Textbox 

var autoAim bool =true// Drücke "x" um automatisch zum nächsten Bug zu springen
var quit chan bool = make(chan bool)		// beendet die Animation von Amoebius und dem Bugs

var wg sync.WaitGroup

var level, anzahlBugsImLevel uint16
var levelSchloss sync.Mutex
var lvlSpeed, lvlNervosität int
var lvlLäuft bool
var lvlZeit uint16 
var lvlMaxPunkte [3]uint16 = [3]uint16{4800,4500,4000}
var lvlMinPunkte [3]uint16 = [3]uint16{4000,3000,2000}
var LevelArray [3]func() = [3]func(){Level1,Level2,Level3}

const maxPunkteProLevel uint16 = 5000
var punkteArray [3]uint16 // Punktestand für jedes Level 
var EndN float32
var EndP uint32
var SpielBeendet bool		// Spiel geht zum Endbildschirm wenn true

