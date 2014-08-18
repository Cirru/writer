
package writer

import (
  "github.com/Cirru/json-loader"
  "fmt"
)

func MakeCode(content []byte) string {
  tree := loader.ParseText(content)
  expr := convertSlice(tree)
  markTree(expr)
  return "fake"
}

func convertSlice(tree []interface{}) *expression {
  list := []interface{}{}
  for _, child := range(tree) {
    if expr, ok := child.([]interface{}); ok {
      list = append(list, convertSlice(expr))
    } else if tok, ok := child.(string); ok {
      list = append(list, &token{tok, byAppend})
    }
  }
  return &expression{list, byAppend}
}

func markTree(tree *expression) {
  // last := tree.len() - 1
  tree.place = byIndent
  for _, child := range(tree.list) {
    if expr, ok := child.(*expression); ok {
      expr.place = byIndent
      markTree(expr)
    } else if tok, ok := child.(*token); ok {
      tok.place = byIndent
    } else {
      panic("got something else")
    }
  }
}
