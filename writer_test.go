
package writer

import (
  "testing"
  "io/ioutil"
)

func TestMakeCode(t *testing.T) {
  code := ExampleMakeCode("json/demo.json")
  println(code)
}