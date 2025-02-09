package main

import "fmt"

func main() {
	// var colors map[string]string
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"white": "#ffffff",
	}

	// colors := make(map[string]string)

	// colors["green"] = "#fe3434"

	// delete(colors, "green")

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hexcode for", color, "is", hex)
	}
}
