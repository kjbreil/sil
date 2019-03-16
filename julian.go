package sil

import (
	"fmt"
	"time"
)

// JulianDate takes a time.Time and turns it into a julian date - just formatting
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
	return fmt.Sprintf("%v %s", JulianNow(), JulianTimePart(n))
}

// JulianTimePart returns the time part of a julian formatted time
func JulianTimePart(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}
