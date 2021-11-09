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

func calcWrapper(IntStr string, Operations []string) string {
  IntToCalc, _ := strconv.ParseFloat(IntStr, 32)
  Result := doCalculatios(IntToCalc, Operations)
  Result = cutOffTrailingZeros(Result)
  return Result
}

func doCalculatios(Value float64, Operations []string) string {
  var Result string = fmt.Sprintf("%f", Value)
  for _, Operation := range Operations {
    if strings.HasPrefix(Operation, "*") {
      Arg1, _ := strconv.ParseFloat(Result, 32)
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "*"), 32)
      Result = fmt.Sprintf("%f", Arg1 * Arg2)
    } else if strings.HasPrefix(Operation, "/") {
      Arg1, _ := strconv.ParseFloat(Result, 32)
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "/"), 32)
      Result = fmt.Sprintf("%f", Arg1 / Arg2)
    } else if strings.HasPrefix(Operation, "+") {
      Arg1, _ := strconv.ParseFloat(Result, 32)
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "+"), 32)
      Result = fmt.Sprintf("%f", Arg1 + Arg2)
    } else if strings.HasPrefix(Operation, "-") {
      Arg1, _ := strconv.ParseFloat(Result, 32)
      Arg2, _ := strconv.ParseFloat(strings.Trim(Operation, "-"), 32)
      Result = fmt.Sprintf("%f", Arg1 - Arg2)
    }
  }
  return Result
}

func cutOffTrailingZeros(Value string) string {
  Chars := []rune(Value)
  var ResultRev string
  var TrailingZeros bool = true
  // Loop over string reverse
  for i := len(Chars) - 1; i >= 0; i-- {
    if (string(Chars[i]) == "." || string(Chars[i]) == "0") && TrailingZeros {
      continue
    }
    if string(Chars[i]) != "0" {
      TrailingZeros = false
    }
    ResultRev = ResultRev + string(Chars[i])
  }

  CharsRev := []rune(ResultRev)
  var Result []rune
  for i := len(CharsRev) -1; i >= 0; i-- {
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
  var IntStr string = ""
  // Loop over INPUT
  for _, CHAR := range INPUT {
    if unicode.IsDigit(CHAR) {
      IntStr = fmt.Sprintf("%s%s", string(CHAR), IntStr)
    } else {
      // If whole number is over
      if len(IntStr) > 0 {
        Result := calcWrapper(IntStr, Operations)
        fmt.Print(Result)
      }
      fmt.Print(string(CHAR))
      IntStr = ""
    }
  }
  // If the last char is a number calculate it too
  if len(IntStr) > 0 {
    Result := calcWrapper(IntStr, Operations)
    fmt.Print(Result)
  }
  fmt.Println()
}
