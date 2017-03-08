package main

import (
	"fmt"
	"strings"
)

func main() {
	request := Request{agent: "Dee Fercloze", mission: "Operation Go"}

	request.agent = "Β β aΒβ"
	location := pingSattelite(request)
	fmt.Println("Location:", location)
}

type Request struct {
	agent string
	mission string
}

func pingSattelite(request Request) string {
	isAccessible := accessible(request.agent)
	return revealLocation(isAccessible)
}


//So we need next:x=true,y=true,z=true.
func accessible(agent string) bool {
	a := strings.Split(agent, " ")
	if len(a) != 3 { return false }
	b := a[0]
	c := a[1]
	d := a[2]
	// EqualFold reports whether b and c, interpreted as UTF-8 strings,
	// are equal under Unicode case-folding.
	x := strings.EqualFold(b,c) //There we had false
	y := b != strings.ToLower(c) // There we had true. Compare two strings.
	// Index returns the index of the first instance of b+c in d, 
	// or -1 if b+c is not present in d.
	//
	// What we need?
	// First: three words in agent name.
	// Second: Third word len=5
	// Third: We need 2 string (b and c), that are equal under Unicode case-folding
	// 			and differ in lowercase comparing.
	//Additionally:unicode symbol had len=2
	z := strings.Index(d, b+c) == 1 && len(d) == 5
	return x && y && z
}

func revealLocation(access bool) string {
	if access {
		return "Congratulations, your location is: x,y,z"
	} else {
		return "Failed to access. Good luck."
	}
}



//Result request.agent = "Β β aΒβ"