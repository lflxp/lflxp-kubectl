package pkg

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

func KeyHelp(g *gocui.Gui) error {
	if err := g.SetKeybinding("", gocui.KeyF1, gocui.ModNone, gethelp); err != nil {
		return err
	}
	return nil
}

func gethelp(g *gocui.Gui, v *gocui.View) error {
	if err = delOtherViewNoBack(g); err != nil {
		return err
	}
	maxX, maxY := g.Size()
	if v, err := g.SetView("gethelp", 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.Title = "Keybindings(F1: quit)"
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F1", "yellow", "", true, true), "Show keybinding help"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F2", "yellow", "", true, true), "Dashboard View & back to Dashboard View && Refresh"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F3", "yellow", "", true, true), "Pod View & back to Pod View && Refresh"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F4", "yellow", "", true, true), "Deployment View && Refresh"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F5", "yellow", "", true, true), "Service View && Refresh"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  F6", "yellow", "", true, true), "Node View && Refresh"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  Space", "yellow", "", true, true), "search current view information"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  Ctrl+C", "yellow", "", true, true), "Exit"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  Ctrl+L", "yellow", "", true, true), "Show log -tail 200 on F3 Pod View || msg refresh Log view"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  ↑ ↓", "yellow", "", true, true), "Move View"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  Enter", "yellow", "", true, true), "Commit Input/Quit current view(msg)"))
		fmt.Fprintln(v, fmt.Sprintf("%s: %s", Colorize("*  Tab", "yellow", "", true, true), "Next View"))

		if _, err := setCurrentViewOnTop(g, "gethelp"); err != nil {
			return err
		}
	}

	return nil
}
