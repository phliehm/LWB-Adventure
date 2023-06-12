// Annalena Cyriacus
// 12.06.2023
// Testdatei für die Klasse Viereck

package main

import ( "fmt" ; "./vierecke" ; "gfx")

func main () {
	
	v1:= vierecke.New(200,200,100,150,255,10,10,255)
	v2:= vierecke.New(400,200,400,350,200,350,100,150)
	
	xA,yA,xB,yB,xC,yC,xD,yD := v1.GetKoordinaten()
	fmt.Println ("Koordinaten von v1:",xA,yA,xB,yB,xC,yC,xD,yD)
	
	v1.SetKoordinaten(300,300,200,500,700,500,600,300)
	
	fmt.Println ("v1 - kodiert:",v1.Kodieren())
	v3:=v1.Kopie().(vierecke.Viereck)
	fmt.Println ("v3 - kodiert:",v3.Kodieren())
	
	gfx.Fenster(1000,600)
	
	fmt.Println ("v1 und v2 werden gezeichnet!")
	v1.Zeichnen()
	v2.Zeichnen()
	
	r1,g1,b1 := v1.GibFarbe()
	fmt.Println ("Farbe von v1:",r1,g1,b1)
    
    v3.SetzeFarbe(255,0,0)
    v3.Zeichnen()
    
	r3,g3,b3 := v3.GibFarbe()
	fmt.Println ("Farbe von v3:",r3,g3,b3)
	
	fmt.Println ("String-Repräsentation von v2:",v2.String())

	fmt.Println ("1. Versuch: Mauszeiger im Bereich von v2?",v2.Angeklickt(50,50))
	fmt.Println ("2. Versuch: Mauszeiger im Bereich von v2?",v2.Angeklickt(470,250))
	
	v2.DeaktiviereKlickbar()
	fmt.Println ("neuer Versuch: Mauszeiger im Bereich von v2?",v2.Angeklickt(470,250))
	v2.AktiviereKlickbar()
	fmt.Println ("und nochmal: Mauszeiger im Bereich von v2?",v2.Angeklickt(470,250))
	
	gfx.TastaturLesen1()
	
}

