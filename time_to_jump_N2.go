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

	deployChuteAtElevation := 2000

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
	deployed := false
/*
So we had 400 tacts. In one tact we have +1 or +4 on variable locationNorth and locationWest. 1*x+4*y = 1000 and x+y=400
x = 200 y = 200
After 200 tacts we should deployed. Thats all - deployChuteAtElevation := 200
*/
	for elevation > 0 {
		if elevation == deployChuteAtElevation {
			deployed = true
		}
		if !deployed {
//Because we have a negative number in windNorth (-4) and apply decrement operation(-=) in result we had a positive number in locationNorth
			locationNorth -= windNorth / 4  
			locationWest -= windWest / 4
		} else {
			locationNorth -= windNorth
			locationWest -= windWest
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




//Result deployChuteAtElevation := 2000