package main

import "fmt"

func main() {
	//var colors map[string]string
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
	}
	printMap(colors)
	//	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Index: %s,\tValue: %s\n", color, hex)
	}
}
