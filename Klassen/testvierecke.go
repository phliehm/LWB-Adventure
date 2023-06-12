// Annalena Cyriacus
// 12.06.2023
// Testdatei für die Klasse vierecke

package main

import ( "fmt" ; "./vierecke" ; "gfx")

func main () {
	
	var farbwert uint8 = 255
	var v1,v2,v3,v4,v5 vierecke.Viereck
	
	v1 = vierecke.New(300,300,200,500,700,500,600,300)
	v2 = vierecke.New(935,290,935,520,1010,545,1020,255)
	v3 = vierecke.New(425,330,430,470,465,460,460,350)
	v4 = vierecke.New(720,355,710,455,740,460,750,340)
	v5 = vierecke.New(570,350,570,435,625,435,625,350)
	
	v1.AktiviereKlickbar()
	fmt.Println("v1 ist jetzt klickbar!")
	v2.AktiviereKlickbar()
	fmt.Println("v2 ist jetzt klickbar!")
	v3.AktiviereKlickbar()
	fmt.Println("v3 ist jetzt klickbar!")
	v4.AktiviereKlickbar()
	fmt.Println("v4 ist jetzt klickbar!")
	v5.AktiviereKlickbar()
	fmt.Println("v5 ist jetzt klickbar!")
	
		
	xA,yA,xB,yB,xC,yC,xD,yD := v1.GetKoordinaten()
	fmt.Println ("Koordinaten von v1:",xA,yA,xB,yB,xC,yC,xD,yD)
	
	v1.SetKoordinaten(45,235,40,595,220,595,220,235)
	
	fmt.Println ("v1 - kodiert:",v1.Kodieren())
	v5 = v1.Kopie().(vierecke.Viereck)
	fmt.Println ("v5 - kodiert:",v5.Kodieren())
	
	gfx.Fenster(1200,700)
	
	fmt.Println ("v1, v2 und v5 werden gezeichnet!")
	v1.Zeichnen()
	v2.Zeichnen()
	v5.Zeichne()
	
	r1,g1,b1 := v1.GibFarbe()
	fmt.Println ("Farbe von v1:",r1,g1,b1)
    
    v3.SetzeFarbe(255,0,0)
    v3.Zeichnen()
    v4.SetzeFarbe(0,255,0)
    v4.Zeichnen()
    
	r3,g3,b3 := v3.GibFarbe()
	fmt.Println ("Farbe von v3:",r3,g3,b3)
	
	fmt.Println ("String-Repräsentation von v2:",v2.String())
	
	for {
		taste, status, mausX, mausY := gfx.MausLesen1()
		
		if taste==1 && status==1 {
			fmt.Println("hi")
			if v1.Angeklickt(mausX,mausY) {
				fmt.Println("v1 wurde angeklickt!")
				v1.SetzeFarbe(150,farbwert+50,farbwert-100)
				v1.Zeichnen()
				farbwert = farbwert + 66
			}
			if v2.Angeklickt(mausX,mausY) {
				fmt.Println("v2 wurde angeklickt!")
				v2.SetzeFarbe(255,farbwert-50,farbwert+100)
				v2.Zeichnen()
				farbwert = farbwert + 111
			}
			if v3.Angeklickt(mausX,mausY) {
				fmt.Println("v3 wurde angeklickt!")
				v2.DeaktiviereKlickbar()
				fmt.Println("Klickbarkeit von v2 wurde DEaktiviert!")
			}
			if v4.Angeklickt(mausX, mausY) {
				v2.AktiviereKlickbar()
				fmt.Println("Klickbarkeit von v2 wurde aktiviert!")
			}
			if v5.Angeklickt(mausX, mausY) {
				fmt.Println("v5 wurde angeklickt, sodass das gfx-Fenster nun mit einem Tastendruck geschlossen werden kann!")
			}
		}
	}
	
	fmt.Println("hi")
	gfx.TastaturLesen1()
	
}

