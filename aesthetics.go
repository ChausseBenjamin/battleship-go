package main

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
func (plyr player) PrimaryDisplay() string {
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

// printTarget displays an ASCII version of the target battleship board
// The target board is the one which shows a player what he knows about
// his opponent. This is the board a player would play on.
// Here is an example
// A B C D E F G H I J
// 0 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 1 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 2 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 3 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 4 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 5 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 6 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 7 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// 8 ~ ~ ~ ~ ▣ ▣ ▣ ~ ~ ~
// 9 ~ ~ ~ ~ ~ ~ ~ ~ ~ ~
// Since the boat in F8, E8, G8 is not sunk, The player does not see
// it's shape. Upon sinking it, he will be able to see it.
func (plyr player) TargetDisplay() string {
	text := "\n  A B C D E F G H I J \n"
	for i := 0; i < 10; i++ {
		text += strconv.Itoa(i)
		text += " "
		for j := 0; j < 10; j++ {
			switch plyr.target[i][j][0] {
			case 0:
				text += boatchars[1][0]
			case 1:
				if plyr.gains[plyr.prey.primary[i][j][0]] {
					text += boatchars[0][plyr.prey.primary[i][j][1]]
				} else {
					switch plyr.prey.primary[i][j][0] {
					case 0:
						text += boatchars[0][0]
					default:
						text += mistery_hit
					}
				}
			}
			text += " "
		}
		text += "\n"
	}
	// fmt.Println(text)
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
	{`◌`, `▲`, `▼`, `◀`, `►`, `▮`, `▬`},
	{`~`, `△`, `▽`, `◁`, `▷`, `▯`, `▭`},
}

// This constant keeps information about boats that aren't totally sunk secret.
// It therefore substitues the shape of a boat on the target board when it is unsunk.
const mistery_hit = `▣`
