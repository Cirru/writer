
package writer

import (
  "regexp"
)

type token struct {
  text string
  place placement
}

func (t *token) format() string {
  r, _ := regexp.Compile("[\\ \\(\\)\\\"]")
  if r.MatchString(t.text) {
    ret := "\""
    for _, char := range(t.text) {
      switch string(char) {
      case "\"": ret += "\\\""
      case "\n": ret += "\\\n"
      case "\t": ret += "\\\t"
      default: ret += string(char)
      }
    }
    return (ret + "\"")
  }
  return t.text
}