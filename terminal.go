package robologger

import (
  "bufio"
  "fmt"
  "os"
  "os/exec"
  "regexp"
  "strconv"
  "strings"
)

// The escape code for the ANSI terminal. Defaults to 27, which is "\033".
const (
  ESC = 27
)

// Terminal is a structure which exports methods to directly interact with the
// terminal. It mainly prints ANSI codes to the terminal corresponding to
// specific commands.
type Terminal struct {
  ESC int
}

func NewTerminal(esc int) *Terminal {
  return &Terminal{esc}
}

var term = NewTerminal(ESC)

// The functions below provide simple terminal manipulation to move the cursor.
func (t Terminal) Clear() {
  fmt.Printf("%c[K", t.ESC)
}

func (t Terminal) ClearScreen() {
  fmt.Printf("%c[2J", t.ESC)
}

// Should not be used, except in certain circumstances.
func (t Terminal) Move(x int, y int) {
	fmt.Printf("%c[%d;%dH", t.ESC, x, y)
}

func (t Terminal) MoveUp(n int) {
  fmt.Printf("%c[%dA", t.ESC, n)
}

func (t Terminal) MoveDown(n int) {
  fmt.Printf("%c[%dB", t.ESC, n)
}

func (t Terminal) MoveForward(n int) {
  fmt.Printf("%c[%dC", t.ESC, n)
}

func (t Terminal) MoveBackward(n int) {
  fmt.Printf("%c[%dD", t.ESC, n)
}

func (t Terminal) MoveToBeginning() {
  fmt.Printf("\r")
}

func (t Terminal) HideCursor() {
  fmt.Printf("%c[?25l", t.ESC)
}

func (t Terminal) ShowCursor() {
  fmt.Printf("%c[?25h", t.ESC)
}

func (t Terminal) SaveCursorPosition() {
  fmt.Printf("%c[s", t.ESC)
}

func (t Terminal) RestoreCursorPosition() {
  fmt.Printf("%c[u", t.ESC)
}

func (t Terminal) EnterRawMode() {
  cmd := exec.Command("/bin/stty", "raw")
  cmd.Stdin = os.Stdin
  _ = cmd.Run()
  cmd.Wait()
}

func (t Terminal) ExitRawMode() {
  cmd := exec.Command("/bin/stty", "-raw")
  cmd.Stdin = os.Stdin
  _ = cmd.Run()
  cmd.Wait()
}

func (t Terminal) GetCursorPosition() (int, int) {
  t.EnterRawMode()
  defer t.ExitRawMode()

  // cmd := exec.Command("echo", fmt.Sprintf("%c[6n", 27))
	// randomBytes := &bytes.Buffer{}
	// cmd.Stdout = randomBytes
  // _ = cmd.Start()

  fmt.Printf(fmt.Sprintf("\r%c[6n", t.ESC))

	reader := bufio.NewReader(os.Stdin)
  // cmd.Wait()

  // fmt.Print(randomBytes)

	text, _ := reader.ReadSlice('R')

	re := regexp.MustCompile(`\d+;\d+`)
	res := re.FindString(string(text))

	if res != "" {
		parts := strings.Split(res, ";")
		line, _ := strconv.Atoi(parts[0])
		col, _ := strconv.Atoi(parts[1])
    return line, col
	}

  return 0, 0
}
