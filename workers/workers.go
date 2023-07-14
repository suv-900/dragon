package workers

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/eiannone/keyboard"
)

func AddtoFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Add note and Press Ctrl+s to save:")
	err = keyboard.Open()
	if err != nil {
		fmt.Println("Couldnt initialise keyboard events ")
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	channel := make(chan string)
	var msg string
	go func() {
		for {
			events, err := keyboard.GetKeys(100) //returns a recieve only channel of keyboard inputs
			if err != nil {
				log.Fatalf("Error occured in recieve-only keyboard channel\n%v", err)
			}
			for event := range events {
				if event.Err != nil {
					log.Fatalf("Error occured in keyboard receive only channel %v", err)
				} else if event.Key == keyboard.KeyCtrlS {
					channel <- "save"
				}
			}
		}
	}()
	msg = <-channel

	var input string
	for {
		scanner.Scan()
		input += scanner.Text()
		if msg == "save" {
			break
		}
	}
	fmt.Println("Saving...")

	fileWriter := bufio.NewWriter(file)
	_, err = fileWriter.WriteString(input)
	if err != nil {
		fmt.Println("Error occured while writing to file")
		fmt.Print(err)
		return
	}

	fmt.Println("File saved successfully :)")
}
