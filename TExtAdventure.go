package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var location string
var input string
var finishedGame bool

type Room struct {
	RoomName         string
	Visited          bool
	LongDescription  string
	ShortDescription string
	Directions       map[string]string
}

var courtyard = &Room{RoomName: "courtyard",
	LongDescription: `
	A once-graceful courtyard, 
	now overtaken by wild growth and gentle decay. 
	The cobbled ground, uneven and cracked, 
	is threaded with stubborn grass and creeping thyme, 
	their roots lifting stones that once formed a tidy path. 
	Broken remnants of a low stone wall encircle the space, 
	half-swallowed by brambles and tangled vines.
	At its center, a dry, moss-covered fountain sits like a forgotten relic, 
	its basin filled with rainwater and fallen leaves. 
	The statue that once adorned it, a weathered saint or perhaps a noble figure, 
	has lost its head to time, its features worn smooth by centuries.`,
	ShortDescription: "You are in the Courtyard again",

	Directions: map[string]string{"north": "lab", "east": "chapel"}}

var chapel = &Room{RoomName: "chapel",
	LongDescription: `
	A ruined medieval chapel stands in solemn silence, 
	its crumbling stone walls half-consumed by ivy and time. 
	The roof has long since collapsed, leaving the interior exposed 
	to the elements; open to sky, rain, and creeping moss. 
	Arched windows, once filled with stained glass, now gape empty, 
	their tracery delicate but fractured. The altar lies broken, 
	worn smooth by centuries of wind and weather, 
	while fragments of sacred carvings: angels, saints, 
	and worn Latin inscriptions, peek out from the rubble.`,
	ShortDescription: "You are in the Chapel again",
	Directions:       map[string]string{"west": "courtyard"},
}

/*
we want to support an arbitrary number of rooms.
with an arbitrary number of idetifiers to get to those rooms: direction or noun.
*/

var roomMap = map[string]*Room{
	"courtyard": courtyard,
	"chapel":    chapel,
}

func main() {

	//counts: a=1,b=2,c=5

	location = roomMap["courtyard"].RoomName
	for !finishedGame {

		LocationDescription()

		fmt.Print("What would you like to do?: ")

		// because fmt.scan stops at blank spaces
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.ToLower(input)

		// if strings.Contains(input, "look") {
		// 	input = strings.ReplaceAll(input, " ", "")
		// 	input = strings.ReplaceAll(input, "look", "")
		// 	LookDirection(input)
		// } else
		if strings.Contains(input, "go") {
			input = strings.ReplaceAll(input, " ", "")
			input = strings.ReplaceAll(input, "go", "")
			GoDirection(input)
		} else {
			fmt.Print("Please input 'Look' or 'Go' and a direction")
		}
		continue
	}

}

func LocationDescription() {

	room := roomMap[location]

	if !room.Visited {
		room.Visited = true
		fmt.Println(room.LongDescription)
	} else {
		fmt.Println(room.ShortDescription)
	}

}

func GoDirection(newinput string) {

	oldlocation := location

	for k, v := range roomMap[location].Directions {
		if strings.Contains(newinput, k) {
			location = v
			NewLocation(newinput, oldlocation)
		} else {
			fmt.Println("You don't know where that is!")
		}
	}

}

// func LookDirection(newinput string) {

// 	if strings.Contains(newinput, "north") {
// 		fmt.Println(roomMap[location].North)
// 	} else if strings.Contains(newinput, "south") {
// 		fmt.Println(roomMap[location].South)
// 	} else if strings.Contains(newinput, "east") {
// 		fmt.Println(roomMap[location].East)
// 	} else if strings.Contains(newinput, "west") {
// 		fmt.Println(roomMap[location].West)
// 	} else {
// 		if rand.Intn(100) != 99 {
// 			fmt.Println("You aren't sure which direction that is")
// 		} else {
// 			fmt.Println(`
// 	The world ignites in silence, bright and wide,
// 	a second sun that tears the morning thin.
// 	I cannot scream; the breath has left my side,
// 	and all I was dissolves before the din.

// 	The trees bow down, then vanish into light,
// 	the city folds like paper in a flame.
// 	No hero's end, no final, noble fight—
// 	just dust, and sky, and none to speak my name.

// 	You have died from an atomic bomb.

// 	The end`)
// 			os.Exit(0)

// 		}
// 	}

// }

func NewLocation(input string, newlocation string) {
	if location == newlocation {
		fmt.Print("There's nothing that way!")
	} else {
		fmt.Printf("You went %s to the ", input)
		fmt.Print(location)
		fmt.Println()
	}
}
