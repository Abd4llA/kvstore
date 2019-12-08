package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	data = make(map[string]string)
)

func main() {
	fmt.Println("KV-Store :: Starting")
	fmt.Println("Welcome to KV-Store")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("1) Add/Set Key.\n2) Get/Find Value.\n3) List all.\nEnter choice (e.g. 1 for ADD): ")
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
				data[key] = value
			case "2":
				fmt.Print("Enter key: ")
				key, _ := reader.ReadString('\n')
				key = strings.Replace(key, "\n", "", -1)
				value, ok := data[key]
				if !ok {
					fmt.Printf(" !! Key %s not found\n", key)
				} else {
					fmt.Printf("-> Found 1 entry\n\tK: %s -> V: %s\n", key, value)
				}
			case "3":
				for key, value := range data {
					fmt.Printf("\tK: %s -> V: %s\n", key, value)
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
