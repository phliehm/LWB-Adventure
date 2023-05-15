package main
// Autor: B. Schneider
// Datum: 20.03.2023
// Zweck: SWP - Minigame Moorhuhn
//--------------------------------------------------------------------

import (. "gfx"
		"./MiniGames/4_Moorhuhn"
		"fmt")
	
func main () {
	Fenster (1200,700)										// Öffnet GFX-Fenster
	i,j := moorhuhn.Moorhuhn()
	// i,j := float32(1.30000),uint32(123)
	
	fmt.Println( fmt.Sprint(i) )
	
	fmt.Printf("Du hast %d Punkte und die Note %.1f erreicht!",j,i)
}
