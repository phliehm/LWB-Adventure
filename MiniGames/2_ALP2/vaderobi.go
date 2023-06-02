//Annalena Cyriacus
//Juni 2023
//Minigame "Vaderobi" (ALP2-Game)

package main

import (
	//"fmt"
	//"gfx"
	"./vaderobi"
	"../../Klassen/texteditoren"
)

//func Vaderobi() (note float32, punkte uint32) {
func Vaderobi() {
	vaderobi.Schrittmodus(false)
	//vaderobi.NewWelt(700,700,50)
	//vaderobi.Melden("Welt erstellt!",0)
	//fmt.Println(editor.Editor())
	var ted texteditoren.Texteditor
	ted = texteditoren.New(700,350,475,325,20,true)
	for {
		switch ted.GibString() {
			case "Laufen1()":
			vaderobi.Laufen1()
			case "LinksDrehen()":
			vaderobi.LinksDrehen()
			case "RechtsDrehen()":
			vaderobi.RechtsDrehen()
			case "AmRand()":
			vaderobi.AmRand()
		}
		//gfx.UpdateAus()
		ted = texteditoren.New(700,350,475,325,20,true)
		//vaderobi.Melden("Neuer Texteditor!",0)
		//gfx.Stiftfarbe(255,255,255)
		//gfx.Vollrechteck(700,675,475,325)
		//gfx.Vollrechteck(700,665,475,4)
		//gfx.UpdateAn()
	}
	vaderobi.Fertig()
	//gfx.TastaturLesen1()
	//fmt.Println(ted.GibString())
	
	//fmt.Println(ted.GibPosition())
	//return 0,0
}

//func eingabe

func main() {
	Vaderobi()
}
