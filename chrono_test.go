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
func TestChrono_Start(t *testing.T) {
	c := New()
	time.Sleep(time.Millisecond * 30)
	c.Stop()
	time.Sleep(time.Millisecond * 60)
	c.Start()
	time.Sleep(time.Millisecond * 30)
	ms := int(RoundFloat(float64(c.ElapsedTime()/time.Millisecond), 0))
	if ms != 60 {
		t.Errorf("Start: got: %d expected: %d\n", ms, 60)
	}
}
func TestChrono_ElapsedTime(t *testing.T) {
	c := New()
	elapsedDurations := make([]time.Duration, 0)
	for i := 1; i < 10; i++ {
		time.AfterFunc(time.Millisecond*10*time.Duration(i), func() {
			elapsedDurations = append(elapsedDurations, c.ElapsedTime())
		})
	}
	time.Sleep(time.Second)
	for i, elapsed := range elapsedDurations {
		ms := int(RoundFloat(float64(elapsed/time.Millisecond), 0))
		n := (i + 1) * 10
		if ms != n {
			t.Errorf("ElapsedTime: got: %d expected: %d\n", ms, n)
		}
	}
}
func TestChrono_Stop(t *testing.T) {
	c := New()
	time.Sleep(time.Millisecond * 30)
	c.Stop()
	time.Sleep(time.Millisecond * 20)

	ms := int(RoundFloat(float64(c.ElapsedTime()/time.Millisecond), 0))
	if ms != 30 {
		t.Errorf("Stop: got: %d expected: %d\n", ms, 30)
	}
}
func TestChrono_Pause(t *testing.T) {}
func TestChrono_Reset(t *testing.T) {
	c := New()
	c.Reset()
	if !c.start.IsZero() {
		t.Error("Reset should reset the initial start timer")
	}
	c.Start()
	time.Sleep(time.Millisecond * 40)
	ms := int(RoundFloat(float64(c.ElapsedTime()/time.Millisecond), 0))
	if ms != 40 {
		t.Errorf("Reset: got: %d expected: %d\n", ms, 40)
	}
}
func TestChrono_Lap(t *testing.T) {
	c := New()
	time.Sleep(time.Millisecond * 10)
	lap1 := c.Lap()
	time.Sleep(time.Millisecond * 20)
	lap2 := c.Lap()
	time.Sleep(time.Millisecond * 30)
	lap3 := c.Lap()
	ms1 := int(RoundFloat(float64(lap1/time.Millisecond), 0))
	ms2 := int(RoundFloat(float64(lap2/time.Millisecond), 0))
	ms3 := int(RoundFloat(float64(lap3/time.Millisecond), 0))

	if ms1 != 10 && ms2 != 20 && ms3 != 30 {
		t.Errorf("Lap: got: %d %d %d, expecting: %d %d %d\n",
			ms1, ms2, ms3, 10, 20, 30)
	}
	if len(c.laps) != 3 {
		t.Error("Lap: number of laps should be 3")
	}
	c.Stop()
	l := c.Lap()
	if l != time.Duration(0) {
		t.Errorf("Lap: stopwatch is stopped but lap returns %d\n", l)
	}

	n := New()
	n.Reset()

	u := n.Lap()
	if u != time.Duration(0) {
		t.Errorf("Lap: stopwatch is resetted but lap returns %d\n", u)
	}

}

// return rounded version of x with prec precision
func RoundFloat(x float64, prec int) float64 {
	frep := strconv.FormatFloat(x, 'g', prec, 64)
	f, _ := strconv.ParseFloat(frep, 64)
	return f
}
