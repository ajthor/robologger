package robologger

import (
  "fmt"
)

// The functions below provide simple terminal manipulation to move the cursor.
func (l *Logger) Clear() {
  fmt.Printf("\033[K")
}

func (l *Logger) ClearScreen() {
  fmt.Printf("\033[2J")
}

// Should not be used, except in certain circumstances.
func (l *Logger) Move(x int, y int) {
	fmt.Printf("\033[%d;%dH", x, y)
}

func (l *Logger) MoveUp(n int) {
  fmt.Printf("\033[%dA", n)
}

func (l *Logger) MoveDown(n int) {
  fmt.Printf("\033[%dB", n)
}

func (l *Logger) MoveForward(n int) {
  fmt.Printf("\033[%dC", n)
}

func (l *Logger) MoveBackward(n int) {
  fmt.Printf("\033[%dD", n)
}

func (l *Logger) MoveToBeginning() {
  fmt.Printf("\r")
}

func (l *Logger) SaveCursorPosition() {
  fmt.Printf("\033[s")
}

func (l *Logger) RestoreCursorPosition() {
  fmt.Printf("\033[u")
}
