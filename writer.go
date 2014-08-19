
package writer

import (
  "github.com/Cirru/json-loader"
  // "fmt"
)

func MakeCode(content []byte) string {
  tree := loader.ParseText(content)
  expr := convertSlice(tree)
  markLines(expr)
  return formatLines(expr)
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

func markLines(tree *expression) {
  for _, child := range(tree.list) {
    if expr, ok := child.(*expression); ok {
      expr.place = bySentence
      markTree(expr)
    }
  }
}

func markTree(tree *expression) {
  end := tree.len() - 1
  okForParens := true
  lastPlace := byIndent
  for index, child := range(tree.list) {
    if expr, ok := child.(*expression); ok {
      if index == end {
        expr.place = byDollar
      } else if index == 0 {
        expr.place = byAppend
      } else if expr.isNested() {
        expr.place = byIndent
      } else if (lastPlace == byAppend) && okForParens {
        expr.place = byAppend
        okForParens = false
      } else {
        expr.place = byIndent
      }
      markTree(expr)
      lastPlace = expr.place
      if lastPlace == byIndent {
        okForParens = true
      }
    } else if tok, ok := child.(*token); ok {
      if lastPlace == byIndent {
        tok.place = byComma
      } else {
        tok.place = byAppend
      }
      lastPlace = tok.place
    } else {
      panic("got something else")
    }
  }
}

func formatLines(tree *expression) string {
  buffer := ""
  for _, child := range(tree.list) {
    if expr, ok := child.(*expression); ok {
      buffer = buffer + "\n\n"
      formatTree(expr, "", &buffer)
    } else {
      panic("should get expression")
    }
  }
  return buffer
}

func formatTree(tree *expression, ss string, cursor *string) {
  atLineHead := true
  for _, child := range(tree.list) {
    if expr, ok := child.(*expression); ok {
      switch expr.place {
      case byAppend:
        if !atLineHead {
          *cursor += " "
        }
        *cursor += expr.format()
      case byIndent:
        *cursor += "\n"
        sss := ss + "  "
        *cursor += sss
        formatTree(expr, sss, cursor)
      case byDollar:
        if !atLineHead {
          *cursor += " "
        }
        *cursor += "$ "
        formatTree(expr, ss, cursor)
      default: panic("should not have anoter option for expressions")
      }
    } else if tok, ok := child.(*token); ok {
      switch tok.place {
      case byAppend:
        if !atLineHead {
          *cursor += " "
        }
        *cursor += tok.format()
      case byComma:
        *cursor += "\n"
        *cursor += ss
        *cursor += ", "
        *cursor += tok.format()
      default: panic("should not have anoter option for tokens")
      }
    }
    atLineHead = false
  }
}