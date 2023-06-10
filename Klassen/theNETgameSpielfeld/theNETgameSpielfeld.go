

//  ADT	- theNETgameSpielfeld für theNETgame

//	Spezifikation

//	Martin Seiß		29.5.2023  (Start)


package theNETgameSpielfeld


// ------------  importierte Pakete und Klassen   --------------------//
//import "gfx"
//import "fmt"
//import "time"
//import "../textboxen"
import "../netze"
//import "../buttons"


type Spielfeld interface {

	netze.Netz			// Vererbung von netze

	// Vor: -
	// Erg: Die aktuelle Punktzahl ist geliefert.
	GibPunktzahl() uint16

	// Vor: -
	// Eff: Die Distanz wird um den angegebenen Wert erhöht.
	ErhoeheDistanz(punkte uint16)
	
	// Vor: -
	// Erg: Gibt die aktuelle Note aus.
	GibNote() float32
	
	// Vor: -
	// Eff. Setzt die Knotenid, wo sich das Paket befindet.
	SetzePaketID(id uint32)

	// Vor: Ein gfx-Grafikfenster ist geöffnet.
	// Eff: Das Spielfeld wird in einer Schleife wiederholt gezeichnet,
	//		solange das Spiel läuft, also weder gewonnen noch verloren
	//		ist.
	Zeichnen()

	
	// Vor: Ein gfx-Grafikfenster ist geöffnet.
	// Eff: Die Hintergrundmusik wird in einer Schleife wiederholt
	//		gestartet.
	Hintergrundmusik()
	
	
	// Vor: -
	// Eff: Das Spiel und die Uhr wird gestartet und das Paket kann bewegt werden.
	StartGame()


	// Vor: -
	// Eff: True ist geliefert, wenn aktuelles Level gestartet wurde und
	// 		das Spiel noch gewonnen oder verloren ist.
	SpielLaeuft() bool

	
	// Vor: -
	// Eff: True - das nächste Level wird geladen, sonst das
	//		alte Level wird mit neuem Graphen geladen.
	NeuesLevel(neu bool)

}
