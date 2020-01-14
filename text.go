package main

const keybindings = `General:
	Esc: quit
Selection mode:
	h:	Left
	j:	Down
	k:	Up
	l:	Right
	":"	Command mode
Command mode:
	Enter a coordinate
	q:	quit
`

const letters = "ABCDEFGHIJ"

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

const boat0 = `Carrier:
 ◁ ▭ ▭ ▭ ▷ `
const boat1 = `Battleship:
 ◁ ▭ ▭ ▷ `
const boat2 = `Destroyer:
 ▭ ▭ ▷ `
const boat3 = `Submarine:
 ◁ ▭ ▷ `
const boat4 = `Patrol Boat:
 ▭ ▷ `
