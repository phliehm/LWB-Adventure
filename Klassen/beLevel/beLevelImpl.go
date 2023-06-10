// Klasse zum Bau der Level zum Bauelementespiel

// Implementierung

// Martin Seiß	31.3.2023



package beLevel


import sch "../schaltungen"


// Structure für die Parameter eines Einzellevels
type EinzelLevel struct {
	sk 			sch.Schaltung 		// Schaltkreis
	maxPunkte	uint16				// maximal erreichbare Punktzahl
	punkte		uint16				// Punkte ereicht
	minSchalter uint16				// maximale Anzahl von Schaltungen,
									// um maximale Punkte zu erreichen
	xSize		uint16				// Größe der Bauelemente in x-Richtung
	text		[]string			// Textausgabe für den Levelbeginn
}



type data struct {
	elevel []EinzelLevel	// Liste von Eizelleveln
}


func New() *data {
	var lev *data = new(data)
	lev.elevel = make([]EinzelLevel,0)
	lev.elevel = append(lev.elevel,baueLevel1())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel2())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel3())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel4())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel5())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel6())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel7())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel9())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel9b())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel8())	// Füge Level hinzu
	lev.elevel = append(lev.elevel,baueLevel10())	// Füge Level hinzu
	return lev
}


func (lev *data) GibSchaltkreis(nummer uint16) sch.Schaltung {
	return lev.elevel[nummer].sk
}

	
func (lev *data) GibXSize(nummer uint16) uint16 {
	return lev.elevel[nummer].xSize
}


func (lev *data) GibPunktzahl(nummer uint16) uint16 {
	return lev.elevel[nummer].punkte
}
	
	
func (lev *data) GibMaxPunktzahl(nummer uint16) uint16 {
	return lev.elevel[nummer].maxPunkte
}

	
func (lev *data) SetzePunktzahl(nummer, punkte uint16) {
	lev.elevel[nummer].punkte = punkte
}


func (lev *data) GibMinSchalter(nummer uint16) uint16 {
	return lev.elevel[nummer].minSchalter
}


func (lev *data) AnzahlLevel() uint16 {
	return uint16(len(lev.elevel))
}


func (lev *data) GibText(nummer uint16) []string {
	return lev.elevel[nummer].text
}

