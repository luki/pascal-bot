package logic

import (
  "regexp"
  "fmt"
  "strconv"
  "errors"
)

func GetCalculation(equation string) (float64, error) {
  num      := "(?:\\+|-)?(?:\\d+)(?:\\.\\d+)?"
  operator := "\\+|-|\\*|\\/"

  exp := fmt.Sprintf("(%s)\\s?(?:(%s)\\s?(%s))?", num, operator, num)

  // Setup
  r, e := regexp.Compile(exp)
  if e != nil { }

  res := r.FindStringSubmatch(equation)

  n1, err := strconv.ParseFloat(res[1], 64)
  if err != nil { return n1, errors.New("Could not parse the first number") }

  // TODO: Check if the 2nd value is really supplied

  // Get Operation

  n2, err := strconv.ParseFloat(res[3], 64)
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
