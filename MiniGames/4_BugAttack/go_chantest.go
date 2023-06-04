package main


import (
		"fmt"
		"time"
		)

var quit chan bool = make(chan bool)
		
func main() {

	go func() {
		for {
			select {
			case <- quit:
				return
			default:
				fmt.Println("YAY!")
				// Do other stuff
			}
			time.Sleep(5e8)
		}
	}()
	
	// Do stuff
	time.Sleep(2e9)
	// Quit goroutine
	quit <- true
	
}
