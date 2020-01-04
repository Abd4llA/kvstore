package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/abd4lla/kvstore/internal/keystore"
)

var (
	ks *keystore.KeyStore
)

func main() {
	fmt.Println("Starting KVStore")
	ks = keystore.NewKeyStore()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("\n1) Add/Set Key.\n2) Get/Find Value.\n3) List all.\nEnter choice (e.g. 1 for ADD): ")
		for {
			cmd, _ := reader.ReadString('\n')
			cmd = strings.Replace(cmd, "\n", "", -1)
			switch cmd {
			case "1":
				fmt.Print("Enter key: ")
				key, _ := reader.ReadString('\n')
				key = strings.Replace(key, "\n", "", -1)
				fmt.Print("Enter value: ")
				value, _ := reader.ReadString('\n')
				value = strings.Replace(value, "\n", "", -1)
				ks.Put(key, value)
			case "2":
				fmt.Print("Enter key: ")
				key, _ := reader.ReadString('\n')
				key = strings.Replace(key, "\n", "", -1)
				value := ks.Get(key)
				if value == "" {
					fmt.Printf(" !! Key %s not found\n", key)
				} else {
					fmt.Printf("-> Found 1 entry\n\tK: %s -> V: %s\n", key, value)
				}
			default:
				fmt.Printf("You entered: %s\n", cmd)
				fmt.Println("Ok")
			}
			fmt.Println("************************")
			fmt.Print("1) Add/Set Key.\n2) Get/Find Value.\n3) List all.\nEnter choice (e.g. 1 for ADD): ")
		}
	}
}
