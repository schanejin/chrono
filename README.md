# chrono 
Chrono implements the common functions of a stop-watch.
For usage see examples below or click on the godoc badge.

## Install
```bash
go get github.com/schanejin/chrono
```
 ## Example
```go
// create a new chrono, the timer start immediately
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
