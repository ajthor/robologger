package robologger

import "fmt"

type ColorType int

const (
  C_RESET ColorType = 1 << iota
  C_BOLD
  C_ITALIC
  C_UNDERLINE
  C_INVERSE
  C_STRIKETHROUGH
  C_BOLD_OFF
  C_ITALIC_OFF
  C_UNDERLINE_OFF
  C_INVERSE_OFF
  C_STRIKETHROUGH_OFF

  C_BLACK_FG
  C_RED_FG
  C_GREEN_FG
  C_YELLOW_FG
  C_BLUE_FG
  C_MAGENTA_FG
  C_CYAN_FG
  C_GRAY_FG
  C_DEFAULT_FG

  C_DARK_GRAY_FG
  C_LIGHT_RED_FG
  C_LIGHT_GREEN_FG
  C_LIGHT_YELLOW_FG
  C_LIGHT_BLUE_FG
  C_LIGHT_MAGENTA_FG
  C_LIGHT_CYAN_FG
  C_WHITE_FG

  C_BLACK_BG
  C_RED_BG
  C_GREEN_BG
  C_YELLOW_BG
  C_BLUE_BG
  C_MAGENTA_BG
  C_CYAN_BG
  C_WHITE_BG
  C_DEFAULT_BG
)

// String returns the ColorType as a string.
func (flags ColorType) String() string {
  s := Color(flags)
  return s
}

// Color returns the ANSI string corresponding to a ColorType.
func Color(flags ColorType) string {
  var s string

  for flags > 0 {
    switch {
    case flags&C_RESET != 0:
      s = s + fmt.Sprintf("%c[0m", term.ESC)
      flags ^= C_RESET
    case flags&C_BOLD != 0:
      s = s + fmt.Sprintf("%c[1m", term.ESC)
      flags ^= C_BOLD
    case flags&C_ITALIC != 0:
      s = s + fmt.Sprintf("%c[3m", term.ESC)
      flags ^= C_ITALIC
    case flags&C_UNDERLINE != 0:
      s = s + fmt.Sprintf("%c[4m", term.ESC)
      flags ^= C_UNDERLINE
    case flags&C_INVERSE != 0:
      s = s + fmt.Sprintf("%c[7m", term.ESC)
      flags ^= C_INVERSE
    case flags&C_STRIKETHROUGH != 0:
      s = s + fmt.Sprintf("%c[9m", term.ESC)
      flags ^= C_STRIKETHROUGH
    case flags&C_BOLD_OFF != 0:
      s = s + fmt.Sprintf("%c[22m", term.ESC)
      flags ^= C_BOLD_OFF
    case flags&C_ITALIC_OFF != 0:
      s = s + fmt.Sprintf("%c[23m", term.ESC)
      flags ^= C_ITALIC_OFF
    case flags&C_UNDERLINE_OFF != 0:
      s = s + fmt.Sprintf("%c[24m", term.ESC)
      flags ^= C_UNDERLINE_OFF
    case flags&C_INVERSE_OFF != 0:
      s = s + fmt.Sprintf("%c[27m", term.ESC)
      flags ^= C_INVERSE_OFF
    case flags&C_STRIKETHROUGH_OFF != 0:
      s = s + fmt.Sprintf("%c[29m", term.ESC)
      flags ^= C_STRIKETHROUGH_OFF
    // Foreground colors.
    case flags&C_BLACK_FG != 0:
      s = s + fmt.Sprintf("%c[30m", term.ESC)
      flags ^= C_BLACK_FG
    case flags&C_RED_FG != 0:
      s = s + fmt.Sprintf("%c[31m", term.ESC)
      flags ^= C_RED_FG
    case flags&C_GREEN_FG != 0:
      s = s + fmt.Sprintf("%c[32m", term.ESC)
      flags ^= C_GREEN_FG
    case flags&C_YELLOW_FG != 0:
      s = s + fmt.Sprintf("%c[33m", term.ESC)
      flags ^= C_YELLOW_FG
    case flags&C_BLUE_FG != 0:
      s = s + fmt.Sprintf("%c[34m", term.ESC)
      flags ^= C_BLUE_FG
    case flags&C_MAGENTA_FG != 0:
      s = s + fmt.Sprintf("%c[35m", term.ESC)
      flags ^= C_MAGENTA_FG
    case flags&C_CYAN_FG != 0:
      s = s + fmt.Sprintf("%c[36m", term.ESC)
      flags ^= C_CYAN_FG
    case flags&C_GRAY_FG != 0:
      s = s + fmt.Sprintf("%c[37m", term.ESC)
      flags ^= C_GRAY_FG
    case flags&C_DEFAULT_FG != 0:
      s = s + fmt.Sprintf("%c[39m", term.ESC)
      flags ^= C_DEFAULT_FG
    // Light foreground colors.
  case flags&C_DARK_GRAY_FG != 0:
      s = s + fmt.Sprintf("%c[90m", term.ESC)
      flags ^= C_DARK_GRAY_FG
    case flags&C_LIGHT_RED_FG != 0:
        s = s + fmt.Sprintf("%c[91m", term.ESC)
        flags ^= C_LIGHT_RED_FG
    case flags&C_LIGHT_GREEN_FG != 0:
      s = s + fmt.Sprintf("%c[92m", term.ESC)
      flags ^= C_LIGHT_GREEN_FG
    case flags&C_LIGHT_YELLOW_FG != 0:
      s = s + fmt.Sprintf("%c[93m", term.ESC)
      flags ^= C_LIGHT_YELLOW_FG
    case flags&C_LIGHT_BLUE_FG != 0:
      s = s + fmt.Sprintf("%c[94m", term.ESC)
      flags ^= C_LIGHT_BLUE_FG
    case flags&C_LIGHT_MAGENTA_FG != 0:
      s = s + fmt.Sprintf("%c[95m", term.ESC)
      flags ^= C_LIGHT_MAGENTA_FG
    case flags&C_LIGHT_CYAN_FG != 0:
      s = s + fmt.Sprintf("%c[96m", term.ESC)
      flags ^= C_LIGHT_CYAN_FG
    case flags&C_WHITE_FG != 0:
      s = s + fmt.Sprintf("%c[97m", term.ESC)
      flags ^= C_WHITE_FG
    // Background colors.
    case flags&C_BLACK_BG != 0:
      s = s + fmt.Sprintf("%c[40m", term.ESC)
      flags ^= C_BLACK_BG
    case flags&C_RED_BG != 0:
      s = s + fmt.Sprintf("%c[41m", term.ESC)
      flags ^= C_RED_BG
    case flags&C_GREEN_BG != 0:
      s = s + fmt.Sprintf("%c[42m", term.ESC)
      flags ^= C_GREEN_BG
    case flags&C_YELLOW_BG != 0:
      s = s + fmt.Sprintf("%c[43m", term.ESC)
      flags ^= C_YELLOW_BG
    case flags&C_BLUE_BG != 0:
      s = s + fmt.Sprintf("%c[44m", term.ESC)
      flags ^= C_BLUE_BG
    case flags&C_MAGENTA_BG != 0:
      s = s + fmt.Sprintf("%c[45m", term.ESC)
      flags ^= C_MAGENTA_BG
    case flags&C_CYAN_BG != 0:
      s = s + fmt.Sprintf("%c[46m", term.ESC)
      flags ^= C_CYAN_BG
    case flags&C_WHITE_BG != 0:
      s = s + fmt.Sprintf("%c[47m", term.ESC)
      flags ^= C_WHITE_BG
    case flags&C_DEFAULT_BG != 0:
      s = s + fmt.Sprintf("%c[49m", term.ESC)
      flags ^= C_DEFAULT_BG
    }
  }

  return s
}

// IsColor returns the length of the color string.
func IsColor(r []rune) bool {

  return false
}
