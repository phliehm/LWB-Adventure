package fachjargonPackage

import (
		"time"
		"gfx"
		"fmt"
		)

// Liest nur die Taste aus
func TastaturEingabe(c chan bool) {
	for {
		select {
			case <-c :
				return
			default:
				taste, gedrueckt, tiefe = gfx.TastaturLesen1()		
		}
	}
}


func leseTastaturStartbildschirm() {
	for SpielBeendet == false {
		switch taste{
			case 's': return 					// Oder Zwischenbildschirm? Count down?	
			default : time.Sleep(1e7)
		}
	}
}

func monitorSpielBeendet() {
	for SpielBeendet == false {
		switch taste {
			case 'q':
				SpielBeendet = true		// geht zum Endbildschirm
				fmt.Println("Spiel wird beendet")
				
			default : time.Sleep(1e8)
		}
	}
}
