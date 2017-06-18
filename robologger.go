package robologger

import (
  "bufio"
  "fmt"
  "io"
  "os"
  "regexp"
)

// Stdin, Stdout, and Stderr are files opened by the "os" package.
var (
  Stdin  = os.Stdin
  Stdout = os.Stdout
  Stderr = os.Stderr
)

// Message is the interface that implements methods for working with messages.
type Message interface {
  fmt.Stringer
  Format() string
  getPrintLength() (n int)
  setPrintLength(n int)
}

// Scanner is the interface that implements the ScanLine method.
type Scanner interface {
  ScanLine() string
}

// ScanLine reads the next line of input from the terminal. It implements a
// Scan method from the bufio Scanner.
func ScanLine() string {
  scanner := bufio.NewScanner(Stdin)
  scanner.Scan()
  return scanner.Text()
}

// Writer is the interface that implements the WriteMessage method.
type Writer interface {
  WriteMessage(msg Message) (n int, err error)
}

// WriteMessage writes the contents of the message msg to w, which accepts a
// slice of bytes. If w implements a WriteMessage method, it is invoked
// directly. Otherwise, w.writeRune is called exactly once.
//
// WriteMessage returns the number of lines printed to the terminal and an
// error (if any).
// func WriteMessage(w bufio.Writer, msg Message) (n int, err error) {
//   if ok := w.(MessageWriter); ok {
//     return w.WriteMessage(msg)
//   }
//   return w.writeRune([]rune(msg.String()))
// }

// removeNewlines uses regular expressions to remove newlines from a string.
func removeNewlines(s string) string {
  rx := regexp.MustCompile(`\r?\n`)
  return rx.ReplaceAllString(s, "")
}

// log is the default logger.
var log = NewHistory()

// printer flags.
const (
  PR_NO_COLOR = 1 << iota
)

// printer is the default printer to the terminal.
type printer struct {
  // length is the total message length in the terminal.
  length int
  flags int

  // out is the io.Writer to use when printing messages.
  out io.Writer
}

// newPrinter creates a new printer to write to the terminal.
func newPrinter(out io.Writer, length int, flags int) *printer {
  return &printer{
    out: out,
    length: length,
    flags: flags,
  }
}

// pr is the default implementation of the printer.
var pr = newPrinter(Stdout, 80, 0)

// SetPrinterFlags sets the flags for the printer.
func SetPrinterFlags(flags int)  {
  pr.flags = flags
}

// SetPrintLength sets the maximum message length for the printer.
func (p *printer) SetPrintLength(l int) {
  pr.length = l
}

// SetOutput sets the output stream to use for the printer.
func (p *printer) SetOutput(out io.Writer)  {
  p.out = out
}

// WriteMessage is the implementation of the WriteMessage method.
func (p printer) WriteMessage(msg Message) (n int, err error) {
  // ll holds the line length of the message.
  var ll = -1
  var esc = rune(term.ESC)
  n = 1
  err = nil

  wr := bufio.NewWriter(p.out)
  defer wr.Flush()

  // Get the rune slice of the message.
  rmsg := []rune(msg.Format())

  // For cleanliness, we clear the lines before we print. This way, if we are
  // updating a log or if we are overwriting an existing log in the terminal,
  // we will have a clean line to work with.
  term.Clear()

  for i := 0; i < len(rmsg); i++ {
    switch rmsg[i] {
    case esc:
      // Flush the buffer up to this point.
      wr.Flush()

      // Write the escape code.
      wr.WriteRune(rmsg[i])

      // Write the ANSI color code that immediately follows the escape code.
      for j := i; rmsg[j] != 'm'; j++ {
        wr.WriteRune(rmsg[j])

        // Cut out the runes.
        i++
      }

      // Write the 'm' rune, completing the color code.
      wr.WriteRune(rmsg[i])

      // If PR_NO_COLOR is specified, we clear the buffer and reset it.
      if p.flags&PR_NO_COLOR != 0 {
        wr.Reset(p.out)
      }

    default:
      // If the line length equals the maximum length, start a new line.
      if ll > 0 && (ll + 1) % p.length == 0 {
        wr.WriteRune('\n')
        wr.WriteString(fmt.Sprintf("%c[K", term.ESC))
        n = n + 1;
      }

      // Write the rune.
      wr.WriteRune(rmsg[i])
      // Increment the line length.
      ll = ll + 1
    }
  }

  return
}
