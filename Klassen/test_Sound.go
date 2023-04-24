package main

import (
	"gfx"
	"./audioloops"
	"fmt"
)




func main() {
	gfx.Fenster(1000,800)
	go audioloops.LoopDuration("../Sounds/Retro Sounds/Death Screams/Robot/sfx_deathscream_robot1.wav",1000,9000)
	//gfx.SpieleSound("Sounds/Retro Sounds/Death Screams/Robot/sfx_deathscream_robot1.wav")
	//gfx.SpieleSound("Sounds/Music/30s_Surf.wav")
	fmt.Println("Spiele Musik!")
	gfx.TastaturLesen1()
}
