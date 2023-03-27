package objekte
// Autor: B. Schneider
// Datum: 20.03.2023

// Dieses Paket stellt einen ADT Objekt gemäß der unten angegebenen 
// Spezifikation zur Verfügung. Ein Objekt hat eine Koordinate x,y ,
// eine Quadrat-Größe qua in Pixel, einen Typ typ

// Alle Methoden des ADT können nebenläufig aufgerufen werden !!
// -------------------------------------------------------------


// Vor.: --
// Eff.: Der neuen Kugel ist eine Identifikationsnummer zugeordnet worden.
// Erg.: Ein neues Obekt mit den Koordinaten x,y und der Quadrat-Größe qua
//       mit dem Typ typ 	ist geliefert.
// New (x,y, qua uint16, typ uint8) *data // *data erfüllt das Interface Objekt

	
type Objekt interface {
	SetzeKoordinaten(x,y uint16)
	
	GibKoordinaten() (uint16,uint16)
	
	SetzeTyp(t uint8)
	
	GibTyp() (uint8)
	
	Zeichnen()
	
	Getroffen(x,y uint16) bool
}
