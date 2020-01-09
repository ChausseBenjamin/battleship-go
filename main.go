package main

import (
	"fmt"
	tv "github.com/rivo/tview"
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
		initBoat(&playerOne, [4]int{0, 0, 1, 1})
		// Battleship: (ID=1), horizontal, (0,9)
		initBoat(&playerOne, [4]int{1, 0, 0, 9})
		// Destroyer: (ID=2), vertical, (5,6)
		initBoat(&playerOne, [4]int{2, 1, 5, 6})
		// Submarine: (ID=3), horizontal, (6,2)
		initBoat(&playerOne, [4]int{3, 0, 6, 2})
		// Patrol Boat: (ID=4), vertical, (1,5)
		initBoat(&playerOne, [4]int{4, 1, 1, 5})

		// PLACING PLAYER TWO BOATS:
		// Carrier: (ID=0), vertical, (9,0)
		initBoat(&playerTwo, [4]int{0, 1, 9, 0})
		// Battleship: (ID=1), horizontal, (1,8)
		initBoat(&playerTwo, [4]int{1, 0, 1, 8})
		// Destroyer: (ID=2), vertical, (5,3)
		initBoat(&playerTwo, [4]int{2, 1, 5, 3})
		// Submarine: (ID=3), horizontal, (2,2)
		initBoat(&playerTwo, [4]int{3, 0, 2, 2})
		// Patrol Boat: (ID=4), vertical, (7,6)
		initBoat(&playerTwo, [4]int{4, 1, 6, 6})

		// HITTING PLAYER ONE AT DIFFERENT COORDINATES
		playerTwo.Hit([2]int{5, 6}) // (F,6)
		playerTwo.Hit([2]int{5, 7}) // (F,7)
		playerTwo.Hit([2]int{5, 8}) // (F,8)
		playerTwo.Hit([2]int{3, 4}) // (D,4)
		playerTwo.Hit([2]int{3, 9}) // (D,9)
		playerTwo.Hit([2]int{6, 6}) // (G,6)
		playerTwo.Hit([2]int{2, 9}) // (C,9)

		// HITTING PLAYER TWO AT DIFFERENT COORDINATES
		playerOne.Hit([2]int{1, 4}) // (B,4)
		playerOne.Hit([2]int{1, 8}) // (B,8)
		playerOne.Hit([2]int{2, 7}) // (C,7)
		playerOne.Hit([2]int{2, 8}) // (C,8)
		playerOne.Hit([2]int{3, 8}) // (D,8)
		playerOne.Hit([2]int{4, 8}) // (E,8)
		playerOne.Hit([2]int{9, 2}) // (I,2)

		// Display both primary boards in stdout
		fmt.Println("Player One:", playerOne.DisplayPrimary())
		fmt.Println("Player Two:", playerTwo.DisplayPrimary())

		// Setting the Player header box
		headerBox := tv.NewBox().SetTitle(playerOne.name).
			SetBorder(true)

		// Setting up the keybindings box
		// Where the keybindings list will be
		keybindingsBox := tv.NewBox().
			SetTitle("Keybindings:").
			SetBorder(true)
		// TODO: Add text/documentation to box

		// Setting the log box whoch shows a history of past moves
		logBox := tv.NewBox().
			SetTitle("Log:").
			SetBorder(true)

		// Setting up the target box
		targetBox := tv.NewTable().
			SetBorder(true).
			SetTitle("The ennemy:")

		// Setting up the primary box
		primaryBox := tv.NewTable().
			SetBorder(true).
			SetTitle("You:")

		// Setting up the gains box
		gainsBox := tv.NewList().
			SetBorder(true).
			SetTitle("Gains:")

		// Setting up the losses box
		lossesBox := tv.NewList().
			SetBorder(true).
			SetTitle("Losses:")

		// Setting up the command prompt box
		commandBox := tv.NewInputField().
			SetBorder(true).
			SetTitle("Command:")

		// Setting up the target flex
		targetFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(targetBox, 26, 0, true).
			AddItem(gainsBox, 26, 0, false)

		// Setting up the primary flex
		primaryFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(primaryBox, 26, 0, false).
			AddItem(lossesBox, 26, 0, false)

		// // Setting up the play flex
		playFlex := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(targetFlex, 13, 0, false).
			AddItem(primaryFlex, 13, 0, false).
			AddItem(commandBox, 0, 1, false)

		// // Setting up the info flex
		infoFlex := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(keybindingsBox, 0, 3, false).
			AddItem(logBox, 0, 1, false)

		// Setting up the bottom flex
		bottomFlex := tv.NewFlex().SetDirection(tv.FlexColumn).
			AddItem(infoFlex, 0, 1, false).
			AddItem(playFlex, 52, 0, false)

		// Setting up the application layout
		app := tv.NewApplication()
		// Dashboard is the Flex containing everything.
		dashboard := tv.NewFlex().SetDirection(tv.FlexRow).
			AddItem(headerBox, 2, 1, false).
			AddItem(bottomFlex, 0, 1, false)

		if err := app.SetRoot(dashboard, true).Run(); err != nil {
			panic(err)
		}
	}
}
