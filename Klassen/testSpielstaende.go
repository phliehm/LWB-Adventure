
//  Test von ADT spielstand

//	zum Speichern vom Spielstand 

//  mit Punkte, Note und Datum (wann gespeichernt wurde)

//	Martin Seiß		5.6.2023


package main

import "fmt"
import "./spielstaende"


func main() {

	var path string = "../SAVE"	// Pfad in dem sich die Speicher-
									// standsdateien befinden 
	//var dateiname string
	var spielername string = "Martin"

	var sp1,sp2 spielstaende.Spielstand
	
	
	//  ----------  Beispiel 1: Lade Spielstand  -------------------//	
	sp1 = spielstaende.New(spielername,path)
	fmt.Println("Spielstand vorhanden? ",sp1.GibVorhanden())
	fmt.Println("Punkte",sp1.GibPunkte())
	fmt.Println("Noten",sp1.GibNoten())
	
	//  ---------- Speichere Spielstand ----------------------------//
	var punkte []uint32 = []uint32{14,6,1978,24062023}
	var noten []float32 = []float32{6.0,4.3,1.7,3.3}
	sp1.Speichern(noten,punkte)
	fmt.Println("Punkte",sp1.GibPunkte())
	fmt.Println("Noten",sp1.GibNoten())
	
	
	// -----------    Laden in Spielstand nochmal ------------------//	
	sp2 = spielstaende.New(spielername,path)		
	fmt.Println("Datei vorhanden? ",sp2.GibVorhanden())
	fmt.Println("Punkte",sp2.GibPunkte())
	fmt.Println("Noten",sp2.GibNoten())
	//  ---------- Speichere Spielstand ----------------------------//
	punkte = []uint32{200,600,12978,240023}
	noten  = []float32{5.0,4.7,2.7,2.3}
	sp2.Speichern(noten,punkte)
	fmt.Println("Punkte",sp2.GibPunkte())
	fmt.Println("Noten",sp2.GibNoten())

}
