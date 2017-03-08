package main

import (
	"fmt"
)

func main(){

	elevation := 4000
	destinationNorth := 1000
	destinationWest := 1000

	windNorth := -4
	windWest := -4

	deployChuteAtElevation := 1000

	ladingLocation := simDrop(elevation,
							destinationNorth,
							destinationWest,
							windNorth,
							windWest,
							deployChuteAtElevation)

	fmt.Println("Initiate deploy at:",
				deployChuteAtElevation,
				"meters -", ladingLocation)
}

func simDrop(elevation int,
			dest_north int,
			dest_west int,
			windNorth int,
			windWest int,
			deployChuteAtElevation int) string {

	locationNorth := 0
	locationWest := 0
	locationNorth -=-1
	fmt.Println(locationNorth,locationWest)
	deployed := false
	for elevation > 0 {
		if elevation == deployChuteAtElevation {
			deployed = true
		}
		if !deployed {
			locationNorth -= windNorth / 4
			locationWest -= windWest / 4
		} else {
			//fmt.Println("Before wind: North, West -",locationNorth, locationWest)
			locationNorth -= windNorth
			locationWest -= windWest
			//fmt.Println("After wind: North, West -",locationNorth, locationWest)
		}
		elevation -= 10
	}
	fmt.Println("North:", locationNorth, 
				"West:", locationWest)
	ladingLocation := "You missed the target!"
	if locationNorth == dest_north && locationWest == dest_west {
		ladingLocation = "You hit the target!"
	}

	return ladingLocation
}