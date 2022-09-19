package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"black": "#000000",
		"green": "#7820323",
	}
	// fmt.Println(colors)

	// var colors map[string]string

	// colors["red"] := "ff0000"

	// colors := make(map[string]string)

	// colors["red"] = "#ffffff"

	// delete(colors, "red")

	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {

		fmt.Println(color, hex)
	}
}
