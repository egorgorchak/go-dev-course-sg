package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for c, h := range c {
		fmt.Println("key is -", c)
		fmt.Println("value is -", h)
	}
}
