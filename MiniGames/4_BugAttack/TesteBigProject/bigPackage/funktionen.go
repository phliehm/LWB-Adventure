package bigPackage


func Addiere2Zahlen(a,b uint32) uint32 {
	return a+b
}

func Multipliziere2Zahlen(a,b uint32) uint32 {
	return a*b
}

func GibSummeMitStruct(z ZweiZahlen) uint32 {
	return z.A+z.B
}
