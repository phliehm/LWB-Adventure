// Testprogramm für theNetgame - Spiel zur NET-Vorlesung

// Martin Seiß    25.5.2023


package main


import ga "./MiniGames/2_EthI"
//import "os"
//import "strconv"
import "fmt"
import "gfx"




func main() {

	gfx.Fenster(1200,700)
	note,punkte := ga.Getraenkeautomat()
	fmt.Println("Note: ",note) 
	fmt.Println("Punkte: ",punkte )
	
}
