# read

Convenient wrappers for reading inputs.

## Examples

```go
fmt.Print("input> ")
line, err := read.Line()
if err != nil {
	panic(err)
}
fmt.Println(line)
```

## Usage

Use shortcut functions like `read.Int()`, `read.Line()` for simple programs.
Use `read.New()`, which is an alias for `read.NewBuffered()` for better performance.
