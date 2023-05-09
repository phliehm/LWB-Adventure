package graphen

// Vor.: keine
// Erg.: Ein frisch initialisierter, leerer Graph ist geliefert. War der Parameter
//       -gerichtet- true, so ist er gerichtet, andernfalls ungerichtet.
// New (gerichtet bool) *data // *data erfüllt das Interface Graph

type Graph interface {

// Vor.: keine
// Erg.: True ist geliefert, gdw. der Graph gerichtet ist.
Gerichtet () bool

// Vor.: keine
// Erg.: True ist geliefert, gdw. der Graph keine Knoten enthält.
Leer () bool

// Vor.: keine
// Erg.: Die aktuelle Anzahl der Knoten im Graphen ist geliefert.
Knotenanzahl () uint32

// Vor.: keine
// Erg.: Wenn der Graph gerichtet ist, ist die aktuelle Anzahl der gerichteten
//       Kanten im Graphen geliefert. Andernfalls ist die Anzahl der ungerichteten
//       Kanten geliefert.
Kantenanzahl () uint32

// Vor.: Es gibt noch keinen Knoten mit der angegebenen ID id.
//       Es gilt: id < 100.
// Eff.: Der neuer Knoten hat die ID id und die Koordinaten (x,y) und
//       eine (vom Anwender verwendbare) Knoteninformation und ist in den Graphen eingefügt. 
KnotenEinfuegen (id uint32, x,y uint16, knoteninfo uint32)

// Vor.: index1 und index2 sind gültige ID-Nummern von existierenden Knoten im Graphen.
//       Es gibt (noch) keine Kante vom Knoten mit dem Index index1 zum Knoten mit dem Index index2.
// Eff.: Vom Knoten mit dem Index index1 gibt es nun eine Kante zum Knoten
//       mit dem Index index2. Dabei wird kanteninfo( z.B. eine Länge oder
//       eine Zeitangabe) mit der Kante verknüpft.
KanteEinfuegen (index1, index2 uint32, kanteninfo uint32)

// Vor.: Es gibt einen Knoten mit der angegebenen ID id.
// Erg.: Die Hintergrundfarbe des Knotens ist mit drei Farbwerten 
//       rot-green-blue geliefert. 
Knotenfarbe (id uint32) (r,g,b uint8)

// Vor.: Es gibt einen Knoten mit der angegebenen ID id.
// Erg.: Die Koordinaten des Knotens sind geliefert.
KnotenKoordinaten (id uint32) (x,y uint16)

// Vor.: Es gibt einen Knoten mit der angegebenen ID id.
// Erg.: Die Knoteninformation des Knoten ist geliefert.
Knoteninfo (id uint32) uint32

// Vor.: vonID und nachID sind gültige ID-Nummern von existierenden Knoten einer existierenden Kante.
// Erg.: Die Kantenfarbe der Kante vom Knoten vonID zum Knoten nachID
//       ist mit drei Farbwerten rot-green-blue geliefert. 
Kantenfarbe (vonID, nachID uint32) (r,g,b uint8)

// Vor.: Es gibt im Graphen eine Kante vom Knoten mit der ID vonID zum
//       Knoten mit der ID nachID.
// Erg.: Die Kanteninformation der Kante, die vom Knoten mit der ID vonID
//       zum Knoten mit der ID nachID geht, ist geliefert, falls
//       es eine solche Kante gibt.
Kanteninfo (vonID,nachID uint32) uint32

// Vor.: Es gibt einen Knoten mit der angegebenen Id id.
// Eff.: Die Eigenschaft 'Hintergrundfarbe' des Knotens in der graphischen
//       Darstellung ist auf den angegebenen Wert geändert.
KnotenFaerben (id uint32, r,g,b uint8)

// Vor.: vonID und nachID sind gültige ID-Nummern von existierenden Knoten im Graphen.
// Erg.: Die Kantenfarbe der Kante vom Knoten vonID zum Knoten nachID
//       ist mit den drei Farbwerten r,g,b im RGB-Modell gesetzt, falls es
//       diese Kante gab. 
KanteFaerben (vonID, nachID uint32, r,g,b uint8)

// Vor.: Es gibt einen Knoten mit der angegebenen Id id.
// Eff.: Die Knoteninformation des Knotens ist auf den neuen 
//       Wert info geändert.
KnoteninfoSetzen (id uint32, info uint32)

// Vor.: vonID und nachID sind gültige ID-Nummern von existierenden Knoten im Graphen.
// Erg.: Die Kanteninformation der Kante vom Knoten vonID zum Knoten nachID
//       ist auf denneuen Wert info geändert, falls es
//       diese Kante gab. 
KanteninfoSetzen (vonID, nachID uint32, info uint32)

// Vor.: keine
// Erg.: Eine 'Liste' aller Knoten-IDs ist geliefert.
KnotenID_Liste () []uint32

// Vor.: keine
// Erg.: Eine 'Liste' aller Kantenverbindungen ist geliefert. Jedes Element
//       ist ein Feld mit 2 Komponenten: die ID des Startknotens und
//       die ID des Zielknotens. 
Kantenliste () [][2]uint32

// Vor.: keine
// Erg.: True ist geliefert, gdw. vom Knoten mit dem Index index1 eine
//       Kante zum Knoten mit dem Index index2 existiert. 
Benachbart (index1, index2 uint32) bool

// Vor.: keine
// Erg.: True ist geliefert, gdw. ein Knoten mit der angegebenen ID id
//       im Graphen enthalten war.
Enthalten (id uint32) bool

// Vor.: keine
// Eff.: Der Knoten mit der k ist nicht mehr im Graphen enthalten. Ebenso sind
//       alle Kanten entfernt,die zu diesem Knoten führen oder von diesem
//       Knoten kamen.
KnotenEntfernen (id uint32) 

// Vor.: keine
// Eff.: Die Kante, die vom Knoten mit dem Index index1 zum Knoten mit dem 
//       Index2 führt, ist entfernt. 
KanteEntfernen (index1, index2 uint32) 

// Vor.: keine
// Eff.: In einem Grafikfenster (640x480 Pixel) ist der Graph dargestellt. War das Fenster
//       schon vor dem Aufruf offen, so ist sein alter Inhalt verloren.
Darstellen ()
}
