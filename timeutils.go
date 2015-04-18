package timeutils

import "time"

//Constants
var stdDaysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

const (
	februaryLeapYearLength = 29
)

//Timeutils a simple wrapper of time.Time
type Timeutils struct {
	time.Time
}

//New return a new instance of Timeutils
func New(n time.Time) *Timeutils {
	return &Timeutils{n}
}

//Tomorrow return a time instance for tomorrows date at the same time
func (t *Timeutils) Tomorrow() time.Time {
	return t.AddDate(0, 0, 1)
}

//TomorrowAtTime return a time instance for tomorrow's date at a given time of day
func (t *Timeutils) TomorrowAtTime(w time.Time) time.Time {
	tn := t.Tomorrow()
	return time.Date(tn.Year(), tn.Month(), tn.Day(), w.Hour(), w.Minute(), w.Second(), w.Nanosecond(), w.Location())
}

//TomorrowDayOfWeek return the day of the week of tomorrow
func (t *Timeutils) TomorrowDayOfWeek() time.Weekday {
	return t.Tomorrow().Weekday()
}

//Yesterday return a time instance for yesterday's date at the same time
func (t *Timeutils) Yesterday() time.Time {
	//return t.Add(-time.Duration(24) * time.Hour)
	return t.AddDate(0, 0, -1)
}

//YesterdayAtTime return a time instance for yesterday's date at a given time of day
func (t *Timeutils) YesterdayAtTime(w time.Time) time.Time {
	tn := t.Yesterday()
	return time.Date(tn.Year(), tn.Month(), tn.Day(), w.Hour(), w.Minute(), w.Second(), w.Nanosecond(), w.Location())
}

//YesterdayDayOfWeek return the day of week of yesterday
func (t *Timeutils) YesterdayDayOfWeek() time.Weekday {
	return t.Yesterday().Weekday()
}

//DayOfWeekAfterNDays return the day of week after a specified number of days from the given time instance
func (t *Timeutils) DayOfWeekAfterNDays(n int) time.Weekday {
	return t.Add(time.Duration(24*n) * time.Hour).Weekday()
}

//IsLeap return true if the Year in t is a leap year
func (t *Timeutils) IsLeap() bool {
	return (t.Year()%400 == 0 || (t.Year()%100 != 0 && t.Year()%4 == 0))
}

//DaysInMonth return the number of days in the month which t is in
func (t *Timeutils) DaysInMonth() int {
	if t.Month() == time.February && t.IsLeap() {
		return februaryLeapYearLength
	}
	return stdDaysInMonth[t.Month()]
}

//AddDays add a number of days at Timeutils
func (t *Timeutils) AddDays(n int) time.Time {
	//return t.Add(time.Duration(24*n) * time.Hour)
	return t.AddDate(0, 0, n)
}

//Since return a time instance representing the time a given seconds since the instance time
func (t *Timeutils) Since(seconds int64) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

//Ago return a time instance representing the time a given number of seconds ago
func (t *Timeutils) Ago(seconds int64) time.Time {
	return t.Add(-time.Duration(seconds) * time.Second)
}

//NextMonth return the next month (time.Month)
func (t *Timeutils) NextMonth() time.Month {
	return t.AddDate(0, 1, 0).Month()
}

//PreviousMonth return the next month (time.Month)
func (t *Timeutils) PreviousMonth() time.Month {
	return t.AddDate(0, -1, 0).Month()
}

//NewTimeAfterInterval return a new time instance obtained adding "length" seconds to the original time instance excluding
//an interval of time
func (t *Timeutils) NewTimeAfterInterval(istart time.Time, iend time.Time, length int64) time.Time {
	td := iend.Sub(istart)
	return t.Add(td + (time.Duration(length) * time.Second))
}

//BeginningOfDay return a time instance for the beginning of the day
func (t *Timeutils) BeginningOfDay() time.Time {
	d := time.Duration(-t.Hour()) * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

//EndOfDay return a time instance for the end of the day
func (t *Timeutils) EndOfDay() time.Time {
	return t.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

//SecondsToEndOfDay return the number of seconds to the end of the day
func (t *Timeutils) SecondsToEndOfDay() int32 {
	dt := t.EndOfDay().Add(-time.Nanosecond).Sub(t.Time)
	return int32(time.Duration(dt) / time.Second)
}

//Between check the given time instance if between to time instances
func (t *Timeutils) Between(start time.Time, end time.Time) bool {
	return t.Before(end) && t.After(start)
}

func mod(a, n int) int {
	r := a % n
	if r < 0 {
		return r + n
	}
	return r
}

//Easter the Easter day for the year of t
func (t *Timeutils) Easter() time.Time {
	y := t.Year()
	c := y / 100
	n := mod(y, 19)
	i := mod(c-c/4-(c-(c-17)/25)/3+19*n+15, 30)
	i -= (i / 28) * (1 - (i/28)*(29/(i+1))*((21-n)/11))
	l := i - mod(y+y/4+i+2-c+c/4, 7)
	m := 3 + (l+40)/44
	d := l + 28 - 31*(m/4)
	return time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.UTC)
}
