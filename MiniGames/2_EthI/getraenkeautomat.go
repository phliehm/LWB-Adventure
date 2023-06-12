
// Paket zum Getränkeautomatenspiel

// Spezifikation

// 4.6.2023		Martin Seiß


// Folgende Getränkezuteilung wird verwendet:
// ******************************************
//  Dark Schmidther	-	Extraschwarzer Kaffee
//	The Herk		-	Melissentee
//  J.EthI			-	Kaffee mit Milch und 2x Zucker
//  FabFour			-	Cappuccino
//	Amoebi			-	Grüner Tee
//  WtheK			-	Bier


// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 1
// verwendet:
// ******************************************
//  A - Grüner Tee			-	[10,20,50]
//	B - Pfefferminztee		-	[20,50,10]
//  C - Melissentee			-	[50,10,20]
//  D - Schwarzer Tee 		-	[20,10,50]
//	E - Kräutertee			-	[10,50,20]
//  F - Früchtetee			-	[50,20,10]


// Folgende Getränkezuteilung mit Name und Nummer wird für Automat 2
// verwendet:
// ******************************************
//  B - Bier					-	[50,50,10,20]
//	R - Radler					-	[50,50,10,10]
//  E - Extrascharzer Kaffee		[50,20,20,10]
//  L - Latte Macchiato 		-	[50,10,20,20]
//	C - Cappuccino				-	[50,20,10,20]
//  K - Kaffee					-	[50,(10,10,20)]
//		mit Zucker				-   [10,50,(10,20)]
// 		mit 2xZucker			-	[10,10,50,20]	
//		mit Milch				-   [20,50,10,10)]
// 		mit Milch und Zucker	-	[20,10,50,10]
// 		mit Milch und 2xZucker	-	[20,10,50,10]
// in ()-Klammern - Reihenfolge egal


// Folgende Getränkezuteilung mit Name und Nummer wird verwendet:
// ******************************************
//  1 - Extradunkle Schokolade			 	-	[]
//	2 - Melissentee							-	[]
//  3 - Kaffee mit Milch und 2x Zucker		-   []
//  4 - Kaffe mit 1x Zucker					-	[]
//	5 - Grüner Tee							-	[]
//  6 - doppelter Expresso mit 2x Zucker	-   []

// ()-Klammern - Reihenfolge egal


package getraenkeautomat


// ------------  importierte Pakete ------------------------//

//import "fmt"
//import "gfx"
//import "../../Klassen/textboxen"
//import "../../Klassen/buttons"


// ------------ exportierte Variablen -------------------- //

// var noten float32
// var punkte uint32


// -----------------   Funktion -------------------------- //

// Vor: Ein gfx-Fenster (1200x700) ist geöffnet.
// Erg: Der Spielstand mir Note und Punkte ist geliefert.
// Eff: Das Getraenkeautomatenspiel ist gestartet.
// func Getraenkeautomat() (float32, uint32)
	
