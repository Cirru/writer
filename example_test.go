
package writer

import "io/ioutil"

func ExampleMakeCode(filename string) string {
  content, _ := ioutil.ReadFile(filename)
  return MakeCode(content)
}