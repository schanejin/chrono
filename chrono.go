// Package crhono provides a timer that implements common funcionality
package chrono

import (
	"fmt"
	"time"
)

type Chrono struct {
	start   time.Time
	stop    time.Time
	elapsed time.Duration
}

func New() *Chrono {
	return &Chrono{
		start: time.Now(),
	}
}
func (c *Chrono) ElapsedTime() time.Duration {
	c.elapsed = time.Since(c.start)
	return c.elapsed
}

// Stop stops the timer. To resume the timer Start() needs to be called again
func (c *Chrono) Stop() {
	c.elapsed = c.ElapsedTime()
	c.stop = time.Now()
}

// Start resumes or starts the timer. if a Stop() was invoked it resumes the
// timer. if a Reset() was invoked it starts a nes session.
func (c *Chrono) Start() {
	if c.start.IsZero() { //reseted
		c.start = time.Now()
	} else { //stopped
		c.start = c.start.Add(time.Since(c.stop))
	}
}
func (c *Chrono) Pause() {}
func (c *Chrono) Reset() {
	c.start = time.Time{}
}
func (c *Chrono) Lap() {}

func (c *Chrono) String() string {
	return fmt.Sprintf("[start: %s current: %s elapsed: %s]\n",
		c.start.Format(time.Stamp), time.Now().Format(time.Stamp), c.ElapsedTime())
}
