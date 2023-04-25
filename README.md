# LWB-Adventure

Mindmap (**ACHTUNG: Löscht sich automatisch alle 30 Tage, immer mal wieder runterladen und neu hochladen!!!**):
https://map.kits.blog/map/d25cd268-72d7-46ee-94c8-445c4e0ad085

## Regeln für das Repository: 

1. Jedes Minigame ist in einem Ordner organisiert.
2. Keiner ändert die Dateien von anderen.
3. Bilder: Es gibt einen Überordner und dann Unterordner für die Minigames
--> Bildpfade sollten also immer relativ zum Mainpfad angegeben werden z.B. "/images/Minigame1/Testbild.bmp"

## Hinweise zum Coden: 

1. Kommentieren!!! ;) 

## Verwendete Pakete:

- gfx (Schmidt)
### Ben:
- objekte
- raeume

### Ideen für gemeinsame Projektordner:
- sounds
- bilder

### Ideenspeicher
- Musik + Soundeffekte in den Minigames, auch im Hauptgame --> Wie werden die Übergänge gemacht?

### Vereinbarungen
- Speichern zwischen den MiniGames
- Note float32 (z.B. 1,3, Noten zwischen 1,0 und 5,0, gerundet ,3 / ,7 / ,0), Punktzahl uint32
- Fenstergröße: 1200 x 700 
- Trennung von Eingabe, Routine, Darstellung?

? Highscore-Ausgabe am Ende jedes Minigames in Form eines Notenspiegels mit Notendurchschnitt
    und Namensliste der besten 5/10 SpielerInnen ??? --> Klasse "Notenspiegel"?
? Zwischenzeugnis (eigenes) anzeigen lassen ?

### Maingame
- Point&Click im Comic Style
- Startbildschirm ((Laden) oder Neu)
- Name eingeben (Schmidt Editor) und (Bild) 
- Hauptraum, Türen für jedes Semester, Beenden Tür
- Räume für jedes Semester (auf Türen klicken, auch für zurück
- Auf Dozent klicken (erstmal)
- Minigame/Semester muss erfolgreich (<=4,0) bestanden wurden sein
- Wenn Semester 1 bestanden --> bekommt man einen Schlüssel für Semester 2
- rechts unten immer verfügbare Schlüssel anzeigen
- Tür zur Zeugnisübergabe (man braucht 4 Schlüssel)
  - Alle Noten
  - Bilder von Dozenten, Heidi
  - Text
  - Spiel Beenden Tür
  - Zurück Tür, weiterspielen, Noten verbessern

### Klasse SpielerIn
Parameter:
- Punktestand [Semester][Minigame]uint32
- Notenfeld [Semester][Minigame]float32 
- AnzahlSchlüssel uint8
- Name
- Bild/Avatar


Viel Spaß uns!:) 
