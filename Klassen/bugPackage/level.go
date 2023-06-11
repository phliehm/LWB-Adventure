/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 *
 * BugLevel:
 * 
 * Beinhaltet Funktionen zum Ausführen der verschiedenen Level 
 * 
 */

package bugPackage

import (
	"gfx"
	"time"
	"../../Klassen/textboxen"
	"fmt"
	//"math/rand"
	//"os"
	)

// Texte für die Zwischensequenzen
var level1Text string = "... Die einfachen Bugs zuerst ... \n\n\nPass auf, dass du nicht verschlimmbesserst!\n\n" + 
	"Bewege dich mit den Pfeiltasten und nutze das hBT mit LEERTASTE. \n\n"+
	"Wenn dir das ganze über den Kopf wächst, drücke einfach 'q'."
var level2Text string = "... Ich frag mich was passiert wenn man nicht in der Mitte bugfixt? ...\n\n" +
						"'Haben Sie es ausprobiert?' \n\n"+
						"Den Spruch kennst du doch, oder? Also mach mal, was kann schon schief gehen!\n\n"+
						"Das mit der SHIFT-Taste hast du mitbekommen? Steht auch oben...aber ihr Studenten lest ja nie."
						 
var level3Text string = "... Gleich hab ich den Bug gefunden ...\n\n" +
						"Diese verflixten Parallelitäts-Bugs! Man denkt man hat sie ...und dann ...\n"+
						"dann tauchen diese kleinen Käfer einfach irgendwo anders wieder auf.\n\n" +
						"Ach NSP hast du noch gar nicht gehört? Na kein Wunder!!! Ok, ich hab noch eine kleine Hilfe für dich."+
						"Aber erzähls dem FabWeb nicht!\n\n" +
						"Ich hab das hBT verbessert!!! Das findet jetzt sogar automatisch einen Bug...irgendeinen.\n\n"+
						"Nutze die Zielautomatik und BAAAAAMMM!"
var level4Text string = "... Einen gefixt, zwei neue gemacht ...\n\n"+
						"In einer Woche ist Abgabe! Die Zeit arbeitet gegen dich, als beeil dich lieber.\n\n" +
						"Die Zielautomatik ist jetzt nicht mehr ganz so oft verfügbar, deine KomilitonInnen brauchen die auch.\n\n"+
						"Eingentlich braucht ihr ein Upgrade mit dem alle Bugs einfach verschwinden ..."
						
var level5Text string = "... Was war denn das??? ...\n\n"+
						"Hast du das programmiert??? Dieses blaue Ding? Nee, das kann nicht von dir sein.\n\n"+
						"Sieht nach FabWeb aus, als würden die rekursiv einfach nacheinander geBUGfixt werden.\n\n" + 
						"Jetzt werd aber deswegen nicht schlampig! Du musst trotzdem noch ein paar per Hand fixen.\n"+
						"Scheinbar hat Mr. FabWeb den Akku begrenzt." 
var level6Text string = "... Schule + LWB ... kein Problem ...\n\n"+
						"Erinnert dich das hier auch irgendwie an Schule? Ein Problem gelöst... \n\n"+
						"In einer Ecke sorgt man für Ruhe, ZACK, fängt es woanders wieder an.\n" +
						"Sind einfach zu viele in so einer Klasse. ... Klasse...HAHA...verstehste?" 

// Startbildschirm wenn man aus dem MainGame kommt
func BugAttackIntro() {
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
	//time.Sleep(1e9) // Wichtig damit die ZeichneWeltIntro() die Chance hat zu beenden
}
	
// Einleitungstext
func Startbildschirm() {
	punkteTB = textboxen.New(200,10,1000,20)
	punkteTB.SetzeFarbe(255,255,255)
	punkteTB.SetzeFont("Schriftarten/ltypeb.ttf")
	
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	gfx.UpdateAus()
	gfx.LadeBild(5,5,"Bilder/BugAttack/Amoebius_klein.bmp")
	gfx.LadeBildMitColorKey(1050,530,"Bilder/BugAttack/FebWebK_red_gespiegelt.bmp",255,0,0)
	HelloTB := textboxen.New(130,50,800,500)
	HelloTB.SchreibeText("Willkomen beim Softwarepraktikum!\n\n" +
						"Du hast ja schon begonnen! Bevor wir euch eine Einweisung gegeben haben?\n\nVerständlich!\n\n"+
						"Endlich darfst du mal so richtig programmieren, genau das was du ja eigentlich wolltest, "+
						"ohne diesen ganzen Theoriekram.\n\nDoch scheinbar bist du doch nicht so cool wie du dachtest..."+
						"Oder hast du wieder während der Vorlesung programmiert??\n"+
						"Jedenfalls wimmelt es hier nur so von BUGs in deinem Code, deine TeamollegInnen werden nicht erfreut sein."+
						"\nWer will denn schon so eine verBUGgte Klasse benutzen?"+
						"\n\nDie StudentInnen werden ja hier alle nicht besser, wir vom Dozententeam haben euch extra ein "+
						"halbautomatisches BUGFIXING-TOOL (hBT) programmiert. Damit solltest selbst du in der Lage sein die Bugs zu beseitigen. "+
						"Sogar auf einem Apfel.\n\n"+
						"Gehe einfach mit deinem Cursor auf einen Bug, drücke LEERTASTE....und BAAAAMM!" )
	HelloTB.SetzeFont("Schriftarten/ltypeb.ttf")
	HelloTB.SetzeFarbe(0,255,0)
	HelloTB.SetzeSchriftgröße(25)
	HelloTB.Zeichne()
	
	gfx.UpdateAn()
	gfx.TastaturLesen1()
}


// Erstes Level
func Level1(){
	LevelTutorial(level1Text)
	erhöheLevel()
	anzahlBugsImLevel = 1
	lvlSpeed = 0
	lvlNervosität = 1
	//var l ladebalken
	levelStart()
	
}

// 
func LevelTutorial(text string) {
	
	gfx.UpdateAus()
	gfx.Stiftfarbe(0,0,0)
	gfx.Cls()
	
	//gfx.LadeBild(10,20,"Bilder/Amoebius_klein.bmp")
	Level1TB:= textboxen.New(200,150,500,200)
	Level1TB.SchreibeText("Level " +fmt.Sprint(level+1))
	Level1TB.SetzeFont("Schriftarten/ltypeb.ttf")
	Level1TB.SetzeSchriftgröße(40)
	Level1TB.SetzeFarbe(255,0,0)
	Level1TB.Zeichne()
	
	Level1StartTB:= textboxen.New(200,250,700,500)
	Level1StartTB.SchreibeText(text)
	Level1StartTB.SetzeFont("Schriftarten/ltypeb.ttf")
	Level1StartTB.SetzeSchriftgröße(26)
	Level1StartTB.SetzeFarbe(0,255,0)
	Level1StartTB.Zeichne()
	gfx.UpdateAn()
	gfx.Archivieren()		// nötig?
	go amoebiusAndBugAnimation()
	//go bugLevelAnimation()
	gfx.TastaturLesen1()
	quit <- true				// Beende Amoebius und Bug Animation
}

// Level 2
func Level2(){
	LevelTutorial(level2Text)
	erhöheLevel()
	anzahlBugsImLevel = 3
	lvlSpeed = 2
	lvlNervosität = 5
	levelStart()
}

// Level 3
func Level3() {
	LevelTutorial(level3Text)
	erhöheLevel()
	anzahlBugsImLevel = 5
	lvlSpeed = 10
	lvlNervosität = 50
	lautoaim := newLadebalken(&autoAimCD,xposAutoAimBalken,yposAutoAimBalken,255,0,255,"x",1,"Sounds/Retro Sounds/General Sounds/Coins/sfx_coin_double1.wav")
	alleLadebalken = append(alleLadebalken,lautoaim)
	levelStart()

}

// Level 4
func Level4() {
	LevelTutorial(level4Text)
	erhöheLevel()
	lvlNervosität = 5
	anzahlBugsImLevel = 10
	lvlSpeed = 5
	lautoaim := newLadebalken(&autoAimCD,xposAutoAimBalken,yposAutoAimBalken,255,0,255,"x",5,"Sounds/Retro Sounds/General Sounds/Coins/sfx_coin_double1.wav")
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := newLadebalken(&killNBugsCD,xposkillNBugs,yposkillNBugs,0,255,255,"k",20,"Sounds/Retro Sounds/General Sounds/Fanfares/sfx_sounds_fanfare1.wav")
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
}

// Level 5
func Level5() {
	LevelTutorial(level5Text)
	erhöheLevel()
	lvlNervosität = 5
	anzahlBugsImLevel = 15
	lvlSpeed = 5
	lautoaim := newLadebalken(&autoAimCD,xposAutoAimBalken,yposAutoAimBalken,255,0,255,"x",4,"Sounds/Retro Sounds/General Sounds/Coins/sfx_coin_double1.wav")
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := newLadebalken(&killNBugsCD,xposkillNBugs,yposkillNBugs,0,255,255,"k",10,"Sounds/Retro Sounds/General Sounds/Fanfares/sfx_sounds_fanfare1.wav")
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
	// Letztes Level vorbei
	//erhöheLevel()
}

// Level 6
func Level6() {
	LevelTutorial(level6Text)
	erhöheLevel()
	lvlNervosität = 20
	anzahlBugsImLevel = 20
	lvlSpeed = 5
	lautoaim := newLadebalken(&autoAimCD,xposAutoAimBalken,yposAutoAimBalken,255,0,255,"x",2,"Sounds/Retro Sounds/General Sounds/Coins/sfx_coin_double1.wav")
	alleLadebalken = append(alleLadebalken,lautoaim)
	l := newLadebalken(&killNBugsCD,xposkillNBugs,yposkillNBugs,0,255,255,"k",10,"Sounds/Retro Sounds/General Sounds/Fanfares/sfx_sounds_fanfare1.wav")
	alleLadebalken = append(alleLadebalken,l)
	levelStart()
	// Letztes Level vorbei
	//erhöheLevel()
}




// Ergebnisbildschirm / Level
func Endbildschirm() {
	var path string
	path = ""
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	
	// Zertifikat Hintergrund
	gfx.LadeBild(150,100,path + "Bilder/Zertifikat/sprechblase_flipped_400.bmp")
	gfx.LadeBild(230,390,path+"Bilder/BugAttack/FabWeb_fullBody_gespiegelt.bmp")
	gfx.LadeBildMitColorKey(250,350,path + "Bilder/BugAttack/Amoebius_klein.bmp",0,0,0)
	gfx.LadeBild(620,80,path + "Bilder/Zertifikat/paper_500.bmp")
	gfx.LadeBild(960,520,path + "Bilder/Zertifikat/certified_100.bmp")
	//gfx.LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
	
	// Exit Anweisung
	gfx.LadeBildMitColorKey(1080,30,path + "Bilder/BugAttack/Bug.bmp",0,0,0)
	gfx.Stiftfarbe(125,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	gfx.SchreibeFont(1070,140,"2 mal [q]")
	gfx.SchreibeFont(1070,170,"für Exit")
		
	
	// Überschrift
	gfx.Stiftfarbe(0,255,0)
	gfx.SetzeFont(path + "Schriftarten/ComputerTypewriter.ttf",80)
	gfx.SchreibeFont(330,10,"Bug  ATTACK")
	
	// Sprechblase
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
	
	// Inhalt des Zertifikats
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	for i:=uint16(1); i<=maxLevel; i++ {
		level = i
		gfx.SchreibeFont(710,150+uint16((i-1)*68), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(punkteArray[i-1]) + " Punkte")
		gfx.SchreibeFont(710,175+uint16((i-1)*68),"           Note " + fmt.Sprintf("%2.1f",berechneNote()))
	}
	gfx.SchreibeFont(700,130+uint16(6*70),"----------------------")
	
	gfx.SchreibeFont(710,160+uint16(6*70),"Gesamt:    " + fmt.Sprint(EndP) + " Punkte")

	gfx.TastaturLesen1()
}


// Erhöht das Level, sicher
func erhöheLevel() {
	levelSchloss.Lock()
	level++
	//fmt.Println("level: ", level)
	levelSchloss.Unlock()
}

// Startet ein neues Level mit den gegebenen Parametern
func levelStart(){
	if SpielBeendet == true {
		punkteArray[level-1] = 0	// setze Punkte im Level auf Null weil das Spiel ja beendet wird
		return
	}
	//fmt.Println("Starte Level")
	wg.Add(3)				// warte später auf 3 Go-Routinen
	beschreibeArray()		// Füllt das Level mit Zahlen
	createNBugs(anzahlBugsImLevel,lvlSpeed,lvlNervosität)	// Füge Bugs hinzu
	go ZeichneWelt()		// Starte mit dem Zeichnen der Welt
	lvlZeit = 0				// Die Zeit im Level ist 0
	lvlLäuft = true			// Das Level beginnt
	go zählePunkte()	// für die Berechnung der Punktzahl
	go lvlTimer()			// Zählt Zeit nach unten
	for _,l:=range alleLadebalken {	// starte den Cooldown aller Ladebalken
		if l!= nil {go l.cooldown()}
	}
	// Warte bis keine Bugs mehr da sind
	for howManyBugs()>0 {
		time.Sleep(1e9)
	}
	lvlLäuft = false			// Signalisiert go-Routingen, dass das Level vorbei ist
	wg.Wait()				// Wartet auf alle zu beendenden Go-Routinen
	time.Sleep(1e9)			// Wartet damit die Ladebalken schließen können.
	entferneAlleLadebalken()	// Löscht alle Ladebalken eines Levels
	//ergebnisLevel()				// Gibt das Ergebnis des Levels in der Konsole aus
	
}

// Zeigt Punkte und Note nach Level an
func ergebnisLevel() {
	fmt.Println("Level: ",level)
	fmt.Println("Punkte: ", punkteArray[level-1])
	fmt.Println("Note: ", berechneNote())
}

