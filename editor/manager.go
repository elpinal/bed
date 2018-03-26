package editor

import (
	. "github.com/itchyny/bed/common"
	"github.com/itchyny/bed/event"
	"github.com/itchyny/bed/layout"
)

// Manager defines the required window manager interface for the editor.
type Manager interface {
	Init(chan<- event.Event, chan<- struct{})
	Open(string) error
	SetSize(int, int)
	Resize(int, int)
	Run()
	Emit(event.Event)
	State() (map[int]*WindowState, layout.Layout, int, error)
	Close()
}
