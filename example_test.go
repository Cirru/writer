
package writer

import (
  "io/ioutil"
  "github.com/Cirru/json-loader"
)

func ExampleMakeCode(filename string) string {
  content, _ := ioutil.ReadFile(filename)
  tree := loader.ParseText(content)
  return MakeCode(tree)
}