/* Philipp Liehm
 * Juni 2023
 * LWB-Adventure: BugAttack
 */

package main


import (
		"./MiniGames/4_BugAttack"
		"fmt"
		//"./MiniGames/test"
		)

func main() {
		endN,endP := bugAttack.BugAttack()
		//fmt.Println("Gewonnen!!!")
		fmt.Println(endN,endP)
		//test.TesteEtwas()
}                

