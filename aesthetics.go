package main

import (
	"strconv"
)

// printPrimary displays using ASCII art the primary battleship board
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

// printPrimary displays using ASCII art the primary battleship board
func (plyr player) PrimarySlice() []string {
	board := []string{
		" ",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
	}
	for i := 0; i < 10; i++ {
		board = append(board, strconv.Itoa(i))
		for j := 0; j < 10; j++ {
			switch plyr.primary[i][j][2] {
			case 0: // That coordinate was not hit
				board = append(board, boatchars[1][plyr.primary[i][j][1]])
			case 1: // That coordinates was hit
				board = append(board, boatchars[0][plyr.primary[i][j][1]])
				// default:
				// 	return errors.New("Unknown State (hit/unhit) at a given coordinate")
			}
		}
	}
	// fmt.Println(text)
	return board
}

func (plyr player) TargetSlice() []string {
	board := []string{
		" ",
		"A",
		"B",
		"C",
		"D",
		"E",
		"F",
		"G",
		"H",
		"I",
		"J",
	}
	for i := 0; i < 10; i++ {
		board = append(board, strconv.Itoa(i))
		for j := 0; j < 10; j++ {
			switch plyr.target[i][j][0] {
			case 0:
				board = append(board, boatchars[1][0])
			case 1:
				if plyr.gains[plyr.prey.primary[i][j][0]] {
					board = append(board, boatchars[0][plyr.prey.primary[i][j][1]])
				} else {
					switch plyr.prey.primary[i][j][0] {
					case 0:
						board = append(board, boatchars[0][0])
					default:
						board = append(board, mistery_hit)
					}
				}
			}
		}
	}
	return board
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
