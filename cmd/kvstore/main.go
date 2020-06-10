package main

import (
	"bufio"
	"fmt"
	"github.com/abd4lla/kvstore/internal/tui"
	"os"
	"strings"

	"github.com/abd4lla/kvstore/internal/keystore"
)

const (
	MENU_TXT = "(1) Add/Set Key\n(2) Get/Find Value\n(h) Print help menu\n" +
		"(q) Exit\nEnter choice (e.g. 1 for ADD): "
)

var (
	ks *keystore.KeyStore
)

func main() {
	fmt.Println("KV-Store 0.01")
	fmt.Print("Starting in shell mode\n\n\n")
	ks = keystore.NewKeyStore()
	reader := bufio.NewReader(os.Stdin)
	m := tui.NewMenu(os.Stdin, os.Stdout)
	m.Add("add_key", "Add/Set Key", "1", func(){
		fmt.Print("Enter key: ")
		key, _ := reader.ReadString('\n')
		key = strings.Replace(key, "\n", "", -1)
		fmt.Print("Enter value: ")
		value, _ := reader.ReadString('\n')
		value = strings.Replace(value, "\n", "", -1)
		ks.Put(key, value)
	})
	m.Add("get_value", "Get/Find Value", "2", func() {
		fmt.Print("Enter key: ")
		key, _ := reader.ReadString('\n')
		key = strings.Replace(key, "\n", "", -1)
		value := ks.Get(key)
		if value == "" {
			fmt.Printf(" !! Key %s not found\n", key)
		} else {
			fmt.Printf("-> Found 1 entry\n\tK: %s -> V: %s\n", key, value)
		}
	})
	m.Add("quit", "Quit", "q", func(){
		os.Exit(0)
	})
	m.Run()
}
