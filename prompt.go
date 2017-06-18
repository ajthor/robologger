package robologger

import "fmt"

// PromptFlag defines the prompt message flags.
type PromptFlag int

const (
  P_STRING PromptFlag = 1 << iota
  P_YES
  P_NO
  P_CANCEL
  P_ALL
)

// PromptMessage implements the Message interface.
type PromptMessage struct {
  // printLength refers to how many lines it takes up on the screen.
  printLength int

  flags PromptFlag

  format *string
  a []interface{}
}

// NewPromptMessage returns a new Message.
func NewPromptMessage(flags PromptFlag, format *string, a ...interface{}) *PromptMessage {
  return &PromptMessage{
    flags: flags,
    format: format,
    a: a,
  }
}

// String is the implementation of the io.Stringer interface.
func (pm PromptMessage) String() string {
  switch pm.format {
  case nil:
    return fmt.Sprint(pm.a...)
  default:
    return fmt.Sprintf(*pm.format, pm.a...)
  }
}

func (pm PromptMessage) getPrintLength() (n int) {
  return pm.printLength
}

func (pm *PromptMessage) setPrintLength(n int) {
  pm.printLength = n
}

func (pm PromptMessage) Format() (fmsg string) {
  flags := pm.flags

  s := pm.String()

  // Remove newlines from the args and from the format string.
  s = removeNewlines(s)

  // Add choices to the prompt.
  switch {
  case flags&P_STRING != 0:
    s = s + ": "
  default:
    s = s + " [" + Color(C_YELLOW_FG)

    for flags > 0 {
      switch {
      case flags&P_STRING != 0:
        flags ^= P_STRING
      case flags&P_YES != 0:
        s = s + "y"
        flags ^= P_YES
      case flags&P_NO != 0:
        s = s + "N"
        flags ^= P_NO
      case flags&P_CANCEL != 0:
        s = s + "c"
        flags ^= P_CANCEL
      case flags&P_ALL != 0:
        s = s + "a"
        flags ^= P_ALL
      }
    }

    s = s + Color(C_RESET) + "] "
  }

  return s
}

func Prompt(flags PromptFlag, args ...interface{}) (string, Message) {
  msg := NewPromptMessage(flags, nil, args...)
  log.Add(msg)

  // Do write stuff here.
  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  // Do read stuff here.
  res := ScanLine()

  return res, msg
}

func Promptf(flags PromptFlag, format string, args ...interface{}) (string, Message) {
  msg := NewPromptMessage(flags, &format, args...)
  log.Add(msg)

  // Do write stuff here.
  n, _ := pr.WriteMessage(msg)
  msg.printLength = n

  // Do read stuff here.
  res := ScanLine()

  return res, msg
}

// //
// // type DefaultReader struct {}
// //
// // type ReaderOptions struct {
// //   MessageWidth int
// // }
// //
// // func NewDefaultReader(opts ReaderOptions) *DefaultReader {
// //   return &DefaultReader{}
// // }
// //
// // // `ParseResponse` takes the input read from `ReadInput` and converts it into
// // // a `ResponseType`. This makes checking for a particular response easier.
// // // NOTE: Remove the error return from this function, potentially? Would not be
// // // calling this if the user types in a string, but this helps if the user types
// // // in nonsense.
// // func (r *DefaultReader) ParseResponse(response string) (ResponseType, error) {
// //   switch response {
// //   case "y", "Y", "yes", "Yes":
// //     return YES, nil
// //   case "n", "N", "no", "No":
// //     return NO, nil
// //   case "c", "C", "cancel", "Cancel":
// //     return CANCEL, nil
// //   case "":
// //     return EMPTY, nil
// //   default:
// //     return EMPTY, fmt.Errorf("Cannot parse response: %s", response)
// //   }
// // }
// //
// // // `ReadInput` uses a bufio scanner instead of fmt.Scan functions in order to
// // // read a single line of text from the user. This may change in the future, but
// // // for now, it works out well.
// // func (r *DefaultReader) ReadInput() string {
// //   scanner := bufio.NewScanner(os.Stdin)
// //   scanner.Scan()
// //   return scanner.Text()
// // }
// //
// // func (l *Logger) prompt(p interface{}, format *string, args ...interface{}) string {
// //   // Create a new message.
// //   msg := NewMessage(PRINT, format, args...)
// //
// //   // Format the message as a prompt.
// //   fmsg := l.Writer.FormatPrompt(p)
// //
// //   msg.Append(fmsg)
// //
// //   // Write the message.
// //   l.Writer.WriteMessage(msg)
// //
// //   // Get the input from the user.
// //   response := l.Reader.ReadInput()
// //
// //   // NOTE: ONLY do this if the type of the message is not a string or a
// //   // password. Perhaps we need to include a default value mechanism?
// //   if response == "" {
// //     response = "n"
// //   }
// //   // We need to move up one line here to compensate for the new line entered by
// //   // the user.
// //   term.MoveToBeginning()
// //   term.MoveUp(1)
// //   msg.lineLength = msg.lineLength + 1
// //
// //   // msg.Append(Cyan(response))
// //
// //   l.Writer.WriteMessage(msg)
// //   // l.Modify(msg)
// //
// //   fmt.Printf("%s", Cyan(response))
// //
// //   // We include a newline for all log messages.
// //   fmt.Print("\n")
// //
// //
// //   // term.SaveCursorPosition()
// //   // term.MoveToBeginning()
// //   // term.MoveUp(1)
// //   //
// //   // term.Clear()
// //   //
// //   // l.Writer.WriteMessage(msg)
// //   //
// //   // fmt.Printf("%s", Cyan(response))
// //   // fmt.Print("\n")
// //   //
// //   // term.RestoreCursorPosition()
// //
// //   return response
// // }
// //
// // // These two functions are part of the `Logger` interface. We include them here
// // // because they have special functions which relate closely to the `Reader`
// // // interface.
// // func (l *Logger) Prompt(t PromptType, args ...interface{}) string {
// //   return l.prompt(t, nil, args...)
// // }
// //
// // func (l *Logger) Promptf(t PromptType, format string, args ...interface{}) string {
// //   return l.prompt(t, &format, args...)
// // }
// //
// // // func (r *DefaultReader) ReadInputMultiline() (lines []string) {
// // //   scanner := bufio.NewScanner(os.Stdin)
// // //   for scanner.Scan() {
// // //     lines = append(lines, scanner.Text())
// // //   }
// // //   return
// // // }
// //
// // func Prompt(t PromptType, args ...interface{}) string {
// //   return std.Prompt(t, args...)
// // }
// //
// // func Promptf(t PromptType, format string, args ...interface{}) string {
// //   return std.Promptf(t, format, args...)
// // }
// //
// // // These last two declarations are utility functions for working with the
// // // reader. If a user wants to read input for their own purposes, for example,
// // // they can interact with these two functions from the std logger.
// // func ParseResponse(response string) (ResponseType, error) {
// //   return std.Reader.ParseResponse(response)
// // }
// //
// // func ReadInput() string {
// //   return std.Reader.ReadInput()
// // }
// //
// // // ---------------------------------------
// //
// // // PromptFormatter is the standard formatter for prompts.
// // type PromptFormatter struct {}
// //
// // // NewPromptFormatter returns a new instance of the PromptFormatter.
// // func NewPromptFormatter() *PromptFormatter {
// //   return new(PromptFormatter)
// // }
// //
// // func (pf PromptFormatter) Format(p []byte, flags int) (n int, fmsg []byte) {
// //   // Apply all flags to the prompt.
// //   for flags > 0 {
// //     switch {
// //     case flags&R_STRING != 0:
// //       flags ^= R_STRING
// //     case flags&R_YES != 0:
// //       flags ^= R_YES
// //     case flags&R_NO != 0:
// //       flags ^= R_NO
// //     case flags&R_CANCEL != 0:
// //       flags ^= R_CANCEL
// //     case flags&R_ALL != 0:
// //       flags ^= R_ALL
// //     }
// //   }
// // }
// //
// // type PromptReader interface {
// //   io.Reader
// // }
// //
// // func NewPromptReader() *PromptReader {
// //   return new(PromptReader)
// // }
// //
// // func (pr PromptReader) Read(p []byte) (n int, err error) {
// //
// // }
