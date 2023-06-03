package bugPackage

import (
	"gfx"
	"time"
	"../../Klassen/textboxen"
	"fmt"
	"math/rand"
	)






	
func Startbildschirm() {
	punkteTB = textboxen.New(200,10,1000,20)
	punkteTB.SetzeFarbe(255,255,255)
	punkteTB.SetzeFont("../../Schriftarten/ltypeb.ttf")
	
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	gfx.UpdateAus()
	gfx.LadeBild(5,5,"../../Bilder/Amoebius_klein.bmp")
	gfx.LadeBildMitColorKey(1050,530,"../../Bilder/FebWebK_red_gespiegelt.bmp",255,0,0)
	HelloTB := textboxen.New(130,50,800,500)
	HelloTB.SchreibeText("Willkomen beim Softwarepraktikum!\n\n" +
						"Du hast ja schon begonnen! Bevor wir euch eine Einweisung gegeben haben?\n\nVerständlich!\n\n"+
						"Endlich darfst du mal so richtig programmieren, genau das was du ja eigentlich wolltest, "+
						"ohne diesen ganzen Theoriekram.\n\nDoch scheinbar bist du doch nicht so cool wie du dachtest..."+
						"Oder hast du wieder während der Vorlesung programmiert??\n"+
						"Jedenfalls wimmelt es hier nur so von BUGs in deinem Code, deine TeamollegInnen werden nicht erfreut sein."+
						"\nWer will den schon so eine verBUGgte Klasse benutzen?"+
						"\n\nDie StudentInnen werden ja hier alle nicht besser, wir vom Dozententeam haben euch extra ein "+
						"halbautomatisches BUGFIXING-TOOL programmiert. Damit solltes selbst du in der Lage sein die Bugs zu beseitigen. "+
						"Sogar auf einem Apfel.\n\n"+
						"Gehe einfach mit deinem Cursor an genau die Stelle des Bugs, drücke LEERTASTE....und BAAAAMM!" )
	HelloTB.SetzeFont("../../Schriftarten/ltypeb.ttf")
	HelloTB.SetzeFarbe(0,255,0)
	HelloTB.SetzeSchriftgröße(25)
	HelloTB.Zeichne()
	
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}


func LevelIntro() {
	
	beschreibeArrayIntro()
	for i:=0;i<len(bugArray);i++ {
		bugArray[i] = NewBug(uint16(rand.Intn(100)),uint16(rand.Intn(40)))
		bugArray[i].nervosität = rand.Intn(10)+1
		bugArray[i].b = uint8(25*bugArray[i].nervosität)
		go bugArray[i].bugAnimation()
		go bugArray[i].startMoving()
	}		
	go cleanBugArray()		// Läuft dann für alle weiteren Level
	//go ShowBugs()
	go ZeichneWeltIntro()
	gfx.TastaturLesen1()
	for i:=0;i<10;i++ {
		bugArray[i] = nil
	}
	time.Sleep(1e9) // Wichtig damit die ZeichneWeltIntro() die Chance hat zu beenden
}

func Level0(){
	gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
	beschreibeArray()
	b1 := NewBug(50,30)
	
	bugArray[0] = b1
	
	b1.speed = 0
	
	go b1.bugAnimation()
	go b1.startMoving()
	
	//go cleanBugArray()
	//go ShowBugs()
	
	go ZeichneWelt(punkteTB)
	
	for howManyBugs() >0 {
		time.Sleep(1e9)
	}
	
	punkteArray[0] = zählePunkte()
}

func Level1(){
	beschreibeArray()
	//gfx.SpieleSound("../../Sounds/Music/bugWars.wav")
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
	
	go ZeichneWelt(punkteTB)
	
	//go cleanBugArray()
	//go ShowBugs()
	
	for howManyBugs()>0 {
		time.Sleep(1e9)
	}
	
	punkteArray[1] = zählePunkte()
	
	
}

func Level1Start() {
	
	gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	Level1StartTB:= textboxen.New(300,100,500,200)
	Level1StartTB.SchreibeText("Level 1 \n\n\nDie einfachen Bugs zuerst.\n\n\nPass auf, dass du nicht verschlimmbesserst!")
	Level1StartTB.SetzeFont("../../Schriftarten/ltypeb.ttf")
	Level1StartTB.SetzeSchriftgröße(30)
	Level1StartTB.SetzeFarbe(0,255,0)
	Level1StartTB.Zeichne()
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}


func Endbildschirm() {
	gfx.Cls()
	BugAttackTB := textboxen.New(200,100,700,500)
	BugAttackTB.SetzeZentriert()
	BugAttackTB.SchreibeText(
		"BUG ATTACK\nNote: "+fmt.Sprint(BerechneNote()))
	BugAttackTB.SetzeFont("../../Schriftarten/ltypeb.ttf")
	BugAttackTB.SetzeFarbe(0,255,0)
	BugAttackTB.SetzeSchriftgröße(100)
	BugAttackTB.Zeichne()
	fmt.Println(BerechneNote())
	gfx.TastaturLesen1()
}

func EndbildschirmReal() {
	var path string
	path = "../../"
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	
	
	gfx.LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	gfx.LadeBildMitColorKey(100,350,path + "Bilder/Amoebius_klein.bmp",0,0,0)
	gfx.LadeBild(620,80,path + "Bilder/paper_500.bmp")
	gfx.LadeBild(960,520,path + "Bilder/certified_100.bmp")
	gfx.LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
		
	gfx.Stiftfarbe(0,255,0)
	gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	gfx.SchreibeFont(330,10,"Bug  ATTACK")
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	gfx.SchreibeFont(295,140,"Du hast die")
	gfx.SchreibeFont(310,260,"erreicht!")
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	gfx.SchreibeFont(285,170,"Gesamtnote")
	gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	gfx.SchreibeFont(325,195,fmt.Sprintf("%2.1f",BerechneNote()))
	
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=1; i<7; i++ {
		gfx.SchreibeFont(710,150+uint16((i-1)*68), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(BerechneNote()) + " Punkte")
		gfx.SchreibeFont(710,175+uint16((i-1)*68),"           Note " + fmt.Sprintf("%2.1f",BerechneNote()))
	}
	gfx.SchreibeFont(700,130+uint16(6*70),"----------------------")
	gfx.SchreibeFont(710,160+uint16(6*70),"Gesamt:    " + fmt.Sprint(berechneSummeVonSlice(punkteArray[:])) + " Punkte")

	gfx.TastaturLesen1()
	//return gesamtnote, gesamtpunkte
}
