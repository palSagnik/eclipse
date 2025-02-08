package tui

import (
	"fmt"

	"github.com/palsagnik2102/eclipse/pkg/commands"
	"github.com/rivo/tview"
)

type App struct {
	app *tview.Application
}

// NewApp creates an instance of a new application
func NewApp() *App {
	app:= tview.NewApplication()
	return &App{
		app: app,
	}
}

func (a *App) ListImage(images ...commands.Image) (error) {
	list := tview.NewList()

	// TODO(sagnik): Add on click functionality
	for _, img := range images {
		a := fmt.Sprintf("Name: %s", img.Name)
		b := fmt.Sprintf("ID: %s \nCreated: %d", img.ID[:16], img.Created)
		list.AddItem(a, b, '-', nil)
	}

	if err := a.app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		return err
	}

	return nil
}