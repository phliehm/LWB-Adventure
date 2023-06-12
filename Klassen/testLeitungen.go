// ADT leitungen - Test zusammen mit Bauelemente

// Martin Seiß    21.3.2023


package main

import l "./leitungen"
//import b "./bauelemente"
import "fmt"
//import "gfx"



func main() {

	var l1 l.Leitung = l.New(1, 3, 1, 175)
	fmt.Println("Leitung 1 von ",l1.GibVonID()," nach ", l1.GibNachID(),
		" an Eingang ",l1.GibEinNr())
	fmt.Println("Knick an x-Position: ",l1.GibXPos())
	
	var l2 l.Leitung = l.New(2, 3, 2, 175)
	fmt.Println("Leitung 2 von ",l2.GibVonID()," nach ", l2.GibNachID(),
		" an Eingang ",l2.GibEinNr())
	fmt.Println("Knick an x-Position: ",l2.GibXPos())

	var l3 l.Leitung = l.New(3, 4, 1, 325)
	fmt.Println("Leitung 3 von ",l3.GibVonID()," nach ", l3.GibNachID(),
		" an Eingang ",l3.GibEinNr())
	fmt.Println("Knick an x-Position: ",l3.GibXPos())


	fmt.Scanln()

	
}
