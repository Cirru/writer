
package writer

type placement int

const (
  byDollar placement = iota
  byIndent
  byComma
  byAppend
  bySentence
)
