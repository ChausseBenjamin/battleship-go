package main

import (
	"fmt"
	tc "github.com/gdamore/tcell"
	tv "github.com/rivo/tview"
	"strconv"
	"strings"
)

func RedrawTarget(plyr *player, table *tv.Table) {
	// generating slice string for the table:
	// We initialize a slice containing all the cells
	// The first row will be the label of the columns
	boardData := strings.Split("  /A/B/C/D/E/F/G/H/I/J", "/")
	// For every row (r)
	for r := 0; r < 10; r++ {
		// Each row starts with the row label/number
		// A space makes the table centered by indenting it...
		str := " " + strconv.Itoa(r)
		boardData = append(boardData, str)
		// For every column (c)
		for c := 0; c < 10; c++ {
			// First thing: is the coordinate hit or not?
			switch plyr.target[r][c][0] {
			// The coordinate is NOT hit:
			case 0:
				// Add the `~` symbol
				boardData = append(boardData, boatchars[1][0])
			// If the coordinate is NOT hit:
			default:
				// Is the coordinate water?
				switch plyr.prey.primary[r][c][0] {
				// It IS water:
				case 6:
					// Add the `◌` symbol
					boardData = append(boardData, boatchars[0][0])
				// It's NOT water:
				default:
					// Is the ID of the boat at that coordinate marked as a gain?
					switch plyr.gains[plyr.prey.primary[r][c][0]] {
					// It IS marked as a gain:
					case true:
						// Show the boat as it's meant to be:
						// (hit boatchars: with `position` marked in the opponents primary)
						boardData = append(boardData, boatchars[0][plyr.prey.primary[r][c][1]])
					default:
						// Nope, you're getting a censored tile: `▣`
						boardData = append(boardData, misteryHit)
					}
				}
			}
		}
	}
	table.Clear()
	for r := 0; r < 11; r++ {
		for c := 0; c < 11; c++ {
			color, selectable := tc.ColorDarkCyan, true
			if r < 1 || c < 1 {
				color, selectable = tc.ColorPurple, false
			}
			if boardData[r*11+c] != `~` && !(r < 1 || c < 1) {
				selectable, color = false, tc.ColorRed
			}
			table.SetCell(r, c,
				tv.NewTableCell(boardData[r*11+c]).
					SetTextColor(color).
					SetAlign(tv.AlignCenter).
					SetSelectable(selectable))
		}
	}
}

func RedrawPrimary(plyr *player, table *tv.Table) {
	// generating slice string for the table:
	// We initialize a slice containing all the cells
	// The first row will be the label of the columns
	boardData := strings.Split("  /A/B/C/D/E/F/G/H/I/J", "/")
	// For every row (r)
	for r := 0; r < 10; r++ {
		// Each row starts with the row label/number
		// A space makes the table centered by indenting it...
		str := " " + strconv.Itoa(r)
		boardData = append(boardData, str)
		// For every column (c)
		for c := 0; c < 10; c++ {
			// Is the coordinate hit or not?
			switch plyr.primary[r][c][2] {
			// It is NOT
			case 0:
				// boatchars selects is character from the not-hit slice
				boardData = append(boardData, boatchars[1][plyr.primary[r][c][1]])
			// It IS
			case 1:
				// boatchars selects is character from the hit slice
				boardData = append(boardData, boatchars[0][plyr.primary[r][c][1]])
			}
		}
	}
	table.Clear()
	for r := 0; r < 11; r++ {
		for c := 0; c < 11; c++ {
			color := tc.ColorDarkCyan
			if r < 1 || c < 1 {
				color = tc.ColorPurple
			}
			if boardData[r*11+c] != `~` && !(r < 1 || c < 1) {
				color = tc.ColorRed
			}
			table.SetCell(r, c,
				tv.NewTableCell(boardData[r*11+c]).
					SetTextColor(color).
					SetSelectable(false).
					SetAlign(tv.AlignCenter))
		}
	}
}

func RedrawGains(plyr *player, gains *tv.TextView) {
	gains.Clear()
	boatlist := []string{
		boat0,
		boat1,
		boat2,
		boat3,
		boat4,
	}
	for i := 0; i < 5; i++ {
		if plyr.gains[i] {
			fmt.Fprintf(gains, "%v\n", boatlist[i])
		} else {
			fmt.Fprintf(gains, "\n\n")
		}
	}
}

func RedrawLosses(plyr *player, losses *tv.TextView) {
	losses.Clear()
	boatlist := []string{
		boat0,
		boat1,
		boat2,
		boat3,
		boat4,
	}
	for i := 0; i < 5; i++ {
		if plyr.prey.gains[i] {
			fmt.Fprintf(losses, "%v\n", boatlist[i])
		} else {
			fmt.Fprintf(losses, "\n\n")
		}
	}
}
