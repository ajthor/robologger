package robologger

import (
  "fmt"
  "os"
)

type logger interface {
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

// The following functions are the standard functions for interacting with the
// logger. We have functions for each message type, and special conditions when
// we have an error.
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
