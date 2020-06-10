package tui

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Action func()

type Entry struct {
	Name      string
	HelpText  string
	ActionKey string
	Callback  Action
}

type Menu struct {
	entries map[string]*Entry
	r       io.Reader
	w       io.Writer
}

func NewMenu(r io.Reader, w io.Writer) *Menu {
	m := Menu{
		make(map[string]*Entry),
		r,
		w,
	}
	return &m
}

func (m *Menu) Add(name string, help string, actionKey string, action Action) error {
	m.entries[actionKey] = &Entry{
		Name:      name,
		HelpText:  help,
		ActionKey: actionKey,
		Callback:  action,
	}
	//TODO Validate input and return error if the key is not unique
	return nil
}

func (m *Menu) Run() {
	//TODO Panic if we don't have any entries
	w, r := m.w, m.r
	for {
		m.printMenu()
		fmt.Fprint(w)
		r := bufio.NewReader(r)
		a, _ := r.ReadString('\n')
		e, ok := m.entries[strings.TrimSpace(a)]
		if !ok {
			fmt.Fprintln(w, "Invalid choice.")
			continue
		}
		e.Callback()
	}
}

func (m *Menu) printMenu() {
	w := m.w
	for _, e := range m.entries {
		fmt.Fprintf(w, "(%s) %s\n", e.ActionKey, e.HelpText)
	}
	fmt.Fprint(w, "Enter choice: ")
}
