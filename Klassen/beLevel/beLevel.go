// Klasse zum Bau der Level zum Bauelementespiel

// Spezifikation

// Martin Seiß	31.3.2023

// Generiert die Levelparameter und gibt sie zurück 
// werden können.
// Level können in der New-Funktion vertauscht werden, daher ist die
// Bezeichnung vorläufig. 

// Erg: Ein Instanz der Klasse Level ist geliefert.
// func New() *data {


package beLevel

// ------------  importierte Pakete und Klassen  -----------------//
import sch "../schaltungen"


type Level interface {
	
	// Erg: Die Nummer des Schaltkreises entsprechend des Levels ist
	//		geliefert.
	GibSchaltkreis(nummer uint16) sch.Schaltung
	
	// Erg: Die Größe der Bauelemente in x-Richtung ist geliefert.
	GibXSize(nummer uint16) uint16

	// Erg: Die maximal zu erreichende Punktzahl im Level ist geliefert.
	GibMaxPunktzahl(nummer uint16) uint16

	// Erg: Die erreichte Punktzahl im Level ist geliefert.
	GibPunktzahl(nummer uint16) uint16
	
	// Eff: Die erreichte Punktzahl im Level ist neu gesetzt.
	SetzePunktzahl(nummer, punkte uint16)

	// Erg: Die minimle Anzahl an Schalterbetätigungen, nötig um das Level
	//		zu gewinnen, ist geliefert.
	GibMinSchalter(nummer uint16) uint16
	
	// Erg: Die Gesamtanzahl der Level ist geliefert.
	AnzahlLevel() uint16
	
	// Erg: Die Ausgabetext für das Level ist geliefert.
	GibText(nummer uint16) []string

}
