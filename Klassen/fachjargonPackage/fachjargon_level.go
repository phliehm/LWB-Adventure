package fachjargonPackage

import (
		"fmt"
		"gfx"
		"math/rand"
		"strconv"
		"time"
		)


// Zeichnet die Antworten des jeweiligen Spiels
func ZeichneAntworten(text []string) {
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 20)
	var zeilenAbstand uint16 = 35 // Zeilenabstand
	for i:=uint16(0);i<uint16(len(text));i++ {gfx.SchreibeFont(50,20 + i*zeilenAbstand,text[i])} 
}

// Startbildschirm
func StartBildschirm() {
	
	slicesNullen()
	initialisiereGlobal()
	level = 0
	
	gfx.Stiftfarbe(255, 255, 255)
	gfx.Vollrechteck(0, 0, breite, höhe)
	gfx.LadeBildMitColorKey(100, 190, bilderPfad+"WtheK_smart.bmp", 255, 255, 255)
	
	gfx.Stiftfarbe(0, 100,0)
	gfx.SetzeFont("Schriftarten/ComputerTypewriter.ttf",80)
	gfx.SchreibeFont(100, 30, "FachJargon")
	gfx.LadeBildMitColorKey(500, 200, bilderPfad+"Tafel.bmp", 255, 255, 255)
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 25)
	gfx.SchreibeFont(700, 650, "Drücken Sie doch mal ne Taste!")
	//gfx.SchreibeFont(200, 500, "Hallo! Ich begrüsse Sie zur heutigen Fachdidaktik Veranstaltung!")
	gfx.TastaturLesen1()
}

// Introbildschirm 
func IntroBildschirm(){
	go monitorSpielBeendet()
	gfx.Stiftfarbe(255, 255, 255)
	gfx.Vollrechteck(0, 0, breite, höhe)
	gfx.Stiftfarbe(0, 0, 0)
	gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 25)
	gfx.SchreibeFont(150, 400, "Hallo! Ich begrüsse Sie zur heutigen Fachdidaktik Veranstaltung!")
	//time.Sleep(1e9)
	gfx.SchreibeFont(150, 450, "Kennen Sie diese 8 Schritte noch? Bestimmt!")
	gfx.SchreibeFont(150, 500, "Ich zeige Ihnen jetzt Bilder und Sie drücken die zugehörigen Tasten 1-8.")
	gfx.LadeBild(0, 0, bilderPfad+"3_Fachdidaktik_Planung.bmp")
	gfx.SchreibeFont(150, 600, "Drücken Sie 'q' zum Beenden oder 's' um das nächste Level zu starten")
	
	ZeichneAntworten(textEntwicklung)
	
	leseTastaturStartbildschirm()
	level++
	return 
} 

func level2Intro() {
	gfx.Stiftfarbe(255, 255, 255)
	gfx.Vollrechteck(0, 0, breite, höhe)
	gfx.Stiftfarbe(0, 0, 0)
	gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 25)
	gfx.SchreibeFont(150, 400, "War das nicht toll? Ich hoffe Sie können noch! ")
	gfx.SchreibeFont(150, 450, "Jetzt wird es erst richtig spannend!")
	gfx.SchreibeFont(150, 500, "Sie haben ja alle so wundervolle Lernwerkzeuge vorgestellt.")
	gfx.SchreibeFont(150, 560, "Haben Sie da auch aufgepasst? Sind schließlich Ihre KomilitonInnen!!")
	ZeichneAntworten(textTools)
	gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 20)
	gfx.SchreibeFont(150, 620, "Drücken Sie 'q' zum Beenden oder 's' um das nächste Level zu starten")
	
	leseTastaturStartbildschirm()
	return 	
}


func levelAbspielen(bilderArray,text []string) {
	time.Sleep(1e9)
	//fmt.Println("Level: ",level)
	var zähler uint16 = uint16(len(bilderArray))			// Anzahl der Fragen
	//gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 20)
	var zeit int
	for SpielBeendet != true && zähler > 0{
		//fmt.Println("Zähler: ",zähler)
		t_start = time.Now().UnixNano()
		
		gfx.UpdateAus()
		gfx.Stiftfarbe(255,255, 255)
		gfx.Vollrechteck(0, 0, breite, höhe)
		
		gfx.Stiftfarbe(0, 100, 0)
		gfx.SetzeFont("Schriftarten/Ubuntu-B.ttf", 80)
		gfx.SchreibeFont(450, 30, "LEVEL " + fmt.Sprint(level))
		
		gfx.Stiftfarbe(0, 0, 0)
		
		Taste = LoseTaste(len(bilderArray))
		
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
		
		ZeichneAntworten(text)
		gfx.LadeBild(50,300,bilderPfad+bilderArray[Taste-'0'-1])		// Konvertiert die Rune in eine Zahl zwisch 0-7
		gfx.UpdateAn()

		for gedrueckt != 1 && SpielBeendet==false {
			gfx.UpdateAus()
			zeit =int(time.Now().UnixNano()-t_start)/1e6
			gfx.Stiftfarbe(255, 255, 255)
			gfx.Vollrechteck(550,150,170,50)
			gfx.Stiftfarbe(255,0,0)
			gfx.SchreibeFont(550,150,"Zeit: "+ strconv.Itoa(zeit)+" ms")
			gfx.UpdateAn()
			time.Sleep(1e8)
		}
		
		richtig:=TesteTaste(taste)
		
		zeitMax := uint16(5000)
		zeitMin := uint16(1000)
		//zeit :=int(time.Now().UnixNano()-t_start)/1e6
		
		zeichneRichtigFalsch(richtig,uint16(zeit),zeitMax)			// Gibt Reaktion von WtheK
		
		//fmt.Println("Zeit nach Berechnung: ",zeit)
		if  richtig == true {
			punkteArray[level-1] += konvertiereZeitInPunkte(zeitMax,zeitMin,uint16(zeit))
		}
		//fmt.Println("Punkte: ",punkteArray[level-1])
		gfx.SchreibeFont(650, 550, "Das hat "+strconv.Itoa(zeit)+" ms gedauert!")
		time.Sleep(1e9)
		
		zähler--	// gehe zur nächsten Frage
	}
	notenArray[level-1]=berechneNoteAusProzent(durchschnittsProzent(punkteArray[level-1],uint16(len(bilderArray))))
	//fmt.Println("Prozent: ",durchschnittsProzent(punkteArray[level-1],uint16(len(bilderArray))))
	//fmt.Println("Note: ",notenArray[level-1])
	
}


// Unterrichtsreihen

func Level1() {
	levelAbspielen(bilderArrayEntwicklung,textEntwicklung)
	level++
}


// Tools
func Level2() {
	level2Intro()
	levelAbspielen(bilderArrayTools,textTools)
	level++
	
}

/*
// Unterrichtsentwicklung
func Level3() {
	fmt.Println("Level 3")
}
*/

// Anzeige für den Endbildschirm
func Endbildschirm() {
	var path string
	path = ""
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	
	gfx.LadeBild(150,100,path + "Bilder/Zertifikat/sprechblase_flipped_400.bmp")
	gfx.LadeBild(230,390,path+"Bilder/FachJargon/winni-the-k.bmp")
		
	gfx.LadeBild(620,80,path + "Bilder/Zertifikat/paper_500.bmp")
	gfx.LadeBild(960,520,path + "Bilder/Zertifikat/certified_100.bmp")
	//gfx.LadeBild(1080,30,path + "Bilder/Zurück-Symbol.bmp")
	
	gfx.LadeBildMitColorKey(1080,30,path + "Bilder/FachJargon/Zurück-Symbol.bmp",0,0,0)
	gfx.Stiftfarbe(255,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",20)
	gfx.SchreibeFont(1070,150,"[q] für Exit")

		
	gfx.Stiftfarbe(0,100,0)
	gfx.SetzeFont(path + "Schriftarten/ComputerTypewriter.ttf",80)
	gfx.SchreibeFont(330,10,"FachJargon")
	gfx.Stiftfarbe(0,0,0)
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",24)
	gfx.SchreibeFont(295,140,"Du hast die")
	gfx.SchreibeFont(310,260,"erreicht!")
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",32)
	gfx.SchreibeFont(285,170,"Gesamtnote")
	gfx.SetzeFont(path + "Schriftarten/Starjedi.ttf",42)
	//fmt.Println("Final Level: ",level)
	EndN, EndP = berechneEndNoteUndGesamtPunktzahl()
	gfx.SchreibeFont(325,195,fmt.Sprintf("%2.1f",EndN))
	
	gfx.SetzeFont(path + "Schriftarten/terminus-font/TerminusTTF-Bold-4.49.2.ttf",22)
	//fmt.Println("level: ",level)
	for i:=uint16(1); i<=maxLevel; i++ {
		//fmt.Println(i)
		level = i
		gfx.SchreibeFont(710,150+uint16((i-1)*68), "Level "+ fmt.Sprint(i) + ":   "+ fmt.Sprint(punkteArray[i-1]) + " Punkte")
		gfx.SchreibeFont(710,175+uint16((i-1)*68),"           Note " + fmt.Sprintf("%2.1f",notenArray[i-1]))
	}
	gfx.SchreibeFont(700,130+uint16(6*70),"----------------------")
	
	gfx.SchreibeFont(710,160+uint16(6*70),"Gesamt:    " + fmt.Sprint(EndP) + " Punkte")

	gfx.TastaturLesen1()

}

// Gibt zurück, ob die gedrückte Taste die richtige war
func TesteTaste(taste uint16) bool {
	switch taste {
		case Taste:
			return true
		default:
			return false
	}
	//time.Sleep(1e9)
}


// Lost aus welches Bild gezeigt wird und welche Taste entsprechend gedrückt werden müsste
func LoseTaste(n int) uint16 {
	z := rand.Intn(n)
	return TastenArray[z]
}



// zeichnet Reaktion auf richtige oder Falsche Antworten
func zeichneRichtigFalsch(richtig bool,zeit,zeitMax uint16) {
	if zeit>zeitMax {
		gfx.Stiftfarbe(255,0,0)
		gfx.SpieleSound("Sounds/Retro Sounds/General Sounds/Weird Sounds/sfx_sound_refereewhistle.wav")
		gfx.SchreibeFont(650, 500, "Schlafen Sie? Das war zu langsam!")
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black_sad.bmp", 255, 255, 255)
		gfx.LadeBildMitColorKey(150,450,bilderPfad+"zzzzz.bmp",255,255,255)
		return
	}
	if richtig == true { 
		gfx.Stiftfarbe(0,0,255)
		gfx.SpieleSound("Sounds/Retro Sounds/General Sounds/Pause Sounds/sfx_sounds_pause6_in.wav")
		gfx.SchreibeFont(650, 500, "Richtig! Das haben Sie aber toll geloest!!")
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black.bmp", 255, 255, 255)
		gfx.LadeBildMitColorKey(150,450,bilderPfad+"Richtig.bmp",255,255,255)
	}
	if richtig == false {
		gfx.Stiftfarbe(255,0,0)
		gfx.SpieleSound("Sounds/Retro Sounds/General Sounds/Negative Sounds/sfx_sounds_error3.wav")
		gfx.SchreibeFont(650, 500, "Ohjeee, das ist ja falsch, wie kommen Sie denn darauf???")
		gfx.LadeBildMitColorKey(750, 50, bilderPfad+"WtheK_black_sad.bmp", 255, 255, 255)
		gfx.LadeBildMitColorKey(150,450,bilderPfad+"Falsch.bmp",255,255,255)
	}
}

func konvertiereZeitInPunkte(maxZeit,minZeit,zeit uint16) uint16{		// Zeit in ms
	//fmt.Println("Zeit in konvertiereZeitInPunkte: ",zeit)
	if zeit>maxZeit {return 0}
	return uint16(float32((maxZeit-zeit))/float32((maxZeit-minZeit))*100)			// Prozentualer Anteil der verbleibenden Zeit an der Gesamtzeit = Punkte
}

func durchschnittsProzent(punkte,n uint16) float32{		// gesamtpunkte im Level, Anzahl der Fragen
	durchschnitt := float32(punkte)/float32(n)
	return durchschnitt
}

func berechneNoteAusProzent(p float32) float32{	// bekommt Prozent, gibt Note
	n := p/100.0
	//fmt.Println("Prozent: ",n*100)
	if n>=0.90 {return 1.0}
	if n>=0.85 {return 1.3}
	if n>=0.80 {return 1.7}
	if n>=0.75 {return 2.0}
	if n>=0.70 {return 2.3}
	if n>=0.65 {return 2.7}
	if n>=0.60 {return 3.0}
	if n>=0.55 {return 3.3}
	if n>=0.50 {return 3.7}
	if n>=0.45 {return 4.0}
	if n>=0.40 {return 4.3}
	if n<0.40  {return 5.0}
	
	return 5.0
} 

// Erg.: Endnote für das MainGame
func berechneEndNoteUndGesamtPunktzahl() (float32,uint32) {
	var summePunkte uint32
	var summeNoten float32
	for i,punkte:= range punkteArray {	// für alle Level
		level = uint16(i+1)		// ändere "level" weil "berechneNote()" die globale Variable "level" verwendet.
		summeNoten+= notenArray[i]
		summePunkte+=uint32(punkte)
	}
	return rundeNote(summeNoten/float32(len(punkteArray))),summePunkte
}


func GibErgebnis() (float32,uint32) {
	//fmt.Println("Ergebnis: ",EndN,EndP)
	return EndN,EndP
}

func rundeNote(n float32) float32{
	if n<1.15 {return 1.0}
	if n<1.5 {return 1.3}
	if n<1.85 {return 1.7}
	if n<2.15 {return 2.0}
	if n<2.5 {return 2.3}
	if n<2.85 {return 2.7}
	if n<3.15 {return 3.0}
	if n<3.5 {return 3.3}
	if n<3.85 {return 3.7}
	if n<4.15 {return 4.0}
	if n<4.5 {return 4.3}
	return 5.0
}

// Damit man am Anfang des Spiels einen evtl. gespeicherten Punktestand löschen kann
func slicesNullen() {
	for i:=0;i<len(punkteArray);i++ {
		punkteArray[i] = 0
	}
}
