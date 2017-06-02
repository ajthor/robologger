package robologger

import (
  "fmt"
  "regexp"
)

type MessageType int

const (
  DEFAULT MessageType = iota
  FATAL
  ERROR
  WARN
  INFO
  DEBUG
  STATUS
  INPUT
)

type message interface {
  fmt.Stringer

  Append(arg string)
}

type Message struct {
  Type MessageType
  Format *string
  Args []interface{}
}

// The `String` function is an implementation of the `Stringer` interface. The
// function also parses the message text to remove any newlines in the log
// text. This is an opinionated choice that helps keep the log clean. If the
// message is long enough to warrant newlines, it would be better to just have
// multiple lines in the log.
func (m *Message) String() string {
  var spmsg string

  if m.Format != nil {
    spmsg = fmt.Sprintf(*m.Format, m.Args...)
  } else {
    spmsg = fmt.Sprint(m.Args...)
  }

  // Remove newlines from the args and from the format string.
  re := regexp.MustCompile(`\r?\n`)
  spmsg = re.ReplaceAllString(spmsg, "")

  return spmsg
}

// `NewMessage` returns a pointer to a new Message. The message can be passed
// into further logger functions if needed to perform updates on the log.
func NewMessage(t MessageType, format interface{}, args ...interface{}) *Message {
  msg := &Message{
    Type: t,
    Args: args,
  }

  switch f := format.(type) {
  case string:
    msg.Format = &f
  case *string:
    msg.Format = f
  }

  return msg
}

func (m *Message) Append(arg string) {
  m.Args = append(m.Args, arg)
}
