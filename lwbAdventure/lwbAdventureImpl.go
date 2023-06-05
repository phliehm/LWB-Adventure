// Autor: A. Cyriacus
// Datum: 05.06.20
// Zweck: Implementierung ADO lwbAdventure

package lwbAdventure

/* Diese Komponente steuert den Ablauf den Spiels. Sie stellt damit den 
 * Controller in der MCV-Architektur dar.
 */

import (
	"gfx"
	"../darstellung"
	"../eingabe"
)

/*
import "../level"
import "../eingabe"
import "../spieler"
import "../kisten"
import "../scores"
import "../darstellung"
import "../config"
import "time"
*/


// globale Variablen
// ----------------


/*
var l level.Level	// aktuelles Level-Objekt
var maxlevel uint8  // Anzahl der vorhandenen Level
var kante uint16
*/

// Funktionen
// ----------

func Start() {
	
	gfx.Fenster(1200,700)
	darstellung.MainfloorDarstellen()
	eingabe.Eingabe()
	
}

/*
//func Start(name string) (bool, scores.Score) {

	var levelnr uint8
	var wert uint8
	var highscore scores.Score
	var geschafft bool = false
	
	maxlevel = config.GetMaxLevel()
	kante = config.GetKante()
	highscore = scores.New(maxlevel, name)
	
	// Spielschleife
A:	for levelnr = 1; levelnr <= maxlevel; levelnr++ {
		l = level.New(levelnr, name)
		darstellung.SpielfeldDarstellen(l)
		//l.SetLevelnr(levelnr)
		
		for !l.LevelGeschafft() {
			
			wert = eingabe.Eingabe()
			switch wert {
				case 0,1,2,3:
					ziehen(wert, kante)
				case 4:
					levelnr--
					l.SetAnzZuege(0)
					continue A
				case 5:
					break A
				default:
					panic("Unerwarteter Rückgabewert!")
			}
			
			if l.LevelGeschafft() {
				highscore.Eintragen(levelnr, l.GetAnzZuege())
				if levelnr == maxlevel { 
					geschafft = true
					highscore.Aktualisieren()					
				} else {
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
		l.SetAnzZuege(0)
		
	}
	
	darstellung.Ende()
	return geschafft, highscore
	
}
*/




// interne Hilfsfunktionen
// -----------------------

/*
func ziehen(richtung uint8, kante uint16) {
	var sp spieler.Spieler
	var k kisten.Kiste
	var aktspalte, aktzeile uint8
	
	sp = l.GetSpieler()
	aktspalte, aktzeile = sp.GetPosition()
	
	// nach oben
	if richtung == 0 {
		if l.FeldFrei(aktspalte, aktzeile-1) {
			darstellung.BewegungDarstellen(l, sp, nil, 0, -1, kante)
			sp.SetPosition(aktspalte, aktzeile-1)
			l.IncAnzZuege()
		}
		if l.KisteVorhanden(aktspalte, aktzeile-1) && l.FeldFrei(aktspalte, aktzeile-2) {
			k, _ = l.GetKiste(aktspalte, aktzeile-1)
			darstellung.BewegungDarstellen(l, sp, k, 0, -1, kante)
			sp.SetPosition(aktspalte, aktzeile-1)
			k.SetPosition(aktspalte, aktzeile-2)
			l.IncAnzZuege()
		}
	}
	
	// nach links
	if richtung == 1 {
		if l.FeldFrei(aktspalte-1, aktzeile) {
			darstellung.BewegungDarstellen(l, sp, nil, -1, 0, kante)
			sp.SetPosition(aktspalte-1, aktzeile)
			l.IncAnzZuege()
		}
		if l.KisteVorhanden(aktspalte-1, aktzeile) && l.FeldFrei(aktspalte-2, aktzeile) {
			k, _ = l.GetKiste(aktspalte-1, aktzeile)
			darstellung.BewegungDarstellen(l, sp, k, -1, 0, kante)
			sp.SetPosition(aktspalte-1, aktzeile)
			k.SetPosition(aktspalte-2, aktzeile)
			l.IncAnzZuege()
		}
	}
	
	// nach unten
	if richtung == 2 {
		if l.FeldFrei(aktspalte, aktzeile+1) {
			darstellung.BewegungDarstellen(l, sp, nil, 0, 1, kante)
			sp.SetPosition(aktspalte, aktzeile+1)
			l.IncAnzZuege()
		}
		if l.KisteVorhanden(aktspalte, aktzeile+1) && l.FeldFrei(aktspalte, aktzeile+2) {
			k, _ = l.GetKiste(aktspalte, aktzeile+1)
			darstellung.BewegungDarstellen(l, sp, k, 0, 1, kante)
			sp.SetPosition(aktspalte, aktzeile+1)
			k.SetPosition(aktspalte, aktzeile+2)
			l.IncAnzZuege()
		}
	}
	
	// nach rechts
	if richtung == 3 {
		if l.FeldFrei(aktspalte+1, aktzeile) {
			darstellung.BewegungDarstellen(l, sp, nil, 1, 0, kante)
			sp.SetPosition(aktspalte+1, aktzeile)
			l.IncAnzZuege()
		}
		if l.KisteVorhanden(aktspalte+1, aktzeile) && l.FeldFrei(aktspalte+2, aktzeile) {
			k, _ = l.GetKiste(aktspalte+1, aktzeile)
			darstellung.BewegungDarstellen(l, sp, k, 1, 0, kante)
			sp.SetPosition(aktspalte+1, aktzeile)
			k.SetPosition(aktspalte+2, aktzeile)
			l.IncAnzZuege()
		}
	}
	
	darstellung.SpielfeldDarstellen(l)
}
*/


