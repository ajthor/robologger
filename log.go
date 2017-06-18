package robologger

import "fmt"
import "os"

// LogFlag defines the log message flags.
type LogFlag int

const (
  L_PRINT LogFlag = 1 << iota
  L_FATAL
  L_ERROR
  L_WARN
  L_INFO
  L_DEBUG
)

// LogMessage implements the Message interface.
type LogMessage struct {
  // printLength refers to how many lines it takes up on the screen.
  printLength int

  flags LogFlag

  format *string
  a []interface{}
}

// NewLogMessage returns a new Message.
func NewLogMessage(flags LogFlag, format *string, a ...interface{}) *LogMessage {
  return &LogMessage{
    flags: flags,
    format: format,
    a: a,
  }
}

// String is the implementation of the io.Stringer interface.
func (lm LogMessage) String() string {
  switch lm.format {
  case nil:
    return fmt.Sprint(lm.a...)
  default:
    return fmt.Sprintf(*lm.format, lm.a...)
  }
}

func (lm LogMessage) getPrintLength() (n int) {
  return lm.printLength
}

func (lm *LogMessage) setPrintLength(n int) {
  lm.printLength = n
}

// FormatMessage adds a prefix to each line of a particular message type. If
// the message has a type other than L_PRINT, we will have a colored prefix.
func (lm LogMessage) Format() (fmsg string) {
  s := lm.String()

  // Remove newlines from the args and from the format string.
  s = removeNewlines(s)

  // Preferentially apply prefixes to the log messages based upon the severity
  // of the message.
  var prefix string

  switch {
  case lm.flags&L_PRINT != 0:
    prefix = "        "
  case lm.flags&L_FATAL != 0:
    prefix = Color(C_RED_FG) + "[FATAL] " + Color(C_RESET)
  case lm.flags&L_ERROR != 0:
    prefix = Color(C_RED_FG) + "[ERROR] " + Color(C_RESET)
  case lm.flags&L_WARN != 0:
    prefix = Color(C_YELLOW_FG) + "[WARN]  " + Color(C_RESET)
  case lm.flags&L_INFO != 0:
    prefix = Color(C_GREEN_FG) + "[INFO]  " + Color(C_RESET)
  case lm.flags&L_DEBUG != 0:
    prefix = Color(C_CYAN_FG) + "[DEBUG] " + Color(C_RESET)
  }

  fmsg = prefix + s
  return
}

// The following functions are the standard functions that print messages to
// the terminal without any formatting.
func Print(args ...interface{}) {
  msg := NewLogMessage(L_PRINT, nil, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Fatal(args ...interface{}) {
  msg := NewLogMessage(L_FATAL, nil, args...)
  log.Add(msg)

  pr.SetOutput(Stderr)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
  os.Exit(1)
}

func Error(args ...interface{}) {
  msg := NewLogMessage(L_ERROR, nil, args...)
  log.Add(msg)

  pr.SetOutput(Stderr)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
  panic(fmt.Sprint(args...))
}

func Warn(args ...interface{}) {
  msg := NewLogMessage(L_WARN, nil, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Info(args ...interface{}) {
  msg := NewLogMessage(L_INFO, nil, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Debug(args ...interface{}) {
  msg := NewLogMessage(L_DEBUG, nil, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

// The following functions are the "formatted" functions that print messages
// using a format string.
func Printf(format string, args ...interface{}) {
  msg := NewLogMessage(L_PRINT, &format, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Fatalf(format string, args ...interface{}) {
  msg := NewLogMessage(L_FATAL, &format, args...)
  log.Add(msg)

  pr.SetOutput(Stderr)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
  os.Exit(1)
}

func Errorf(format string, args ...interface{}) {
  msg := NewLogMessage(L_ERROR, &format, args...)
  log.Add(msg)

  pr.SetOutput(Stderr)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
  panic(fmt.Sprintf(format, args...))
}

func Warnf(format string, args ...interface{}) {
  msg := NewLogMessage(L_WARN, &format, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Infof(format string, args ...interface{}) {
  msg := NewLogMessage(L_INFO, &format, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}

func Debugf(format string, args ...interface{}) {
  msg := NewLogMessage(L_DEBUG, &format, args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")
}
