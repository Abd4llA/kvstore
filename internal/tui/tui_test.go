package tui

import (
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

func TestTUI(t *testing.T) {
	t.Log("Given a TUI menu")
	{
		b := new(bytes.Buffer)
		r,w := io.Pipe()

		// When used in production, it should take stdin and stdout
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
