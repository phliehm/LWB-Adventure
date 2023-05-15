package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import (. "gfx"
		"./MiniGames/4_Moorhuhn/moorhuhn"
		"fmt")
	
func main () {
	Fenster (1200,700)										// Öffnet GFX-Fenster
	
	fmt.Println("Du hast ",moorhuhn.Moorhuhn()," Punkte erreicht!")
}
