
Cirru Writer
------

Format Cirru AST into intented format.

[![GoDoc](https://godoc.org/github.com/Cirru/json-loader?status.png)][godoc]

### Usage

Read [GoDoc][godoc] for details.

[godoc]: https://godoc.org/github.com/Cirru/json-loader

### Rules

* short expressions are inlined with parentheses

```cirru
a (b c)
```

* complicated ones may need to start in new lines

```cirru
a
  b (c d)
  e f
```

* last item in expression can be appended with `$`

```cirru
a $ b $ c d
```

* only first pair of parentheses are suggested in the same line

```cirru
a (b c)
  d e
```

* use `,` to simplify structures

```cirru
a
  b c
  , d
  e f
```

* start new lines when theres too many parameters

```cirru
a b c d
  , e f g h
  , i j k
```

* add empty lines between top-level expressions

```cirru
a

b
```

### License

MIT