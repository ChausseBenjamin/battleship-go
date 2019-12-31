package main

import (
	"fmt"
)

func main() {

	settings := parameters{
		debug:    true,
		ssh_game: false,
	}

	if settings.debug {

		// SETUP:
		var player_one = player{}
		var player_two = player{}
		// Setting up prey for when using Hit function
		player_one.prey = &player_two
		player_two.prey = &player_one

		if settings.debug {
			fmt.Println("# #---TESTING SEQUENCE---# #")

			player_one.primary = [10][10][3]int{} // Empty plyr1 board
			// Initialising boats on the board work* (see initBoat TODOS)
			initBoat(&player_one, [4]int{0, 0, 0, 0}) // Index 0 -> Carrier
			initBoat(&player_one, [4]int{1, 0, 6, 9}) // Index 1 -> Battleship
			initBoat(&player_one, [4]int{2, 1, 4, 3}) // Index 2 -> Destroyer
			initBoat(&player_one, [4]int{3, 1, 7, 1}) // Index 3 -> Submarine
			initBoat(&player_one, [4]int{4, 1, 1, 8}) // Index 4 -> PatrolBoat
			// fmt.Println("Player 1:")
			// fmt.Print(player_one.PrimaryDisplay())

			// Initialising the board works
			// fmt.Println("Empty:")
			// fmt.Print(player_one.PrimaryDisplay())

			player_two.primary = [10][10][3]int{}     // Empty plyr2 board
			initBoat(&player_two, [4]int{0, 0, 3, 3}) // Index 0 -> Carrier
			initBoat(&player_two, [4]int{1, 0, 4, 8}) // Index 1 -> Battleship
			initBoat(&player_two, [4]int{2, 1, 2, 4}) // Index 2 -> Destroyer
			initBoat(&player_two, [4]int{3, 1, 9, 0}) // Index 3 -> Submarine
			initBoat(&player_two, [4]int{4, 1, 7, 4}) // Index 4 -> PatrolBoat

			fmt.Println("Player 2:")
			fmt.Print(player_two.PrimaryDisplay())

			// fmt.Println("Hit B2:")
			// hit_coord := [2]int{1, 2} // Water hit at B2
			// fmt.Print("There was a boat: ")
			// fmt.Println(player_one.Hit(hit_coord))
			// fmt.Println(player_two.PrimaryDisplay())
			// fmt.Println("Hit H4:")
			// hit_coord = [2]int{7, 4} // PatrolBoat hit at H4
			// fmt.Print("There was a boat: ")
			// fmt.Println(player_one.Hit(hit_coord))
			// fmt.Println(player_two.PrimaryDisplay())
			// fmt.Println("Player 1 TargetDisplay:")
			// fmt.Println(player_one.TargetDisplay())
			// fmt.Println("Hit H5:")
			// hit_coord = [2]int{7, 5} // PatrolBoat hit at H4
			// fmt.Print("There was a boat: ")
			// fmt.Println(player_one.Hit(hit_coord))
			// fmt.Println(player_two.PrimaryDisplay())
			// fmt.Println("Player 1 TargetDisplay:")
			// fmt.Println(player_one.TargetDisplay())

			fmt.Println("Hit: E8, F8, G8")
			player_one.Hit([2]int{4, 8})
			player_one.Hit([2]int{5, 8})
			player_one.Hit([2]int{6, 8})
			fmt.Println(player_two.PrimaryDisplay())
			fmt.Println(player_one.TargetDisplay())
			player_one.Hit([2]int{6, 8})
			fmt.Println(player_two.PrimaryDisplay())
			fmt.Println(player_one.TargetDisplay())

		}
	}
}
