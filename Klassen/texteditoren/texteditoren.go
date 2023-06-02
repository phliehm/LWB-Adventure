//Annalena Cyriacus
//Juni 2023
//Klasse Texteditoren
//basierend auf dem Quelltext von Stefan Schmidt (LWB Informatik, ALP 3)

package texteditoren

type Texteditor interface {
	
	GibPosition() (x,y uint16) 
	
	SetzePosition(xneu,yneu uint16)
	
	//GibFont() (font string)
	
	//SetzeFont(fontneu string)
	
	GibHöhe() (höhe uint16)
	
	SetzeHöhe(höheneu uint16)
	
	GibBreite() (breite uint16)
	
	SetzeBreite(breiteneu uint16)
	
	GibSchriftgröße() (schriftgr int)
	
	SetzeSchriftgröße(schriftgrneu int)
	
	IstEinzeilig() bool
	
	NeuerTexteditor(posx,posy,breite,höhe uint16, schriftgr int, einzeilig bool)
	
	GibString() string
	
}

