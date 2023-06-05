package bugPackage

import (
		"gfx"
		"fmt"
		"time"
		"math/rand"
		"math"
		)


var bug1Shape [21][2] uint16 = [21][2]uint16{{0,0},{6*zB,0},
												{1*zB,1*zH},{5*zB,1*zH},
												{2*zB,2*zH},{3*zB,2*zH},{4*zB,2*zH},
												{0,3*zH},{1*zB,3*zH},{2*zB,3*zH},{3*zB,3*zH},{4*zB,3*zH},{5*zB,3*zH},{6*zB,3*zH},
												{2*zB,4*zH},{3*zB,4*zH},{4*zB,4*zH},
												{1*zB,5*zH},{5*zB,5*zH},
												{0,6*zH},{6*zB,6*zH}}


type bug struct {
	x,y uint16		// x,y geben die linke obere Ecke des Bugs an
	alive bool 
	dying uint16
	a uint16 // AnimatiosSchritt
	ende bool 
	typ uint8 // Art des Bugs
	//mu sync.Mutex	// Zum synchronisieren der Animation und Bewegung
	speed int // Geschwindigkeit des Bugs, also wie weit bewegt er sich pro Schritt
	nervosität int // wie oft bewegt sich der bug
	r,g,b uint8
}


func NewBug(x,y uint16) *bug{
	var b *bug
	b = new(bug)
	b.alive=true
	b.speed = 5
	b.nervosität = 1
	b.x = x*zB
	b.y = y*zH+y_offset*zH
	b.r,b.g,b.b = 255,0,0
	return b
}

func (b *bug) zeichneBug() {
	x:= b.x
	y:= b.y
	// Bug Größe: Breite: 7*9 , Höhe 7*14
	animation1 := func() {
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
		sr,sg,sb = b.r,b.g,b.b		// Wie ändere ich global die Farbe?
		male_1(x+1*zB,y+1*zH)
		male_1(x+5*zB,y+1*zH)
		
		male_1(x+1*zB,y+3*zH)
		male_1(x+5*zB,y+3*zH)
		
		male_1(x+1*zB,y+5*zH)
		male_1(x+5*zB,y+5*zH)
	}
	animation3 := func() {
		sr,sg,sb = b.r,b.g,b.b		//?
		male_1(x,y)
		male_1(x+6*zB,y)

		male_1(x,y+3*zH)
		male_1(x+6*zB,y+3*zH)
		male_1(x,y+6*zH)
		male_1(x+6*zB,y+6*zH)
	}

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

// Hilfsfunktion um einen Bug bei Kollision zu versetzen
func versetzeBug(x,y uint16, bu *bug) uint16{	
	// Wenn sich zwei Bugs überdecken
	if math.Abs(float64(x)-float64(bu.x))<float64(7*zB) && math.Abs(float64(y)-float64(bu.y))<float64(7*zH) {
		//fmt.Println("Versetzt!")
		return uint16(int(x)-int(bu.x)+int(x))
		
	}
	return x
}

func (b *bug) eatCode() {
		fmt.Println("Bug: ",b.y,b.x)
		welt[(b.y)/zH-3][(b.x)/zB+3] = 2
		/*welt[(b.y)/zH-2][(b.x)/zB+3] = 2
		welt[(b.y)/zH-4][(b.x)/zB+3] = 2
		welt[(b.y)/zH-3][(b.x)/zB+2] = 2
		welt[(b.y)/zH-2][(b.x)/zB+2] = 2
		welt[(b.y)/zH-4][(b.x)/zB+2] = 2
		welt[(b.y)/zH-3][(b.x)/zB+4] = 2
		welt[(b.y)/zH-2][(b.x)/zB+4] = 2
		welt[(b.y)/zH-4][(b.x)/zB+4] = 2
		*/
}

// Bewegung des Bugs
func (b *bug) startMoving() {
	
	for b.ende == false {		
		time.Sleep(time.Duration(1e9+rand.Intn(5000/b.nervosität)*1e6)) // Warte zufällige Zeit bevor sich Position ändert
		bugArraySchloss.Lock()
		var new_x, new_y uint16
		new_x = uint16(int(b.x)+((-b.speed/2 +rand.Intn(1+b.speed)))*int(zB))
		new_y = uint16(int(b.y)+((-b.speed/2 +rand.Intn(1+b.speed)))*int(zH))
		
		//fmt.Println("x: ",((-b.speed/2 +-1+rand.Intn(1+b.speed)))*int(zB)," y: ",((-b.speed/2 +-1+rand.Intn(1+b.speed)))*int(zH))
		xposWrite = append(xposWrite,(((-b.speed-1)/2 +rand.Intn(1+b.speed)))*int(zB))
		// Kollision mit anderen Bugs
		for _,bu:= range bugArray {
			if bu==nil {continue}
			if b!=bu {
				new_x = versetzeBug(new_x,new_y,bu)
			}
				
		}
		
		
		// Randbedingungen
		if 	new_x < 0{
			b.x =weltB*zB-7*zB		// 7 ist die Breite eines Bugs
		}else if  new_x > weltB*zB-zB*7{
			b.x=0
		}else {b.x = new_x}
		
		if  new_y < y_offset*zH {
			b.y=weltH*zH+y_offset*zH-7*zH
		}else if new_y > weltH*zH+y_offset*zH- 7*zH{
			b.y=y_offset*zH
		}else {b.y = new_y} 
		
 		bugArraySchloss.Unlock()
	}
	
}

// Belebt und zeichnet Bug
func (b *bug) bugAnimation() {
	b.alive=true				// belebt Bug
	//if b.typ == 5 {bug.startMoving()}
	for b.alive{				// Zeichne Bug nur wenn am leben

		if b.ende {
			b.a = 4
			break
		}
		// Animation startet

		b.a=1
		time.Sleep(4e8)
		if b.ende {
			b.a = 4
			break
		}
		b.a=2
		time.Sleep(4e8)
		if b.ende {
			b.a = 4
			break
		}
		b.a=3
		time.Sleep(4e8)
		

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

func rundeAufGeradeZahlen(z int) int{
	if z%2 == 0 {return z}
	return z+1
}

// Erzeugt einen neuen Bug an der Stelle eines anderen Bug
func babyBugs(b *bug) {
	
	if howManyBugs() < len(bugArray) {
		bugArraySchloss.Lock()
		// Prüfe ob es weniger als maxAnzahlBugs gibt
		for index,bu:= range bugArray {
			// Wenn noch Platz ist erzeuge neuen Bug
			if bu==nil {
				if b.x%zB!=0 || b.y%zH!=0 {fmt.Println("WRONG!!!!!",b.x,b.y)}
				bugArray[index]=NewBug(b.x/zB,b.y/zH-y_offset) 
				bugArray[index].nervosität=5+rand.Intn(5)	
				bugArray[index].speed = rundeAufGeradeZahlen(rand.Intn(5)) // Bei ungeraden Zahlen bewegen sich die Bugs ungleichmäßig
				bugArray[index].b = uint8(25*bugArray[index].nervosität)
				go bugArray[index].bugAnimation()
				go bugArray[index].startMoving()
				break
			}
		}
		bugArraySchloss.Unlock()
	}
	
	
}

// Gibt Anzahl der Bugs zurück
func howManyBugs() int {
	var anzahl int
	bugArraySchloss.Lock()
	for _,b:= range bugArray {
		if b!=nil {
			anzahl++
			b.eatCode()
			}
	}
	bugArraySchloss.Unlock()	
	//time.Sleep(1e6)
	return anzahl
}

// Löscht tote Bugs
func cleanBugArray() {
	for {
		bugArraySchloss.Lock()
		for index,b:= range bugArray {
			// entferne Bug wenn er tot ist
			if b!=nil && b.alive == false { 
				bugArray[index] = nil
				
			}
		}
		bugArraySchloss.Unlock()
		time.Sleep(1e9)
	}
}

func killAllBugs() {
	gfx.SpieleSound("../../Sounds/Retro Sounds/Explosions/Long/sfx_exp_long3.wav")
	bugArraySchloss.Lock()
	for _,b:= range bugArray {
		if b!=nil {
			b.ende = true
		}
	}
	bugArraySchloss.Unlock()
}
// Zum debugging, anzeigen wie viele Bugs es gibt
func ShowBugs(){
	for{
		time.Sleep(1e9)
		fmt.Println(howManyBugs())
	}
}

//  Es werden n Bugs erzeugt, speed und nervosität werden zufällig bis 
//	zur maximalen speed, nervosität gewählt
func createNBugs(n uint16,speed,nervosität int) {
	bugArraySchloss.Lock()
	for i:=uint16(0);i<n;i++ {
		fmt.Println(i)
		b := NewBug(uint16(rand.Intn(130)),uint16(rand.Intn(41)))
		b.speed = rand.Intn(speed + 1)
		b.g = uint8(25*b.speed)
		b.nervosität = rand.Intn(nervosität)+1
		b.b = uint8(25*b.nervosität)
		go b.bugAnimation()
		go b.startMoving()
		bugArray[i] = b
		
	}
	bugArraySchloss.Unlock()
}
