// Package crhono provides a timer that implements common funcionality
package chrono

import "time"

type Chrono struct {
	start time.Time
}

func New() *Chrono {
	return &Chrono{
		start: time.Now(),
	}
}
func (c *Chrono) ElapsedTime() time.Duration {
	return time.Since(c.start)
}
func (c *Chrono) Stop()  {}
func (c *Chrono) Pause() {}
func (c *Chrono) Reset() {}
func (c *Chrono) Lap()   {}
