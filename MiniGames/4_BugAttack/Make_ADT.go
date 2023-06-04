package main

import (
	"gfx"
	"time"
	"math/rand"
	"fmt"
	"sync"
	"math"
	"../../Klassen/textboxen"
	)

// Tastatur
var taste uint16
var gedrueckt uint8
var tiefe uint16

// Fenstermaße
const breite uint16 = 1200
const höhe uint16 = 700

const weltB,weltH uint16 = 133,44
var welt [3][weltH][weltB] uint8 // Welt: Animation 1-3, Zeile, Spalte, Zahl
const y_offset uint16 = 6 // 5* zH offset damit oben ein schwarzer Balken ist
const zB, zH uint16 = 9,14			// Maße für die Zahlen 0 und 1 (Zellengröße)

var cursor_x, cursor_y uint16  = 0,y_offset*zH

var a uint8 =1 // Bug Animation
var bugArray [10]*bug

var punkte uint32

var manual string = "Bewegen mit Pfeiltasten. Mit SHIFT und Pfeiltasten längere Schritte. 'q' beendet das Programm\n\n" 

var sr,sg,sb uint8 = 0,0,0



												


// Box hat Höhe von 14, Breite 9 (mit schwarzer Umrandung von überall 1px)
func male_1(x,y uint16){
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
	gfx.Stiftfarbe(sr,sg,sb)
	gfx.Punkt(x+2,y+4)
	gfx.Punkt(x+3,y+3)
	gfx.Linie(x+4,y+2,x+4,y+11)
	gfx.Linie(x+2,y+11,x+6,y+11)
	
}

func male_0(x,y uint16) {
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
	gfx.Stiftfarbe(sr,sg,sb)
	gfx.Linie(x+3,y+2,x+5,y+2)
	gfx.Linie(x+3,y+11,x+5,y+11)
	gfx.Linie(x+2,y+3,x+2,y+10)
	gfx.Linie(x+6,y+3,x+6,y+10)	
}

func male_schwarz(x,y uint16) {
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,zB,zH)
}
// Zeichne eine 0 oder 1
func male_Zahl(x,y uint16,z uint8) {
	switch z {
		case 0: male_0(x,y)
		case 1: male_1(x,y)
		case 2: male_schwarz(x,y)
	}
}

// Zeichne den Hintergrund/Code
func zeichneHintergrund(z uint8) {
	gfx.Sperren()
	gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	var x,y uint16
	for y=0;y<700;y+=zH {
		for x=0;x<1200;x+=zB {
			z = uint8(rand.Intn(2))
			male_Zahl(x,y,z)
		}
	}
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(x,y,7*zB,7*zH)
	gfx.UpdateAn()
	gfx.Entsperren()
}

// Zeichne Bug nr. 1
func male_bug1(x,y uint16,a uint8) {
	// Bug Größe: Breite: 7*9 , Höhe 7*14
	animation1 := func() {
		gfx.Stiftfarbe(0,0,0)
		gfx.Vollrechteck(x,y,7*zB,7*zH)
		gfx.Stiftfarbe(0,255,0)
		male_0(x+2*zB,y+2*zH)
		male_0(x+3*zB,y+2*zH)
		male_0(x+4*zB,y+2*zH)
		male_0(x+2*zB,y+3*zH)
		male_0(x+3*zB,y+3*zH)
		male_0(x+4*zB,y+3*zH)
		male_0(x+2*zB,y+4*zH)
		male_0(x+3*zB,y+4*zH)
		male_0(x+4*zB,y+4*zH)
	}
			
	animation2 := func() {
		male_1(x+1*zB,y+1*zH)
		male_1(x+5*zB,y+1*zH)
		
		male_1(x+1*zB,y+3*zH)
		male_1(x+5*zB,y+3*zH)
		
		male_1(x+1*zB,y+5*zH)
		male_1(x+5*zB,y+5*zH)
	}
	animation3 := func() {
		male_1(x,y)
		male_1(x+6*zB,y)

		male_1(x,y+3*zH)
		male_1(x+6*zB,y+3*zH)
		male_1(x,y+6*zH)
		male_1(x+6*zB,y+6*zH)
	}
	
	switch a { 
		case 1: 
			animation1()
		case 2: 
			animation1()
			animation2()
		case 3: 
			animation1()
			animation2()
			animation3()
			
		default: return
	}
}

// Zeichne den Cursor
func cursorZeichnen() {
		gfx.Stiftfarbe(0,255,0)
		gfx.Vollrechteck(cursor_x+2,cursor_y+2,zB-3,zH-3)
		gfx.Stiftfarbe(0,0,0)
}

// Update Cursor-Position
func cursorPos() {
	var step uint16 = 1
	for {
		gfx.Stiftfarbe(0,255,0)
		taste, gedrueckt, tiefe = gfx.TastaturLesen1()
		if tiefe==1 {
			step=10
		}else {step=1}
		if gedrueckt == 1 {
			
			switch taste {
				case 273:	
							cursor_y -= step*zH
							if cursor_y<y_offset*zH {cursor_y = y_offset*zH+weltH*zH-zH}	
							
				case 274:  
							cursor_y += step*zH
							if cursor_y>y_offset*zH+weltH*zH-zH {cursor_y = y_offset*zH}
				case 275:	
							cursor_x += step*zB
							if cursor_x>weltB*zB-zB {cursor_x = 0}
				case 276:	
							cursor_x -= step*zB
							if cursor_x>weltB*zB-zB {cursor_x = weltB*zB-zB}
				case 32 : 	
							welt[0][(cursor_y-y_offset*zH)/zH][cursor_x/zB] = 0
							
							bugGetroffen()
				case 'q':
					gfx.FensterAus()
					return
			
				default:
					continue				
			}
		fmt.Println(taste,tiefe,cursor_x,cursor_y)	
		}
		
		time.Sleep(1e7)
	}
}



// Zeichne den nächsten Frame
func zeichnen() {
	for {
		gfx.UpdateAus()
		gfx.Restaurieren(0,0,1200,700)
		male_bug1(50*9,6*14,a)
		cursorZeichnen()
		gfx.UpdateAn()
		time.Sleep(1e8)
	}
}

// füllt Welt mit Zahlen (Code)
func beschreibeArray(){
	var s,z uint16
	for a:=0;a<3;a++ {
		for z=0;z<weltH;z++ {
			for s=0;s<weltB;s++ {
				welt[a][z][s] = uint8(rand.Intn(2))
			}
		}
	}
}


func zeichneArray() {
	var s,z uint16
	//gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	//gfx.Stiftfarbe(0,255,0)
	sr,sg,sb = 0,255,0
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			male_Zahl(s*zB,y_offset*zH+z*zH,welt[0][z][s])
		}
	}
	//gfx.UpdateAn()
}


func zeichneWelt(punkteTB textboxen.Textbox) {
	for {
		gfx.UpdateAus()
		//gfx.Stiftfarbe(0,255,0)
		sg = 255
		zeichneArray()
		gfx.LadeBild(0,0,"../../Bilder/Amoebius_klein.bmp")
		punkteTB.SchreibeText(manual+"Punkte: "+fmt.Sprint((ZählePunkte())))
		punkteTB.Zeichne()
		for index,_ := range bugArray {
			if bugArray[index]==nil {continue}
			bugArray[index].zeichneBug()
		}
		//bugArray[0].zeichneBug()
		cursorZeichnen()
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
}

////////////////////
//// KLASSE bug ////
////////////////////










// Prüft ob ein Bug getroffen wurde und zerstört Bug oder lässt neuen wachsen, oder macht nix
func BugGetroffen() {
	for _,b:= range bugArray {
		if b==nil{continue}
		//fmt.Println("x: ",cursor_x, b.x, b.x+7*zB,"y: ",cursor_y, b.y,b.y + 7*14)
		if (cursor_x-3*zB == b.x && cursor_y-3*zH==b.y) {
			fmt.Println(cursor_x-3*zB,b.x,cursor_y-3*zH,b.y)
			fmt.Println("Getroffen!!!")
			gfx.SpieleSound("../../Sounds/Retro Sounds/Explosions/Long/sfx_exp_long1.wav")
			b.ende=true
			punkte+=1
		}else if (cursor_x > b.x+zB && cursor_x<b.x+6*zB) && (cursor_y> b.y+zH && cursor_y<b.y+6*zH) {
			fmt.Println("Oh nein!! Der Bug ist provoziert")
			gfx.SpieleSound("../../Sounds/Retro Sounds/General Sounds/Negative Sounds/sfx_sounds_damage1.wav")
			// Neuen Bug generieren
			babyBugs(b)
		}
				
	}
}

// Erzeugt einen neuen Bug an der Stelle eines anderen Bug
func BabyBugs(b *bug) {
	if howManyBugs() < len(bugArray) {
		// Prüfe ob es weniger als maxAnzahlBugs gibt
		for index,bu:= range bugArray {
			// Wenn noch Platz ist erzeuge neuen Bug
			if bu==nil {
				if b.x%zB!=0 || b.y%zH!=0 {fmt.Println("WRONG!!!!!",b.x,b.y)}
				bugArray[index]=NewBug(b.x/zB,b.y/zH) 
				bugArray[index].nervosität=5+int(rand.Intn(5))
				bugArray[index].speed = 1
				bugArray[index].b = uint8(25*bugArray[index].nervosität)
				go bugArray[index].bugAnimation()
				go bugArray[index].startMoving()
				break
			}
		}
	}
}

// Gibt Anzahl der Bugs zurück
func howManyBugs() int {
	var anzahl int
	for _,b:= range bugArray {
		if b!=nil {
			anzahl++
			b.eatCode()
			}
		
	}	
	return anzahl
}

// Löscht tote Bugs
func cleanBugArray() {
	for {
		for index,b:= range bugArray {
			// entferne Bug wenn er tot ist
			if b!=nil && b.alive == false { 
				bugArray[index] = nil
				
			}
		}
		time.Sleep(1e9)
	}
}

// Zum debugging, anzeigen wie viele Bugs es gibt
func ShowBugs(){
	for{
		time.Sleep(1e9)
		fmt.Println(howManyBugs())
	}
}


///////////////
//// LEVEL ////
///////////////

func Level1(){
	
	gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
	b1 := NewBug(50,30)
	b2 := NewBug(20,10)
	b3 := NewBug(40,15)
	bugArray[0] = b1
	bugArray[1] = b2
	bugArray[2] = b3
	b2.g = 50
	b2.speed=1
	b1.speed = 1
	b3.speed = 1
	go b1.bugAnimation()
	go b1.startMoving()
	go b2.bugAnimation()
	go b2.startMoving()
	go b3.bugAnimation()
	go b3.startMoving()
	go cleanBugArray()
	go ShowBugs()
	
	
}

// Zählt die Punkte im Array
func ZählePunkte() uint32 {
	var arrayPunkte uint32
	var z,s uint16 
	for z=0;z<weltH;z++ {
		for s=0;s<weltB;s++ {
			if welt[0][z][s] !=2 {arrayPunkte++}
		}
	}
	return arrayPunkte
}

func Level0(){
	gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
	b1 := NewBug(50,30)
	
	bugArray[0] = b1
	
	b1.speed = 0
	
	go b1.bugAnimation()
	go b1.startMoving()
	
	go cleanBugArray()
	go ShowBugs()
}



func main() {
	rand.Seed(time.Now().UnixNano())
	gfx.Fenster(1200,700)
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	//var z uint8 =0
	//zeichneHintergrund(z)
	//gfx.Archivieren()
	//go zeichnen()
	
	punkteTB := textboxen.New(200,10,1000,20)
	punkteTB.SchreibeText("HALLOOOOO")
	punkteTB.SetzeFarbe(255,255,255)
	punkteTB.SetzeFont("../../Schriftarten/collegeb.ttf")
	
	
	Level1()
	
	time.Sleep(1e9)
	go cursorPos()
	
	
	beschreibeArray()
	//zeichneArray()
	
	go zeichneWelt(punkteTB)
	
	
	
	
	for{time.Sleep(3e7)}
	
	
	
}
