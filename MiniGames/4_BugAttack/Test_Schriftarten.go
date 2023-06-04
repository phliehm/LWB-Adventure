package main

import (
		"../../Klassen/textboxen"
		"fmt"
		"gfx"
	)
func main() {
	var a float32
	a = 1.24214
	gfx.Fenster(1200,700)
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	
	t1:=textboxen.New(100,200,500,500)
	t1.SchreibeText("Note: 1.0 "+fmt.Sprint(a))
	t1.SetzeFont("../../Schriftarten/ltypeb.ttf")
	t1.SetzeFarbe(0,255,0)
	t1.Zeichne()
	gfx.TastaturLesen1()
	
}
