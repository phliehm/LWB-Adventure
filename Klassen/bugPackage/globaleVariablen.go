/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 */

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

const weltHIntro uint16 = 50		// Höhe der Intro Welt (größer als sonst)
const weltH,weltB uint16 = 44,133	// Höhe und Breite der Welt in Zahlenpixeln
var welt [weltH][weltB] uint8 // Welt: Zeile, Spalte, Zahl (0,1,2)
var weltIntro [weltHIntro][weltB] uint8 
const y_offset uint16 = 6 // 6* zH offset damit oben ein schwarzer Balken ist
const zB, zH uint16 = 9,14			// Maße für die Zahlen 0 und 1 (Zellengröße)

var cursor_x, cursor_y uint16  = 0,y_offset*zH		// Cursor-Position

//var a uint8 =1 // Bug Animation						

var bugArray [20]*bug				// maximal 20 Bugs
var bugArraySchloss sync.Mutex		// Schloss zum Schützen des Arrays

// String für die Hinweise im Level
var manual string = "Bewegen :  Pfeiltasten  |  Größere Schritte:  SHIFT + Pfeiltasten  |  Aufgeben:  'q'\n\n" 

// Globale Farbe
var sr,sg,sb uint8 = 0,0,0

// Texbox für die Punkte
var punkteTB textboxen.Textbox 

//var autoAim bool =true		// Drücke "x" um automatisch zum nächsten Bug zu springen
var killNBugsCD uint16		// Cooldown der Fähigkeit Kill-N-Bugs
var autoAimCD uint16		// Cooldown der Fähigkeit autoAim

var alleLadebalken []*ladebalken 	// Array welcher die Ladebalken enthält

var quit chan bool = make(chan bool)		// beendet die Animation von Amoebius und dem Bugs

var wg sync.WaitGroup					// Wait-Group, nötig damit auf Go-Routinen am Ende eines Levels gewartet wird

var level, anzahlBugsImLevel uint16		// Welches Level, wieviele Bugs sind am Anfang im Level
var levelSchloss sync.Mutex				// Schloss zum Schutz der level Variable
var lvlSpeed, lvlNervosität int			// Speed und Nervosität der Bugs ist vom Level anhängig
var lvlLäuft bool						// Läuft das Level noch
var lvlZeit uint16 // Zeit im Level	
const maxZeit uint16 = 30 // maximale Zeit
const maxLevel = 6 						// Anzahl der Level
var lvlMaxPunkte [maxLevel]uint16 = [maxLevel]uint16{4850,4700,4000,3500,3000,1500}	// Maximal mögliche Punktzahl, entsprich 1.0
var lvlMinPunkte [maxLevel]uint16 = [maxLevel]uint16{4500,3500,3000,1500,0,0}		// Minimale Punktzahl, entspricht 5.0
var LevelArray [maxLevel]func() = [maxLevel]func(){Level1,Level2,Level3,Level4,Level5,Level6}	// Level-Funktionen die von der Game-Funktion gestartet werden

const maxPunkteProLevel uint16 = 5000			// Maximal mögliche Punkte in jedem Level
var punkteArray [maxLevel]uint16 				// Punktestand für jedes Level 
var EndN float32								// Endnote
var EndP uint32									// Endpunktestand
var SpielBeendet bool		// Spiel geht zum Endbildschirm wenn true

const xposAutoAimBalken,yposAutoAimBalken uint16 = 500,50	// Position des autoAim Balkens
const xposkillNBugs,yposkillNBugs uint16 = 650,50			// Position des killNBugs Balkens
