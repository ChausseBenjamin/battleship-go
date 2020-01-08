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

			// Making an empty player_one board canvas to place boats:
			player_one.primary = [10][10][3]int{}
			// Index 0 -> Carrier // (x,y) = (0,0) // Horizontal:
			initBoat(&player_one, [4]int{0, 0, 0, 0})
			// Index 1 -> Battleship // (x,y) = (6,9) // Horizontal:
			initBoat(&player_one, [4]int{1, 0, 6, 9})
			// Index 2 -> Destroyer // (x,y) = (4,3) // Vertical:
			initBoat(&player_one, [4]int{2, 1, 4, 3})
			// Index 3 -> Submarine // (x,y) = (7,1) // Vertical:
			initBoat(&player_one, [4]int{3, 1, 7, 1})
			// Index 4 -> PatrolBoat // (x,y) = (1,8) // Vertical:
			initBoat(&player_one, [4]int{4, 1, 1, 8})

			// Making an empty player_one board canvas to place boats:
			player_two.primary = [10][10][3]int{}
			// Index 0 -> Carrier // (x,y) = (3,3) // Horizontal
			initBoat(&player_two, [4]int{0, 0, 3, 3})
			// Index 1 -> Battleship // (x,y) = (4,8) // Horizontal
			initBoat(&player_two, [4]int{1, 0, 4, 8})
			// Index 2 -> Destroyer // (x,y) = (2,4) // Vertical
			initBoat(&player_two, [4]int{2, 1, 2, 4})
			// Index 3 -> Submarine // (x,y) = (9,0) // Vertical
			initBoat(&player_two, [4]int{3, 1, 9, 0})
			// Index 4 -> PatrolBoat // (x,y) = (7,4) // Vertical
			initBoat(&player_two, [4]int{4, 1, 7, 4})

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
