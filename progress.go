package robologger

import "fmt"

// ProgressMessage implements the Message interface.
type ProgressMessage struct {
  // printLength refers to how many lines it takes up on the screen.
  printLength int

  progress int

  a []interface{}
}

// NewProgressMessage returns a new Message.
func NewProgressMessage(a ...interface{}) *ProgressMessage {
  return &ProgressMessage{
    a: a,
  }
}

// String is the implementation of the io.Stringer interface.
func (pm ProgressMessage) String() string {
  return fmt.Sprint(pm.a...)
}

func (pm ProgressMessage) getPrintLength() (n int) {
  return pm.printLength
}

func (pm *ProgressMessage) setPrintLength(n int) {
  pm.printLength = n
}

func (pm ProgressMessage) Format() (fmsg string) {
  s := pm.String()

  // Remove newlines from the args and from the format string.
  s = removeNewlines(s)

  // Calculate the percent complete.
  var percent int

  percent = int(100 * (float32(pm.progress) / float32(100)))
  if percent > 100 {
    percent = 100
  }

  fmsg = fmt.Sprintf("  %3d%%  ", percent)
  bar := make([]rune, 40)

  // Modify the progress bar to show the 'progress' character in the bar.
  for i := range bar {
    if float32(pm.progress)/float32(100) > float32(i)/float32(40) {
      bar[i] = '='
    } else {
      bar[i] = ' '
    }
  }

  // Here, we append the carat symbol to the end of the progress bar if the bar
  // is not full.
  for i := range bar {
    if bar[i] == ' ' {
      bar[i] = '>'
      break
    }
  }

  // Append the progress bar to the formatted string.
  // fmsg = fmsg + fmt.Sprintf("%c[97m[%s]%c[0m", term.ESC, string(bar), term.ESC)
  fmsg = fmsg + Color(C_WHITE_FG) + "["
  fmsg = fmsg + string(bar)
  fmsg = fmsg + "]" + Color(C_RESET)

  // 21 is the width of all of the characters in the progress bar that are not
  // inclusive of the status. This allows for the color codes necessary to
  // print the progress bar in the nice color we have.
  // statusWidth := 80 - (40 + 21)
  // statusWidth := 19
  if len(s) > 29 {
    fmsg = fmsg + " " + string(s[:26]) + "..."
  } else {
    fmsg = fmsg + " " + s
  }

  // If the status message is longer than the width of the writer, we need to
  // truncate the output so that it all fits on one line.
  // fmsg = fmsg + fmt.Sprintf(" %-*s", statusWidth, s)
  // \033[0m

  return
}

// Progress prints a progress bar to the terminal.
func Progress(args ...interface{}) (func(int, ...interface{})) {
  // Create a blank message and add it to the history.
  msg := NewProgressMessage(args...)
  log.Add(msg)

  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  // We include a newline for all log messages.
  fmt.Print("\n")

  // This function is returned to the user as a callable closure which will
  // update the status bar.
  Update := func (progress int, args ...interface{}) {
    msg.progress = progress
    msg.a = args

    // Update the message in the log.
    log.Update(msg)
  }

  // We call the update function once to print the message to the log.
  Update(0, args...)

  return Update
}
