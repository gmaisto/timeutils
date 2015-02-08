package timeutils

import "time"

//Constants
var stdDaysInMonth = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

const (
	februaryLeapYearLength = 29
)

//Now a simple wrapper of time.Time
type Now struct {
	time.Time
}

//New return a new instance of Now
func New(n time.Time) *Now {
	return &Now{n}
}

//Tomorrow return a time instance for tomorrows date at the same time
func (t *Now) Tomorrow() time.Time {
	return t.AddDate(0, 0, 1)
}

//TomorrowAtTime return a time instance for tomorrow's date at a given time of day
func (t *Now) TomorrowAtTime(w time.Time) time.Time {
	tn := t.Tomorrow()
	return time.Date(tn.Year(), tn.Month(), tn.Day(), w.Hour(), w.Minute(), w.Second(), w.Nanosecond(), w.Location())
}

//TomorrowDayOfWeek return the day of the week of tomorrow
func (t *Now) TomorrowDayOfWeek() time.Weekday {
	return t.Tomorrow().Weekday()
}

//Yesterday return a time instance for yesterday's date at the same time
func (t *Now) Yesterday() time.Time {
	//return t.Add(-time.Duration(24) * time.Hour)
	return t.AddDate(0, 0, -1)
}

//YesterdayAtTime return a time instance for yesterday's date at a given time of day
func (t *Now) YesterdayAtTime(w time.Time) time.Time {
	tn := t.Yesterday()
	return time.Date(tn.Year(), tn.Month(), tn.Day(), w.Hour(), w.Minute(), w.Second(), w.Nanosecond(), w.Location())
}

//YesterdayDayOfWeek return the day of week of yesterday
func (t *Now) YesterdayDayOfWeek() time.Weekday {
	return t.Yesterday().Weekday()
}

//DayOfWeekAfterNDays return the day of week after a specified number of days from the given time instance
func (t *Now) DayOfWeekAfterNDays(n int) time.Weekday {
	return t.Add(time.Duration(24*n) * time.Hour).Weekday()
}

//IsLeap return true if the Year in t is a leap year
func (t *Now) IsLeap() bool {
	return (t.Year()%400 == 0 || (t.Year()%100 != 0 && t.Year()%4 == 0))
}

//DaysInMonth return the number of days in the month which t is in
func (t *Now) DaysInMonth() int {
	if t.Month() == time.February && t.IsLeap() {
		return februaryLeapYearLength
	}
	return stdDaysInMonth[t.Month()]
}

//AddDays add a number of days at Now
func (t *Now) AddDays(n int) time.Time {
	//return t.Add(time.Duration(24*n) * time.Hour)
	return t.AddDate(0, 0, n)
}

//Since return a time instance representing the time a given seconds since the instance time
func (t *Now) Since(seconds int64) time.Time {
	return t.Add(time.Duration(seconds) * time.Second)
}

//Ago return a time instance representing the time a given number of seconds ago
func (t *Now) Ago(seconds int64) time.Time {
	return t.Add(-time.Duration(seconds) * time.Second)
}

//NextMonth return the next month (time.Month)
func (t *Now) NextMonth() time.Month {
	return t.AddDate(0, 1, 0).Month()
}

//PreviousMonth return the next month (time.Month)
func (t *Now) PreviousMonth() time.Month {
	return t.AddDate(0, -1, 0).Month()
}

//NewTimeAfterInterval return a new time instance obtained adding "length" seconds to the original time instance excluding
//an interval of time
func (t *Now) NewTimeAfterInterval(istart time.Time, iend time.Time, length int64) time.Time {
	td := iend.Sub(istart)
	return t.Add(td + (time.Duration(length) * time.Second))
}

//BeginningOfDay return a time instance for the beginning of the day
func (t *Now) BeginningOfDay() time.Time {
	d := time.Duration(-t.Hour()) * time.Hour
	return t.Truncate(time.Hour).Add(d)
}

//EndOfDay return a time instance for the end of the day
func (t *Now) EndOfDay() time.Time {
	return t.BeginningOfDay().Add(24*time.Hour - time.Nanosecond)
}

//SecondsToEndOfDay return the number of seconds to the end of the day
func (t *Now) SecondsToEndOfDay() int32 {
	dt := t.EndOfDay().Add(-time.Nanosecond).Sub(t.Time)
	return int32(time.Duration(dt) / time.Second)
}
