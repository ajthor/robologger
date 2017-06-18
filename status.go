package robologger

import "fmt"

// StatusMessage implements the Message interface.
type StatusMessage struct {
  // printLength refers to how many lines it takes up on the screen.
  printLength int

  symbol rune

  format *string
  a []interface{}
}

// NewStatusMessage returns a new Message.
func NewStatusMessage(format *string, a ...interface{}) *StatusMessage {
  return &StatusMessage{
    format: format,
    a: a,
  }
}

// String is the implementation of the io.Stringer interface.
func (sm StatusMessage) String() string {
  switch sm.format {
  case nil:
    return fmt.Sprint(sm.a...)
  default:
    return fmt.Sprintf(*sm.format, sm.a...)
  }
}

func (sm StatusMessage) getPrintLength() (n int) {
  return sm.printLength
}

func (sm *StatusMessage) setPrintLength(n int) {
  sm.printLength = n
}

func (sm StatusMessage) Format() (fmsg string) {
  s := sm.String()

  // Remove newlines from the args and from the format string.
  s = removeNewlines(s)

  // s = fmt.Sprintf("%c[90m%s", term.ESC, s) + Color(C_RESET)
  // if len(s) > 80 {
  //   fmsg = string(s[:77]) + "..."
  // } else {
  //   fmsg = s
  // }

  return s
}

func Status(args ...interface{}) (func(...interface{})) {
  msg := NewStatusMessage(nil, args...)
  log.Add(msg)

  // Do write stuff here.
  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  // This function is returned to the user as a callable closure which will
  // update the status bar.
  Update := func (args ...interface{}) {
    msg.a = args

    // Update the message in the log.
    log.Update(msg)
  }

  return Update
}

func Statusf(format string, args ...interface{}) (func(...interface{})) {
  msg := NewStatusMessage(&format, args...)
  log.Add(msg)

  // Do write stuff here.
  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  fmt.Print("\n")

  // This function is returned to the user as a callable closure which will
  // update the status bar.
  Update := func (args ...interface{}) {
    msg.a = args

    // Update the message in the log.
    log.Update(msg)
  }

  return Update
}
