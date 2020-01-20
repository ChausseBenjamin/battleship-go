package main

import (
	"fmt"
	tc "github.com/gdamore/tcell"
	tv "github.com/rivo/tview"
)

func main() {

	settings := parameters{
		debug:   true,
		sshGame: false,
	}

	// SETUP:
	var playerOne = player{name: "P1"}
	var playerTwo = player{name: "P2"}
	// Setting up prey for when using Hit function
	playerOne.InitBoard(&playerTwo)
	playerTwo.InitBoard(&playerOne)

	formApp := tv.NewApplication()

	form := tv.NewForm().
		AddInputField("Player 1 name:", playerOne.name, 30, nil, func(text string) {
			playerOne.name = text
		}).
		AddInputField("Player 2 name:", playerTwo.name, 30, nil, func(text string) {
			playerTwo.name = text
		}).
		AddButton("Start", func() {
			formApp.Stop()
		}).
		AddButton("Quit", func() {
			for i := 0; i < 5; i++ {
				playerOne.gains[i] = true
				playerTwo.gains[i] = true
			}
			formApp.Stop()
		})

	if err := formApp.SetRoot(form, true).Run(); err != nil {
		panic(err)
	}

	if settings.debug {
		// PLACING PLAYER ONE BOATS:
		// Carrier: (ID=0), horizontal, (1,1)
		playerOne.InitBoat(0, horizontal, 1, 1)
		// Battleship: (ID=1), horizontal, (0,9)
		playerOne.InitBoat(1, horizontal, 0, 9)
		// Destroyer: (ID=2), vertical, (5,6)
		playerOne.InitBoat(2, vertical, 5, 6)
		// Submarine: (ID=3), horizontal, (6,2)
		playerOne.InitBoat(3, horizontal, 6, 2)
		// Patrol Boat: (ID=4), vertical, (1,5)
		playerOne.InitBoat(4, vertical, 1, 5)

		// PLACING PLAYER TWO BOATS:
		// Carrier: (ID=0), vertical, (9,0)
		playerTwo.InitBoat(0, vertical, 9, 0)
		// Battleship: (ID=1), horizontal, (1,8)
		playerTwo.InitBoat(1, horizontal, 1, 8)
		// Destroyer: (ID=2), vertical, (5,3)
		playerTwo.InitBoat(2, vertical, 5, 3)
		// Submarine: (ID=3), horizontal, (2,2)
		playerTwo.InitBoat(3, horizontal, 2, 2)
		// Patrol Boat: (ID=4), vertical, (7,6)
		playerTwo.InitBoat(4, vertical, 6, 6)

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

		// // Display both primary boards in stdout
		// fmt.Println("Player One (vue de ses propres pièces):", playerOne.DisplayPrimary())
		// fmt.Println("Player Two (vue de ses propres pièces):", playerTwo.DisplayPrimary())
		// fmt.Println("Player One (vue des pièces de son ennemi):", playerOne.DisplayTarget())
		// fmt.Println("Player Two (vue des pièces de son ennemi):", playerTwo.DisplayTarget())

	}

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

	*/

	currentPlayer := &playerTwo
	var log string = "Game Started!"

	// Until Somedody Wins:
	for winner := currentPlayer; !(winner.gains == [5]bool{true, true, true, true, true}); {
		currentPlayer.Hit(9, 9)

		// Make the loop toggle between both players
		if currentPlayer == &playerTwo {
			currentPlayer = &playerOne
		} else {
			currentPlayer = &playerTwo
		}

		// Initializing the application
		app := tv.NewApplication()
		// HEADERBOX:
		// 	box which displays in it's title the current player.
		headerBox := tv.NewBox().SetTitle(currentPlayer.name).
			SetBorder(true)
		// KEYBINDINGSBOX:
		// 	simple box containing a list of all the keybindings one can use.
		keybindingsBox := tv.NewTextView()
		fmt.Fprintf(keybindingsBox, keybindings)
		keybindingsBox.SetTitle("Keybindings:").
			SetBorder(true)
		// LOGBOX:
		// 	box which shows a log of the past moves each player made.
		logBox := tv.NewTextView()
		fmt.Fprintln(logBox, log)
		logBox.SetTitle("Log:").
			SetBorder(true)
		// TARGETBOX:
		// 	box containing the board where the player attacks his opponent
		// 	this box is focused by default and is of type table as it can be navigated.
		targetBox := tv.NewTable()
		RedrawTarget(currentPlayer, targetBox)
		targetBox.SetFixed(1, 1).
			SetSelectable(true, true).
			SetSelectedFunc(func(row, column int) {
				// Hit the target at the chosen coordinate
				fmt.Fprintf(logBox, log)
				currentPlayer.Hit(column-1, row-1)
				// The coordinate where the player is hitting will no longer be selectable
				targetBox.SetSelectable(false, false)
				logEntry := fmt.Sprintf("%v: %c%v\n", currentPlayer.name, letters[column-1], row-1)
				log = fmt.Sprintf("%v%v", logEntry, log)
				RedrawTarget(currentPlayer, targetBox)
				app.Stop()
			}).
			SetDoneFunc(func(key tc.Key) {
				if key == tc.KeyEscape {
					// TODO: change this for a dopdown prompt "Are you sure? (Y/N)"
					for i := 0; i < 5; i++ {
						currentPlayer.prey.gains[i] = true
						currentPlayer.gains[i] = true
					}
					app.Stop()
				}
			}).
			SetBorder(true).
			SetTitle("The enemy:")

		// PRIMARYBOX:
		// 	box containing the board showing the players boat layout.
		primaryBox := tv.NewTable()
		RedrawPrimary(currentPlayer, primaryBox)
		primaryBox.SetFixed(1, 1).
			SetBorder(true).
			SetTitle("You:")
		// GAINSBOX:
		// 	box showing the names of all the ennemies boats which are sunk.
		// 	each sunk boat has a small display of the boat going with it.
		gainsBox := tv.NewTextView()
		RedrawGains(currentPlayer, gainsBox)
		gainsBox.SetBorder(true).
			SetTitle("Gains:")
		// LOSSESBOX:
		// 	box showing the names of all the players boats which are sunk.
		// 	each sunk boat has a small display of the boat going with it.
		lossesBox := tv.NewTextView()
		RedrawLosses(currentPlayer, lossesBox)
		lossesBox.SetBorder(true).
			SetTitle("Losses:")
		// COMMANDBOX:
		// 	box containing an input field which is to be used as a command prompt.
		// 	coordinates can be inputed directly and typing quit will exit the game.
		commandBox := tv.NewInputField().
			SetBorder(true).
			SetTitle("Command:")
		// Setting up the flexbox hell!
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
