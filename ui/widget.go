package ui

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type NumericalEntry struct {
	widget.Entry
}

type ReadOnlyEntry struct {
	widget.Entry
}

// An extended Fyne.io Entry that will only allow numeric (0-9) inputs
func NewNumericalEntry() *NumericalEntry {
	entry := &NumericalEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewNumericalEntryWithData(data binding.String) *NumericalEntry {
	entry := NewNumericalEntry()
	entry.Bind(data)
	return entry
}

// Override standard Entry.TypedRune to only allow numbers
func (e *NumericalEntry) TypedRune(r rune) {
	if r >= '0' && r <= '9' {
		e.Entry.TypedRune(r)
	}
}

// Override the standard copy-paste shortcut for Entry to only allow numbers to be pasted
func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseInt(content, 10, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func NewReadOnlyEntry() *ReadOnlyEntry {
	entry := &ReadOnlyEntry{}
	entry.ExtendBaseWidget(entry)
	return entry
}

func (e *ReadOnlyEntry) TypedKey(key *fyne.KeyEvent) {
	// do nothing
}

func (e *ReadOnlyEntry) TypedRune(r rune) {
	// do nothing
}

func (e *ReadOnlyEntry) TypedShortcut(shortcut fyne.Shortcut) {
	// do nothing
}
