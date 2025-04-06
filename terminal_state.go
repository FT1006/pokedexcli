package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

// Global variables to hold terminal state information
var (
	terminalState *term.State
	stdinFd       int
)

// SaveTerminalState stores the terminal state globally
func SaveTerminalState(state *term.State, fd int) {
	terminalState = state
	stdinFd = fd
}

// RestoreTerminalState restores the terminal to its original state
func RestoreTerminalState() error {
	if terminalState != nil {
		return term.Restore(stdinFd, terminalState)
	}
	return nil
}

// EnterRawMode puts the terminal in raw mode and returns the previous state
func EnterRawMode() error {
	fd := int(os.Stdin.Fd())
	state, err := term.MakeRaw(fd)
	if err != nil {
		return fmt.Errorf("error setting raw mode: %w", err)
	}
	// Save terminal state globally
	SaveTerminalState(state, fd)
	return nil
}

// EnterStandardMode temporarily restores the terminal to standard mode
// This should be called before any command that needs standard terminal behavior
func EnterStandardMode() error {
	if terminalState == nil {
		return fmt.Errorf("terminal state not saved")
	}
	return term.Restore(stdinFd, terminalState)
}
