package main

import (
	"fmt"
)

//const (
//	prizeDay = "Wednesday"
//	prizeFund = 10000
//)
var (
	prizeDay = "Wednesday"
	prizeFund = 10000
)

func main() {
	prizeDay = "Thursday"
	prizeFund = 50000
	msg := "%s's prize-winning entry is %d and wins %d !!!\n"
	winner := 9
	fmt.Printf(msg, prizeDay, winner, prizeFund)
}
