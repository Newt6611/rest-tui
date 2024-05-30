package ui

import (
	"os"

	"github.com/charmbracelet/x/term"
)

func GetTerminalSize() (int, int) {
	width, height, _ := term.GetSize(os.Stdout.Fd())
	return width, height
}
