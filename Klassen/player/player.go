package player

// abstrakter Datentyp für die LWB-Adventure-Player

// Vor.: keine
// Erg.: Ein neuer Player ist initialisiert, der sich im Mainfloor (Raum 0) befindet,
//       keine Schlüssel besitzt, den Namen "unknown LWB-Player" und einen Standard-Avatar
//       hat sowie nullinitialisierte Tabellen für Punktestand und Noten
// New (n uint) *data // *data erfüllt das Interface Player
type Player interface {
	//TODO: Spezifikationen
	
	GetName() string
	SetName(string)
	//ChangeName(string) ???
	
	GetBild() string
	SetBild(string)
	
	GetRaum() uint
	SetRaum(uint)
	
	//GetAllPunkte() [4][3]uint32
	GetPunkte(sem,game int) uint32
	SetPunkte(sem,game int, punkte uint32)
	
	//GetAllNoten() [4][3]float32
	GetNote(sem,game int) float32
	SetNote(sem, game int, note float32)
		
	GetKeys() uint8
	//SetKeys(uint8)
	IncrKeys()
	
}

