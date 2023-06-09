package bugPackage

import (
	"gfx"
	"time"
	"../../Klassen/textboxen"
	"fmt"
	//"math/rand"
	//"os"
	)

	
func Startbildschirm() {
	punkteTB = textboxen.New(200,10,1000,20)
	punkteTB.SetzeFarbe(255,255,255)
	punkteTB.SetzeFont("Schriftarten/ltypeb.ttf")
	
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	gfx.UpdateAus()
	gfx.LadeBild(5,5,"Bilder/BugAttack/Amoebius_klein.bmp")
	gfx.LadeBildMitColorKey(1050,530,"Bilder/FebWebK_red_gespiegelt.bmp",255,0,0)
	HelloTB := textboxen.New(130,50,800,500)
	HelloTB.SchreibeText("Willkomen beim Softwarepraktikum!\n\n" +
						"Du hast ja schon begonnen! Bevor wir euch eine Einweisung gegeben haben?\n\nVerständlich!\n\n"+
						"Endlich darfst du mal so richtig programmieren, genau das was du ja eigentlich wolltest, "+
						"ohne diesen ganzen Theoriekram.\n\nDoch scheinbar bist du doch nicht so cool wie du dachtest..."+
						"Oder hast du wieder während der Vorlesung programmiert??\n"+
						"Jedenfalls wimmelt es hier nur so von BUGs in deinem Code, deine TeamollegInnen werden nicht erfreut sein."+
						"\nWer will den schon so eine verBUGgte Klasse benutzen?"+
						"\n\nDie StudentInnen werden ja hier alle nicht besser, wir vom Dozententeam haben euch extra ein "+
						"halbautomatisches BUGFIXING-TOOL (hBT) programmiert. Damit solltes selbst du in der Lage sein die Bugs zu beseitigen. "+
						"Sogar auf einem Apfel.\n\n"+
						"Gehe einfach mit deinem Cursor an genau die Stelle des Bugs, drücke LEERTASTE....und BAAAAMM!" )
	HelloTB.SetzeFont("Schriftarten/ltypeb.ttf")
	HelloTB.SetzeFarbe(0,255,0)
	HelloTB.SetzeSchriftgröße(25)
	HelloTB.Zeichne()
	
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}


func LevelIntro() {
	gfx.SpieleSound("Sounds/Music/bugWars.wav")
	wg.Add(1)
	beschreibeArrayIntro()
	createNBugs(20,5,5)
	go cleanBugArray()		// Läuft dann für alle weiteren Level
	//go ShowBugs()
	go ZeichneWeltIntro()
	//time.Sleep(2e9)
	
	gfx.TastaturLesen1()
	bugArraySchloss.Lock()
	// tötet alle Bugs
	for i:=0;i<len(bugArray);i++ {
		bugArray[i] = nil
	}
	bugArraySchloss.Unlock()
	wg.Wait()
	gfx.StoppeAlleSounds()
	//time.Sleep(1e9) // Wichtig damit die ZeichneWeltIntro() die Chance hat zu beenden
}

// Erhöht das Level, sicher
func erhöheLevel() {
	levelSchloss.Lock()
	level++
	fmt.Println("level: ", level)
	levelSchloss.Unlock()
}

// Startet ein neues Level mit den gegebenen Parametern
func levelStart(){
	if SpielBeendet == true {
		punkteArray[level-1] = 0	// setze Punkte im Level auf Null weil das Spiel ja beendet wird
		return
	}
	fmt.Println("Starte Level")
	wg.Add(1)
	beschreibeArray()
	createNBugs(anzahlBugsImLevel,lvlSpeed,lvlNervosität)
	go ZeichneWelt()
	go zählePunkte()
	lvlZeit = 0			// für die Berechnung der Punktzahl
	lvlLäuft = true
	go lvlTimer()		
	for _,l:=range alleLadebalken {	// starte den Cooldown aller Ladebalken
		if l!= nil {go l.cooldown()}
	}
	// Warte bis keine Bugs mehr da sind
	for howManyBugs()>0 {
		time.Sleep(1e9)
	}
	lvlLäuft = false			// Signalisiert go-Routingen, dass das Level vorbei ist
	wg.Wait()				// nötig?
	entferneAlleLadebalken()
	ergebnisLevel()
	
}

// Zeigt Punkte und Note nach Level an
func ergebnisLevel() {
	fmt.Println("Level: ",level)
	fmt.Println("Punkte: ", punkteArray[level-1])
	fmt.Println("Note: ", berechneNote())
}

// Tutorial
func Level1(){
	LevelTutorial()
	erhöheLevel()
	anzahlBugsImLevel = 1
	lvlSpeed = 0
	lvlNervosität = 1
	gfx.SpieleSound("Sounds/Music/bugWars.wav")
	//var l ladebalken
	
	
	levelStart()
	
}

func Level2(){
	LevelTutorial()
	erhöheLevel()
	anzahlBugsImLevel = 3
	lvlSpeed = 2
	lvlNervosität = 2
	lautoaim := NewLadebalken(&autoAimCD,600,50,255,0,255,"x",5)
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := NewLadebalken(&killAllBugsCD,400,50,0,255,255,"k",20)
	alleLadebalken = append(alleLadebalken,l)

	//gfx.SpieleSound("Sounds/Music/bugWars.wav")
	levelStart()
}

func Level3() {
	LevelTutorial()
	erhöheLevel()
	lvlNervosität = 5
	anzahlBugsImLevel = 5
	lvlSpeed = 2
	lautoaim := NewLadebalken(&autoAimCD,600,50,255,0,255,"x",5)
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := NewLadebalken(&killAllBugsCD,400,50,0,255,255,"k",20)
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
	// Letztes Level vorbei
	//erhöheLevel()
}


func Level4() {
	LevelTutorial()
	erhöheLevel()
	lvlNervosität = 5
	anzahlBugsImLevel = 10
	lvlSpeed = 2
	lautoaim := NewLadebalken(&autoAimCD,600,50,255,0,255,"x",5)
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := NewLadebalken(&killAllBugsCD,400,50,0,255,255,"k",20)
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
	// Letztes Level vorbei
	//erhöheLevel()
}

func Level5() {
	LevelTutorial()
	erhöheLevel()
	lvlNervosität = 5
	anzahlBugsImLevel = 10
	lvlSpeed = 2
	lautoaim := NewLadebalken(&autoAimCD,600,50,255,0,255,"x",5)
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := NewLadebalken(&killAllBugsCD,400,50,0,255,255,"k",20)
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
	// Letztes Level vorbei
	//erhöheLevel()
}


func LevelTutorial() {
	
	gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	
	//gfx.LadeBild(10,20,"Bilder/Amoebius_klein.bmp")
	Level1TB:= textboxen.New(300,150,500,200)
	Level1TB.SchreibeText("Level 1")
	Level1TB.SetzeFont("Schriftarten/ltypeb.ttf")
	Level1TB.SetzeSchriftgröße(40)
	Level1TB.SetzeFarbe(255,0,0)
	Level1TB.Zeichne()
	
	Level1StartTB:= textboxen.New(300,250,500,200)
	Level1StartTB.SchreibeText("Die einfachen Bugs zuerst.\n\n\nPass auf, dass du nicht verschlimmbesserst!\n\n" + 
	"Bewege dich mit den Pfeiltasten und nutze das hBT mit LEERTASTE. Das hBT muss aber zentriert werden!! Also nicht einfach "+
	"irgendwo Bugfixen. Sonst machst du alles nur noch schlimmer!\n\n"+
	"Wenn dir das ganze über den Kopf wächst, drücke einfach 'q'." )
	Level1StartTB.SetzeFont("Schriftarten/ltypeb.ttf")
	Level1StartTB.SetzeSchriftgröße(26)
	Level1StartTB.SetzeFarbe(0,255,0)
	Level1StartTB.Zeichne()
	gfx.UpdateAn()
	gfx.Archivieren()		// nötig?
	go amoebiusAndBugAnimation()
	//go bugLevelAnimation()
	gfx.TastaturLesen1()
	quit <- true
}


/*
 * Nicht mehr benötigt?
func Endbildschirm() {
	gfx.Cls()
	BugAttackTB := textboxen.New(200,100,700,500)
	BugAttackTB.SetzeZentriert()
	BugAttackTB.SchreibeText(
		"BUG ATTACK\nNote: "+fmt.Sprint(BerechneNote()))
	BugAttackTB.SetzeFont("Schriftarten/ltypeb.ttf")
	BugAttackTB.SetzeFarbe(0,255,0)
	BugAttackTB.SetzeSchriftgröße(100)
	BugAttackTB.Zeichne()
	fmt.Println(BerechneNote())
	gfx.TastaturLesen1()
}
*/

// Ergebnisbildschirm / Level
func Endbildschirm() {
	var path string
	path = ""
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	
	
	gfx.LadeBild(150,100,path + "Bilder/sprechblase_flipped_400.bmp")
	gfx.LadeBildMitColorKey(100,350,path + "Bilder/BugAttack/Amoebius_klein.bmp",0,0,0)
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
	fmt.Println("Final Level: ",level)
	EndN, EndP = berechneEndNoteUndGesamtPunktzahl()
	gfx.SchreibeFont(325,195,fmt.Sprintf("%2.1f",EndN))
	
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	//fmt.Println("level: ",level)
	for i:=uint16(1); i<=maxLevel; i++ {
		//fmt.Println(i)
		level = i
		gfx.SchreibeFont(710,150+uint16((i-1)*68), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(punkteArray[i-1]) + " Punkte")
		gfx.SchreibeFont(710,175+uint16((i-1)*68),"           Note " + fmt.Sprintf("%2.1f",berechneNote()))
	}
	gfx.SchreibeFont(700,130+uint16(6*70),"----------------------")
	
	gfx.SchreibeFont(710,160+uint16(6*70),"Gesamt:    " + fmt.Sprint(EndP) + " Punkte")

	gfx.TastaturLesen1()
	//return gesamtnote, gesamtpunkte
}


/*
 * Funktion um einen stillsitzenden Bug zu generieren und anzeigen zu lassen, für Fotos;)
func BugFoto() {
	gfx.Stiftfarbe(0,0,0)
	beschreibeArraySchwarz()
	//beschreibeArray()
	b1 := NewBug(50,30)
	b1.speed=0
	go b1.bugAnimation()
	go b1.startMoving()
	//beschreibeArray()
	bugArray[0]=b1
	go ZeichneWelt(punkteTB)
	for howManyBugs()>0 {
		time.Sleep(1e9)
	}
	gfx.TastaturLesen1()
}*/
