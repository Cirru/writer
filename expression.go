
package writer

type expression struct {
  list []interface{}
  place placement
}

func (e *expression) len() int {
  return len(e.list)
}