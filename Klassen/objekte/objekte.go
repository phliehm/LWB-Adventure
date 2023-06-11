package objekte

// Autor: B. Schneider
// Datum: 20.03.2023

// Dieses Paket stellt einen ADT Objekt gemäß der unten angegebenen 
// Spezifikation zur Verfügung. Ein Objekt (für das Paket gfx) hat 
// eine Koordinate "x,y" , je nach Aussehen einen eigenen "typ",
// eine dargestellte Quadrat-Größe "qua" (in Pixeln), den Bool "aktiv",
// der entscheidet, ob mit einem Objekt interagiert werden kann, einen
// String für Text "inhalt" und die Nanosekunden seiner "erstellung". 

// !!!   Die Methoden des ADT sind nicht nebenläufig abgeschirmt   !!!
// Im Hauptprogramm müssen also entsprechende Vorkehrungen getroffen werden !!
// -------------------------------------------------------------

/* Vor.: --
 *  
 *  Erg.:Eine neue Textbox mit den Koordinaten x,y und der Breite und Höhe 
 * 		ist geliefert. Default Werte sind: 
 * 		Schriftfarbe - schwarz
 * 		ZeilenAbstand - 0
 * 		Schriftgröße - 20
 * 		Font - "../Schriftarten/Ubuntu-B.ttf"
 * 		Inhalt - Leer ("")
 *  New (x,y, b, h uint16) *data // *data erfüllt das Interface Textbox
*/

// Vor.: --
// Eff.: Das neue Objekt enthält die Erstellungszeit und ist aktiv.
// Erg.: Ein neues Obekt mit den Koordinaten x,y und der Quadrat-Größe qua
//       mit dem Typ typ ist geliefert. Sein "inhalt" ist leer.
// func New (x,y, qua uint16, typ uint8) *data	 // *data erfüllt das Interface Objekt

	
type Objekt interface {
	
	// Vor.: --
	// Eff.: Das Objekt hat nun die neuen Koordinaten.
	SetzeKoordinaten(x,y uint16)
	
	// Vor.: --
	// Erg.: Die Koordinaten den Objekts sind geliefert.
	GibKoordinaten() (uint16,uint16)
	
	// Vor.: --
	// Eff.: Das Objekt hat nun den neuen Typ.  
	SetzeTyp(t uint8)
	
	// Vor.: --
	// Erg.: Der Typ des Objekts ist geliefert.
	GibTyp() (uint8)
	
	// Vor.: --
	// Eff.: Das Objekt ist nun (nicht [mehr]) aktiv.
		// (Nur) aktive Objekte werden in den Methoden gezeichnet und getroffen. 
	SetzeAkt(akt bool)
	
	// Vor.: --
	// Erg.: True ist geliefert, genau dann wenn das Objekt aktiv ist.  	
	GibAkt() (bool)
	
	// Vor.: --
	// Eff.: Das Objekt hat nun den eingegebenen Inhalt. 
	SetzeInhalt(inhalt string)
	
	// Vor.: --
	// Eff.: Das Objekt hat nun die eingegebene Erstellung. 
	SetzeErstellung(erstellt int64)
	
	// Vor.: --
	// Erg.: Liefert die Erstellung des Objekts.  	
	GibErstellung() (int64)
	
	// Vor.: Das Objekt ist aktiv. Das gfx-Fenster ist offen.
	// Eff.: Zeichnet das Objekt im gfx-Fenster.
	Zeichnen()
	
	// Vor.: Das Objekt ist aktiv.
	// Eff.: Das Objekt erhält Koordinaten (z.B. eines Maus-Klicks) und optional eine 
		// Spezifizierung opt der Art des Treffers (z.B. linke/recht Maustaste).
	// Erg.: True ist geliefert gdw die definierten Trefferbedingungen des entsprechenden Objekts erfüllt sind.
		// Außerdem ist der Erstellungswert des Objekts geliefert (um z.B. seine Lebenszeit zu bestimmen).
	Getroffen(x,y uint16, opt uint8) (bool, int64)
}
