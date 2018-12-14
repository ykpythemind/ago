package ago

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

type ti struct {
	start time.Time
	now   time.Time
}

func TestFromNow(t *testing.T) {
	start := parse("2018-12-14T12:34:56")
	fmt.Println(FromNow(start))
	start = parse("2018-12-11T12:34:56")
	fmt.Println(FromNow(start))
	start = parse("2018-10-11T12:34:56")
	fmt.Println(FromNow(start))
}

func TestDetectDistance(t *testing.T) {
	var tests = []struct {
		name     string
		expected distance
		given    ti
	}{
		{"1min", distance{1, MinutesAgo}, ti{parse("2018-12-14T12:34:05"), parse("2018-12-14T12:35:06")}},
		{"12hour", distance{12, HoursAgo}, ti{parse("2018-12-14T10:34:05"), parse("2018-12-14T22:40:03")}},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			start := tt.given.start
			now := tt.given.now
			actual := detectDistance(start, now)
			if actual != tt.expected {
				t.Errorf("expected %s, actual %s", tt.expected.str(), actual.str())
			}

		})
	}
}

func parse(str string) time.Time {
	if !strings.Contains(str, "+") {
		str = str + "+09:00"
	}
	t, err := time.Parse(time.RFC3339, str) // "2006-01-02T15:04:05+07:00"
	if err != nil {
		panic(err)
	}
	return t
}
