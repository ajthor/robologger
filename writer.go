package robologger

import (
  "fmt"
//  "strings"

  . "github.com/logrusorgru/aurora"
)

type Writer interface {
  FormatMessage(m *Message) string
  WriteMessage(m *Message)
}

type DefaultWriter struct {
  MessageWidth int
}

type WriterOptions struct {
  MessageWidth int
}

func NewDefaultWriter(opts WriterOptions) *DefaultWriter {
  return &DefaultWriter{
    MessageWidth: opts.MessageWidth,
  }
}

// The `FormatMessage` function adds a prefix to each line of a particular
// message type. If the message has a type other than DEFAULT, we will have a
// colored prefix.
func (w *DefaultWriter) FormatMessage(m *Message) string {
  spmsg := m.String()

  switch m.Type {
  case FATAL:
    return fmt.Sprintf("%s %s", Red("[FATAL]"), spmsg)
  case ERROR:
    return fmt.Sprintf("%s %s", Red("[ERROR]"), spmsg)
  case WARN:
    return fmt.Sprintf("%s  %s", Brown("[WARN]"), spmsg)
  case INFO:
    return fmt.Sprintf("%s  %s", Green("[INFO]"), spmsg)
  case DEBUG:
    return fmt.Sprintf("%s %s", Cyan("[DEBUG]"), spmsg)
  default:
    return spmsg
  }
}

// The `WriteMessage` function prints the message to the terminal. We
// intrinsically have two different types of messages: messages with a prefix
// and those without. If we have a prefix, we display the message type, which
// corresponds to the log level, and then we print out the message. However,
// messages without a prefix are inherently 8 characters shorter than their
// counterparts. Since we are breaking the log at 80 characters (just
// convention), we need to take the relative length of the messages into
// account.
func (w *DefaultWriter) WriteMessage(m *Message) {
  var n int = w.MessageWidth
  // if m.Type == DEFAULT { n = 80 }

  fmsg := w.FormatMessage(m)
  rmsg := []rune(fmsg)

  for i, r := range rmsg {
    fmt.Print(string(r))
    if i > 0 && (i + 1) % n == 0 {
      n = w.MessageWidth - 8
      fmt.Print("\n")
      if m.Type != DEFAULT { fmt.Print("        ") }
    }
  }
}

// func (w *DefaultWriter) WriteMessage(m *Message) {
//   var n int = 72
//   if m.Type == DEFAULT { n = 80 }
//
//   fmsg := w.FormatMessage(m)
//   words := strings.Fields(fmsg)
//
//   line := []rune{}
//
//   // printLine := func(l []rune) {
//   //   fmt.Print(string(line))
//   //   fmt.Print("\n")
//   //   if m.Type != DEFAULT { fmt.Print("        ") }
//   // }
//
//   for _, word := range words {
//     rword := []rune(word)
//     rword = append(rword, ' ')
//     lineLen := len(line)
//     lenWord := len(word)
//     // fmt.Printf("LL: %d, FL: %d", lineLen, lenField)
//
//     if (lineLen + lenWord) < n {
//       for _, r := range rword {
//         line = append(line, r)
//       }
//     } else {
//       fmt.Print(string(line))
//       fmt.Print("\n")
//       if m.Type != DEFAULT { fmt.Print("        ") }
//       line = []rune{}
//       for _, r := range rword {
//         line = append(line, r)
//       }
//     }
//
//   }
//
//   fmt.Print(string(line))
// }



// func (w *TerminalWriter) Update(m *Message)  {
//
// }
//
//
// func (l *Logger) update(m *Message) {
//   histLen := len(l.history)
//   var offset int
//
//   // offset = histLen - offset
//   if index < 0 {
//     offset = -1*index
//   } else {
//     offset = histLen - index
//   }
//
//   for _, v := range l.History {
//     if m == v {
//
//     }
//   }
//
//
//   l.SaveCursorPosition()
//   l.MoveToBeginning()
//   l.MoveUp(offset)
//
//   l.Clear()
//
//   rmsg := l.format(msg)
//   fmt.Printf(rmsg)
//
//   l.RestoreCursorPosition()
// }
