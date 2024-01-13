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

// or stop the stopWatch
c.Stop()

// resume the timer afeter a reset/stop
c.Start()

// create a lap
lap1 := c.Lap()
lap2 := c.Lap()
lap3 := c.Lap()

// get a list of all lap durations
list := c.Laps()

//string representation of stopwatch
fmt.Printf("chrono: %s", c)
```
#### Basic Operations

## Credits

* [Schanejin](https://github.com/schanejin)
