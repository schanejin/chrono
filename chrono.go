// Package crhono provides a timer that implements common funcionality
package chrono

import (
	"fmt"
	"time"
)

type Chrono struct {
	start, stop, lap time.Time
	laps             []time.Duration
}

func New() *Chrono {
	c := &Chrono{}
	c.init()
	return c
}
func (c *Chrono) init() {
	c.start, c.lap = time.Now(), time.Now()
	c.laps = make([]time.Duration, 0)
}
func (c *Chrono) ElapsedTime() time.Duration {
	if c.stop.After(c.start) {
		return c.stop.Sub(c.start)
	}
	return time.Since(c.start)
}

// Stop stops the timer. To resume the timer Start() needs to be called again
func (c *Chrono) Stop() {
	c.stop = time.Now()
}

// Start resumes or starts the timer. if a Stop() was invoked it resumes the
// timer. if a Reset() was invoked it starts a nes session.
func (c *Chrono) Start() {
	if c.start.IsZero() { //reseted
		c.init()
	} else { //stopped
		c.start = c.start.Add(time.Since(c.stop))
	}
}
func (c *Chrono) Pause() {}
func (c *Chrono) Reset() {
	c.start, c.stop, c.lap = time.Time{}, time.Time{}, time.Time{}
	c.laps = nil
}

// Lap takes and stores the current lap time and returns the elapsed time
// since the latest lap.
func (c *Chrono) Lap() time.Duration {
	if c.stop.After(c.start) || c.lap.IsZero() {
		return time.Duration(0)
	}
	lap := time.Since(c.lap)
	c.lap = time.Now()
	c.laps = append(c.laps, lap)
	return lap

}

func (c *Chrono) Laps() []time.Duration {
	return c.laps
}

// String returns the representation of a single Chrono instance
func (c *Chrono) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]\n",
		c.start.Format(time.Stamp), time.Now().Format(time.Stamp), c.ElapsedTime())
}
