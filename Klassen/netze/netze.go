// Klasse Netz für theNETgame wird von Graph abgeleitet.

// Spezifikation

// Martin Seiß			12.5.2023	(Start)

//  !!!!!! Möglicher Bug im Dijkstra-Algorithmus. Erspielte Distanz manchmal
//		geringer als über den Algorthmus bestimmte, was nicht sein 
//		darf.

// Vor: Die Wahrscheinlichkeiten für die Router- und Kantensperren
//		leigen zwischen 0 und 1.
// Erg: Ein Netz von Routen mit Verbindungen ist erzeugt und geliefert.
// func New(pKnotensperre,pKantensperre float64) *data

package netze

// ------------  importierte Pakete und Klassen   --------------------//
//import "math/rand"
//import "time"
//import "fmt"
import "../graphen"


type Netz interface{
	
	graphen.Graph				// Ableitung von ADT graph von St. Schmidt
	
	// Vor: Der Knoten mit der id existiert.
	// Eff: Die Nachbarknoten zur angegebenen ID werden grün gesetzt.
	// Erg: Eine Liste von IDs der Nachbarknoten ist geliefert. 
	GibNachbarIDs(id uint32) []uint32
	
	// Vor: .
	// Eff: -
	// Erg: Die größte ID (gleich der ID des Ziels) ist geliefert. 
	GibZielID() uint32
	
	// Vor: -
	// Eff: -
	// Erg: Die minimale Distanz auf Grundlage der Kantenlängen ist
	//		geliefert. 	
	GibMinDist() uint32
	
	// Vor: -
	// Eff: Setzt die Wahrscheinlichkeiten, dass eine Kante oder eine 
	//		Knoten gesperrt rot ist.
	// Erg: - 	
	//GibNetz() uint32
	SetzeWkeitHindernisse(pKnotensperre,pKantensperre float64)
	
	// Vor: -
	// Eff: Setzt gemäß den Vorgaben die Hindernisse im Spiel in einer
	//		wiederholenden Schleife neu. Kanten und Knoten werden damit
	//		für eine bestimmte Zeit gesperrt. 
	// Erg: -
	Hindernisse()
	
	// Vor: verloren:   1 = Kante gesperrt, 2 Knoten gesperrt,
	// 					3 = Bugget zu Ende
	// Eff: Setzt, ob verloren wurde.
	SetzeVerloren(verloren uint16)
	
	// Vor: -
	// Eff: Setzt, ob gewonnen wurde.
	SetzeGewonnen(gewonnen bool)
		
	// Vor: -
	// Erg: True, wenn Spiel gewonnen.
	GibGewonnen() bool

	// Vor: -
	// Erg: 0, wenn Spiel nicht verloren. 1 = Kante war gesperrt,
	//		2 Knoten war gesperrt, 3 = Bugget war zu Ende
	GibVerloren() uint16
	
}
