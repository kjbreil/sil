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

// JulianTime is the julien date with a time formatted after
func JulianTime() string {
	n := time.Now()
	return fmt.Sprintf("%v %02d:%02d:%02d", JulianNow(), n.Hour(), n.Minute(), n.Second())
}
