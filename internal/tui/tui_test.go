package tui

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"
)

const (
	SUCCEEDED = "\u2713"
	FAILED    = "\u2717"
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
	r io.Reader
	w io.Writer
}

func NewMenu(r io.Reader, w io.Writer) *Menu {
	m := Menu {
		make(map[string]*Entry),
		r,
		w,
	}
	return &m
}

func (m *Menu) Add(name string, text string, actionKey string, action Action) error {
	m.entries[actionKey] = & Entry{
		Name: name,
		HelpText: text,
		ActionKey: actionKey,
		Callback: action,
	}
	//TODO Validate input and return error if the key is not unique
	return nil
}

func (m *Menu) Run(){
	//TODO Panic if we don't have any entries
	w, r := m.w, m.r
	for _ , e := range m.entries {
		fmt.Fprintf(w, "(%s) %s\n", e.ActionKey, e.HelpText)
	}
	fmt.Fprintln(w, "Enter choice:")
	for {
		fmt.Fprint(w)
		r := bufio.NewReader(r)
		a, _ := r.ReadString('\n')
		e, ok := m.entries[strings.TrimSpace(a)]
		if ! ok {
			fmt.Fprintln(w, "Invalid choice.")
			continue
		}
		e.Callback()
	}
}

func TestTUI(t *testing.T) {
	t.Log("Given a TUI menu")
	{
		b := new(bytes.Buffer)
		r,w := io.Pipe()

		// In normal mode, it should take stdin and stdout
		menu := NewMenu(r, b)

		t.Logf("\tTest 0:\tWhen adding a new entry")
		{
			if err := menu.Add("Add", "Add a new key/value pair", "1", func() {fmt.Fprint(b,"yippee ki yay")}); err != nil {
				t.Fatalf("\t%s\tShould be added successfully : %v", FAILED, err)
			}
			t.Logf("\t%s\tShould be added successfully.", SUCCEEDED)
		}
		t.Logf("\tTest 1:\tWhen printing the menu")
		{
			go menu.Run()
			var correctOut bool
			var output string
			for i := 0 ; i <= 3 && !correctOut; i++ {
				time.Sleep(200*time.Millisecond)
				output = b.String()
				correctOut = strings.Contains(output, "Add a new key/value pair") &&
					strings.Contains(output, "1")
			}
			if ! correctOut {
				t.Fatalf("\t%s\tShould be printed correctly. Got: %s", FAILED, output)
			}
			t.Logf("\t%s\tShould be printed correctly.", SUCCEEDED)
		}
		t.Logf("\tTest 2:\tWhen choosing an entry")
		{
			w.Write([]byte("1\n"))
			defer w.Close()
			output := b.String()
			if ! strings.Contains(output, "yippee ki yay") {
				t.Fatalf("\t%s\tEntry callback should be executed correctly. Step callback wasn't executed.", FAILED)
			}
			t.Logf("\t%s\tEntry callback should be executed correctly.", SUCCEEDED)
		}
	}
}