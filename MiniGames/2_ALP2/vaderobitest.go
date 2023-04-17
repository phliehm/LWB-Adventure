// Einfaches Testprogramm zu Robis Fähigkeiten.
package main

import . "./vaderobi"

func main () {
	Melden("Welt erstellt!",0)
	Markieren ()
	Legen1 ()
	Laufen1()
	Markieren ()
	Legen1 ()
	Laufen1()
	LinksDrehen()
	LinksDrehen() 
	if NachbarMarkiert () {Melden ("Das Nachbarfeld ist markiert!",0)}
	if Leer() {Melden ("Kein Klotz auf dem Feld!",0)}
	Laufen1 ()
	LinksDrehen ()
	LinksDrehen ()
	if !Leer() {Melden ("Klotz vorhanden!",0)}
	Leeren1 ()
	Laufen1()
	LinksDrehen ()
	if !NachbarMarkiert () {Melden ("Das Nachbarfeld ist nicht markiert!",0)}
	Laufen1()
	Laufen1()
	if Leer() {Melden ("Kein Klotz auf dem Feld!",0)}
	Mauern1 ()
	Mauern1 ()
	RechtsDrehen()
	Laufen1 ()
	RechtsDrehen ()
	Laufen1()
	Mauern1()
	Mauern1()
	LinksDrehen ()
	LinksDrehen ()
	if VorMauer() {Melden("Stehe vor einer Mauer!",0)}
	Entmauern1 ()
	Entmauern1 ()
	if !VorMauer() {Melden("Keine Mauer vor mir!",0)}
	Fertig ()
}
