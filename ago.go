package ago

import (
	"fmt"
	"time"
)

type Ago struct {
	start *time.Time
	now   *time.Time
}

type AgoUnit int

const (
	SecondsAgo AgoUnit = iota
	MinutesAgo
	HoursAgo
	DaysAgo
	// WeeksAgo
	MonthsAgo
	YearsAgo
)

const hoursInYear = 24 * 365
const hoursInMonth = 24 * 30

var unitMap map[AgoUnit]string = map[AgoUnit]string{
	SecondsAgo: "秒前",
	MinutesAgo: "分前",
	HoursAgo:   "時間前",
	DaysAgo:    "日前",
	MonthsAgo:  "ヶ月前",
	YearsAgo:   "年前",
}

type distance struct {
	count int
	unit  AgoUnit
}

func FromNow(start time.Time) string {
	ago := newAgo(&start, nil)
	distance := detectDistance(*ago.start, *ago.now)
	return distance.String()
}

func newAgo(start *time.Time, now *time.Time) Ago {
	if start == nil {
		n := time.Now()
		start = &n
	}
	if now == nil {
		n := time.Now()
		now = &n
	}

	return Ago{
		start: start,
		now:   now,
	}
}

func detectDistance(start time.Time, now time.Time) distance {
	duration := start.Sub(now)

	if 24 < duration.Hours() {
		if hoursInYear < duration.Hours() {
			years := duration.Hours() / hoursInYear
			return newDistance(int(years), YearsAgo)
		} else if hoursInMonth < duration.Hours() {
			months := duration.Hours() / hoursInMonth
			return newDistance(int(months), MonthsAgo)
		} else {
			days := duration.Hours() / 24
			return newDistance(int(days), DaysAgo)
		}
	} else {
		if 1 <= duration.Hours() {
			return newDistance(int(duration.Hours()), HoursAgo)
		} else if 1 <= duration.Minutes() {
			return newDistance(int(duration.Minutes()), MinutesAgo)
		} else {
			return newDistance(int(duration.Seconds()), SecondsAgo)
		}
	}
}

func newDistance(count int, unit AgoUnit) distance {
	return distance{count, unit}
}

func (d *distance) String() string {
	unit := unitMap[d.unit]
	return fmt.Sprintf("%d%s", d.count, unit)
}

func (d *distance) str() string {
	return fmt.Sprintf("count = %d : unit = %+v", d.count, d.unit)
}
