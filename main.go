package main

import (
	"fmt"
	tc "github.com/gdamore/tcell"
	tv "github.com/rivo/tview"
	"strconv"
	"strings"
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
		fmt.Println("Player One:", playerOne.DisplayPrimary())
		fmt.Println("Player Two:", playerTwo.DisplayPrimary())

		/* TVIEW UI SETUP:
		┌------------------------------------------------------┐
		|dashboard                                             |
		|┌----------------------------------------------------┐|
		||headerBox                                           ||
		|└----------------------------------------------------┘|
		|┌----------------------------------------------------┐|
		||bottomFlex                                          ||
		||┌-------------------┐ ┌----------------------------┐||
		||| infoFlex          | | playFlex                   |||
		|||┌-----------------┐| |┌--------------------------┐|||
		|||| keybindingsBox  || || targetFlex               ||||
		||||                 || ||┌------------------------┐||||
		||||                 || ||| targetBox | gainsBox   |||||
		||||                 || |||           |            |||||
		||||                 || |||           |            |||||
		||||                 || ||└------------------------┘||||
		||||                 || |└--------------------------┘|||
		||||                 || |┌--------------------------┐|||
		|||└-----------------┘| || primaryFlex              ||||
		|||┌-----------------┐| ||┌------------------------┐||||
		|||| logBox          || ||| primaryBox | lossesBox |||||
		||||                 || |||            |           |||||
		||||                 || |||            |           |||||
		||||                 || ||└------------------------┘||||
		|||└-----------------┘| |└--------------------------┘|||
		||└-------------------┘ └----------------------------┘||
		|└----------------------------------------------------┘|
		└------------------------------------------------------┘

		DASHBOARD:
			flex structure containing everything.
		- Direction: rows

		BOTTOMFLEX:
			flex structure containing everything but the headerBox
		- Direction: columns

		INFOFLEX:
			flex structure containing general info related boxes:
				- keybindingsBox
				- logBox
		- Direction: rows

		PLAYFLEX:
			flex structure containing everything related to playing the game:
				- targetFlex
				- primaryFlex
			- Direction: rows

		TARGETFLEX:
			flex structure containing everything the player knows about his target:
				- targetBox
				- gainsBox
		- Direction: columns

		PRIMARYFLEX:
			flex structure containing everything the player knows about himself:
				- primaryBox
				- lossesBox
		- Direction: columns

		HEADERBOX:
			box which displays in it's title the current player.

		KEYBINDINGSBOX:
			simple box containing a list of all the keybindings one can use.

		LOGBOX:
			box which shows a log of the past moves each player made.

		TARGETBOX:
			box containing the board where the player attacks his opponent
			this box is focused by default and is of type table as it can be navigated.

		GAINSBOX:
			box showing the names of all the ennemies boats which are sunk.
			each sunk boat has a small display of the boat going with it.

		PRIMARYBOX:
			box containing the board showing the players boat layout.

		LOSSESBOX:
			box showing the names of all the players boats which are sunk.
			each sunk boat has a small display of the boat going with it.

		COMMANDBOX:
			box containing an input field which is to be used as a command prompt.
			coordinates can be inputed directly and typing quit will exit the game.

		*/

		// Initializing the application
		app := tv.NewApplication()

		headerBox := tv.NewBox().SetTitle(playerOne.name).
			SetBorder(true)

		keybindingsBox := tv.NewBox().
			SetTitle("Keybindings:").
			SetBorder(true)
		// TODO: Add text/documentation to box

		logBox := tv.NewBox().
			SetTitle("Log:").
			SetBorder(true)

		targetBox := tv.NewTable()
		RedrawTarget(&playerOne, targetBox)
		targetBox.SetFixed(1, 1).
			SetSelectable(true, true).
			SetDoneFunc(func(key tc.Key) {
				if key == tc.KeyEscape {
					// TODO: change this for a dopdown prompt "Are you sure? (Y/N)"
					app.Stop()
				}
			}).
			SetBorder(true).
			SetTitle("The enemy:")

		primaryBox := tv.NewTable().
			SetBorder(true).
			SetTitle("You:")

		gainsBox := tv.NewList().
			SetBorder(true).
			SetTitle("Gains:")

		lossesBox := tv.NewList().
			SetBorder(true).
			SetTitle("Losses:")

		commandBox := tv.NewInputField().
			SetBorder(true).
			SetTitle("Command:")

		targetFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(targetBox, 26, 0, true).
			AddItem(gainsBox, 26, 0, false)

		primaryFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(primaryBox, 26, 0, false).
			AddItem(lossesBox, 26, 0, false)

		playFlex := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(targetFlex, 13, 0, true).
			AddItem(primaryFlex, 13, 0, false).
			AddItem(commandBox, 0, 1, false)

		infoFlex := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(keybindingsBox, 0, 3, false).
			AddItem(logBox, 0, 1, false)

		bottomFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(infoFlex, 0, 1, false).
			AddItem(playFlex, 52, 0, true)

		dashboard := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(headerBox, 2, 1, false).
			AddItem(bottomFlex, 0, 1, true)

		if err := app.SetRoot(dashboard, true).Run(); err != nil {
			panic(err)
		}
	}
}

func RedrawTarget(plyr *player, table *tv.Table) {
	// generating slice string for the table:
	boardData := strings.Split("  /A/B/C/D/E/F/G/H/I/J", "/")
	for i := 0; i < 10; i++ {
		str := " " + strconv.Itoa(i)
		boardData = append(boardData, str)
		for j := 0; j < 10; j++ {
			switch plyr.target[i][j][0] {
			case 0: // Tile is unhit (water since we don't know)
				boardData = append(boardData, boatchars[1][0])
			case 1: // Tile has already been hit
				if plyr.gains[plyr.prey.primary[i][j][0]] { // Boat is sunk:
					// Boat ID at that coordinate coresponds to ID of a sunk boat
					boardData = append(boardData, boatchars[0][plyr.prey.primary[i][j][1]])
				} else {
					switch plyr.prey.primary[i][j][0] {
					case 6:
						boardData = append(boardData, boatchars[0][0])
					default:
						boardData = append(boardData, mistery_hit)
					}
				}
			}
		}
	}

	table.Clear()
	for r := 0; r < 11; r++ {
		for c := 0; c < 11; c++ {
			color, selectable := tc.ColorDarkBlue, true
			if r < 1 || c < 1 {
				color, selectable = tc.ColorPurple, false
			}
			if boardData[r*11+c] != `~` {
				selectable = false
			}
			table.SetCell(r, c,
				tv.NewTableCell(boardData[r*11+c]).
					SetTextColor(color).
					SetAlign(tv.AlignCenter).
					SetSelectable(selectable))
		}
	}
}
