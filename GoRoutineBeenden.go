package main

// Beenden einer Go-Routine

import (
	"fmt"
	"time"
)

// Schreibt so lange "Stuff" bis die Go-Routine beendet wird
func schreibeStuff(c chan bool) {
	
	for {
		select {
			case <-c:
				return
			default:
				fmt.Println("Stuff")
				time.Sleep(5e8)
		}
	}
}

func main() {
	c := make(chan bool)
	go schreibeStuff(c)
	time.Sleep(2e9)
	c<-true
	go schreibeStuff(c)
	time.Sleep(1e9)
	c<-false		// true oder false macht hier keinen Unterschied, der Wert wird gar nicht verwendet. 
					// Es ist nur wichtig, dass eine Wertzuweisung geschehen ist

	
}
