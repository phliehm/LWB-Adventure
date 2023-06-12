package vaderobi
// Autor:(c) Stefan Schmidt, St.Schmidt@online.de
// Datum: 09.04.2015 ; letzte Änderung: 18.02.2019
// Zweck:
/* Das Paket verwaltet eine rechteckige Welt aus schachbrettförmig
 angeordenten Plätzen (Standard: 12 Plätze hoch und 16 Plätze breit)
 und einem Roboter namens Robi. 
 Robi steht immer auf genau einem der Plätze in der Welt (Robis
 Platz) und schaut immer in genau eine Himmelsrichtung (Nord, Ost, Süd,
 West), die als "Robis Richtung" bezeichnet wird.
 In jeder Himmelsrichtung gibt es genau einen Nachbarplatz oder keinen.
 Ein Platz kann durch Robi markiert oder entmarkiert werden.
 Robi hat eine Tasche mit Klötzen, die er einzeln auf seinen Platz
 ablegen oder von dort einzeln aufnehmen kann.
 Robi hat außerdem Zugriff auf Mauersteine und kann Plätze zumauern
 oder entmauern.
 Zu Beginn steht Robi in einer leeren Welt in der linken oberen Ecke
 mit Blick nach Süden und hat 999 Klötze in seiner Tasche.
 Wird eine bereits existierende Welt geladen, so kann Robi an einer
 anderen Position starten. In der Welt befinden sich ebendo insgesamt 
 999 Klötze, allerdings nicht alle unbedingt in seiner Tasche.
 Mit dem Import des Pakets wird automatisch ein Grafikfenster geöffnet,
 indem die Welt von Robi sichtbar gemacht wird. Der Grafikbereich des
 Fensters ist standardmäßig 800 x 600 Pixel. Ein Platz hat dann 50 x 50 Pixel.
 
 In der Grundeinstellung befindet sich Robi im sogenannten Schrittmodus,
 d.h., jede einzelne „Grundaktion“, die von Robi ausgeführt werden soll 
 und die Welt verändert, muss mit einem Druck auf die Entertaste bestätigt
 werden (siehe Funktion Schrittmodus).
 */


/* ----------------LAUF-ANWEISUNGEN ---------------------*/

// Vor.: Robi hat in seiner Richtung einen betretbaren Nachbarplatz.
// Eff.: Robi steht in der gleichen Richtung wie vorher auf diesem
// Nachbarplatz.
// func Laufen1 () 
	

// Vor.: keine 
// Eff.: Robis Richtung ist um 90 Grad nach links gedreht.
// func LinksDrehen () 


// Vor.: keine 
// Eff.: Robis Richtung ist um 90 Grad nach rechts gedreht.
// func RechtsDrehen () 


// Vor.: keine 
// Erg.: -true- ist geliefert, genau dann, wenn Robi in seiner Richtung
// keinen Nachbarplatz hat. 
// func AmRand () bool 


// Vor.: keine 
// Erg.: -true- ist geliefert, genau dann, wenn Robi in der linken 
// oberen Ecke der Welt (nordwestlichste Ecke) steht.
// func InLinkerObererEcke () bool 


/* ----------------MELDE-ANWEISUNGEN ---------------------*/

// Vor.: keine 
// Eff.: -text- und -n- sind in der letzten Zeile des Grafikfensters 
// ausgegeben. Der aufrufende Prozess wurde solange angehalten, bis der 
// Benutzer in der Ausgabe mit einem Druck auf <ESC> quittiert hatte.
// Nun ist die Meldung wieder aus dem Grafikfenster entfernt.
// func FehlerMelden (text string, n uint)  //alter Funktionsname
// func Melden (text string, n uint)

// Vor.: keine 
// Eff.: Das Programm ist mit der Fehlermeldung "Programm beendet! 
// Druecken Sie <ESC> !" angehalten. Nachdem der Benutzer mit einem 
// Druck auf <ESC> quittiert hat, ist das Grafikfenster geschlossen und
// das Programm beendet.
// func Fertig () 


/* ----------------MARKIER-ANWEISUNGEN ---------------------*/

// Vor.: keine 
// Eff.: Robis Platz ist markiert.
// func Markieren () 


// Vor.: keine 
// Eff.: Robis Platz ist nicht markiert.
// func Demarkieren () 


// Vor.: keine 
// Erg.: -true- ist geliefert, falls Robis Platz markiert ist, 
// andernfalls ist -false- geliefert.
// func Markiert () bool 


// Vor.: In Robis Richtung gibt es einen Nachbarplatz. 
// Erg.: -true- ist geliefert, falls der Nachbarplatz in Robis Richtung
// markiert ist. Andernfalls ist -false- geliefert.
//func NachbarMarkiert () bool 


/* ----------------LEGE-ANWEISUNGEN ---------------------*/

// Vor.: keine 
// Erg.: -true- ist geliefert, falls auf Robis Platz kein Klotz liegt,
// anderfalls ist -false- geliefert.
// func Leer () bool 


// Vor.: Auf Robis Platz liegt mindestens ein Klotz.
// Eff.: Auf Robis Platz liegt nun ein Klotz weniger und in Robis Tasche
// ist ein Klotz mehr. In der linken oberen Ecke des Platzes ist die
// aktuelle Anzahl der Klötze auf dem Platz angezeigt. Ist kein Klotz auf 
// dem Platz,so ist keine Anzahl der Klötze angezeigt.
// func Leeren1 () 


// Vor.: keine
// Erg.: -true- ist geliefert, falls Robi in seiner Tasche noch 
// mindestens einen Klotz hat. Andernfalls ist -false- geliefert.
// func HatKloetze () bool 


// Vor.: Robis Tasche mit Klötzen ist nicht leer.
// Eff.: In Robis Tasche ist nun ein Klotz weniger und auf seinem Platz 
// ist nun einer mehr. In der linken oberen Ecke des Platzes ist die
// aktuelle Anzahl der Klötze auf dem Platz angezeigt. 
// func Legen1 () 


/* ----------------MAUER-ANWEISUNGEN ---------------------*/

// Vor.: keine
// Erg.: -true- is geliefert, falls in Robis Richtung ein Nachbarfeld
// ist, das zugemauert ist. Andernfalls ist -false- geliefert.
// func VorMauer () bool 


// Vor.: In Robis Richtung gibt es ein Nachbarfeld, das nicht 
// zugemauert ist. 
// Eff.: Robi steht jetzt auf diesem Nachbarfeld und sein ursprünglicher
// Standort ist nun zugemauert. Klötze, die sich auf dem zugemauerten
// Feld befunden haben, sind nun in Robis Tasche. Eine auf dem 
// zugemauerten Feld ggf. vorhandene Markierung ist jetzt entfernt.
//func Mauern1 () 


// Vor.: In Robis Richtung gibt es ein Nachbarfeld, das zugemauert ist. 
// Eff.: Dieses Nachbarfeld ist nun nicht mehr zugemauert und Robi steht
// mit gleicher Richtung auf diesem  entmauerten Platz. 
// func Entmauern1 () 


/* ----------------SPEZIAL-ANWEISUNGEN ---------------------*/

// Vor.: keine 
// Eff.: Wurde Schrittmodus (true) aufgerufen, so muss nun jede einzelne
// Aktion von Robi, die die Welt ändert, mit einem Druck auf die 
// Entertaste aktiviert werden, andernfalls nicht. 
// Standardeinstellung ist -true-.
// func Schrittmodus (mode bool) 


// Vor.: -name- ist ein gültiger Dateiname inkl. Pfad und die zugehörige
// Datei existiert dort.
// Eff.: Die Welt aus der Datei -name- wurde aktiviert. Dazu gehört auch
// Robis Platz und seine Blickrichtung.
// func WeltLaden (name string) 


// Vor.: -name- ist ein gültiger Dateiname inkl. Pfad und diese Datei
// existiert bis jetzt noch nicht.
// Eff.: Diese Datei existiert nun und die aktuelle Welt und Robis
// Eigenschaften sind in der Datei -name- abgespeichert.
// func WeltSpeichern (name string) 


// Vor.: -breite- und -hoehe- sind nicht kleiner als -feldbreite-.
// Eff.: Im Grafikfenster erscheint eine neue leere Welt mit Robi in der
// linken oberen Ecke mit Blickrichtung nach Süden. Die vorher 
// vorhandene Welt ist verloren.
// func NeueWelt (breite, hoehe, feldbreite uint16) 


// Vor.: -p- ist ein Wert zwischen 50 und 300 und gibt den Zoomfaktor
//       in Prozent an.
// Eff.: Robis Welt ist entsprechend dem Zoomfaktor auf dem Bildschirm
//       angezeigt. Das Grafikfenster hat seine Größe geändert. Die Feld-
//       breite der Felder von Robis Welt wurde entsprechend angepasst.
// func WeltZoomen (p uint16)


// Vor.: keine
// Eff.: Solange der Benutzer nicht <ESC> drückt ist er im Baumodus
// gefangen und kann Robi mit den Pfeiltasten durch die Welt steuern.
// Außerdem kann er mit <1> Robis Platz markieren bzw. demarkieren,
// mit <2> jeweils einen Klotz auf Robis Platz ablegen, solange Robi
// Klötze in der Tasche hat, und mit <3> jeweils einen Klotz von Robis
// Platz aufnehmen und in die Tasche stecken, solange Klötze auf Robis
// Platz liegen. Mit <4> kann gemauert und mit <5> entmauert werden.
// func Baumodus () 
