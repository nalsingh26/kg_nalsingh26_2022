package main

import (
  "fmt"
  "os"
)

var charNum = map[rune]string{
  48: "Zero",
  49: "One",
  50: "Two",
  51: "Three",
  52: "Four",
  53: "Five",
  54: "Six",
  55: "Seven",
  56: "Eight",
  57: "Nine",
}

func main() {
  nums_string := os.Args[1:]
  result := ""
  for _, num := range nums_string {
    for _, char := range num {
      result += charNum[char]
    }
    result += ","
  }
  fmt.Println(result[:len(result)-1])
}
