package tui

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
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
	entries map[string]Entry
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
				t.Fatal("\t%s\tShould be added successfully : %v", FAILED, err)
			}
			t.Logf("\t%s\tShould be added successfully.", SUCCEEDED)
		}
		t.Logf("\tTest 1:\tWhen printing the menu")
		{
			menu.Display()
			output := b.String()
			correctOut := strings.Contains(output, "Add a new key/value pair") &&
				strings.Contains(output, "1")
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
				t.Fatalf("\t%s\tEntry callback should be executed correctly. Step callback wasn't executed", FAILED, output)
			}
			t.Logf("\t%s\tEntry callback should be executed correctly.", SUCCEEDED)
		}
	}
}
