// Paket zum Bauelementespiel

// Implementation

// Martin Seiß    12.6.2023

//"BUG: Gattersymbol wird nicht gezeichet, da Testdatei "+
//		"im falschen Verzeichnis relativ zum Fontverzeichnis."

package main


import "fmt"
import "gfx"
import sch "./schaltungen"
import b "./bauelemente"


var path string = "" //"../"


func main()  {

	gfx.Fenster(700,600)
	var xSize uint16 = 75


	// ---------------- Erzeuge Schaltkreis ----------------------- //
	var sk sch.Schaltung = sch.New()		// neuer leerer Schaltkreis
	
		// --------------   Schalter einfügen   -----------------------//
	sk.BauteilEinfuegen(1,100,100,b.Schalter)
	sk.BauteilEinfuegen(2,100,400,b.Schalter)
	
	// --------------   Logische Gatter einfügen   ----------------//
	sk.BauteilEinfuegen(3,300,250,b.AND)	
	
	// --------------   Lampen einfügen   -------------------------//
	sk.BauteilEinfuegen(4,500,250,b.Lampe)

	// --------------   Leitungen einfügen   ----------------------//
	sk.VerbindungEinfuegen(1,3,1,200)
	sk.VerbindungEinfuegen(2,3,2,200)
	sk.VerbindungEinfuegen(3,4,1,400)


	fmt.Println("Schalter 1: ",sk.GibSchalterwert(1))
	fmt.Println("Schalter 2: ",sk.GibSchalterwert(2))
	sk.SchaltungBerechnen()
	fmt.Println("Lampe an?: ",sk.GibLampenwert(4))	
	
	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)		
	sk.Zeichnen(xSize)
		
	fmt.Println()
	fmt.Scanln()

	sk.SchalteSchalterAn(1,true)
	fmt.Println("Schalter 1: ",sk.GibSchalterwert(1))
	fmt.Println("Schalter 2: ",sk.GibSchalterwert(2))
	sk.SchaltungBerechnen()
	fmt.Println("Lampe an?: ",sk.GibLampenwert(4))

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)
	sk.Zeichnen(xSize)
	
	fmt.Println()
	fmt.Scanln()

	sk.SchalteSchalterAn(2,true)
	fmt.Println("Schalter 1: ",sk.GibSchalterwert(1))
	fmt.Println("Schalter 2: ",sk.GibSchalterwert(2))
	sk.SchaltungBerechnen()
	fmt.Println("Lampe an?: ",sk.GibLampenwert(4))

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)
	sk.Zeichnen(xSize)
	
	fmt.Println()
	fmt.Scanln()

	sk.SchalteSchalterAn(1,false)
	fmt.Println("Schalter 1: ",sk.GibSchalterwert(1))
	fmt.Println("Schalter 2: ",sk.GibSchalterwert(2))
	sk.SchaltungBerechnen()
	fmt.Println("Lampe an?: ",sk.GibLampenwert(4))

	gfx.Stiftfarbe(255,255,255)
	gfx.Cls()
	gfx.Stiftfarbe(0,0,0)
	sk.Zeichnen(xSize)
	
	fmt.Println()
	fmt.Scanln()
	
	fmt.Println("BUG: Gattersymbol wird nicht gezeichet, da Testdatei "+
		"im falschen Verzeichnis relativ zum Fontverzeichnis.")

}
