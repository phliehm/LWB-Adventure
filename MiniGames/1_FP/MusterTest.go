package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import (. "gfx"
		"./muster"
		"fmt")
	
func main () {
	Fenster (1200,700)										// Öffnet GFX-Fenster
	
	fmt.Println("Du hast ",muster.Muster()," Punkte erreicht!")
}
