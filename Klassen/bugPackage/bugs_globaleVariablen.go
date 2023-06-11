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
var killNBugsCD uint16
var autoAimCD uint16

var alleLadebalken []*ladebalken 

var quit chan bool = make(chan bool)		// beendet die Animation von Amoebius und dem Bugs

var wg sync.WaitGroup

var level, anzahlBugsImLevel uint16
var levelSchloss sync.Mutex
var lvlSpeed, lvlNervosität int
var lvlLäuft bool
var lvlZeit uint16 // Zeit im Level
const maxZeit uint16 = 30 // maximale Zeit
const maxLevel = 6 
var lvlMaxPunkte [maxLevel]uint16 = [maxLevel]uint16{4850,4700,4000,3500,3000,1500}
var lvlMinPunkte [maxLevel]uint16 = [maxLevel]uint16{4500,3500,3000,1500,0,0}
var LevelArray [maxLevel]func() = [maxLevel]func(){Level1,Level2,Level3,Level4,Level5,Level6}

const maxPunkteProLevel uint16 = 5000
var punkteArray [maxLevel]uint16 // Punktestand für jedes Level 
var EndN float32
var EndP uint32
var SpielBeendet bool		// Spiel geht zum Endbildschirm wenn true

const xposAutoAimBalken,yposAutoAimBalken uint16 = 500,50
const xposkillNBugs,yposkillNBugs uint16 = 650,50
