package fachjargon

import (
	"fmt"
	"gfx"
	"math/rand"
	"strconv"
	"time"
	"os"
)

const bilderPfad string = "Bilder/FachJargon/"

const breite uint16 = 1200
const höhe uint16 = 700

var Taste uint16  // Variable für zu drückende Taste
var TastenArray [5]uint16 = [5]uint16{'4','5','6','7','8'}

var BilderArray [5]string = [5]string{"4_RLP_start.bmp","5_Basiskonzepte.bmp","6_Planung.bmp",
										"7_Stundenplanung.bmp","8_Unterricht_nf.bmp"}

func TesteTaste(taste uint16) {
	switch taste {
	case Taste:
		gfx.SchreibeFont(650, 500, "Richtig! Das haben Sie aber toll geloest!!")
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
	default:
		gfx.SchreibeFont(650, 500, "Ohjeee, das ist ja falsch, wie kommen Sie denn darauf???")
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black_sad.bmp", 255, 255, 255)
	}
	//time.Sleep(1e9)
}

func LoseTaste() uint16 {

	return TastenArray[rand.Intn(5)]

}

// Startbildschirm 
func startBildschirm() uint16{
	gfx.Stiftfarbe(230, 255, 230)
	gfx.Vollrechteck(0, 0, breite, höhe)
	gfx.Stiftfarbe(0, 0, 0)
	gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
	gfx.SchreibeFont(200, 500, "Hallo! Ich begrüsse Sie zur heutigen Fachdidaktik Veranstaltung!")
	time.Sleep(1e9)
	gfx.SchreibeFont(200, 540, "Kennen Sie diese 8 Schritte noch? Bestimmt!")
	AchtSchritteText()
	gfx.SchreibeFont(200, 560, "Ich zeige Ihnen jetzt Bilder und Sie drücken die zugehörigen Tasten 1-8.")
	gfx.LadeBild(0, 0, bilderPfad+"3_Fachdidaktik_Planung.bmp")
	gfx.SchreibeFont(200, 600, "Drücken Sie 'q' zum Beenden oder 's' um das Spiel zu starten")
	for {
		taste, gedrueckt, _ := gfx.TastaturLesen1()
		if gedrueckt == 1 {
			switch taste {
			case 'q':
				return 'q'
			case 's':
				return 's'
			default:
				continue
			}
		}
	}
	
	return 's'
} 

// Schreibe 8 Schritte zur Kompetenzentwicklung
func AchtSchritteText() {
	var zeilenAbstand uint16 = 20 // Zeilenabstand
	var text [8]string = [8]string{"1. Bildung","2. Bildungsstandards","3. Kompetenzmodell",
		"4. Rahmen(lehr)plan","5. fundamentale Ideen, Basiskonzepte","6. Planung der Unterichtseinheit",
		"7. Stundenplanung","8. Stundendurchführung"}
	var i uint16
	for i=0;i<uint16(len(text));i++ {gfx.SchreibeFont(50,20 + i*zeilenAbstand,text[i])} 
}

func FachJargon() (float32,uint32){
	if ! gfx.FensterOffen() {
			gfx.Fenster(breite,höhe)
		}

	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 20)
	
	if startBildschirm()  == 'q' {os.Exit(0)}
	time.Sleep(1e9)

A:
	for {
		var t_start int64 = time.Now().UnixNano()
		
		gfx.UpdateAus()
		gfx.Stiftfarbe(230, 255, 230)
		gfx.Vollrechteck(0, 0, breite, höhe)
		gfx.Stiftfarbe(0, 0, 0)
		//gfx.Rechteck(100, 100, 400, 200)
		Taste = LoseTaste()
		fmt.Println(Taste-'0'-4)
		//gfx.SchreibeFont(150, 350, "Druecke "+string(Taste))
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
		//gfx.LadeBildMitColorKey(50, 400, "LWB-A-BMP/3_Fachdidaktik_Planung.bmp", 255, 255, 255)
		AchtSchritteText()
		gfx.LadeBild(50,200,bilderPfad+BilderArray[Taste-'0'-4])
		gfx.UpdateAn()

		taste, gedrueckt, _ := gfx.TastaturLesen1()
		if gedrueckt == 1 {
			switch taste {
			case 'q':
				break A

			default:
				TesteTaste(taste)
			}
		}
		gfx.SchreibeFont(650, 550, "Das hat "+strconv.Itoa(int(time.Now().UnixNano()-t_start)/1e6)+" ms gedauert!")
		time.Sleep(2e9)
	}
	return 1.0,5000

}
