package robologger

import (
  "fmt"
  "os"
)

type logger interface {
  // These functions are wrappers for the `fmt` functions, but manipulate the
  // history of the logger.
  Print(args ...interface{}) *Message
  // Println(args ...interface{}) *Message
  Printf(format *string, args ...interface{}) *Message
  // The following functions add a prefix to the messages that are printed to
  // the terminal.
  Fatal(args ...interface{}) *Message
  Fatalf(format *string, args ...interface{}) *Message
  Error(args ...interface{}) *Message
  Errorf(format *string, args ...interface{}) *Message
  Warn(args ...interface{}) *Message
  Warnf(format *string, args ...interface{}) *Message
  Info(args ...interface{}) *Message
  Infof(format *string, args ...interface{}) *Message
  Debug(args ...interface{}) *Message
  Debugf(format *string, args ...interface{}) *Message

  Find(arg interface{}) (int, *Message)
  // Definitions for `Status` are found in `status.go`
  Status()
  // Definitions for `Prompt` and `Promptf` are found in `reader.go`
  Prompt(t PromptType, args ...interface{})
  Promptf(t PromptType, format *string, args ...interface{})
  DisplayInput(input string)
}

type Logger struct {
  LogLevel MessageType
  history *History
  Reader Reader
  Writer Writer
}

// Creates a new instance of the logger.
func New(level MessageType) *Logger {
  r := NewDefaultReader(ReaderOptions{})
  w := NewDefaultWriter(WriterOptions{
    MessageWidth: 80,
  })

  h := NewHistory()

  return &Logger{
    LogLevel: level,
    history: h,
    Reader: r,
    Writer: w,
  }
}

// This function is called as a proxy for every standard function in the logger
// interface. It creates a message, adds it to the history, and then tells the
// writer to write the message.
func (l *Logger) log(t MessageType, format *string, args ...interface{}) *Message {
  // Create a new message and add it to the history.
  msg := NewMessage(t, format, args...)

  l.history.PostMessage(msg)

  // Write the message.
  l.Writer.WriteMessage(msg)

  // We include a newline for all log messages.
  fmt.Print("\n")

  return msg
}

// The `Find` function accepts either an integer, representing an "offset" from
// the last log item to be posted to the log history, or a pointer to a
// message, returned from another command which prints a message to the
// terminal.
// It then returns a pointer to a message in the log, and can confirm that a message passed to the function is in the log, or
func (l *Logger) Find(arg interface{}) *Message {
  msgs := l.history.messages

  switch a := arg.(type) {
  case *Message:
    // If we are passed a pointer to a message, we can confirm that the message
    // is in the log.
    for _, m := range msgs {
      if a == m {
        return m
      }
    }
  case int:
    // If we have a negative value, we need to return the difference between
    // the history length and the value passed.
    if a < 0 {
      a = len(msgs) + a // (a * -1)
    }

    // If the index is out of bounds, we return nil.
    if a < 0 || a >= len(msgs) {
      return nil
    }

    // Otherwise, we return the message at the index.
    return msgs[a]
  default:
    return nil
  }

  return nil
}

// The `Modify` function updates the log messages that have previously been printed to the log. By passing either a pointer to a message or an integer representing the index of the message,
func (l *Logger) Modify(arg interface{}) {
  var offset, index int
  var msg *Message

  msgs := l.history.messages

  // Count the number of lines from the bottom of the log.
  switch a := arg.(type) {
  case *Message:
    offset = 0
    msg = a
    for i := len(msgs) - 1; i >= 0; i-- {
      offset = offset + msgs[i].lineLength
      if a == msgs[i] {
        index = i
        break
      }
    }
  case int:
    index = a
    for i := 0; i < len(msgs); i++ {
      offset = offset + msgs[i].lineLength
      if a == i {
        msg = msgs[i]
        break
      }
    }
  default:
    offset = 1
  }

  // If we cannot find the message, we default to using the last message in the
  // log. This way, we can still continue, even if the user passes an unknown
  // value for the arguments.
  if msg == nil {
    offset = 1
  }

  ll := msg.lineLength

  term.SaveCursorPosition()
  term.MoveToBeginning()
  term.MoveUp(offset)

  // Write the message.
  l.Writer.WriteMessage(msg)

  // We include a newline for all log messages.
  fmt.Print("\n")

  // If the new message length is longer than the previous line length, we need
  // to rewrite the log from here down, effectively adding a new line in the
  // middle of the log.
  if msg.lineLength > ll {
    for i := index + 1; i < len(l.history.messages); i++ {
      l.Writer.WriteMessage(l.history.messages[i])
    }
  }

  term.RestoreCursorPosition()
}

// The following functions are the standard functions for interacting with the
// logger. We have functions for each message type, and special conditions when
// we have an error.
func (l *Logger) Print(args ...interface{}) (msg *Message) {
  msg = l.log(PRINT, nil, args...)
  return
}

func (l *Logger) Printf(format string, args ...interface{}) (msg *Message) {
  msg = l.log(PRINT, &format, args...)
  return
}

func (l *Logger) Fatal(args ...interface{}) (msg *Message) {
  msg = l.log(FATAL, nil, args...)
  os.Exit(1)
  return
}

func (l *Logger) Fatalf(format string, args ...interface{}) (msg *Message) {
  msg = l.log(FATAL, &format, args...)
  os.Exit(1)
  return
}

func (l *Logger) Error(args ...interface{}) (msg *Message) {
  msg = l.log(ERROR, nil, args...)
  panic(fmt.Sprint(args...))
  return
}

func (l *Logger) Errorf(format string, args ...interface{}) (msg *Message) {
  msg = l.log(ERROR, &format, args...)
  panic(fmt.Sprintf(format, args...))
  return
}

func (l *Logger) Warn(args ...interface{}) (msg *Message) {
  msg = l.log(WARN, nil, args...)
  return
}

func (l *Logger) Warnf(format string, args ...interface{}) (msg *Message) {
  msg = l.log(WARN, &format, args...)
  return
}

func (l *Logger) Info(args ...interface{}) (msg *Message) {
  msg = l.log(INFO, nil, args...)
  return
}

func (l *Logger) Infof(format string, args ...interface{}) (msg *Message) {
  msg = l.log(INFO, &format, args...)
  return
}

func (l *Logger) Debug(args ...interface{}) (msg *Message) {
  msg = l.log(DEBUG, nil, args...)
  return
}

func (l *Logger) Debugf(format string, args ...interface{}) (msg *Message) {
  msg = l.log(DEBUG, &format, args...)
  return
}

// Global declaration of a logger. This logger is available through the exposed
// methods founds below.
var std = New(DEBUG)

// Methods for calling the default 'std' logger. These are exported and
// available for use by importing the package in another file. This effectively
// makes the logger global.
func Print(args ...interface{}) (msg *Message) {
  return std.Print(args...)
}

func Printf(format string, args ...interface{}) (msg *Message) {
  return std.Printf(format, args...)
}

func Fatal(args ...interface{}) (msg *Message) {
  return std.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) (msg *Message) {
  return std.Fatalf(format, args...)
}

func Error(args ...interface{}) (msg *Message) {
  return std.Error(args...)
}

func Errorf(format string, args ...interface{}) (msg *Message) {
  return std.Errorf(format, args...)
}

func Warn(args ...interface{}) (msg *Message) {
  return std.Warn(args...)
}

func Warnf(format string, args ...interface{}) (msg *Message) {
  return std.Warnf(format, args...)
}

func Info(args ...interface{}) (msg *Message) {
  return std.Info(args...)
}

func Infof(format string, args ...interface{}) (msg *Message) {
  return std.Infof(format, args...)
}

func Debug(args ...interface{}) (msg *Message) {
  return std.Debug(args...)
}

func Debugf(format string, args ...interface{}) (msg *Message) {
  return std.Debugf(format, args...)
}
