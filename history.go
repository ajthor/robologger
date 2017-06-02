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

  h.messages = append(h.messages, m)
}
