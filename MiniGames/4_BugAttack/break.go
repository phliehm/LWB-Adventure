
package main

import (
	"time"
	 "gfx"
	//"math/rand"
//	"fmt"
	 )

func main() {
	gfx.Fenster(1200,700)
	var i uint16
	var bilder [13]string = [13]string{"../../Bilder/BugAttack/Amoebius40.bmp","../../Bilder/BugAttack/Amoebius50.bmp",
		"../../Bilder/BugAttack/Amoebius60.bmp","../../Bilder/BugAttack/Amoebius70.bmp","../../Bilder/BugAttack/Amoebius80.bmp",
		"../../Bilder/BugAttack/Amoebius90.bmp","../../Bilder/BugAttack/Amoebius100.bmp",
		"../../Bilder/BugAttack/Amoebius90.bmp","../../Bilder/BugAttack/Amoebius80.bmp",
		"../../Bilder/BugAttack/Amoebius70.bmp","../../Bilder/BugAttack/Amoebius60.bmp","../../Bilder/BugAttack/Amoebius50.bmp",
		"../../Bilder/BugAttack/Amoebius40.bmp"}
	for i=0;i<1000;i+=130{
		for j,b :=  range bilder {
			gfx.UpdateAus()
			gfx.Stiftfarbe(255,0,0)
			gfx.Cls()
			gfx.LadeBildMitColorKey(0+i+uint16(j)*10,100,b,0,0,0)
			gfx.UpdateAn()
			time.Sleep(1e8)
		}
	
	}
	
}
