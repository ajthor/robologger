package robologger

import (
  "sync"
)

type history interface {
  PostMessage(m *Message)
}

type History struct {
  mu sync.Mutex
  messages []*Message
}

func NewHistory() *History {
  return &History{}
}

func (h *History) PostMessage(m *Message) {
  h.mu.Lock()
  defer h.mu.Unlock()

  // m.line, _ = term.GetCursorPosition()
  // term.MoveToBeginning()
  // term.Clear()

  h.addLine()

  m.line = 1

  h.messages = append(h.messages, m)
}

func (h *History) addLine() {
  historyLen := len(h.messages)
  if historyLen == 0 {
    return
  }

  last := h.messages[len(h.messages) - 1]
  last.line = last.line + 1
}
