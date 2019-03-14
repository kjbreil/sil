package sil

import (
	"fmt"
	"time"
)

// JulianDate takes a time.Time and turns it into a julien date - just formatting
func JulianDate(t time.Time) string {
	return fmt.Sprintf("%04d%03d", t.Year(), t.YearDay())
}

// JulianNow returns the julian date for right now
func JulianNow() string {
	return JulianDate(time.Now())
}

// JulianTimeNow returns the JulianDate with time for right now
func JulianTimeNow() string {
	n := time.Now()
	return fmt.Sprintf("%v %02d:%02d:%02d", JulianNow(), n.Hour(), n.Minute(), n.Second())
}

// JulianTimePart takes
func JulianTimePart(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
