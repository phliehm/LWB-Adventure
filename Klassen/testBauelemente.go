// ADT bauelemente - Test

// Martin Seiß    21.3.2023


package main

//import sch "./schaltungen"
import b "./bauelemente"
import "fmt"
import "gfx"



func main() {
	
	var xSize uint16 = 100
	
	gfx.Fenster(500,500)
	//gfx.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf",int(xSize))
	if gfx.GibFont() == "" {
		fmt.Println("Kein Font")
	}
	// fmt.Println(gfx.GibFont())
	fmt.Println("Hallo")
	// Font muss vorher einmal gesetzt werden
	
	//  ------------  Teste Schalter ----------------------  //
	
	var schalter b.Bauelement
	schalter = b.New(1,100,100, false,false,false, b.Schalter)
	
	fmt.Println("Bauteilnummer: ", schalter.GibID() )
	if schalter.GibBauelementtyp() == b.Schalter {
		fmt.Println("Bauteiltyp: Schalter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",schalter.GibEingang(1))
	fmt.Println("Eingang 2: ",schalter.GibEingang(2))
	fmt.Println("Berechnet?: ", schalter.GibBerechnet() )
	fmt.Println("Schalterausgang: ", schalter.BerechneAusgang())
	schalter.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",schalter.GibEingang(1))
	fmt.Println("Eingang 2: ",schalter.GibEingang(2))
	fmt.Println("Schalterausgang: ", schalter.BerechneAusgang())
	x,y := schalter.GibPosXY()
	fmt.Println("x,y: ",x,y)
	
	fmt.Println("Berechnet?: ", schalter.GibBerechnet() )
	schalter.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", schalter.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",schalter.GibVerbindungen())
	
	schalter.ZeichneBauelement(xSize)
	
	fmt.Println()

	
	//  ------------  Teste AND-Gatter ----------------------  //

	var and b.Bauelement
	and = b.New(2,300,100, false,false,false, b.AND)
	fmt.Println("Bauteilnummer: ", and.GibID() )
	and.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	if and.GibBauelementtyp() == b.AND {
		fmt.Println("Bauteiltyp: AND-Gatter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",and.GibEingang(1))
	fmt.Println("Eingang 2: ",and.GibEingang(2))
	fmt.Println("Ausgang: ", and.BerechneAusgang())
	and.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",and.GibEingang(1))
	fmt.Println("Eingang 2: ",and.GibEingang(2))
	fmt.Println("Ausgang: ", and.BerechneAusgang())
	and.SetzeEingang(2, true)
	fmt.Println("Eingang 1: ",and.GibEingang(1))
	fmt.Println("Eingang 2: ",and.GibEingang(2))
	fmt.Println("Ausgang: ", and.BerechneAusgang())
	and.SetzeEingang(1, false)
	fmt.Println("Eingang 1: ",and.GibEingang(1))
	fmt.Println("Eingang 2: ",and.GibEingang(2))
	fmt.Println("Ausgang: ", and.BerechneAusgang())
	
	x,y = and.GibPosXY()
	fmt.Println("x,y: ",x,y)
	fmt.Println("Berechnet?: ", and.GibBerechnet() )
	and.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", and.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",and.GibVerbindungen())
	
	and.ZeichneBauelement(xSize)
	
	fmt.Println()

	
	//  ------------  Teste OR-Gatter ----------------------  //

	var or b.Bauelement
	or = b.New(3,300,250, false,false,false, b.OR)
	fmt.Println("Bauteilnummer: ", or.GibID() )
	or.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	if or.GibBauelementtyp() == b.OR {
		fmt.Println("Bauteiltyp: or-Gatter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",or.GibEingang(1))
	fmt.Println("Eingang 2: ",or.GibEingang(2))
	fmt.Println("Ausgang: ", or.BerechneAusgang())
	or.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",or.GibEingang(1))
	fmt.Println("Eingang 2: ",or.GibEingang(2))
	fmt.Println("Ausgang: ", or.BerechneAusgang())
	or.SetzeEingang(2, true)
	fmt.Println("Eingang 1: ",or.GibEingang(1))
	fmt.Println("Eingang 2: ",or.GibEingang(2))
	fmt.Println("Ausgang: ", or.BerechneAusgang())
	or.SetzeEingang(1, false)
	fmt.Println("Eingang 1: ",or.GibEingang(1))
	fmt.Println("Eingang 2: ",or.GibEingang(2))
	fmt.Println("Ausgang: ", or.BerechneAusgang())
	
	x,y = or.GibPosXY()
	fmt.Println("x,y: ",x,y)
	fmt.Println("Berechnet?: ", or.GibBerechnet() )
	or.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", or.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",or.GibVerbindungen())
	
	or.ZeichneBauelement(xSize)
	
	fmt.Println()

	//  ------------  Teste XOR-Gatter ----------------------  //

	var xor b.Bauelement
	xor = b.New(4,300,400, false,false,false, b.XOR)
	fmt.Println("Bauteilnummer: ", xor.GibID() )
	xor.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	if xor.GibBauelementtyp() == b.XOR {
		fmt.Println("Bauteiltyp: xor-Gatter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",xor.GibEingang(1))
	fmt.Println("Eingang 2: ",xor.GibEingang(2))
	fmt.Println("Ausgang: ", xor.BerechneAusgang())
	xor.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",xor.GibEingang(1))
	fmt.Println("Eingang 2: ",xor.GibEingang(2))
	fmt.Println("Ausgang: ", xor.BerechneAusgang())
	xor.SetzeEingang(2, true)
	fmt.Println("Eingang 1: ",xor.GibEingang(1))
	fmt.Println("Eingang 2: ",xor.GibEingang(2))
	fmt.Println("Ausgang: ", xor.BerechneAusgang())
	xor.SetzeEingang(1, false)
	fmt.Println("Eingang 1: ",xor.GibEingang(1))
	fmt.Println("Eingang 2: ",xor.GibEingang(2))
	fmt.Println("Ausgang: ", xor.BerechneAusgang())
	
	x,y = xor.GibPosXY()
	fmt.Println("x,y: ",x,y)
	fmt.Println("Berechnet?: ", xor.GibBerechnet() )
	xor.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", xor.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",xor.GibVerbindungen())
	
	xor.ZeichneBauelement(xSize)

	fmt.Println()
	
	//  ------------  Teste NOT-Gatter ----------------------  //

	var not b.Bauelement
	not = b.New(5,100,400, false,false,false, b.NOT)
	fmt.Println("Bauteilnummer: ", not.GibID() )
	not.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	if not.GibBauelementtyp() == b.NOT {
		fmt.Println("Bauteiltyp: not-Gatter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",not.GibEingang(1))
	fmt.Println("Eingang 2: ",not.GibEingang(2))
	fmt.Println("Ausgang: ", not.BerechneAusgang())
	not.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",not.GibEingang(1))
	fmt.Println("Eingang 2: ",not.GibEingang(2))
	fmt.Println("Ausgang: ", not.BerechneAusgang())
	not.SetzeEingang(2, true)
	fmt.Println("Eingang 1: ",not.GibEingang(1))
	fmt.Println("Eingang 2: ",not.GibEingang(2))
	fmt.Println("Ausgang: ", not.BerechneAusgang())
	not.SetzeEingang(1, false)
	fmt.Println("Eingang 1: ",not.GibEingang(1))
	fmt.Println("Eingang 2: ",not.GibEingang(2))
	fmt.Println("Ausgang: ", not.BerechneAusgang())
	
	x,y = not.GibPosXY()
	fmt.Println("x,y: ",x,y)
	fmt.Println("Berechnet?: ", not.GibBerechnet() )
	not.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", not.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",not.GibVerbindungen())
	
	not.ZeichneBauelement(xSize)


	//  ------------  Teste Lampe    ----------------------  //

	var lampe b.Bauelement
	lampe = b.New(5,100,250, false,false,false, b.Lampe)
	fmt.Println("Bauteilnummer: ", lampe.GibID() )
	lampe.SetzeFont("../Schriftarten/terminus-font/TerminusTTF-4.49.2.ttf")
	if lampe.GibBauelementtyp() == b.Lampe {
		fmt.Println("Bauteiltyp: lampe-Gatter" )
	} else {
		fmt.Println("Falscher Typ")
	}
	fmt.Println("Eingang 1: ",lampe.GibEingang(1))
	fmt.Println("Eingang 2: ",lampe.GibEingang(2))
	fmt.Println("Ausgang: ", lampe.BerechneAusgang())
	lampe.ZeichneBauelement(xSize)
	fmt.Scanln()
	lampe.SetzeEingang(1, true)
	fmt.Println("Eingang 1: ",lampe.GibEingang(1))
	fmt.Println("Eingang 2: ",lampe.GibEingang(2))
	fmt.Println("Ausgang: ", lampe.BerechneAusgang())
	lampe.ZeichneBauelement(xSize)
	fmt.Scanln()	
	lampe.SetzeEingang(2, true)
	fmt.Println("Eingang 1: ",lampe.GibEingang(1))
	fmt.Println("Eingang 2: ",lampe.GibEingang(2))
	fmt.Println("Ausgang: ", lampe.BerechneAusgang())
	lampe.ZeichneBauelement(xSize)
	fmt.Scanln()	
	lampe.SetzeEingang(1, false)
	fmt.Println("Eingang 1: ",lampe.GibEingang(1))
	fmt.Println("Eingang 2: ",lampe.GibEingang(2))
	fmt.Println("Ausgang: ", lampe.BerechneAusgang())
	lampe.ZeichneBauelement(xSize)
	fmt.Scanln()
	
	x,y = lampe.GibPosXY()
	fmt.Println("x,y: ",x,y)
	fmt.Println("Berechnet?: ", lampe.GibBerechnet() )
	lampe.SetzeBerechnet(true)
	fmt.Println("Berechnet?: ", lampe.GibBerechnet() )
	
	fmt.Println("Gib Verbindungen: ",lampe.GibVerbindungen())
	
	lampe.ZeichneBauelement(xSize)
	
/*
	or = New(1,200,100, false,false,false, b.OR)

	xor = New(1,200,100, false,false,false, b.XOR)
*/

	fmt.Scanln()

	
}
