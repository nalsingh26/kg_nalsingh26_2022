package main

import (
  "fmt"
  "os"
)

func main() {
  // creating character to string map
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
  // taking inline arguments
  nums_string := os.Args[1:]
  // creating an empty result string
  result := ""
  //iterating over the array of numbers
  for _, num := range nums_string {
    //iterating over digits of a number
    for _, char := range num {
      result += charNum[char]
    }
    result += ","
  }
  // printing the output
  fmt.Println(result[:len(result)-1])
}
