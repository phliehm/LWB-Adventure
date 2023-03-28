package main

import (
	"fmt"
	"gfx"
	"math/rand"
	"strconv"
	"time"
)

const breite uint16 = 1000
const höhe uint16 = 800

var Taste uint16 = 's' // Variable für zu drückende Taste
var TastenArray [5]uint16 = [5]uint16{'s', 'o', 'p', 'u', 'b'}

func TesteTaste(taste uint16) {
	//gfx.SetzeFont("CollegiateBlackFLF.ttf",20)
	switch taste {
	case Taste:
		fmt.Println("Richtig! Das haben Sie aber toll geloest!!")
		gfx.SchreibeFont(200, 500, "Richtig! Das haben Sie aber toll geloest!!")
		gfx.LadeBildMitColorKey(650, 50, "../../Bilder/WtheK_black.bmp", 255, 255, 255)
	default:
		fmt.Println("Ohjeee, das ist ja falsch, wie kommen Sie de6nn darauf???")
		gfx.SchreibeFont(200, 500, "Ohjeee, das ist ja falsch, wie kommen Sie denn darauf???")
		gfx.LadeBildMitColorKey(650, 50, "../../Bilder/WtheK_black_sad.bmp", 255, 255, 255)
	}
	//time.Sleep(1e9)
}

func LoseTaste() uint16 {

	return TastenArray[rand.Intn(5)]

}

func main() {

	gfx.Fenster(breite, höhe)
	gfx.SetzeFont("../../Schriftarten/Ubuntu-B.ttf", 20)

A:
	for {
		var t_start int64 = time.Now().UnixNano()
		gfx.UpdateAus()
		gfx.Stiftfarbe(230, 255, 230)
		gfx.Vollrechteck(0, 0, breite, höhe)
		gfx.Stiftfarbe(0, 0, 0)
		gfx.Rechteck(100, 100, 400, 200)
		Taste = LoseTaste()
		gfx.SchreibeFont(150, 150, "Druecke "+string(Taste))
		gfx.LadeBildMitColorKey(650, 50, "../../Bilder/WtheK_black.bmp", 255, 255, 255)
		//gfx.LadeBildMitColorKey(50, 400, "LWB-A-BMP/3_Fachdidaktik_Planung.bmp", 255, 255, 255)
		fmt.Println("Rechteck")
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
		gfx.SchreibeFont(150, 200, "Das hat: "+strconv.Itoa(int(time.Now().UnixNano()-t_start)/1e6)+" ms gedauert!")
		time.Sleep(2e9)
	}

}
