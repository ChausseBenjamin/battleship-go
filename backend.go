package main

import (
// "fmt"
)

type parameters struct {
	vsCPU   bool
	sshGame bool
	debug   bool
}

type userError struct {
	query string
	err   error
}

// boatlist catalogs boat structures for use with icons.
// It is meant to be used with a "boatchars" slice containing said icons (with sunk and unsunk boats).
// Here are the slice values meanings:
// 0: Water
// 1: Vertical north tip (Triangle)
// 2: Vertical south tip (Triangle)
// 3: Horizontal west tip (Triangle)
// 4: Horizontal east tip (Triangle)
// 5: Vertical body section (rectangle)
// 6: Horizontal body section (rectangle)
var boatlist = [5][2][5]int{
	// Horizontal    | Vertical
	// Carrier:
	{{3, 6, 6, 6, 4}, {1, 5, 5, 5, 2}},
	// Battleship
	{{3, 6, 6, 4, 0}, {1, 5, 5, 2, 0}},
	// Destroyer
	{{3, 6, 6, 0, 0}, {1, 5, 5, 0, 0}},
	// Submarine
	{{3, 6, 4, 0, 0}, {1, 5, 2, 0, 0}},
	// PatrolBoat
	{{3, 6, 0, 0, 0}, {1, 5, 0, 0, 0}},
}

/* Boats Info:
   |------------+--------+----------------|
   | BoatName   | BoatID | HorizontalBoat |
   | Carrier    | 0      | ◁ ▭ ▭ ▭ ▷      |
   | Battleship | 1      | ◁ ▭ ▭ ▷        |
   | Destroyer  | 2      | ▭ ▭ ▷          |
   | Submarine  | 3      | ◁ ▭ ▷          |
   | PatrolBoat | 4      | ▭ ▷            |
   |------------+--------+----------------|
*/

// player contains all the environment of a player.
// primary is the grid containing his own boats.
// target stores info the player knows about the ennemy.
// gains keeps track of the ships the player has managed to sink.
type player struct {
	// name of the player
	// 	It is used for scoreboards but also throughout
	// 	the game to announce turns.
	name string
	// primary is a 3 dimensional slice containing information
	// about the players' own boats.
	//	Structure:
	//		[y][x][boatID, position, hitStatus]int
	//		- [y][x]
	//			The last slice index defines his coordinates in the x,y axis.
	//		- BoatID:
	//			n+1: Water
	//			n: ID of the boat
	//		- Position:
	//			0: Water
	//			1: North arrow
	//			2: South arrow
	//			3: West  arrow
	//			4: East  arrow
	//			5: Vertical   middle
	//			6: Horizontal middle
	//		- HitStatus:
	//			0: Unhit
	//			1: Hit
	primary [10][10][3]int
	// target is a 3 dimensional slice containing the information
	// the player knows about his opponent.
	//	Structure:
	//		[ hit, id ]
	//		- hit:
	//			true: Was hit
	//			false: Has not been hit
	//		- id:
	//			0: water or unknown
	//			n: ID of the boat
	target [10][10][2]int
	// gains is a slice containing the list of boats the player
	// has destroyed.
	// 	Structure:
	//		- The index in the slice refers to the boatID
	//		- true if the boat is sunk
	gains [5]bool
	// prey is a pointer targetting the player's opponent.
	// 	It allows to access it's board to retrieve and edit it's
	//  information when the current player makes a move.
	prey *player
}

// InitBoard creates two blank 3 dimensional slice:
// The primary slice and the prey slice
func (plyr *player) InitBoard(opponent *player) {
	// Sets the players' opponent
	plyr.prey = opponent
	for r := 0; r < 10; r++ {
		for c := 0; c < 10; c++ {
			plyr.primary[r][c] = [3]int{6, 0, 0}
			plyr.target[r][c] = [2]int{0, 6}

		}
	}
}

// InitBoat places a boat on a players primary grid.
// boat sizes are defined by the boatlist variable.
// Consult it for more info.
func (plyr *player) InitBoat(boatID, orientation, x, y int) {
	boatLength := len(boatlist[boatID][0])
	if orientation == 0 { // Boat is HORIZONTAL
		for i := 0; i < boatLength; i++ {
			char := boatlist[boatID][0][i]
			if char == 0 { // Compensating for boatlist having
				break //      zeros at the end of slices
			}
			//TODO: Add error handling for stacking boats
			plyr.primary[y][x][0] = boatID
			plyr.primary[y][x][1] = char
			x++ // HORIZONTAL: Therefore the loop increments the x axis
		}
	} else { // Boat is VERTICAL
		for i := 0; i < boatLength; i++ {
			char := boatlist[boatID][1][i]
			if char == 0 {
				break
			}
			plyr.primary[y][x][0] = boatID
			plyr.primary[y][x][1] = char
			y++ // VERTICAL: Therefore the loop increments the y axis
		}
	}
}

func (plyr *player) Hit(x, y int) bool {
	// We change the coord status to hit
	// We add the id of the boat since we know it now.
	plyr.prey.primary[y][x][2] = 1
	plyr.target[y][x] = [2]int{1, plyr.prey.primary[y][x][0]}
	// If that coordinate contains a boat
	if BoatID := plyr.prey.primary[y][x][0]; BoatID < 5 {
		switch plyr.prey.primary[y][x][1] {
		case 1, 2, 5: // If the hit boat was vertical
			// We suppose the boat is sunk since it only takes one unhit coordinate to prove this wrong
			plyr.gains[plyr.prey.primary[y][x][0]] = true
			// We check the entire column iteratively
			for i := 0; i < 10; i++ {
				// If the boat ID of the coordinate on the board is the same as the boat which was hit
				// AND
				// If the boat was not hit at that coordinate
				if plyr.prey.primary[i][x][0] == plyr.prey.primary[y][x][0] && plyr.prey.primary[i][x][2] == 0 {
					// Then reset the boat to being unsunk
					plyr.gains[plyr.prey.primary[y][x][0]] = false
					break
				}
			}
		case 3, 4, 6: // If the hit boat was horizontal
			// We suppose the boat is sunk since it only takes one unhit coordinate to prove this wrong
			plyr.gains[plyr.prey.primary[y][x][0]] = true
			// We check the entire row iteratively
			for j := 0; j < 10; j++ {
				// If the boat ID of the coordinate on the board is the same as the boat which was hit
				// AND
				// If the boat was not hit at that coordinate
				if plyr.prey.primary[y][j][0] == plyr.prey.primary[y][x][0] && plyr.prey.primary[y][j][2] == 0 {
					// Then reset the boat to being unsunk
					plyr.gains[plyr.prey.primary[y][x][0]] = false
					break
				}
			}
		}
		return true // Returns true if there was a hit
	} else {
		return false // Returns false if water was hit
	}
}

const horizontal = 0
const vertical = 1
