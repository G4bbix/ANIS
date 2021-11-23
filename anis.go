package main

import (
  "flag"
  "fmt"
  "os"
  "strconv"
  "strings"
  "unicode"
)

func printUsage() {
  fmt.Println("Usage:")
  fmt.Println("./anis (-s) Operations INPUT")
  fmt.Println("-s: Treat every digit as single number")
  fmt.Println("Operations: (*N | /N | +N | -N)")
  fmt.Println("Input: String containing numbers")
}

func inveseString(input string) string {
  runes := []rune(input)
  for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

func calcWrapper(intStr string, Operations []string) string {
  intStrRev := inveseString(intStr)
  IntToCalc, _ := strconv.ParseFloat(intStrRev, 32)
  Result := doCalculatios(IntToCalc, Operations)
  Result = cutOffTrailingZeros(Result)
  return Result
}

func doCalculatios(Value float64, Operations []string) string {
  var Result string = fmt.Sprintf("%f", Value)
  for _, Operation := range Operations {
    Arg1, _ := strconv.ParseFloat(Result, 32)
    if strings.HasPrefix(Operation, "*") {
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "*"), 32)
      Result = fmt.Sprintf("%f", Arg1 * Arg2)
    } else if strings.HasPrefix(Operation, "/") {
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "/"), 32)
      Result = fmt.Sprintf("%f", Arg1 / Arg2)
    } else if strings.HasPrefix(Operation, "+") {
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "+"), 32)
      Result = fmt.Sprintf("%f", Arg1 + Arg2)
    } else if strings.HasPrefix(Operation, "-") {
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "-"), 32)
      Result = fmt.Sprintf("%f", Arg1 - Arg2)
    }
  }
  return Result
}

func cutOffTrailingZeros(Value string) string {
  Chars := []rune(Value)
  var CharsTrimmed []rune

  // Loop over string reverse
  for i := len(Chars) - 1; i >= 0; i-- {
    // If point (46) is reached stop cutting off
    if Chars[i] == 46 {
      // if no places behind the comma, do not add dot
      var startingIndex int
      if len(CharsTrimmed) == 0 {
        startingIndex = i - 1
      } else {
        startingIndex = i
      }
      // Add remaining places and break
      for j := startingIndex; j >= 0; j-- {
        CharsTrimmed = append(CharsTrimmed, Chars[j])
      }
      break
    }
    // If zero (48), continue and dont add to CharsTrimmed
    if Chars[i] == 48 {
      continue
    }
    // Add places behind the comma
    CharsTrimmed = append(CharsTrimmed, Chars[i])
  }
  CharsRev := []rune(CharsTrimmed)
  var Result []rune
  for i := len(CharsRev) - 1; i >= 0; i-- {
    Result = append(Result, CharsRev[i])
  }

  return string(Result)
}

func main() {
//  SINGLE_CHAR_PTR := flag.Bool("s", false, "Alterate single chars")
  flag.Parse()
  var INPUT string
  Operations := make([]string, len(flag.Args()) - 1)
  if len(flag.Args()) == 1 {
    fmt.Println("No alterations found. exiting")
    printUsage()
    os.Exit(1)

  } else {
    var OperationsIndex int8 = 0
    for _, ELEMENT := range flag.Args() {
      if strings.HasPrefix(ELEMENT, "*") || strings.HasPrefix(ELEMENT, "/") ||
        strings.HasPrefix(ELEMENT, "+") || strings.HasPrefix(ELEMENT, "-") {

        Operations[OperationsIndex] = ELEMENT
        OperationsIndex++
      } else {
        INPUT = ELEMENT
      }
    }
  }
  var intStr string = ""
  var output string = ""
  // Loop over INPUT
  for _, CHAR := range INPUT {
    if unicode.IsDigit(CHAR) {
      intStr = fmt.Sprintf("%s%s", string(CHAR), intStr)
    } else {
      // If whole number is over
      if len(intStr) > 0 {
        Result := calcWrapper(intStr, Operations)
        output += Result
      }
      output += string(CHAR)
      intStr = ""
    }
  }
  // If the last char is a number calculate it too
  if len(intStr) > 0 {
    Result := calcWrapper(intStr, Operations)
    output += Result
  }
  fmt.Println( output)
}
