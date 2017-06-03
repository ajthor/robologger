package robologger

import (
  "fmt"

  . "github.com/logrusorgru/aurora"
)

type StatusType int

const (
  WAITING StatusType = iota
  OK
  DONE
  ERR
)

func (w *DefaultWriter) FormatStatus(arg interface{}) string {
  switch a := arg.(type) {
  case StatusType:
    switch a {
    case WAITING:
      return fmt.Sprintf(" %s", Gray("waiting"))
    case OK:
      return fmt.Sprintf(" %s", Green("ok"))
    case DONE:
      return fmt.Sprintf(" %s", Green("done"))
    case ERR:
      return fmt.Sprintf(" %s", Red("err"))
    }
  default:
    return fmt.Sprintf(" %s", a)
  }

  return ""
}

func (w *DefaultWriter) FormatProgress(progress int, arg interface{}) string {
  var prog string
  var percent int

  // The width is the width of the progress bar.
  width := 40
  // 21 is the width of all of the characters in the progress bar that are not
  // inclusive of the status. This allows for the color codes necessary to
  // print the progress bar in the nice color we have.
  statusWidth := w.MessageWidth - (width + 21)

  bar := make([]rune, width)
  // var bar string

  percent = int(100 * (float32(progress) / float32(100)))
  if percent > 100 {
    percent = 100
  }
  prog = fmt.Sprintf("  %3d%%  ", percent)
  // \033[97m

  // Modify the progress bar to show the 'progress' character in the bar.
  for i := range bar {
    if float32(progress)/float32(100) > float32(i)/float32(width) {
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
  prog = prog + fmt.Sprintf("%c[97m[%s]%c[0m", term.ESC, string(bar), term.ESC)

  fmsg := w.FormatStatus(arg)

  // If the status message is longer than the width of the writer, we need to
  // truncate the output so that it all fits on one line.
  prog = prog + fmt.Sprintf(" %-*s", statusWidth, fmsg)
  // \033[0m

  return prog
}

func (l *Logger) Status(s interface{}, arg interface{}) {
  var msg *Message
  // Get the message from the history.
  msg = l.Find(arg)
  if msg == nil {
    msg = l.Find(-1)
  }

  // Format the message.
  fmsg := l.Writer.FormatStatus(s)

  msg.Append(fmsg)

  // Modify the message in the log.
  l.Modify(msg)
}

func (l *Logger) Progress() (func(int, interface{})) {
  // Create a blank message and add it to the history.
  msg := NewMessage(PRINT, nil, "")
  l.Writer.WriteMessage(msg)

  // We include a newline for all log messages.
  fmt.Print("\n")

  l.history.PostMessage(msg)

  // This function is returned to the user as a callable closure which will
  // update the status bar.
  Update := func (progress int, arg interface{}) {
    fmsg := l.Writer.FormatProgress(progress, arg)

    msg.Args[0] = fmsg

    // Modify the message in the log.
    l.Modify(msg)
  }

  // We call the update function once to print the message to the log.
  Update(0, "")

  return Update
}


func Status(t StatusType, arg interface{}) {
  std.Status(t, arg)
}

func Progress() (func(int, interface{})) {
  return std.Progress()
}

func Update(ch chan int) {
  // std.Update(ch)
}

// func (l *Logger) Status(msg *Message, status string) {
//
//   // Locate the message in the log history and append the status to the message
//   // as an argument.
//   for _, v := range l.History.messages {
//     if msg == v {
//       msg.Append(fmt.Sprintf(" %s", status))
//     }
//   }
//
//   // Now, we need to rewrite the log from
// }
//
// // The Status function is used to update the status of the previous statement
// // by appending either "err" or "ok" to the end and changing the 'logLevel'
// // label if the result is 'falsey'. Changing the label does not call the panic
// // or os.Exit functions, but does provide a visual indicator of whether or not
// // the previous command exited properly. Any panic or exit functions will need
// // to be called manually.
// // 'Falsey' in the context of the Done function is any value which would
// // normally casue the program to stop. Falsey values in this context could be
// // an error that is not nil or a false boolean value.
// func (l *Logger) Status(msg *Message, args ...interface{}) (ok bool) {
//   // Get the last message from history.
//   msg := l.history[len(l.history) - 1]
//
//   // Remove any trailing newlines from the string.
//   last := len(msg.args) - 1
//   msg.args[last] = strings.TrimRight(msg.args[last].(string), "\n")
//
//   ok = true
//
//   // Run through all of the arguments and determine if any of them are falsey.
//   // If they are falsey, we change the `ok` variable to false.
//   for _, arg := range args {
//     switch t := arg.(type) {
//     case error:
//       ok = false
//     case bool:
//       if !t { ok = false }
//     default:
//       // fmt.Printf("type %T\n", t)
//       // fmt.Println(arg)
//     }
//   }
//
//   // Update the log message and modify the history.
//   if ok {
//     // msg.lvl = INFO
//     msg.args = append(msg.args, Sprintf(Green(" ok\n")))
//     // msg.append(Sprintf(Green(" ok\n")))
//   } else {
//     msg.lvl = ERROR
//     msg.args = append(msg.args, Sprintf(Red(" err\n")))
//     // msg.append(Sprintf(Red(" err\n")))
//   }
//
//   l.modify(-1, msg)
//
//   return
// }
//
// func Done(args ...interface{}) bool {
//   return std.Done(args...)
// }
