package fachjargonPackage
import "fmt"

const breite uint16 = 1200
const höhe uint16 = 700

// Tastatur
var taste uint16
var gedrueckt uint8
var tiefe uint16

// Richtig oder Falsch

var Taste uint16  // Variable für zu drückende Taste
var TastenArray []uint16 = []uint16{'1','2','3','4','5','6','7','8'}		

var bilderArrayEntwicklung []string = []string{"1_Bildung.bmp","2_Bildungsstandards.bmp",
						"3_Kompetenzmodell.bmp","4_RLP_start.bmp","5_Basiskonzepte.bmp","6_Planung.bmp",
										"7_Stundenplanung.bmp","8_Unterricht.bmp"}
										
var bilderArrayTools []string = []string{"Tools_1_Standardsoftware.bmp","Tools_2_Netze.bmp",
	"Tools_3_Crypttool.bmp","Tools_4_Kara.bmp","Tools_5_SQLIsland.bmp","Tools_6_Flaci.bmp",
	"Tools_7_Johnny.bmp"}
	
var bilderArrayReihen []string = []string{}	

var textEntwicklung []string = []string{"1. Bildung","2. Bildungsstandards","3. Kompetenzmodell",
		"4. Rahmen(lehr)plan","5. fundamentale Ideen, Basiskonzepte","6. Planung der Unterichtseinheit",
		"7. Stundenplanung","8. Stundendurchführung"}
		
var textTools []string = []string{"1. Standardsoftware ","2. Leben in und mit vernetzten Systemen","3. Information und Daten",
				"4. Algorithmisches Problemlösen","5. Datenbanken","6. Sprachen und Automaten","7. Von-Neumann-Architektur "}
				
var textReihen []string = []string{"1. ","2. ","3. ","4. ","5. ","6. ","7. ","8. "}
								
const bilderPfad string = "Bilder/FachJargon/"
										

// Zeit

var t_start int64 

// Level
var level uint16 
var lvlLäuft bool
var lvlZeit uint16 // Zeit im Level
//const maxZeit uint16 = 30 // maximale Zeit
const maxLevel = 2 


// Punkte
var LevelArray [maxLevel]func() = [maxLevel]func(){Level1,Level2}

var punkteArray [maxLevel]uint16 // Punktestand für jedes Level 
var notenArray [maxLevel]float32 = [maxLevel]float32 {6.0,6.0} // Note für jedes Level

// Ausgabe am Ende
var EndN float32
var EndP uint32

// Beenden
var SpielBeendet bool

func printAlleGlobal() {
	fmt.Println(taste,gedrueckt,tiefe)
	fmt.Println(Taste, TastenArray)
	fmt.Println(t_start,level,lvlLäuft,lvlZeit)
	fmt.Println(EndN,EndP)
	fmt.Println(punkteArray)
	fmt.Println(notenArray)
}

func initialisiereGlobal() {
	taste,gedrueckt,tiefe = 0,0,0
	Taste = 0
	t_start,level = 0,0
	EndN, EndP = 0,0
	notenArray[0] = 6.0
	notenArray[1] = 6.0
}
