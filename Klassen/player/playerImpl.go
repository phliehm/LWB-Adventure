package player

import "gfx"

type data struct {
	name string						//Name
	bild string						//Pfad zum Bild/Avatar
	raum uint						//Raum (0: Mainfloor, 1-4: Semester, 5: Nicht-Zeugnis)
	punkte [4][3]uint32				//Punktestand [Semester][Minigame]uint32
	noten [4][3]float32				//Notenfeld [Semester][Minigame]float32 
	keys uint8						//AnzahlSchlüssel uint8									
}

func New() *data {
	var pl *data = new(data)
	pl.name = "unknown LWB-Player"
	pl.bild = "../../Bilder/unknownLWB-Player"
	return pl
}

//TODO: Methoden schreiben

func (pl *data) GetName() string {
	return pl.name
}

func (pl *data) SetName(name string) {
	pl.name = name
}

//ChangeName(string) ???

func (pl *data) GetBild(x,y uint16) string {
	gfx.LadeBild(x,y,pl.bild)
	return pl.bild
}

func (pl *data) SetBild(bildpfad string) {
	pl.bild = bildpfad
}

func (pl *data) GetRaum() uint {
	return pl.raum
}

func (pl *data) SetRaum(raum uint) {
	pl.raum = raum
}

//GetAllPunkte() [4][3]uint32

func (pl *data) GetPunkte(sem,game int) uint32 {
	return pl.punkte[sem][game]
}

func (pl *data) SetPunkte(sem,game int, punkte uint32) {
	pl.punkte[sem][game] = punkte
}

//GetAllNoten() [4][3]float32

func (pl *data) GetNote(sem,game int) float32 {
	return pl.noten[sem][game]
}

func (pl *data) SetNote(sem, game int, note float32){
	pl.noten[sem][game] = note
}
	
func (pl *data) GetKeys() uint8 {
	return pl.keys
}

//SetKeys(uint8)

func (pl *data) IncrKeys() {
	pl.keys++
}
