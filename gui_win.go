// +build windows

package gocui

import "github.com/gdamore/tcell/termbox"

func (g *Gui) getTermWindowSize() (int, int, error) {
	x, y := termbox.Size()
	return x, y, nil
}
