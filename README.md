# chrono [![Go Reference](https://pkg.go.dev/badge/github.com/schanejin/chrono.svg)](https://pkg.go.dev/github.com/schanejin/chrono)
Chrono implements the common functions of a stop-watch.
For usage see examples below or click on the godoc badge.

## Install
```bash
go get github.com/schanejin/chrono
```
 ## Example
```go
// create a set with zero items
c := chrono.New()

// get elapased duration at any time
duration := c.ElapsedTime()

// reset the stopwatch
c.Reset()

//string representation of stopwatch
fmt.Printf("chrono: %s", c)
```
#### Basic Operations

## Credits

* [Schanejin](https://github.com/schanejin)
