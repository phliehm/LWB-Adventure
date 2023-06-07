
//  ADT spielstand

//	zum Speichern vom Spielstand 

//  mit Punkte, Note und Datum (wann gespeichert wurde)

//	Martin Seiß		5.6.2023


package spielstaende


//import "fmt"
import . "dateien"
//import "bufio"
//import "os"

type Spielstand interface {

	// Vor: Der Pfad (path) existiert.
	// Eff: Lädt den Spielstand in den Speicher, wenn vorhanden,
	//		sonst wird eine neue Datei angelegt. 
	// Erg: Der Spielstand ist geliefert.
	// New(spielername,pfad string) bool 


	// Vor: -
	// Eff: Speichert die Noten und Punkte der Minispiele in die zum
	//		Spielstand gehörige Datei.
	// Erg: -
	Speichern(noten []float32, punkte []uint32)
	
	
	// Vor: -
	// Eff: -
	// Erg: True ist geliefert, wenn Spielstandsdatei vorhanden war, und
	// 		alter Spielstand geladen wurde.
	GibVorhanden() bool

	
	// Vor: -
	// Eff: -
	// Erg: Ein Float-Slice ist geliefert, der die Noten in den Spielen
	//		wiedergibt.
	GibNoten() []float32
	
	// Vor: -
	// Eff: -
	// Erg: Ein Float-Slice ist geleifert, der die Noten in den Spielen
	//		wiedergibt.
	GibPunkte() []uint32

	// Vor: -
	// Eff: - 
	// Erg: Das Datum mit Uhrzeit vom letzten Speichern ist geliefert.
	//Datum() string

}


type data struct {
	spielername 	string			// Name des Spielers für den Spielstand	
	pfad			string			// Pfad für den Speicherstand
	dateiname		string			// Dateiname mit Pfad
	vorhanden		bool			// war der Spielsstand schon vorhanden
									// d.h. war die Datei schon angelegt?
	punkte			[]uint32
	noten			[]float32
}


func New(spielername,pfad string) *data {
	var sp *data = new(data)
	sp.spielername = spielername
	sp.pfad = pfad
	sp.dateiname = sp.pfad + "/" + sp.spielername + ".dat"
	sp.noten = make([]float32,0)
	sp.punkte = make([]uint32,0)
	sp.ladeDaten()
	return sp
}


func (sp *data) Speichern(noten []float32, punkte []uint32) {
	
	var ausgabe []byte			// String der ausgegeben wird
	var b1, b2, b3, b4	byte	// Hilfsbytes zum speichern
	var d Datei
	var nByte = byte('N')
	
	sp.punkte = punkte
	sp.noten = noten
	
	ausgabe = make([]byte,0)
	// kodieren der Punkte
	for i:=0; i<len(sp.punkte); i++ {
		b1 = byte(sp.punkte[i] % 256)
		b2 = byte((sp.punkte[i]/256) % 256)
		b3 = byte((sp.punkte[i]/256/256) % 256)
		b4 = byte(sp.punkte[i]/256/256/256)
		//fmt.Println(b1,b2,b3,b4)
		ausgabe = append(ausgabe,b1)
		ausgabe = append(ausgabe,b2)
		ausgabe = append(ausgabe,b3)
		ausgabe = append(ausgabe,b4)		
	}
	ausgabe = append(ausgabe,nByte)
	
	// kodieren der Note
	for i:=0; i<len(sp.noten); i++ {
		b1 = byte(sp.noten[i]*10)
		//fmt.Println(b1)
		ausgabe = append(ausgabe,b1)
	}
	ausgabe = ausgabe //+ "\n"
	
	//fmt.Println("Ausgabe: \n",ausgabe)
	
	// Schreibe Daten in Datei.
	d = Oeffnen(sp.dateiname,'s')
	for i:=0; i<len(ausgabe); i++ {
		d.Schreiben(ausgabe[i])
	}
	d.Schliessen()
		
}



func (sp *data) GibVorhanden() bool {
	return sp.vorhanden
}

	
func (sp *data) GibNoten() []float32 {
	return sp.noten
}


func (sp *data)	GibPunkte() []uint32 {
	return sp.punkte
}


// Vor: -
// Eff: Die Noten und Punkte werden geladen. Vorhanden wird auf true
// 		gesetzt.
// Erg: -
func (sp *data)	ladeDaten() {
	
	//var byteliste []byte = make([]byte,0)
	var noten []float32 = make([]float32,0)
	var punkte []uint32 = make([]uint32,0)
	var wert uint32			// Hilfswert für Punktzahl
	var d Datei				// Datei-Objekt zum Spielstand
	var b1, b2, b3, b4	byte	// Hilfsbytes zum speichern

	//fmt.Println("Dateiname ",sp.dateiname)
	d = Oeffnen(sp.dateiname,'x')
	if d.Ende() {				// Dateiinhalt existierte noch nicht
		d.Schliessen()
	} else {
		for !d.Ende() {				// Lese die Bytes und umwandeln
			b1 = d.Lesen()
			//fmt.Println(b1,string(b1),"N")
			if string(b1) == "N" {break}
			b2 = d.Lesen()
			b3 = d.Lesen()
			b4 = d.Lesen()
			//fmt.Println(b1,b2,b3,b4)
			wert = uint32(b1) + uint32(b2)*256
			wert = wert + uint32(b3)*256*256 + uint32(b4)*256*256*256
			punkte = append(punkte,wert)
		}
		for !d.Ende() {
			b1 = d.Lesen()
			//fmt.Println(b1,float32(b1)/10)
			noten = append(noten,float32(b1)/10)
		}
		sp.vorhanden = true
		d.Schliessen()
		}
		
		sp.punkte = punkte
		sp.noten = noten

}
