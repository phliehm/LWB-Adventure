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


var bug1Shape [21][2] uint16 = [21][2]uint16{{0,0},{6*zB,0},
												{1*zB,1*zH},{5*zB,1*zH},
												{2*zB,2*zH},{3*zB,2*zH},{4*zB,2*zH},
												{0,3*zH},{1*zB,3*zH},{2*zB,3*zH},{3*zB,3*zH},{4*zB,3*zH},{5*zB,3*zH},{6*zB,3*zH},
												{2*zB,4*zH},{3*zB,4*zH},{4*zB,4*zH},
												{1*zB,5*zH},{5*zB,5*zH},
												{0,6*zH},{6*zB,6*zH}}
												


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
		//fmt.Println(taste,tiefe,cursor_x,cursor_y)	
		}
		time.Sleep(1e7)
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


// Datentyp bug
type bug struct {
	//barray [3][7][7]uint8
	x,y uint16
	alive bool 
	dying uint16
	a uint16 // AnimatiosSchritt
	ende bool 
	typ uint8 // Art des Bugs
	mu sync.Mutex	// Zum synchronisieren der Animation und Bewegung
	speed int // Geschwindigkeit des Bugs, also wie weit bewegt er sich pro Schritt
	nervosität int // wie oft bewegt sich der bug
	r,g,b uint8
}

func NewBug(x,y uint16) *bug {
	var b *bug
	b = new(bug)
	b.alive=true
	b.speed = 6
	b.nervosität = 1
	b.x = x*zB
	b.y = y*zH+y_offset*zH
	b.r,b.g,b.b = 255,0,0
	return b
	
}


func (b *bug) zeichneBug() {
	//fmt.Println("Bug wird gezeichnet")
	x:= b.x
	y:= b.y
	//gfx.UpdateAus()
	// Bug Größe: Breite: 7*9 , Höhe 7*14
	animation1 := func() {
		//gfx.Stiftfarbe(0,0,0)
		//gfx.Vollrechteck(x,y,7*zB,7*zH)
		//gfx.Stiftfarbe(b.r,b.g,b.b)
		sr,sg,sb = b.r,b.g,b.b
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
		sr,sg,sb = b.r,b.g,b.b
		male_1(x+1*zB,y+1*zH)
		male_1(x+5*zB,y+1*zH)
		
		male_1(x+1*zB,y+3*zH)
		male_1(x+5*zB,y+3*zH)
		
		male_1(x+1*zB,y+5*zH)
		male_1(x+5*zB,y+5*zH)
	}
	animation3 := func() {
		sr,sg,sb = b.r,b.g,b.b
		male_1(x,y)
		male_1(x+6*zB,y)

		male_1(x,y+3*zH)
		male_1(x+6*zB,y+3*zH)
		male_1(x,y+6*zH)
		male_1(x+6*zB,y+6*zH)
	}
	/*animation4 := func() {				// Implosion
		gfx.Stiftfarbe(0,0,0)
		for i:=0;i<500;i++ {
			gfx.Punkt(x+uint16(rand.Intn(int(7*zB))),y+uint16(rand.Intn(int(7*zH))))
		}
	}*/
	implosion:=func(ix,iy uint16) {
		gfx.Punkt(ix+uint16(rand.Intn(int(zB))),iy+uint16(rand.Intn(int(zH))))
	}
	animation4 := func() {				// Implosion
		gfx.Stiftfarbe(0,0,0)
		for i:=0;i<10;i++ {
			for k:=range bug1Shape {
				implosion(x+bug1Shape[k][0],y+bug1Shape[k][1])
			}
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




func versetzeBug(x,y uint16, bu *bug) uint16{	
	// Wenn sich zwei Bugs überdecken
	if math.Abs(float64(x)-float64(bu.x))<float64(7*zB) && math.Abs(float64(y)-float64(bu.y))<float64(7*zH) {
		//fmt.Println("Versetzt!")
		return uint16(int(x)-int(bu.x)+int(x))
		
	}
	return x
}

func (b *bug) eatCode() {
		welt[0][(b.y)/zH-3][(b.x)/zB+3] = 2
		
}

// Bewegung des Bugs
func (b *bug) startMoving() {
	for b.ende == false {		// Bug stirbt gerade, hör auf dich zu bewegen
		var new_x, new_y uint16
		// neue Koordinaten, abhängig von Geschwindigkeit (wie viele Felder)
		new_x = uint16(int(b.x)+(-b.speed/2 +(rand.Intn(1+b.speed)))*int(zB)) // speed = 5, --> -5/2 = -2; -2+rand(6)-1
		new_y = uint16(int(b.y)+(-b.speed/2 +(rand.Intn(1+b.speed)))*int(zH)) 
		fmt.Println("Random: ",-b.speed/2 +(rand.Intn(1+b.speed)-1))
		// Kollision mit anderen Bugs
		for _,bu:= range bugArray {
			if bu==nil {continue}		// wenn es ein andere Bug ist
			if b!=bu {					// wenn nicht ich der Bug bin
				new_x = versetzeBug(new_x,new_y,bu)	// verschiebe mich
			}
				
		}
		
		
		// Randbedingungen
		if 	new_x < zB*7{
			b.x =weltB*zB-7*zB
		}else if  new_x > weltB*zB-zB*7{
			b.x=0
		}else {b.x = new_x}
		
		if  new_y < 7*zH+y_offset*zH {
			b.y=weltH*zH+y_offset*zH-7*zH
		}else if new_y > weltH*zH+y_offset*zH- 7*zH{
			b.y=y_offset*zH
		}else {b.y = new_y} 
 		
	time.Sleep(time.Duration(1e9+rand.Intn(5000/b.nervosität)*1e6)) // Warte zufällige Zeit bevor sich Position ändert
	}
}

// Belebt und zeichnet Bug
func (b *bug) bugAnimation() {
	b.alive=true				// belebt Bug
	//if b.typ == 5 {bug.startMoving()}
	for b.alive{				// Zeichne Bug nur wenn am leben

		if b.ende {				
			fmt.Println("No Animation")
			b.a = 4
			break
			}
		// Animation startet

		b.a=1
		time.Sleep(3e8)
		b.a=2
		time.Sleep(3e8)
		b.a=3
		time.Sleep(3e8)

	}
	// Bug wird zersetzt
	for {
		if b.dying>50 {break}
		//fmt.Println("Bug is dying:(")
		b.dying++
		time.Sleep(1e7)
	}
	b.a = 5			// Code Reperatur
	for {
		if b.dying==0 {break}
		b.dying--
		time.Sleep(1e7)
	}
	b.alive = false
	
}

// Prüft ob ein Bug getroffen wurde und zerstört Bug oder lässt neuen wachsen, oder macht nix
func bugGetroffen() {
	for _,b:= range bugArray {
		if b==nil{continue}
		//fmt.Println("x: ",cursor_x, b.x, b.x+7*zB,"y: ",cursor_y, b.y,b.y + 7*14)
		if (cursor_x-3*zB == b.x && cursor_y-3*zH==b.y) {
			fmt.Println(cursor_x-3*zB,b.x,cursor_y-3*zH,b.y)
			fmt.Println("Getroffen!!!")
			gfx.SpieleSound("../../Sounds/Retro Sounds/Explosions/Long/sfx_exp_long1.wav")
			b.ende=true
			punkte+=1
			return
		}else if (cursor_x > b.x+zB && cursor_x<b.x+6*zB) && (cursor_y> b.y+zH && cursor_y<b.y+6*zH) {
			fmt.Println("Oh nein!! Der Bug ist provoziert")
			gfx.SpieleSound("../../Sounds/Retro Sounds/General Sounds/Negative Sounds/sfx_sounds_damage1.wav")
			// Neuen Bug generieren
			babyBugs(b)
			return
		}
				
	}
	// keinen Bug getroffen, mache Feld schwarz
	welt[0][cursor_y/zH-y_offset][cursor_x/zB] = 2
}

// Erzeugt einen neuen Bug an der Stelle eines anderen Bug
func babyBugs(b *bug) {
	if howManyBugs() < len(bugArray) {
		// Prüfe ob es weniger als maxAnzahlBugs gibt
		for index,bu:= range bugArray {
			// Wenn noch Platz ist erzeuge neuen Bug
			if bu==nil {
				if b.x%zB!=0 || b.y%zH!=0 {fmt.Println("WRONG!!!!!",b.x,b.y)}
				bugArray[index]=NewBug(b.x/zB,b.y/zH) 
				bugArray[index].nervosität=20+int(rand.Intn(5))
				bugArray[index].speed = 4
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
///////////////
//// LEVEL ////
///////////////

func Level1(){
	beschreibeArray()
	gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
	b1 := NewBug(50,30)
	b2 := NewBug(20,10)
	b3 := NewBug(40,15)
	bugArray[0] = b1
	bugArray[1] = b2
	bugArray[2] = b3
	b2.g = 50
	b2.speed=2
	b1.speed = 2
	b3.speed = 2
	go b1.bugAnimation()
	go b1.startMoving()
	go b2.bugAnimation()
	go b2.startMoving()
	go b3.bugAnimation()
	go b3.startMoving()
	go cleanBugArray()
	go ShowBugs()
	
	
}




func Level0(){
	gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
	beschreibeArray()
	b1 := NewBug(50,30)
	
	bugArray[0] = b1
	
	b1.speed = 0
	
	go b1.bugAnimation()
	go b1.startMoving()
	
	go cleanBugArray()
	go ShowBugs()
	for howManyBugs() >0 {
		time.Sleep(1e8)
	}
}

func Startbildschirm() {
	gfx.UpdateAus()
	HelloTB := textboxen.New(200,10,800,500)
	HelloTB.SchreibeText("Willkomen beim Softwarepraktikum!\n" +
						"Du hast ja schon begonnen? Bevor wir euch eigentlich eine Einweisung gegeben haben? Verständlich!\n"+
						"Endlich darfst du mal so richtig programmieren, genau das was du ja eigentlich wolltest, "+
						"ohne diesen ganzen Theoriekram.\n\nDoch scheinbar bist du doch nicht so cool wie du dachtest..."+
						"Oder hast du wieder während der Vorlesung programmiert??\n"+
						"Jedenfalls wimmelt es hier nur so von Bugs in deinem Code, deine TeamollegInnen werden nicht erfreut sein."+
						"\nWer will den schon so eine verbuggte Klasse benutzen?"+
						"\n\nDie StudentInnen werden ja hier alle nicht besser, wir vom Dozententeam haben euch extra ein "+
						"halbautomatisches Bugfixing-Tool programmiert." )
	HelloTB.SetzeFont("../../Schriftarten/Ubuntu-B.ttf")
	HelloTB.SetzeFarbe(0,255,0)
	HelloTB.Zeichne()
	gfx.LadeBild(0,0,"../../Bilder/Amoebius_klein.bmp")
	gfx.LadeBildMitColorKey(1000,550,"../../Bilder/FebWebK_red.bmp",255,0,0)
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}

func Level1Start() {
	
	gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	Level1StartTB:= textboxen.New(300,100,500,200)
	Level1StartTB.SchreibeText("Level 1 \n\n\nDie einfachen Bugs zuerst.\n\n\nPass auf, dass du nicht verschlimmbesserst!")
	Level1StartTB.SetzeFont("../../Schriftarten/Ubuntu-B.ttf")
	Level1StartTB.SetzeSchriftgröße(30)
	Level1StartTB.SetzeFarbe(0,255,0)
	Level1StartTB.Zeichne()
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}

func main() {
	rand.Seed(time.Now().UnixNano())		// Seed für Zufallszahlen
	gfx.Fenster(1200,700)
	gfx.Stiftfarbe(0,0,0)
	gfx.Vollrechteck(0,0,1200,700)
	gfx.Stiftfarbe(0,255,0)
	
	Startbildschirm() 
	time.Sleep(5e8)
	Level1Start()
	
	punkteTB := textboxen.New(200,10,1000,20)
	punkteTB.SetzeFarbe(255,255,255)
	punkteTB.SetzeFont("../../Schriftarten/Ubuntu-B.ttf")
	fmt.Println("geht los")
	go cursorPos()
	
	
	
	//zeichneArray()
	
	go zeichneWelt(punkteTB)
	Level0()
	Level1()
	
	//time.Sleep(1e9)
	
	
	
	
	
	for{time.Sleep(3e7)}
	
	
	
}
