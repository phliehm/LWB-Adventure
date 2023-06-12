// ADT objekte - Test

// Benjamin Schneider    29.3.2023

package main


import ("./Klassen/objekte"; . "gfx"; "time" ; "fmt" ; "math/rand")


func main() {
	var ende bool = false
	
	
	maus 		:= objekte.New(0, 0, 0, 0)	
	erstelleObjekte := objekte.New(0, 0, 150, 32)
	erstelleObjekte.SetzeInhalt("     Erstelle")
	leereObjekte := objekte.New(225, 0, 150, 32)
	leereObjekte.SetzeInhalt("    Leere")
	endeButton 	:= objekte.New(975, 550, 150, 32)
	endeButton.SetzeInhalt("        ENDE")

	obj := make([]objekte.Objekt,0)
		
	
	Fenster(1200,700)
	SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)

	go zeichne(&obj, maus, leereObjekte, erstelleObjekte, endeButton, &ende)

	// Mausabfrage
A:	for {
		_, status, mausX, mausY := MausLesen1()
		maus.SetzeKoordinaten(mausX,mausY)
		if status==1 {
			for _,ob := range obj { 									// überprüft Objekte im Array
				if okay,erstellung := ob.Getroffen(mausX,mausY,1); okay {
					fmt.Println("Dieses Objekt hat folgende Erstellung: ",erstellung)
					switch ob.GibTyp() {
						case 32: ob.SetzeTyp(31)
						case 31: ob.SetzeTyp(32)
					}
				}
			}
			if okay,_ := erstelleObjekte.Getroffen(mausX,mausY,1); okay {
				neu := objekte.New(
				obj = append(obj,neu)
			}
			if okay,_ := leereObjekte.Getroffen(mausX,mausY,1); okay {
				neu := objekte.New(
				obj = make([]objekte.Objekt,0)
			}
			if okay,_ := endeButton.Getroffen(mausX,mausY,1); okay {
				ende = true
				break A
			}
		}
	}
	TastaturLesen1()
}

func zeichne (obj *[]objekte.Objekt, maus,leereObjekte,erstelleObjekte,endeButton objekte.Objekt, ende *bool) {   	
	
	var t1 int64 = time.Now().UnixNano() 		//Startzeit
	var anz,anzahl int                  		// zur Bestimmung der Frames pro Sekunde
	var verzögerung = 90
	
	SetzeFont ("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",20)
	
	for {
		UpdateAus () 										// Nun wird alles im nicht sichtbaren "hinteren" Fenster gezeichnet!
		Stiftfarbe(0,0,0)
		Cls()												// Cleart vollständigen Screen
		for _,ob := range *obj { 								// Zeichnet alleweiteren Objekte ein
			ob.Zeichnen()
		}
		leereObjekte.Zeichnen()
		erstelleObjekte.Zeichnen()
		endeButton.Zeichnen()
		maus.Zeichnen()														// Zeichnet Maus
		
		if time.Now().UnixNano() - t1 < 1000000000 { 		//noch in der Sekunde ...
			anz++
		} else {
			t1 = time.Now().UnixNano() 						// neue Sekunde
			anzahl = anz
			anz=0
			if anzahl < 100 { verzögerung--}				//Selbstregulierung der 
			if anzahl > 100 { verzögerung++}				//Frame-Rate :-)		-- dieser 8-zeilige Abschnitt wurde  von Herrn Schmidt übernommen
		}
		UpdateAn () 										// Nun wird der gezeichnete Frame sichtbar gemacht!
		time.Sleep(time.Duration(verzögerung * 1e5)) 		// Immer ca. 100 FPS !!
		if *ende { return }
	}
}
