package main

import (
	"fmt"
	"github.com/mbndr/figlet4go"
	tv "github.com/rivo/tview"
)

func main() {

	settings := parameters{
		debug:    true,
		ssh_game: false,
	}

	// SETUP:
	var player_one = player{}
	var player_two = player{}
	// Setting up prey for when using Hit function
	player_one.prey = &player_two
	player_two.prey = &player_one

	if settings.debug {

		ascii := figlet4go.NewAsciiRender()

		options := figlet4go.NewRenderOptions()
		options.FontName = "TwoPoint"

		// If 'larry3d' wouldn't be included you would have to load your .flf files like that:
		ascii.LoadFont("./ressources/")

		renderStr, _ := ascii.RenderOpts("Hello Fonts", options)
		fmt.Print(renderStr)

		app := tv.NewApplication()
		flex := tv.NewFlex().
			//AddItem(item, fixedSize, proportion, focus)
			AddItem(tv.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
			AddItem(tv.NewFlex().SetDirection(tv.FlexRow).
				AddItem(tv.NewBox().SetBorder(true).SetTitle("Top                       "), 0, 1, false).
				AddItem(tv.NewBox().SetBorder(true).SetTitle("Middle (3 x height of Top)"), 0, 3, false).
				AddItem(tv.NewBox().SetBorder(true).SetTitle("Bottom (5 rows)           "), 5, 1, false), 0, 2, false).
			AddItem(tv.NewFlex().SetDirection(tv.FlexRow).
				AddItem(tv.NewBox().SetBorder(true).SetTitle("Gains"), 0, 1, false).
				AddItem(tv.NewBox().SetBorder(true).SetTitle("Losses"), 0, 1, false), 12, 2, false)
		if err := app.SetRoot(flex, true).SetFocus(flex).Run(); err != nil {
			panic(err)
		}
	}

}
