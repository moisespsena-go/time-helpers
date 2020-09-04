package time_helpers

import (
	"strconv"
	"time"
	"fmt"
)

var t time.Time

// OffsetS returns location offset in format -07:00
func OffsetS(loc *time.Location) string {
	time.Now().String()
	return t.In(loc).Format("-07:00")
}

// OffsetS returns hour and minute location offset
func Offset(loc *time.Location) (hour, minute int) {
	offset := OffsetS(loc)
	hourS, minuteS := offset[0:3], offset[4:]
	hour, _ = strconv.Atoi(hourS)
	minute, _ = strconv.Atoi(minuteS)
	return
}
// OffsetS returns location offset in format -07:00
func OffsetMS(loc *time.Location) string {
	time.Now().String()
	return t.In(loc).Format("-07:00 MST")
}

// OffsetS returns hour and minute location offset
func OffsetM(loc *time.Location) (hour, minute, monotonic int) {
	offset := OffsetS(loc)
	hourS, minuteS := offset[0:3], offset[4:]
	hour, _ = strconv.Atoi(hourS)
	minute, _ = strconv.Atoi(minuteS)
	return
}

// LocToGmtS returns location in format "GMT -07:00"
func LocToGmtS(loc *time.Location) string {
	return "GMT " + OffsetS(loc)
}

// SetLoc retuns new time with location
func SetLoc(t time.Time, loc *time.Location) (r time.Time) {
	r = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), loc)
	return
}

// ParseHtml5Input parses html5 datetime-local field value
func ParseHtml5Input(v string, loc ...*time.Location) (t time.Time, err error) {
	var l *time.Location
	for _, l = range loc {
		if l != nil {
			break
		}
	}
	if l == nil {
		l = time.Local
	}

	fmt.Println(v)
	if t, err = time.ParseInLocation("2006-01-02T15:04", v, l); err != nil {
		return
	}
	fmt.Println(t)
	return
}
