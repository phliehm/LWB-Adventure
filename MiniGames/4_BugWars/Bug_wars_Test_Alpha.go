package main

import (
	"gfx"
	"time"
	"math/rand"
	"fmt"
	)

var taste uint16
var gedrueckt uint8
var tiefe uint16
var breite uint16 = 1200
var höhe uint16 = 700
var cursor_x, cursor_y uint16 
var a uint8 =1 // Bug Animation
var zB, zH uint16 = 9,14
var zeilen, spalten uint16 = 45,120
var welt [3][45][120] uint8 // Welt: Animation 1-3, Zeile, Spalte, Zahl
var bugArray [3]*bug
var ende bool



// Box hat Höhe von 14, Breite 9 (mit schwarzer Umrandung von überall 1px)
func male_1(x,y uint16){
	gfx.Punkt(x+2,y+4)
	gfx.Punkt(x+3,y+3)
	gfx.Linie(x+4,y+2,x+4,y+11)
	gfx.Linie(x+2,y+11,x+6,y+11)
	
}

func male_0(x,y uint16) {
	gfx.Linie(x+3,y+2,x+5,y+2)
	gfx.Linie(x+3,y+11,x+5,y+11)
	gfx.Linie(x+2,y+3,x+2,y+10)
	gfx.Linie(x+6,y+3,x+6,y+10)	
}

func male_Zahl(x,y uint16,z uint8) {
	switch z {
		case 0: male_0(x,y)
		case 1: male_1(x,y)
	}
}

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

func cursorZeichnen() {
		gfx.Stiftfarbe(0,255,0)
		gfx.Vollrechteck(cursor_x+2,cursor_y+2,zB-3,zH-3)
		gfx.Stiftfarbe(0,0,0)
}

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
				case 273:	cursor_y -= step*zH
				case 274:   cursor_y += step*zH
				case 275:	cursor_x += step*zB
				case 276:	cursor_x -= step*zB
				case 32 : 	welt[0][cursor_y/zH][cursor_x/zB] = 0
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

// füllt welt mit Zahlen
func beschreibeArray(){
	var s,z uint16
	for a:=0;a<3;a++ {
		for z=0;z<zeilen;z++ {
			for s=0;s<spalten;s++ {
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
	gfx.Stiftfarbe(0,255,0)
	for z=0;z<zeilen;z++ {
		for s=0;s<spalten;s++ {
			male_Zahl(s*zB,z*zH,welt[0][z][s])
		}
	}
	//gfx.UpdateAn()
}

func zeichneWelt() {
	for {
		gfx.UpdateAus()
		gfx.Stiftfarbe(0,255,0)
		zeichneArray()
		for index,_ := range bugArray {bugArray[index].zeichneBug()}
		//bugArray[0].zeichneBug()
		cursorZeichnen()
		gfx.UpdateAn()
		time.Sleep(1e7)
	}
}

type bug struct {
	//barray [3][7][7]uint8
	x,y uint16
	alive bool 
	dying uint16
	a uint16 // animation
	ende bool
}

func (b *bug) zeichneBug() {
	//fmt.Println("Bug wird gezeichnet")
	x:= b.x
	y:= b.y
	//gfx.UpdateAus()
	// Bug Größe: Breite: 7*9 , Höhe 7*14
	animation1 := func() {
		gfx.Stiftfarbe(0,0,0)
		gfx.Vollrechteck(x,y,7*zB,7*zH)
		gfx.Stiftfarbe(255,0,0)
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
	animation4 := func() {				// Implosion
		gfx.Stiftfarbe(0,0,0)
		for i:=0;i<500;i++ {
			gfx.Punkt(x+uint16(rand.Intn(63)),y+uint16(rand.Intn(98)))
		}
	}
	
	switch b.a { 
		case 1: 
			//fmt.Println("Animation1")
			animation1()
		case 2: 
			//fmt.Println("Animation2")
			animation1()
			animation2()
		case 3: 
			//fmt.Println("Animation3")
			animation1()
			animation2()
			animation3()
		case 4: 
			//fmt.Println("Animation4")
			animation1()
			animation2()
			animation3()
			for i:=uint16(0);i<b.dying;i++{
				animation4()
			}
		case 5:
			for i:=uint16(0);i<b.dying;i++{
				animation4()
			}
		
		default: return
	}
	
}

func (b *bug) bugAnimation() {
	b.alive=true				// belebt Bug
	for b.alive{				// Zeichne Bug nur wenn am leben

		if b.ende {				
			fmt.Println("No Animation")
			b.a = 4
			break
			}
		b.a=1
		time.Sleep(4e8)
		b.a=2
		time.Sleep(4e8)
		b.a=3
		time.Sleep(4e8)
	}
	
	for {
		if b.dying>50 {break}
		fmt.Println("Bug is dying:(")
		b.dying++
		time.Sleep(1e7)
	}
	b.a = 5
	for {
		if b.dying==0 {break}
		b.dying--
		time.Sleep(1e7)
	}
	
}

func bugGetroffen() {
	for _,b:= range bugArray {
		fmt.Println("x: ",cursor_x, b.x, b.x+7*zB,"y: ",cursor_y, b.y,b.y + 7*14)
		if (cursor_x > b.x && cursor_x<b.x+7*zB) && (cursor_y> b.y && cursor_y<b.y+7*14) {
			fmt.Println("Getroffen!!!")
			gfx.SpieleSound("../../Sounds/Retro Sounds/Explosions/Long/sfx_exp_long1.wav")
			b.ende=true
		}
	}
}

func main() {
	gfx.Fenster(1200,700)
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	//var z uint8 =0
	//zeichneHintergrund(z)
	//gfx.Archivieren()
	//go zeichnen()
	b1 := new(bug)
	b1.x=50*9
	b1.y=30*14
	b2 := new(bug)
	b2.x=20*9
	b2.y=10*14
	b3 := new(bug)
	b3.x=40*9
	b3.y=15*14
	bugArray[0] = b1
	bugArray[1] = b2
	bugArray[2] = b3
	time.Sleep(1e9)
	go cursorPos()
	beschreibeArray()
	//zeichneArray()
	go zeichneWelt()
	go b1.bugAnimation()
	go b2.bugAnimation()
	go b3.bugAnimation()
	for{time.Sleep(3e7)}
	
	
	
}
