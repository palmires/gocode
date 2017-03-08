package main

import (
	"fmt"
	"strings"
)

func main() {
	request := Request{agent: "Dee Fercloze", mission: "Operation Go"}

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

func accessible(agent string) bool {
	a := strings.Split(agent, " ")
	if len(a) != 3 { return false }
	b := a[0]
	c := a[1]
	d := a[2]
	x := strings.EqualFold(b,c)
	y := b != strings.ToLower(c)
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