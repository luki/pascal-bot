package logic

import (
  "regexp"
  "fmt"
  "strconv"
  "errors"
  "strings"
)

func GetCalculation(equation string) (float64, error) {
  num      := "(?:\\+|-)?(?:\\d+)(?:\\.\\d+)?|[π]"
  operator := "\\+|-|\\*|\\/"

  exp := fmt.Sprintf("(%s)\\s?(?:(%s)\\s?(%s))?", num, operator, num)

  fmt.Println(exp)

  // Setup
  r, e := regexp.Compile(exp)
  if e != nil { }

  res := r.FindStringSubmatch(equation)

  fmt.Println(res[1])
  fmt.Println(replaceSpecialSymbols(res[1]))
  n1, err := strconv.ParseFloat(replaceSpecialSymbols(res[1]), 64)
  if err != nil { return n1, errors.New("Could not parse the first number") }

  // TODO: Check if the 2nd value is really supplied

  // Get Operation

  n2, err := strconv.ParseFloat(replaceSpecialSymbols(res[3]), 64)
  if err != nil { return n2, errors.New("Could not parse the second number") }

  var calc float64 = 0.0

  switch res[2] {
    case "+":
      calc = n1 + n2
    case "-":
      calc = n1 - n2
    case "*":
      calc = n1 * n2
    case "/":
      calc = n1 / n2
    default:
      return calc, errors.New("The operator isn't valid")
  }

  return calc, nil
}

func replaceSpecialSymbols(term string) string {
  return strings.Replace(term, "π", "3.1415962", 1)
}
