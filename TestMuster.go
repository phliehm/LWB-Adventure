package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Testdatei zum Muster-Minigame 
//--------------------------------------------------------------------

import (. "gfx"
		"./MiniGames/1_FP"
		"fmt")
	
func main () {
	Fenster (1200,700)										// Öffnet GFX-Fenster
	
	note,punkte := muster.Muster()
	fmt.Println("Du hast ",punkte," Punkte erreicht!")
	fmt.Println("Damit erreichst du die Note ",note," .\nHerzlichen Glückwunsch!")
}
