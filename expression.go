
package writer

type expression struct {
  list []interface{}
  place placement
}

func (e *expression) len() int {
  return len(e.list)
}

func (e *expression) isNested() bool {
  for _, child := range(e.list) {
    if _, ok := child.(*expression); ok {
      return true
    }
  }
  return false
}

func (e *expression) format() string {
  ret := "("
  for index, child := range(e.list) {
    if index > 0 {
      ret += " "
    }
    if expr, ok := child.(*expression); ok {
      ret += expr.format()
    } else if tok, ok := child.(*token); ok {
      ret += tok.format()
    }
  }
  return (ret + ")")
}