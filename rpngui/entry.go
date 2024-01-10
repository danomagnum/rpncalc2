package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type enterEntry struct {
	widget.Entry
	history     []string
	history_pos int
	history_max int
}

func (e *enterEntry) onEnter() {
	e.history = append(e.history, e.Entry.Text)
	e.history_pos = 0
}

func (e *enterEntry) onUp() {
	e.history_pos++
	e.updateHistoryPos()
}

func (e *enterEntry) onDown() {
	e.history_pos--
	e.updateHistoryPos()

}

func (e *enterEntry) updateHistoryPos() {
	count := len(e.history)
	if e.history_pos < 0 {
		e.history_pos = -1
		e.Entry.SetText("")
		return
	}
	if e.history_pos > count {
		e.history_pos = count
	}
	if e.history_pos < count {
		t := e.history[count-e.history_pos-1]
		e.Entry.SetText(t)
	}
}

func (e *enterEntry) loadHistory() {
	h := LoadHistory()
	if len(h) > e.history_max {
		e.history = h[:e.history_max]
	} else {
		e.history = h
	}
}

func (e *enterEntry) shutdown() {
	shutdown_complete.Add(1)
	defer shutdown_complete.Done()
	<-shutdown
	h := e.history
	if len(h) > e.history_max {
		h = h[len(h)-e.history_max:]
	}
	SaveHistory(h)
}

func newEnterEntry(cfg configfile) *enterEntry {
	entry := &enterEntry{}
	entry.ExtendBaseWidget(entry)
	entry.history = make([]string, 0)
	entry.history_max = cfg.History
	entry.loadHistory()
	go entry.shutdown()
	return entry
}

func (e *enterEntry) KeyDown(key *fyne.KeyEvent) {
	switch key.Name {
	case fyne.KeyUp:
		e.onUp()
	case fyne.KeyDown:
		e.onDown()
	case fyne.KeyEnter, fyne.KeyReturn:
		e.onEnter()
		e.Entry.KeyDown(key)
	default:
		e.Entry.KeyDown(key)
	}
}
