
package writer

type placement int

const (
  byDollar placement << iota
  byIndent placement
  byComma placement
  byAppend placement
  byParen placement
)
