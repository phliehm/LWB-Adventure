

// test StoppeAlleSounds
// Martin Seiß
// 6.6.2023

package main

import "fmt"
import "./gfx"
import "time"


var soundOn bool


func spiele() {
	
	soundOn = true
	for soundOn {
		gfx.SpieleSound("30s_Surf.wav")
		time.Sleep(40e9)
	}

	//gfx.StoppeAlleSounds()
	
}


func main() {
	
	gfx.Fenster(100,100)
	go spiele()			// Starte Musik-routine
	
	for {
		go spiele()			// Starte Musik-routine
		fmt.Println("ENTER für Stop!")
//		gfx.SpieleSound("30s_Surf.wav")
		fmt.Scanln()
		soundOn = false		// Stoppe Musik
		gfx.StoppeAlleSounds()
		fmt.Println("ENTER für Start!")
		fmt.Scanln()
		soundOn = true		// Stoppe Musik
	}
	
}
