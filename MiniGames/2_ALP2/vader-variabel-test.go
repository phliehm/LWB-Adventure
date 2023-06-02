package main

import (
	"gfx"
	"robi"
)

const b uint16 = 75			//Feldbreite
var sp uint16 = 2
var ze uint16 = 2
var xm uint16 = sp*b+b/2
var ym uint16 = ze*b+b/2

/*	r.welt.feld[r.ze][r.sp].zeichnen()
	xm:=uint16(r.sp)*r.welt.feldbreite+r.welt.feldbreite/2
	ym:=uint16(r.ze)*r.welt.feldbreite+r.welt.feldbreite/2
	b:= r.welt.feldbreite
*/

func vader (r uint8) {
	
	switch r {
		case 2:
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreis(xm,ym-b*7/50,b*15/50)
		gfx.Volldreieck(xm-b*22/50,ym+b*11/50,xm-b*1/5,ym-b*3/10,xm-b*1/5,ym+b*11/50)
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreissektor(xm+b/50,ym-b*7/50,b*15/50,270,90)
		gfx.Volldreieck(xm+b*22/50,ym+b*11/50,xm+b*1/5,ym-b*3/10,xm+b*1/5,ym+b*11/50)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm,ym-b/10,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
		gfx.Volldreieck(xm,ym+b*21/50,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm-(b/50+1),ym+b*1/10,b/50+1,b*8/50)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*4/50,ym+b*7/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm-b*7/50,ym+b*9/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm,ym+b*1/10,b/50+1,b*8/50)									//+1 wegen Division ohne Rest
		gfx.Vollrechteck(xm+b*2/50,ym+b*7/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm+b*5/50,ym+b*9/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm+b*7/50,ym+b*23/50,xm+b*11/50,ym+b*23/50,xm+b*22/50,ym-b*23/50)	
		gfx.Volldreieck(xm+b*7/50,ym+b*23/50,xm+b*20/50,ym-b*23/50,xm+b*22/50,ym-b*23/50)
		case 1:
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreis(xm-b*7/50,ym,b*15/50)
		gfx.Volldreieck(xm+b*11/50,ym-b*22/50,xm-b*3/10,ym-b*1/5,xm+b*11/50,ym-b*1/5)
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreissektor(xm-b*7/50,ym+b/50,b*15/50,180,0)
		gfx.Volldreieck(xm+b*11/50,ym+b*22/50,xm-b*3/10,ym+b*1/5,xm+b*11/50,ym+b*1/5)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm-b/10,ym,xm+b*11/50,ym-b*22/50,xm+b*11/50,ym+b*22/50)
		gfx.Volldreieck(xm+b*21/50,ym,xm+b*11/50,ym-b*22/50,xm+b*11/50,ym+b*22/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm+b*1/10,ym-(b/50+1),b*8/50,b/50+1)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*7/50,ym-b*4/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm+b*9/50,ym-b*7/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm+b*1/10,ym,b*8/50,b/50+1)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*7/50,ym+b*2/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm+b*9/50,ym+b*5/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm+b*23/50,ym-b*7/50,xm+b*23/50,ym-b*11/50,xm-b*23/50,ym-b*22/50)	
		gfx.Volldreieck(xm+b*23/50,ym-b*7/50,xm-b*23/50,ym-b*20/50,xm-b*23/50,ym-b*22/50)
		case 0:
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreis(xm,ym+b*7/50,b*15/50)
		gfx.Volldreieck(xm-b*22/50,ym-b*11/50,xm-b*1/5,ym+b*3/10,xm-b*1/5,ym-b*11/50)
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreissektor(xm+b/50,ym+b*7/50,b*15/50,270,90)
		gfx.Volldreieck(xm+b*22/50,ym-b*11/50,xm+b*1/5,ym+b*3/10,xm+b*1/5,ym-b*11/50)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm,ym+b/10,xm-b*22/50,ym-b*11/50,xm+b*22/50,ym-b*11/50)
		gfx.Volldreieck(xm,ym-b*21/50,xm-b*22/50,ym-b*11/50,xm+b*22/50,ym-b*11/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm-(b/50+1),ym-b*13/50,b/50+1,b*8/50)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*4/50,ym-b*13/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm-b*7/50,ym-b*13/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm,ym-b*13/50,b/50+1,b*8/50)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm+b*2/50,ym-b*13/50,b*2/50,b*6/50)
		gfx.Vollrechteck(xm+b*5/50,ym-b*13/50,b*2/50,b*4/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm-b*7/50,ym-b*23/50,xm-b*11/50,ym-b*23/50,xm-b*22/50,ym+b*23/50)	
		gfx.Volldreieck(xm-b*7/50,ym-b*23/50,xm-b*20/50,ym+b*23/50,xm-b*22/50,ym+b*23/50)
		case 3:
		gfx.Stiftfarbe(50,50,50)				//dunkeldunkelgrau
		gfx.Vollkreis(xm+b*7/50,ym,b*15/50)
		gfx.Volldreieck(xm-b*11/50,ym-b*22/50,xm+b*3/10,ym-b*1/5,xm-b*11/50,ym-b*1/5)
		gfx.Stiftfarbe(100,100,100)				//dunkelhellgrau
		gfx.Vollkreissektor(xm+b*7/50,ym+b/50,b*15/50,180,0)
		gfx.Volldreieck(xm-b*11/50,ym+b*22/50,xm+b*3/10,ym+b*1/5,xm-b*11/50,ym+b*1/5)
		gfx.Stiftfarbe(0,0,0)					//schwarz
		gfx.Volldreieck(xm+b/10,ym,xm-b*11/50,ym-b*22/50,xm-b*11/50,ym+b*22/50)
		gfx.Volldreieck(xm-b*21/50,ym,xm-b*11/50,ym-b*22/50,xm-b*11/50,ym+b*22/50)
		gfx.Stiftfarbe(80,80,80)				//helldunkelgrau
		gfx.Vollrechteck(xm-b*13/50,ym-(b/50+1),b*8/50,b/50+1)							//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*13/50,ym-b*4/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm-b*13/50,ym-b*7/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(130,130,130)				//hellhellgrau
		gfx.Vollrechteck(xm-b*13/50,ym,b*8/50,b/50+1)									//+1 wegen Division ohne Rest!
		gfx.Vollrechteck(xm-b*13/50,ym+b*2/50,b*6/50,b*2/50)
		gfx.Vollrechteck(xm-b*13/50,ym+b*5/50,b*4/50,b*2/50)
		gfx.Stiftfarbe(255,0,0)					//rot
		gfx.Volldreieck(xm-b*23/50,ym+b*7/50,xm-b*23/50,ym+b*11/50,xm+b*23/50,ym+b*22/50)	
		gfx.Volldreieck(xm-b*23/50,ym+b*7/50,xm+b*23/50,ym+b*20/50,xm+b*23/50,ym+b*22/50)
	}
	
	/*
	//Vader variabel
	Stiftfarbe(50,50,50)
	//Vollkreis(325,318,15)
	Vollkreis(xm,ym-b*7/50,b*15/50)
	//Volldreieck(303,336,315,310,315,336)
	Volldreieck(xm-b*22/50,ym+b*11/50,xm-b*1/5,ym-b*3/10,xm-b*1/5,ym+b*11/50)
	
	Stiftfarbe(100,100,100)
	//Vollkreissektor(326,318,15,270,90)
	Vollkreissektor(xm+b/50,ym-b*7/50,b*15/50,270,90)
	//Volldreieck(347,336,335,310,335,336)
	Volldreieck(xm+b*22/50,ym+b*11/50,xm+b*1/5,ym-b*3/10,xm+b*1/5,ym+b*11/50)
	
	Stiftfarbe(0,0,0)
	//Volldreieck(325,320,303,336,347,336)
	Volldreieck(xm,ym-b/10,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
	//Volldreieck(325,346,303,336,347,336)
	Volldreieck(xm,ym+b*21/50,xm-b*22/50,ym+b*11/50,xm+b*22/50,ym+b*11/50)
	
	Stiftfarbe(80,80,80)
	//Vollrechteck(324,330,1,8)
	Vollrechteck(xm-b/50,ym+b*1/10,b/50,b*8/50)
	//Vollrechteck(321,332,2,6)
	Vollrechteck(xm-b*4/50,ym+b*7/50,b*2/50,b*6/50)
	//Vollrechteck(318,334,2,4)
	Vollrechteck(xm-b*7/50,ym+b*9/50,b*2/50,b*4/50)
	
	Stiftfarbe(130,130,130)
	//Vollrechteck(325,330,1,8)
	Vollrechteck(xm,ym+b*1/10,b/50,b*8/50)
	//Vollrechteck(327,332,2,6)	
	Vollrechteck(xm+b*2/50,ym+b*7/50,b*2/50,b*6/50)
	//Vollrechteck(330,334,2,4)
	Vollrechteck(xm+b*5/50,ym+b*9/50,b*2/50,b*4/50)
	
	Stiftfarbe(255,0,0)
	Volldreieck(xm+b*1/5,ym+b*21/50,xm+b*14/50,ym+b*21/50,xm+b*31/50,ym-b*2/5)	
	Volldreieck(xm+b*1/5,ym+b*21/50,xm+b*29/50,ym-b*2/5,xm+b*31/50,ym-b*2/5)
	
	//Linie(335,346,354,305)
	Linie(xm+b*1/5,ym+b*21/50,xm+b*29/50,ym-b*2/5)
	//Linie(336,346,355,305)
	Linie(xm+b*11/50,ym+b*21/50,xm+b*3/5,ym-b*2/5)
	//Linie(337,346,356,305)
	Linie(xm+b*12/50,ym+b*21/50,xm+b*31/50,ym-b*2/5)
	//Linie(338,346,356,305)
	Linie(xm+b*13/50,ym+b*21/50,xm+b*31/50,ym-b*2/5)
	//Linie(339,346,356,305)
	Linie(xm+b*14/50,ym+b*21/50,xm+b*31/50,ym-b*2/5)
	*/
}


func main () {
	robi.NeueWelt(800,600,b)
	
	//Variablen-Test
	//Stiftfarbe(0,255,0)
	//Vollkreis(xm,ym,feldbreite/2-5)
	
	vader(3)
	//vader(1)
	//vader(2)
	//vader(3)
	
	/*
	//Vader 1
	Stiftfarbe(50,50,50)
	Vollkreis(125,118,15)
	Volldreieck(103,136,115,110,115,136)
	//Vollrechteck(123,130,2,10)
	Stiftfarbe(100,100,100)
	Vollkreissektor(126,118,15,270,90)
	Volldreieck(147,136,135,110,135,136)
	//Vollrechteck(125,130,2,10)
	Stiftfarbe(0,0,0)
	Volldreieck(125,120,103,136,147,136)
	Volldreieck(125,146,103,136,147,136)
	Stiftfarbe(80,80,80)
	Vollrechteck(124,130,1,8)
	Vollrechteck(121,132,2,6)
	Vollrechteck(118,134,2,4)
	Stiftfarbe(130,130,130)
	Vollrechteck(125,130,1,8)
	Vollrechteck(127,132,2,6)
	Vollrechteck(130,134,2,4)
	Stiftfarbe(255,0,0)
	Linie(135,145,160,108)
	Linie(136,145,161,108)
	Linie(137,145,162,108)
	Linie(138,145,162,108)
	Linie(139,145,162,108)
	
	//Vader 2
	Stiftfarbe(50,50,50)
	Vollkreis(225,218,15)
	Volldreieck(203,236,215,210,215,236)
	//Vollrechteck(223,230,2,10)
	Stiftfarbe(100,100,100)
	Vollkreissektor(226,218,15,270,90)
	Volldreieck(247,236,235,210,235,236)
	//Vollrechteck(225,230,2,10)
	Stiftfarbe(0,0,0)
	Volldreieck(225,220,203,236,247,236)
	Volldreieck(225,246,203,236,247,236)
	Stiftfarbe(80,80,80)
	Vollrechteck(224,230,1,8)
	Vollrechteck(221,232,2,6)
	Vollrechteck(218,234,2,4)
	Stiftfarbe(130,130,130)
	Vollrechteck(225,230,1,8)
	Vollrechteck(227,232,2,6)
	Vollrechteck(230,234,2,4)
	Stiftfarbe(255,0,0)
	Linie(235,248,244,203)
	Linie(236,248,245,203)
	Linie(237,248,246,203)
	Linie(238,248,246,203)
	Linie(239,248,246,203)

	//Vader 2
	Stiftfarbe(50,50,50)
	Vollkreis(325,318,15)
	Volldreieck(303,336,315,310,315,336)
	//Vollrechteck(323,330,2,10)
	Stiftfarbe(100,100,100)
	Vollkreissektor(326,318,15,270,90)
	Volldreieck(347,336,335,310,335,336)
	//Vollrechteck(325,330,2,10)
	Stiftfarbe(0,0,0)
	Volldreieck(325,320,303,336,347,336)
	Volldreieck(325,346,303,336,347,336)
	Stiftfarbe(80,80,80)
	Vollrechteck(324,330,1,8)
	Vollrechteck(321,332,2,6)
	Vollrechteck(318,334,2,4)
	Stiftfarbe(130,130,130)
	Vollrechteck(325,330,1,8)
	Vollrechteck(327,332,2,6)
	Vollrechteck(330,334,2,4)
	Stiftfarbe(255,0,0)
	Linie(335,346,354,305)
	Linie(336,346,355,305)
	Linie(337,346,356,305)
	Linie(338,346,356,305)
	Linie(339,346,356,305)
*/

	gfx.TastaturLesen1()
}
