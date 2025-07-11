package main

import (
	"bufio"
	"fmt"
	"math/rand"
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
	North            string
	South            string
	East             string
	West             string
	NorthLocation    string
	SouthLocation    string
	EastLocation     string
	WestLocation     string
}

var courtyard = Room{RoomName: "courtyard",
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
	North:            "You see nothing to the North",
	East:             "You see a chapel to the East",
	West:             "you see a Vault to the West",
	South:            "You see a Lab to the South",
	NorthLocation:    "courtyard",
	SouthLocation:    "courtyard",
	EastLocation:     "chapel",
	WestLocation:     "courtyard"}

var chapel = Room{RoomName: "chapel",
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
	North:            "You see nothing to the North",
	East:             "You see nothing to the East",
	West:             "you see the courtyard to the West",
	South:            "You see nothing to the South",
	NorthLocation:    "chapel",
	SouthLocation:    "chapel",
	EastLocation:     "chapel",
	WestLocation:     "courtyard"}

var names = []string{"courtyard", "chapel"}

var rooms = []Room{courtyard, chapel}

var roomMap = make(map[string]Room)

func main() {

	for i, names := range names {

		if i < len(rooms) {
			roomMap[names] = rooms[i]
		} else {
			// error if the lengths don't match
			fmt.Printf("Harry")
			os.Exit(0)
		}
	}

	location = roomMap["courtyard"].RoomName
	for !finishedGame {

		LocationDescription()

		fmt.Print("What would you like to do?: ")

		// because fmt.scan stops at blank spaces
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = strings.ToLower(input)

		if strings.Contains(input, "look") {
			strings.ReplaceAll(input, " ", "")
			strings.ReplaceAll(input, "look", "")
			LookDirection(input)
		} else if strings.Contains(input, "go") {
			strings.ReplaceAll(input, " ", "")
			strings.ReplaceAll(input, "go", "")
			GoDirection(input)
		} else {
			fmt.Print("Please input 'Look' or 'Go' and a direction")
		}
		continue
	}

}

func LocationDescription() {
	for _, v := range rooms {
		if v.RoomName == location {
			if v.Visited == false {

				// roomMap[v.RoomName].Visited = true

				// They made it annoying to try to change the value here

				fmt.Print(v.LongDescription)

				fmt.Println()
			} else {
				fmt.Print(v.ShortDescription)
				fmt.Println()
			}
		}
	}
}

func GoDirection(newinput string) {

	newlocation := location

	if strings.Contains(newinput, "north") {
		location = roomMap[location].NorthLocation
		if location == newlocation {
			fmt.Print("There's nothing that way!")
		} else {
			fmt.Print("You went North to the ")
			fmt.Print(location)
			fmt.Println()
		}
	} else if strings.Contains(newinput, "south") {
		location = roomMap[location].SouthLocation
		if location == newlocation {
			fmt.Print("There's nothing that way!")
		} else {
			fmt.Print("You went South to the ")
			fmt.Print(location)
			fmt.Println()
		}
	} else if strings.Contains(newinput, "east") {
		location = roomMap[location].EastLocation
		if location == newlocation {
			fmt.Print("There's nothing that way!")
		} else {
			fmt.Print("You went East to the ")
			fmt.Print(location)
			fmt.Println()
		}
	} else if strings.Contains(newinput, "west") {
		location = roomMap[location].WestLocation
		if location == newlocation {
			fmt.Print("There's nothing that way!")
		} else {
			fmt.Print("You went West to the ")
			fmt.Print(location)
			fmt.Println()
		}
	} else {
		fmt.Println("You don't know where that is!")
	}

}

func LookDirection(newinput string) {

	if strings.Contains(newinput, "north") {
		fmt.Println(roomMap[location].North)
	} else if strings.Contains(newinput, "south") {
		fmt.Println(roomMap[location].South)
	} else if strings.Contains(newinput, "east") {
		fmt.Println(roomMap[location].East)
	} else if strings.Contains(newinput, "west") {
		fmt.Println(roomMap[location].West)
	} else {
		if rand.Intn(100) != 99 {
			fmt.Println("You aren't sure which direction that is")
		} else {
			fmt.Println(`
	The world ignites in silence, bright and wide,
	a second sun that tears the morning thin.
	I cannot scream; the breath has left my side,
	and all I was dissolves before the din.

	The trees bow down, then vanish into light,
	the city folds like paper in a flame.
	No hero's end, no final, noble fight—
	just dust, and sky, and none to speak my name.

	You have died from an atomic bomb.

	The end`)
			os.Exit(0)

		}
	}

}
