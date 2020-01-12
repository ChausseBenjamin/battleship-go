package main

// aesthetics.go contains all the functions which
// display or setup visuals without the use of tview.

import (
	"strconv"
)

// printPrimary displays an ASCII version of the primary battleship board
// The primary board is the one which belongs to the person calling a hit
// He sees his own boats on this board.
// Here is an example
//   A B C D E F G H I J
// 0 ~ ~ ~ ~ ~ ~ ~ ~ ~ △
// 1 ~ ~ ~ ~ ~ ~ ~ ~ ~ ▯
// 2 ~ ~ ~ ~ ~ ~ ~ ~ ~ ▽
// 3 ~ ~ ~ ◁ ▭ ▭ ▭ ▷ ~ ~
// 4 ~ ~ △ ~ ~ ~ ~ △ ~ ~
// 5 ~ ~ ▯ ~ ~ ~ ~ ▯ ~ ~
// 6 ~ ~ ▯ ~ ~ ~ ~ ~ ~ ~
// 7 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 8 ~ ~ ~ ~ ◀ ▬ ▬ ▷ ~ ~
// 9 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// Only E8, F8, and G8 were hit.
func (plyr player) DisplayPrimary() string {
	text := "\n  A B C D E F G H I J \n"
	for i := 0; i < 10; i++ {
		text += strconv.Itoa(i)
		text += " "
		for j := 0; j < 10; j++ {
			switch plyr.primary[i][j][2] {
			case 0: // That coordinate was not hit
				text += boatchars[1][plyr.primary[i][j][1]]
			case 1: // That coordinates was hit
				text += boatchars[0][plyr.primary[i][j][1]]
				// default:
				// 	return errors.New("Unknown State (hit/unhit) at a given coordinate")
			}
			text += " "
		}
		text += "\n"
	}
	// fmt.Println(text)
	return text
}

func (plyr *player) DisplayTarget() string {
	// First row labelling the columns
	text := "\n  A B C D E F G H I J \n"
	// For every row (r)
	for r := 0; r < 10; r++ {
		// Add the row number at the start of each line
		text += strconv.Itoa(r)
		// For every column (c)
		for c := 0; c < 10; c++ {
			// Separate each character with a space
			text += " "
			// First thing: is the coordinate hit or not?
			switch plyr.target[r][c][0] {
			// The coordinate is NOT hit:
			case 0:
				// Add the `~` symbol
				text += boatchars[1][0]
			// If the coordinate is NOT hit:
			default:
				// Is the coordinate water?
				switch plyr.prey.primary[r][c][0] {
				// It IS water:
				case 6:
					// Add the `◌` symbol
					text += boatchars[0][0]
				// It's NOT water:
				default:
					// Is the ID of the boat at that coordinate marked as a gain?
					switch plyr.gains[plyr.prey.primary[r][c][0]] {
					// It IS marked as a gain:
					case true:
						// Show the boat as it's meant to be:
						// (hit boatchars: with `position` marked in the opponents primary)
						text += boatchars[0][plyr.prey.primary[r][c][1]]
					default:
						// Nope, you're getting a censored tile: `▣`
						text += misteryHit
					}
				}
			}
		}
		text += "\n"
	}
	return text
}

// TODO: Function which returns what was hit as a commentary for the hitter

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
var boatchars = [2][7]string{
	{`◌`, `▲`, `▼`, `◀`, `▶`, `▮`, `▬`},
	{`~`, `△`, `▽`, `◁`, `▷`, `▯`, `▭`},
}

// This constant keeps information about boats that aren't totally sunk secret.
// It therefore substitues the shape of a boat on the target board when it is unsunk.
const misteryHit = `▣`
