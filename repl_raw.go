package main

import (
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	return words
}

type cliCommand struct {
	name        string
	description string
	callback    func(string, *Config) error
}

func repl(cfg *Config) {
	// Enter raw mode at start
	err := EnterRawMode()
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		return
	}

	prevCommands := []string{}
	currentInput := ""
	historyIdx := -1

	fmt.Print("\r\033[2KPokedex > ")

	for {
		// Increase buffer size to handle paste operations
		buffer := make([]byte, 1024)
		n, err := os.Stdin.Read(buffer)
		if err != nil {
			fmt.Println("Error reading input:", err)
			// Restore terminal in case of read error
			err := RestoreTerminalState()
			if err != nil {
				fmt.Println("Error restoring terminal state:", err)
			}
			return
		}

		// Up arrow
		if n == 3 && buffer[0] == 27 && buffer[1] == 91 && buffer[2] == 65 {
			if historyIdx == -1 {
				continue
			}

			if historyIdx == 0 {
				currentInput = prevCommands[historyIdx]
				fmt.Printf("\r\033[2KPokedex > %s", currentInput)
				continue
			}

			if historyIdx > 0 {
				currentInput = prevCommands[historyIdx]
			}

			historyIdx--
			fmt.Printf("\r\033[2KPokedex > %s", currentInput)
			continue
		}

		// Down arrow
		if n == 3 && buffer[0] == 27 && buffer[1] == 91 && buffer[2] == 66 {
			if historyIdx == -1 {
				continue
			}

			historyIdx++

			if historyIdx == len(prevCommands) {
				historyIdx = len(prevCommands) - 1
				currentInput = ""
				fmt.Printf("\r\033[2KPokedex > ")
				continue
			}

			if historyIdx < len(prevCommands) {
				currentInput = prevCommands[historyIdx]
			}

			fmt.Printf("\r\033[2KPokedex > %s", currentInput)
			continue
		}

		// Enter
		if n == 1 && buffer[0] == 13 {
			fmt.Println() // Move to new line after input
			cleanedInput := cleanInput(currentInput)

			if len(cleanedInput) == 0 {
				fmt.Print("\r\033[2KPokedex > ")
				continue
			}

			command := cleanedInput[0]
			additionalInput := ""
			if len(cleanedInput) > 1 {
				additionalInput = cleanedInput[1]
			}
			if len(additionalInput) > 0 {
				prevCommands = append(prevCommands, command+" "+additionalInput)
			} else {
				prevCommands = append(prevCommands, command)
			}

			historyIdx = len(prevCommands) - 1
			currentInput = ""

			// Enter standard mode before executing command
			err := EnterStandardMode()
			if err != nil {
				fmt.Println("Error entering standard mode:", err)
			}

			// Print an explicit carriage return and clear line to ensure proper alignment
			fmt.Print("\r\033[2K")

			// Execute command
			var cmdErr error
			if c, exist := callCommand()[command]; !exist {
				fmt.Println("Unknown command")
			} else {
				cmdErr = c.callback(additionalInput, cfg)
			}

			// Handle other errors
			if cmdErr != nil {
				fmt.Printf("\r\033[2K%s\n", cmdErr.Error())
			}
			// Return to raw mode for input
			err = EnterRawMode()
			if err != nil {
				fmt.Println("Error entering raw mode:", err)
			}

			fmt.Print("\r\033[2KPokedex > ")
			continue
		}

		// Backspace
		if n == 1 && (buffer[0] == 127 || buffer[0] == 8) {
			if len(currentInput) > 0 {
				currentInput = currentInput[:len(currentInput)-1]
				fmt.Print("\b \b")
			}
			continue
		}

		// Handle paste operations (multiple characters at once)
		if n > 1 {
			// Process each character in the pasted content
			for i := 0; i < n; i++ {
				// Only handle printable ASCII characters
				if buffer[i] >= 32 && buffer[i] <= 126 {
					currentInput += string(buffer[i])
					fmt.Print(string(buffer[i]))
				}
			}
			continue
		}

		// Single printable character
		if n == 1 && buffer[0] >= 32 && buffer[0] <= 126 {
			currentInput += string(buffer[0])
			fmt.Print(string(buffer[0]))
		}
	}
}