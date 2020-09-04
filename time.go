package time_helpers

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func TimeF(t time.Time) string {
	if t.Second() > 0 {
		return fmt.Sprintf("%02d:%02d:%02dH", t.Hour(), t.Minute(), t.Second())
	}
	if t.Minute() > 0 {
		return fmt.Sprintf("%02d:%02dH", t.Hour(), t.Minute())
	}
	return fmt.Sprintf("%02dH", t.Hour())
}

func TimeS(t time.Time) string {
	return fmt.Sprintf("%02d:%02d:%02d", t.Hour(), t.Minute(), t.Second())
}

func ParseTime(now time.Time, v string) (t time.Time, err error) {
	defer func() {
		if err != nil {
			err = ErrBadFormat.Cause(err)
		}
	}()
	parts := strings.Split(v, ":")
	if len(parts) == 0 || len(parts) > 3 {
		return
	}
	var (
		h, m, s int
	)
	if h, err = strconv.Atoi(parts[0]); err != nil {
		return
	} else if h < 0 || h > 23 {
		err = ErrBadHour.WithData(h)
		return
	}
	if len(parts) > 1 {
		if m, err = strconv.Atoi(parts[1]); err != nil {
			return
		} else if m < 0 || m > 59 {
			err = ErrBadMinute.WithData(m)
			return
		}
		if len(parts) > 2 {
			if s, err = strconv.Atoi(parts[2]); err != nil {
				return
			} else if s < 0 || s > 59 {
				err = ErrBadSecond.WithData(s)
				return
			}
		}
	}
	t = time.Date(now.Year(), now.Month(), now.Day(), h, m, s, 0, now.Location())
	return
}

func MustParseTime(now time.Time, v string) time.Time {
	t, err := ParseTime(now, v)
	if err != nil {
		panic(err)
	}
	return t
}

