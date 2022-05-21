package goms

import (
	"errors"
	"regexp"
	"strconv"
)

const (
	s = 1000
	m = s * 60
	h = m * 60
	d = h * 24
	w = d * 7
	y = d * 365
)

const re = `^((?:\d+)?\.?\d+) *(microseconds?|microsecond|μs|milliseconds?|msecs?|ms|seconds?|secs?|s|minutes?|mins?|m|hours?|hrs?|h|days?|d|w|weeks?|years?|yrs?|y)?$`

var reg = regexp.MustCompile(re)
var ErrUnsupportedUnit = errors.New("Unsupported time unit")

//Parse export ms func
func Parse(val string) (float64, error) {
	if len(val) > 15 {
		return 0, errors.New("string too long")
	}
	matched := reg.FindStringSubmatch(val)
	if len(matched) < 3 {
		return 0, errors.New("Match Error")
	}
	real, err := strconv.ParseFloat(matched[1], 32)
	if err != nil {
		return 0.0, errors.New("parse to float error")
	}
	switch matched[2] {
	case "years", "year", "yrs", "yr", "y":
		return real * y, nil
	case "weeks", "week", "w":
		return real * w, nil
	case "days", "day", "d":
		return real * d, nil
	case "hours", "hour", "hrs", "hr", "h":
		return real * h, nil
	case "minutes", "minute", "min", "m":
		return real * m, nil
	case "seconds", "second", "secs", "sec", "s":
		return real * s, nil
	case "milliseconds", "millisecond", "msecs", "msec", "ms":
		return real, nil
	default:
		return 0, ErrUnsupportedUnit
	}
}
