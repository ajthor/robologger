package robologger

import (
  "fmt"
  "sync"
)

// History holds the message history of a log. The message history is a slice
// of printed messages, where each message has an "index" and a "length". The
// index of the message corresponds to the index in the History slice.
//
//     [0] First message.
//     [1] Second message.
//     [2] Third message.
//     etc.
//
// Modifying the history is accomplished through editing this slice.
type History struct {
  mu sync.Mutex
  messages []Message
}

// NewHistory returns a new, empty History object for use in the logger.
func NewHistory() *History {
  return new(History)
}

// Add, Remove, Update methods for interacting with the History.

// Add adds a message to the history.
func (h *History) Add(msg Message) {
  h.mu.Lock()
	defer h.mu.Unlock()
  h.messages = append(h.messages, msg)
}

// Remove removes a message from the history.
func (h *History) Remove(msg Message) {
  h.mu.Lock()
	defer h.mu.Unlock()

  // Implemented in a future update.
}

// Update updates a message in the history.
func (h *History) Update(msg Message) {
  h.mu.Lock()
	defer h.mu.Unlock()

  offset := h.getPrintOffset(msg)

  // If the message is not part of this log's history, we default to using the
  // last message in the log.
  if msg == nil {
    msg = h.Get(-1)
    offset = msg.getPrintLength()
  }

  ll := msg.getPrintLength()

  term.SaveCursorPosition()
  term.HideCursor()
  term.MoveToBeginning()
  term.MoveUp(offset)

  // Write the message.
  n, _ := pr.WriteMessage(msg)

  // We include a newline for all log messages.
  fmt.Print("\n")

  // If the new message length does not equal the previous line length, the log
  // is rewritten from the updated message down.
  if n != ll {
    msg.setPrintLength(n)

    index, _ := h.Find(msg)

    for i := index + 1; i < len(h.messages); i++ {
      pr.WriteMessage(h.messages[i])

      fmt.Print("\n")
    }
  }

  term.ShowCursor()
  term.RestoreCursorPosition()
}

// getPrintOffset returns the number of printed lines from the bottom of the
// log.
func (h History) getPrintOffset(msg Message) int {
  var offset int

  for i := len(h.messages) - 1; i >= 0; i-- {
    offset = offset + h.messages[i].getPrintLength()
    if msg == h.messages[i] {
      break
    }
  }

  return offset
}

// Get returns the message in the history referenced by the index.
func (h History) Get(index int) Message {
  if index < 0 {
    index = len(h.messages) + index // (index * -1)
  }

  if index > len(h.messages) {
    index = index % len(h.messages)
  }

  return h.messages[index]
}

// Find finds a message in the history and returns the pointer to that message.
// Returns nil if the message is not in the history.
func (h History) Find(msg Message) (int, Message) {
  for i, m := range h.messages {
    if msg == m {
      return i, m
    }
  }

  return -1, nil
}
