package bugPackage

//import "fmt"

func berechneNote() float32{
	n:= berechneProzent()
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

// a ist x Prozent von b. Ausgabe in 0...1
func berechneProzent() float32{
	//fmt.Println("PUNKTE: ",punkteArray[level-1])
	if punkteArray[level-1]<lvlMinPunkte[level-1] {return 0}
	return float32(punkteArray[level-1]-lvlMinPunkte[level-1])/float32((lvlMaxPunkte[level-1]-lvlMinPunkte[level-1]))
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

// Erg.: Endnote für das MainGame
func berechneEndNoteUndGesamtPunktzahl() (float32,uint32) {
	var summePunkte uint32
	var summeNoten float32
	for i,punkte:= range punkteArray {	// für alle Level
		level = uint16(i+1)		// ändere "level" weil "berechneNote()" die globale Variable "level" verwendet.
		summeNoten+= berechneNote()
		summePunkte+=uint32(punkte)
	}
	return rundeNote(summeNoten/float32(len(punkteArray))),summePunkte
}


func GibErgebnis() (float32,uint32) {
	//fmt.Println("Ergebnis: ",EndN,EndP)
	return EndN,EndP
}
