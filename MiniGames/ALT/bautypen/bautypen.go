// Definition der verwendeten Bautypen

// Martin Seiß    24.3.2023

package bautypen

type Bautyp uint16

const (
	Schalter 	Bautyp = 0
	AND			Bautyp = 1
	OR			Bautyp = 2
	NOT			Bautyp = 3
	Lampe		Bautyp = 4
//	Leitung		Bautyp = 5
)
