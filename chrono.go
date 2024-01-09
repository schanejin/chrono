// Package crhono for time measurements
package chrono

import "time"

type Chrono struct{}

func New() *Chrono                           { return &Chrono{} }
func (c *Chrono) ElapsedTime() time.Duration { return 0 }
func (c *Chrono) Stop()                      {}
func (c *Chrono) Pause()                     {}
func (c *Chrono) Reset()                     {}
func (c *Chrono) Lap()                       {}
