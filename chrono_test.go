package chrono

import (
	"strconv"
	"testing"
	"time"
)

func TestChrono(t *testing.T) {
	c := New()
	if c == nil {
		t.Error("New returns a nil struct")
	} else if c.start.IsZero() {
		t.Error("Start time returns a zero time.Time type")
	}
}
func TestChrono_ElapsedTime(t *testing.T) {
	c := New()
	elapsedDurations := make([]time.Duration, 0)
	for i := 1; i < 10; i++ {
		time.AfterFunc(time.Millisecond*100*time.Duration(i), func() {
			elapsedDurations = append(elapsedDurations, c.ElapsedTime())
		})
	}
	time.Sleep(time.Second)
	for i, elapsed := range elapsedDurations {
		ms := int(RoundFloat(float64(elapsed/time.Millisecond), 0))
		n := (i + 1) * 100
		if ms != n {
			t.Errorf("ElapsedTime: got: %d expected: %d\n", ms, n)
		}
	}
}
func TestChrono_Stop(t *testing.T)  {}
func TestChrono_Pause(t *testing.T) {}
func TestChrono_Reset(t *testing.T) {}
func TestChrono_Lap(t *testing.T)   {}

// return rounded version of x with prec precision
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
