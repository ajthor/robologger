package robologger

import "fmt"

func ExampleUnicode() {
  fmt.Print("🤖")
  // Output:
  // 🤖
}

func ExampleColors() {
  fg := 38
  for color := 0; color < 16; color++ {
    s := fmt.Sprintf("%d;5;%d", fg, color)
    fmt.Printf("%c[%sm %3d%c[0m", 27, s, color, 27)
    if color > 0 && (color + 1) % 36 == 0 {
      fmt.Print("\n")
    }
  }

  fmt.Print("\n")

  for color := 0; color < 240; color++ {
    s := fmt.Sprintf("%d;5;%d", fg, color + 16)
    fmt.Printf("%c[%sm %3d%c[0m", 27, s, color + 16, 27)
    if color > 0 && (color + 1) % 36 == 0 {
      fmt.Print("\n")
    }
  }

  fmt.Print("\n")

  bg := 48
  for color := 0; color < 16; color++ {
    s := fmt.Sprintf("%d;5;%d", bg, color)
    fmt.Printf("%c[%sm %3d%c[0m", 27, s, color, 27)
    if color > 0 && (color + 1) % 36 == 0 {
      fmt.Print("\n")
    }
  }

  fmt.Print("\n")

  for color := 0; color < 240; color++ {
    s := fmt.Sprintf("%d;5;%d", bg, color + 16)
    fmt.Printf("%c[%sm %3d%c[0m", 27, s, color + 16, 27)
    if color > 0 && (color + 1) % 36 == 0 {
      fmt.Print("\n")
    }
  }

  fmt.Print("\n")
}
