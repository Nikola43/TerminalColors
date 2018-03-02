package main

import "fmt"

// Reset
const RESET = "\033[0m";  // Text Reset

// Regular Colors
const BLACK = "\033[0;30m";   // BLACK
const RED = "\033[0;31m";     // RED
const GREEN = "\033[0;32m";   // GREEN
const YELLOW = "\033[0;33m";  // YELLOW
const BLUE = "\033[0;34m";    // BLUE
const PURPLE = "\033[0;35m";  // PURPLE
const CYAN = "\033[0;36m";    // CYAN
const WHITE = "\033[0;37m";   // WHITE

// Bold
const BLACK_BOLD = "\033[1;30m";  // BLACK
const RED_BOLD = "\033[1;31m";    // RED
const GREEN_BOLD = "\033[1;32m";  // GREEN
const YELLOW_BOLD = "\033[1;33m"; // YELLOW
const BLUE_BOLD = "\033[1;34m";   // BLUE
const PURPLE_BOLD = "\033[1;35m"; // PURPLE
const CYAN_BOLD = "\033[1;36m";   // CYAN
const WHITE_BOLD = "\033[1;37m";  // WHITE

// Underline
const BLACK_UNDERLINED = "\033[4;30m";  // BLACK
const RED_UNDERLINED = "\033[4;31m";    // RED
const GREEN_UNDERLINED = "\033[4;32m";  // GREEN
const YELLOW_UNDERLINED = "\033[4;33m"; // YELLOW
const BLUE_UNDERLINED = "\033[4;34m";   // BLUE
const PURPLE_UNDERLINED = "\033[4;35m"; // PURPLE
const CYAN_UNDERLINED = "\033[4;36m";   // CYAN
const WHITE_UNDERLINED = "\033[4;37m";  // WHITE

// Background
const BLACK_BACKGROUND = "\033[40m";  // BLACK
const RED_BACKGROUND = "\033[41m";    // RED
const GREEN_BACKGROUND = "\033[42m";  // GREEN
const YELLOW_BACKGROUND = "\033[43m"; // YELLOW
const BLUE_BACKGROUND = "\033[44m";   // BLUE
const PURPLE_BACKGROUND = "\033[45m"; // PURPLE
const CYAN_BACKGROUND = "\033[46m";   // CYAN
const WHITE_BACKGROUND = "\033[47m";  // WHITE

// High Intensity
const BLACK_BRIGHT = "\033[0;90m";  // BLACK
const RED_BRIGHT = "\033[0;91m";    // RED
const GREEN_BRIGHT = "\033[0;92m";  // GREEN
const YELLOW_BRIGHT = "\033[0;93m"; // YELLOW
const BLUE_BRIGHT = "\033[0;94m";   // BLUE
const PURPLE_BRIGHT = "\033[0;95m"; // PURPLE
const CYAN_BRIGHT = "\033[0;96m";   // CYAN
const WHITE_BRIGHT = "\033[0;97m";  // WHITE

// Bold High Intensity
const BLACK_BOLD_BRIGHT = "\033[1;90m"; // BLACK
const RED_BOLD_BRIGHT = "\033[1;91m";   // RED
const GREEN_BOLD_BRIGHT = "\033[1;92m"; // GREEN
const YELLOW_BOLD_BRIGHT = "\033[1;93m";// YELLOW
const BLUE_BOLD_BRIGHT = "\033[1;94m";  // BLUE
const PURPLE_BOLD_BRIGHT = "\033[1;95m";// PURPLE
const CYAN_BOLD_BRIGHT = "\033[1;96m";  // CYAN
const WHITE_BOLD_BRIGHT = "\033[1;97m"; // WHITE

// High Intensity backgrounds
const BLACK_BACKGROUND_BRIGHT = "\033[0;100m";// BLACK
const RED_BACKGROUND_BRIGHT = "\033[0;101m";// RED
const GREEN_BACKGROUND_BRIGHT = "\033[0;102m";// GREEN
const YELLOW_BACKGROUND_BRIGHT = "\033[0;103m";// YELLOW
const BLUE_BACKGROUND_BRIGHT = "\033[0;104m";// BLUE
const PURPLE_BACKGROUND_BRIGHT = "\033[0;105m"; // PURPLE
const CYAN_BACKGROUND_BRIGHT = "\033[0;106m";  // CYAN
const WHITE_BACKGROUND_BRIGHT = "\033[0;107m";   // WHITE

type Color struct {

}

func NewColor() *Color {
	c := &Color{}
	return c
}

func (c Color) printlnColor(text string, color string) {
	fmt.Println(color + text + RESET)
}

func (c Color) printColor(text string, color string) {
	fmt.Println(color + text + RESET)
}
