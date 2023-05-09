package texte
// Autor: B. Schneider
// Datum: 05.04.2023


var MoorEinl []string = []string{"    S E M I N A R R A U M",
								"",
								"   Zeige wieder Praesenz",
								" in der heutigen Vorlesung",
								"",
								" !!! Triff die Zielscheiben !!!",
								"",
								"  Je schneller du triffst, ",
								" desto mehr Punkte gibt's !"}

var MoorLvl2 []string = []string{"! ! ! G U T   G E M A C H T ! ! !",
								"",
								"Die Konzentration laesst",
								"aber mehr und mehr nach .",
								"",
								"CccCccchhhHhhrRrrRrrr ...",
								"",
								"Irgendwie scheinen die",
								"Ziele nun noch kleiner . . ."}

var MoorLvl3 []string = []string{". . . uiii, wie SCHWIERIG !",
								"",
								"Du brauchst was im Magen :",
								"     KAFFEE  und  PIZZA",
								"",
								"  !!! Sammle die Nahrung !!!",
								"             aber nur so :",
								"Rechts-Klick auf KAFFEE , ",
								"   Links-Klick auf PIZZA !"}
								
var MoorLvl4 []string = []string{" Das   L W B - P R I N Z I P",
								"",
								"Wie das mit der Nahrung geht,",
								"   wusstest du ja schon .",
								" Aber wirst du auch satt ...",
								"   ... in 30 Minuten Pause ?",
								"         ",
								"---  MASSE statt KLASSE  ---",
								" Horte, was das Zeug haelt !"}
								
var MoorLvl51 []string = []string{"Haus- und Raumordnung des ",
								"Studienzentrums fuer Erzie-",
								"hung, Paedagogik und Schule",
								"StEPS - 10.02.2023 - Rev 10 : ",
								". . .",
								"13. Das Verzehren von Spei-",
								"sen und Getraenken ist in al-",
								"len Seminar- und Vorlesungs-",
								"raeumen nicht gestattet .",
								". . . "}
								
var MoorLvl52 []string = []string{"   H  u  h  u ,    H  E  I  D  I ",
								"",
								"Nahrung und StEPS passt",
								"nicht .        Deshalb :",
								"Mach es wie Heidi und iss",
								"draussen in der Natur !",
								"",
								"Erscheint Heidi, druecke H",
								"und Nahrung verschwindet ."}
								
var MoorScore []string = []string{"  H o l a a d i o o o o o o o o o",
								"",
								"Gruene Wiesen im Sonnenschein.",
								"Brauchst du zum Gluecklichsein.",
								"",
								"",
								"  Du hast dich - gemessen an",
								"   deinen derzeitigen Kompe-",
								"    tenzen - ideal geschlagen !"}	
								
var MoorOut1 []string = []string{"Als du kamst, ",
								"         warst du LWB,",
								"                 jetzt bist du ",
								"",
								"                               StEPSi  !  !  !",
								"",
								"       erreichte PUNKTE : ",
								"",
								"veranstaltungs-NOTE : "}											
								

var MoorOut2 []string = []string{"",
								"",
								"   Moege die Macht . . .",
								"           . . . immer mit dir sein !",
								"",
								"",
								" Druecke nicht b, sondern Q ",
								" zum Verlassen dieses Spiels !"}

var MusterEinl []string = []string{"   H e r z l i c h    W i l l k o m m e n",
								"                             zur",
								"     Funktionalen Programmierung",
								" ",
								" Heute beschäftigen wir uns zuerst",
								"          mit der MUSTER-ERKENNUNG",
								"",
								" und schließlich festigen wir das",
								"          Wissen per MUSTER-MEMORY ."}
			
											
var MusterV [6]string = [6]string{"(f:w)",			// (f:w)
								"('f':w)",			// ('f':w)
								"(f:\"w\")",		// (f:"w")
								"[f,w]",			// [f,w]
								"(f,w)",			// (f,w)
								"(f:a:w)"}			// (f:w:w)
									
	
var MusterJ [6][]string = [6][]string {{"[1]","[[1]]","[\"w\"]","\"w\""},						// (f:w)
									{"\"f\"","\"febweb\""},										// ('f':w)
									{"\"fw\"","['v','w']"},										// (f:"w")
									{"\"wf\"","[True,False]","[3,8]"},							// [f,w]
									{"('f',False)","(1,'w')","(\"LWB\",True)"},					// (f,w)
									{"[32,23]","\"LWB\"","[\"so\",\"ist\",\"das\"]","[1,2,3,4,5,6]"} }			// (f:a:w)

var MusterL [6][][2]string = [6][][2]string {{{"1","[]"},{"[1]","[]"},{"\"w\"","[]"},{"'w'","[]"}},			// (f:w)
									{{"","[]"},{"","\"ebweb\""}},											// ('f':w)
									{{"'f'",""},{"'v'",""}},												// (f:"w")
									{{"'w'","'f'"},{"True","False"},{"3","8"}},								// [f,w]
									{{"'f'","False"},{"1","'w'"},{"\"LWB\"","True"}},						// (f,w)
									{{"32","[]"},{"L","\"B\""},{"\"so\"","[\"das\"]"},{"1","[3,4,5,6]"}}}	// (f:a:w)
	
var MusterN [6][]string = [6][]string {{"[1,'a']","\"j\":\"a\"","(True,True)","35"},		// (f:w)
									{"'feuer'","[\"feder\"]"},								// ('f':w)
									{"[\"v\",\"w\"]","\"wo\""},								// (f:"w")
									{"\"wie\"","[[]]"},										// [f,w]
									{"\"er\"","1.44"},										// (f,w)
									{"\"V\"","(\"20\",\"23\")"} }							// (f:a:w)	
	
	
	
	
	
