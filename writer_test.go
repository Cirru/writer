
package writer

import (
  "testing"
  "io/ioutil"
)

func TestMakeCode(t *testing.T) {
  names := []string {"demo",
    "folding",
    "html",
    "indent",
    "line",
    "parentheses",
    "quote",
    "spaces",
    "unfolding",
  }
  for _, name := range(names) {
    jsonPath := "json/" + name + ".json"
    cirruPath := "cirru/" + name + ".cirru"
    code := ExampleMakeCode(jsonPath)
    cirruBytes, _ := ioutil.ReadFile(cirruPath)
    cirruCode := string(cirruBytes)

    if code == cirruCode {
      println("----", name)
    } else {
      println("\n####", name)
      println(code)
    }
  }
}