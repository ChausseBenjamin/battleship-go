package main

import (
	"fmt"
)

func main() {

	settings := parameters{
		debug:    true,
		ssh_game: false,
	}

	// SETUP:
	var playerOne = player{name: "Ben"}
	var playerTwo = player{name: "Hugo"}
	// Setting up prey for when using Hit function
	playerOne.prey = &playerTwo
	playerTwo.prey = &playerOne

	if settings.debug {

		// PLACING PLAYER ONE BOATS:
		// Carrier: (ID=0), horizontal, (1,1)
		playerOne.initBoat(0, horizontal, 1, 1)
		// Battleship: (ID=1), horizontal, (0,9)
		playerOne.initBoat(1, horizontal, 0, 9)
		// Destroyer: (ID=2), vertical, (5,6)
		playerOne.initBoat(2, vertical, 5, 6)
		// Submarine: (ID=3), horizontal, (6,2)
		playerOne.initBoat(3, horizontal, 6, 2)
		// Patrol Boat: (ID=4), vertical, (1,5)
		playerOne.initBoat(4, vertical, 1, 5)

		// PLACING PLAYER TWO BOATS:
		// Carrier: (ID=0), vertical, (9,0)
		playerTwo.initBoat(0, vertical, 9, 0)
		// Battleship: (ID=1), horizontal, (1,8)
		playerTwo.initBoat(1, horizontal, 1, 8)
		// Destroyer: (ID=2), vertical, (5,3)
		playerTwo.initBoat(2, vertical, 5, 3)
		// Submarine: (ID=3), horizontal, (2,2)
		playerTwo.initBoat(3, horizontal, 2, 2)
		// Patrol Boat: (ID=4), vertical, (7,6)
		playerTwo.initBoat(4, vertical, 6, 6)

		// HITTING PLAYER ONE AT DIFFERENT COORDINATES
		playerTwo.Hit(5, 6) // (F,6)
		playerTwo.Hit(5, 7) // (F,7)
		playerTwo.Hit(5, 8) // (F,8)
		playerTwo.Hit(3, 4) // (D,4)
		playerTwo.Hit(3, 9) // (D,9)
		playerTwo.Hit(6, 6) // (G,6)
		playerTwo.Hit(2, 9) // (C,9)

		// HITTING PLAYER TWO AT DIFFERENT COORDINATES
		playerOne.Hit(1, 4) // (B,4)
		playerOne.Hit(1, 8) // (B,8)
		playerOne.Hit(2, 7) // (C,7)
		playerOne.Hit(2, 8) // (C,8)
		playerOne.Hit(3, 8) // (D,8)
		playerOne.Hit(4, 8) // (E,8)
		playerOne.Hit(9, 2) // (I,2)

		// Display both primary boards in stdout
		fmt.Println("Player One (Primary):", playerOne.DisplayPrimary())
		fmt.Println("Player Two (Primary):", playerTwo.DisplayPrimary())
	}
}
